package utils

import (
	"awesomeProject/internal/email/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Parse lee el archivo y crea el objeto Email
func Parse(filePath string) (*models.Email, error) {
	// Abrimos el archivo
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	email := &models.Email{}
	inBody := false

	// Usamos bufio.Reader para leer línea por línea
	reader := bufio.NewReader(file)
	var bodyBuilder strings.Builder

	for {
		// Leemos una línea
		line, err := reader.ReadString('\n')
		if err != nil && err.Error() != "EOF" {
			// Si no es EOF y hay otro error, lo manejamos
			return nil, fmt.Errorf("Error leyendo línea: %v", err)
		}

		// Si llegamos al final del archivo, salimos del bucle
		if err != nil && err.Error() == "EOF" {
			break
		}

		// Limpiamos los espacios innecesarios de la línea
		line = strings.TrimSpace(line)

		// Si encontramos una línea vacía, comenzamos a procesar el cuerpo
		if line == "" {
			inBody = true
			continue
		}

		// Procesamos el cuerpo si estamos en esa sección
		if inBody {
			bodyBuilder.WriteString(line + "\n")
		} else {
			// Procesamos las cabeceras
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				// Asignamos las cabeceras al objeto Email
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

	// Asignamos el cuerpo procesado al email
	email.Body = bodyBuilder.String()

	// Validamos que al menos el Message-ID esté presente
	if email.MessageID == "" {
		return nil, fmt.Errorf("error: el campo 'Message-ID' está vacío")
	}

	return email, nil
}
