package author

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
	authorsURL = "/authors"
	authorURL  = "/authors/:name"
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

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, authorsURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, authorsURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, authorURL, apperror.Middleware(h.GetAuthorByName))
	router.HandlerFunc(http.MethodPut, authorURL, apperror.Middleware(h.UpdateUser))
	router.HandlerFunc(http.MethodPatch, authorURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, authorURL, apperror.Middleware(h.DeleteUser))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("this is API error")
}

func (h *handler) GetAuthorByName(w http.ResponseWriter, r *http.Request) error {

	//nameWithRemPref := strings.ReplaceAll(r.RequestURI, "/authors/", "")
	//
	//
	//result, err := h.repository.FindOne(context.Background(), nameWithRemPref)
	//if err != nil {
	//	w.WriteHeader(http.StatusNoContent)
	//	return fmt.Errorf("author doesnt exist. error: %v", err)
	//}
	//resultJson, err := json.Marshal(result)
	//if err != nil {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return err
	//}
	//w.WriteHeader(200)
	//w.Write(resultJson)

	return nil
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.Write([]byte("UpdateAuthor"))
	w.WriteHeader(204)

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("PartiallyUpdateAuthor"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("DeleteAuthor"))

	return nil
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {

	result, err := h.repository.FindAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return fmt.Errorf("author doesnt exist. error: %v", err)
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
