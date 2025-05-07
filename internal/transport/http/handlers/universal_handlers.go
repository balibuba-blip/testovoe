package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testovoe/internal/interfaces"
	"testovoe/internal/service"
)

type UniversalHandler struct {
	service *service.UniversalService
}

func NewUniversalHandler(s *service.UniversalService) *UniversalHandler {
	return &UniversalHandler{service: s}
}

func (h *UniversalHandler) GetAllEntities(w http.ResponseWriter, r *http.Request) {
	// Парсим параметры
	entityType := r.URL.Query().Get("type")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	if entityType != "product" && entityType != "measure" {
		http.Error(w, "Invalid entity type", http.StatusBadRequest)
		return
	}
	res, err := h.service.GetAllEntities(interfaces.EntityType(entityType), r.Context(), limit, offset)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
