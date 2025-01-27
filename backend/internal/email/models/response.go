package models

type SearchDocumentsResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}

type ListDocumentsResponse struct {
	List []Index `json:"list"`
	Page struct {
		PageNum  int `json:"page_num"`
		PageSize int `json:"page_size"`
		Total    int `json:"total"`
	} `json:"page"`
}

type Index struct {
	Name string `json:"name"`
}

type ErrorResponseType struct {
	Message      string `json:"message"`
	StatusCode   int    `json:"status_code"`
	Error        bool   `json:"error"`
	DetailsError string `json:"details_error,omitempty"`
}
