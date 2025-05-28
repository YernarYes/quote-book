package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"quotes/internal/core/domain/quote"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	log     *slog.Logger
	service quote.Service
}

func NewHandler(service quote.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

// to implement errors

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var view View
	if err := json.NewDecoder(r.Body).Decode(&view); err != nil {
		h.log.Error("failed to decode json", "err", err)
		http.Error(w, "failed to decode json", http.StatusBadRequest)
		return
	}

	if err := h.service.Create(r.Context(), view.ToModel()); err != nil {
		h.log.Error("failed to create a quote", "err", err)
		http.Error(w, "failed to create a quote", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(view); err != nil {
		h.log.Error("error encoding json", "err", err)
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	if author == "" {
		h.log.Warn("field is empty", "author", author)
	}

	filter := quote.Filter{
		Author: author,
	}

	quotes, err := h.service.Get(r.Context(), filter)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		h.log.Error("failed to get quotes", "err", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(quotes); err != nil {
		h.log.Error("error encoding json", "err", err)
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing ID", http.StatusBadRequest)
		h.log.Error("failed to get ID")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "failed to parse quote", http.StatusBadRequest)
		h.log.Error("failed to parse id", "err", err)
		return
	}

	filter := quote.Filter{
		ID: id,
	}

	if err := h.service.Delete(r.Context(), filter); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("Failed to delete", "err", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetRandom(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.Random(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.log.Error("failed to get random quote", "err", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(quote); err != nil {
		h.log.Error("error encoding json", "err", err)
		http.Error(w, "error encoding json", http.StatusInternalServerError)
		return
	}
}
