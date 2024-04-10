package project

import (
	"log/slog"
)

type Handler struct {
	log *slog.Logger
}

func New(log *slog.Logger) *Handler {
	return &Handler{log: log}
}

//
//func (h *Handler) Index() http.HandlerFunc {
//
//}
//
//func (h *Handler) Store() http.HandlerFunc {
//
//}
//
//func (h *Handler) Show() http.HandlerFunc {
//
//}
//
//func (h *Handler) Update() http.HandlerFunc {
//
//}
//
//func (h *Handler) Delete() http.HandlerFunc {
//
//}
