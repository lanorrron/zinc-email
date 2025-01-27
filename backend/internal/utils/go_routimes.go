package utils

import (
	"awesomeProject/internal/email/models"
	"fmt"
	"log"
	"sync"
)

func ProcessEmailInParallel(files []string, nunGoroutines int) ([]*models.Email, error) {
	numFiles := len(files)
	batchSize := (numFiles + nunGoroutines - 1) / nunGoroutines

	var wg sync.WaitGroup
	resultsChan := make(chan []*models.Email, nunGoroutines)
	errorChan := make(chan error, 1)

	processBatch := func(startIndex, endIndex int, goroutineID int) {
		defer wg.Done()
		var emailsBulk []*models.Email

		// process the files from lots
		for i := startIndex; i < endIndex && i < numFiles; i++ {
			emailParsed, err := Parse(files[i])
			if err != nil {
				errorChan <- fmt.Errorf("Error parsing email at index %d: %v", i, err)
				return
			}
			emailsBulk = append(emailsBulk, emailParsed)
		}
		resultsChan <- emailsBulk
	}

	for i := 0; i < nunGoroutines; i++ {
		startIndex := i * batchSize
		endIndex := startIndex + batchSize

		if endIndex > numFiles {
			endIndex = numFiles
		}
		wg.Add(1)
		go processBatch(startIndex, endIndex, i)
	}

	wg.Wait()
	close(resultsChan)
	close(errorChan)

	// collect results
	var allEmails []*models.Email
	for emails := range resultsChan {
		allEmails = append(allEmails, emails...)
	}

	select {
	case err := <-errorChan:
		if err != nil {
			log.Printf("Error processing emails: %v", err)
			return nil, err
		}
	default:
	}
	return allEmails, nil
}
