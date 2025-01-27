package models

type ListDocumentsRequest struct {
	PageNum  int    `json:"page_num"`
	PageSize int    `json:"page_size"`
	SortBy   string `json:"sort_by"`
	Desc     string `json:"desc"`
	Name     string `json:"name"`
}

type SearchRequest struct {
	Query     string `json:"query"`
	Limit     int    `json:"limit"`
	Offset    int    `json:"offset"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	NameIndex string `json:"name_index"`
}
