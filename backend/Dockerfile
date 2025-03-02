FROM golang:1.23-alpine

#Workdir
WORKDIR /backend

#Files the backend
COPY go.mod go.sum ./

#Install dependencies
RUN go mod tidy

#Copy source code
COPY . .

#build the binary
RUN go build -o server ./cmd/server/main.go

#Port
EXPOSE 3000

#Init comand
CMD [ "./server" ]