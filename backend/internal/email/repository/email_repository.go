package repository

import (
	"awesomeProject/internal/email/models"
	"awesomeProject/internal/zincsearch"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type EmailRepository struct {
	client *zincsearch.ZincClient
}

func NewEmailRepository(client *zincsearch.ZincClient) *EmailRepository {
	return &EmailRepository{client: client}
}

func (r *EmailRepository) IndexEmailsToZinInBulk(emails []*models.Email) {
	var bulkRequest []byte

	for _, email := range emails {
		action := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": r.client.IndexName,
			},
		}

		actionJSON, err := json.Marshal(action)
		if err != nil {
			log.Fatalf("Error marshalling action: %v", err)
		}

		// Convert email to json and covert to bulk
		emailJSON, err := json.Marshal(email)
		if err != nil {
			log.Fatalf("Error marshalling email: %v", err)
		}

		bulkRequest = append(bulkRequest, append(actionJSON, '\n')...)
		bulkRequest = append(bulkRequest, append(emailJSON, '\n')...)
	}

	// build ZincSearch url
	url := fmt.Sprintf("%s/api/_bulk", r.client.Host)

	// create request http_request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bulkRequest))
	if err != nil {
		log.Fatalf("error to send bulk request http_request: %s", err.Error())
	}

	// setting headers
	req.SetBasicAuth(r.client.User, r.client.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("error to send bulk request http_request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("Emails indexed successfully", resp.StatusCode)
	} else {
		log.Printf("Error when indexing bulk request http_request: %v\n", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	log.Println("Emails indexed successfully", resp.StatusCode, string(body))
}

func (r *EmailRepository) SearchEmailsInZinc(query io.Reader, nameIndex string) (*models.SearchDocumentsResponse, error) {

	url := fmt.Sprintf("%s/es/%s/_search", r.client.Host, nameIndex)
	req, err := http.NewRequest("POST", url, query)
	if err != nil {
		log.Fatalf("Error to create search request http_request: %v", err)
	}

	req.SetBasicAuth(r.client.User, r.client.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to execute request http_request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error the read error body response: %w", err)
		}

		return nil, fmt.Errorf("%s", string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error the read success body response: %v", err)
	}

	var response models.SearchDocumentsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("error deserializing response %v", err)
	}

	return &response, nil
}

func (r *EmailRepository) ListIndex(req *models.ListDocumentsRequest) ([]string, error) {
	url := fmt.Sprintf("%s/api/index?page_num=%d&page_size=%d&sort_by=%s&desc=%s",
		r.client.Host,
		req.PageNum,
		req.PageSize,
		req.SortBy,
		req.Desc,
	)

	reqHTTPRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error making the GET request: %v", err)
	}

	reqHTTPRequest.SetBasicAuth(r.client.User, r.client.Password)
	reqHTTPRequest.Header.Set("Content-Type", "application/json")
	reqHTTPRequest.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(reqHTTPRequest)
	if err != nil {
		return nil, fmt.Errorf("error making the GET request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error the read body response: %v", err)
	}

	var response *models.ListDocumentsResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("error deserializing response %v", err)
	}

	// Extract the indices name
	var indexNames []string
	for _, index := range response.List {
		if index.Name != "" {
			indexNames = append(indexNames, index.Name)
		}
	}
	return indexNames, nil
}

func (r *EmailRepository) DeleteIndex(indexName string) (interface{}, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/api/index/%s", r.client.Host, indexName), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.client.User, r.client.Password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to execute request http_request: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error the read body response: %v", err)
	}

	var responseBody interface{}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		return nil, fmt.Errorf("error deserializing response %v", err)
	}
	return responseBody, nil
}
