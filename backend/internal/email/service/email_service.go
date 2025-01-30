package service

import (
	"awesomeProject/internal/email/models"
	"awesomeProject/internal/email/repository"
	"awesomeProject/internal/utils"
	"fmt"
	"strings"
)

type EmailService struct {
	repo *repository.EmailRepository
}

func NewEmailService(repo *repository.EmailRepository) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) IndexEmailsInBulk(dir string) error {

	files, err := utils.ReadFileFromDir(dir)
	if err != nil {
		return fmt.Errorf("an error occurred while reading files  %s: %s", dir, err.Error())

	}

	emailsBulk, err := utils.ProcessEmailInParallel(files, 4)
	if err != nil {
		return fmt.Errorf("an error occurred while processing emails: %s", err.Error())
	}

	if len(emailsBulk) > 0 {

		chunkedEmails := chunkEmails(emailsBulk, 1000)

		for _, chunk := range chunkedEmails {
			s.repo.IndexEmailsToZinInBulk(chunk)
		}

		fmt.Println("Emails indexed successfully in bulk")
	} else {
		fmt.Println("No bulk emails found")
	}
	return nil
}

func chunkEmails(emails []*models.Email, chunkSize int) [][]*models.Email {
	var chunks [][]*models.Email
	for i := 0; i < len(emails); i += chunkSize {
		end := i + chunkSize
		if end > len(emails) {
			end = len(emails)
		}
		chunks = append(chunks, emails[i:end])
	}
	return chunks
}

func (s *EmailService) SearchEmailsInZinc(query string, limit, offset int, startDate, endDate, nameIndex string) (*models.SearchDocumentsResponse, error) {

	var mustClauses []string

	if startDate != "" && endDate != "" {
		rangeFilter := fmt.Sprintf(`{
		"range": {
			"@timestamp": {
				"gte": "%s",
				"lt": "%s",
				"format": "2006-01-02T15:04:05Z07:00"
			}
		}
	}`, startDate, endDate)
		mustClauses = append(mustClauses, rangeFilter)
	}

	if query != "" {
		queryBody := fmt.Sprintf(`{
		"query_string": {
			"query": "%s"
		}
	}`, query)
		mustClauses = append(mustClauses, queryBody)
	} else {
		mustClauses = append(mustClauses, `{"match_all": {}}`)
	}

	var mustSection string
	if len(mustClauses) > 0 {
		mustSection = strings.Join(mustClauses, ",")
	} else {
		mustSection = "{}"
	}

	requestBody := fmt.Sprintf(`{
	"query": {
		"bool": {
			"must": [%s]
		}
	},
	"from": %d,
	"size": %d,
	"aggs": {
		"histogram": {
			"date_histogram": {
				"field": "@timestamp",
				"fixed_interval": "30s"
			}
		}
	},
	"sort": [
		{
			"@timestamp": {
				"order": "desc"
			}
		}
	]
}`, mustSection, offset, limit)

	bodyReader := strings.NewReader(requestBody)
	response, err := s.repo.SearchEmailsInZinc(bodyReader, nameIndex)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *EmailService) ListDocuments(req *models.ListDocumentsRequest) ([]string, error) {
	result, err := s.repo.ListIndex(req)
	if err != nil {
		return nil, fmt.Errorf("error gettings documents: %s", err.Error())
	}
	return result, nil
}

func (s *EmailService) DeleteIndex(indexName string) (interface{}, error) {
	return s.repo.DeleteIndex(indexName)
}
