package handler

import (
	"awesomeProject/internal/email/models"
	"awesomeProject/internal/email/service"
	"awesomeProject/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"strconv"
)

type EmailHandler struct {
	service *service.EmailService
}

func NewEmailHandler(service *service.EmailService) *EmailHandler {
	return &EmailHandler{service: service}
}

func (h *EmailHandler) IndexEmailToZinc() {
	dir := os.Getenv("EMAIL_DIRECTORY")

	if dir == "" {
		log.Fatal("The EMAIL_DIRECTORY environment variable is required")
	}

	err := h.service.IndexEmailsInBulk(dir)

	if err != nil {
		log.Fatalf("Error indexing emails: %s", err)
		return
	}
	fmt.Println("Successfully indexed emails")
}

func (h *EmailHandler) SearchEmailInZinc(w http.ResponseWriter, r *http.Request) {
	var request models.SearchRequest
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed decode json", err.Error())
		return
	}
	if request.Limit == 0 {
		request.Limit = 10
	}

	response, err := h.service.SearchEmailsInZinc(request.Query, request.Limit, request.Offset, request.StartDate, request.EndDate, request.NameIndex)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Failed search emails", err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error encoding response", err.Error())
		return
	}
}

func (h *EmailHandler) ListIndex(w http.ResponseWriter, r *http.Request) {
	// obtain query params
	pageNum := r.URL.Query().Get("page_num")
	pageSize := r.URL.Query().Get("page_size")
	sortBy := r.URL.Query().Get("sort_by")
	desc := r.URL.Query().Get("desc")
	name := r.URL.Query().Get("name")

	if pageNum == "" {
		pageNum = "1"
	}
	if pageSize == "" {
		pageSize = "10"
	}

	pageNumInt, err := strconv.Atoi(pageNum)
	if err != nil {
		pageNumInt = 1 // Valor por defecto si no se puede convertir
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10 // Valor por defecto si no se puede convertir
	}

	request := models.ListDocumentsRequest{
		PageNum:  pageNumInt,
		PageSize: pageSizeInt,
		SortBy:   sortBy,
		Desc:     desc,
		Name:     name,
	}

	result, err := h.service.ListDocuments(&request)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error searching documents", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// JSON response
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error encoding response", err.Error())
		return
	}
}

func (h *EmailHandler) DeleteIndex(w http.ResponseWriter, r *http.Request) {
	indexName := chi.URLParam(r, "index_name")
	if indexName == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error encoding response")
		return
	}

	result, err := h.service.DeleteIndex(indexName)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error deleting index", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Error encoding response", err.Error())
	}
}
