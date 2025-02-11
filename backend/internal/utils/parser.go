package utils

import (
	"awesomeProject/internal/email/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Parse(filePath string) (*models.Email, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	email := &models.Email{}
	inBody := false

	reader := bufio.NewReader(file)
	var bodyBuilder strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			// Si no es EOF y hay otro error, lo manejamos
			return nil, fmt.Errorf("Error reading line: %v", err)
		}

		if err != nil && err.Error() == "EOF" {
			break
		}

		line = strings.TrimSpace(line)

		if line == "" {
			inBody = true
			continue
		}

		if inBody {
			if strings.HasPrefix(line, "-----Original Message-----") {
				break
			}

			if strings.HasPrefix(line, "On") || strings.HasPrefix(line, ">") {
				break
			}

			bodyBuilder.WriteString(line + "\n") //invstigar 
		} else {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				switch key {
				case "Message-ID":
					email.MessageID = value
				case "Date":
					email.Date = value
				case "From":
					email.From = value
				case "To":
					email.To = value
				case "Subject":
					email.Subject = value
				case "Mime-Version":
					email.MimeVersion = value
				case "Content-Type":
					email.ContentType = value
				case "Content-Transfer-Encoding":
					email.ContentTransferEncoding = value
				case "X-From":
					email.XFrom = value
				case "X-To":
					email.XTo = value
				case "X-cc":
					email.Xcc = value
				case "X-bcc":
					email.Xbcc = value
				case "X-Folder":
					email.XFolder = value
				case "X-Origin":
					email.XOrigin = value
				case "X-FileName":
					email.XFileName = value
				case "Cc":
					email.Cc = value
				}
			}
		}
	}
	email.Body = bodyBuilder.String()
	return email, nil
}
