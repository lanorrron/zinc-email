package zincsearch

import "net/http"

type ZincClient struct {
	Client    *http.Client
	Host      string
	User      string
	Password  string
	IndexName string
}

func NewZincClient(host, user, password string, indexName ...string) *ZincClient {
	if len(indexName) == 0 {
		indexName = append(indexName, "")
	}
	return &ZincClient{
		Client:    &http.Client{},
		Host:      host,
		User:      user,
		Password:  password,
		IndexName: indexName[0],
	}
}
