package book

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Salavei/golang_advanced_restapi/internal/apperror"
	"github.com/Salavei/golang_advanced_restapi/internal/handlers"
	"github.com/Salavei/golang_advanced_restapi/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var _ handlers.Handler

const (
	booksURL = "/books"
	bookURL  = "/books/:id"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h handler) Register(router *httprouter.Router) {

	router.HandlerFunc(http.MethodGet, booksURL, apperror.Middleware(h.GetList))
	//router.HandlerFunc(http.MethodPost, booksURL, apperror.Middleware(h.CreateBook))
	//router.HandlerFunc(http.MethodGet, bookURL, apperror.Middleware(h.GetBookByUUID))
	//router.HandlerFunc(http.MethodPut, bookURL, apperror.Middleware(h.UpdateBook))
	//router.HandlerFunc(http.MethodPatch, bookURL, apperror.Middleware(h.PartiallyUpdateBook))
	//router.HandlerFunc(http.MethodDelete, bookURL, apperror.Middleware(h.DeleteUser))
}

//func (h handler) GetBookByUUID(w http.ResponseWriter, r *http.Request) error {
//
//}

func (h handler) GetList(w http.ResponseWriter, r *http.Request) error {
	result, err := h.repository.FindAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return fmt.Errorf("book doesnt exist. error: %v", err)
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
	return nil
}
