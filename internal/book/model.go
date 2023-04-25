package book

import "github.com/Salavei/golang_advanced_restapi/internal/author"

type Book struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Authors []author.Author `json:"authors"`
}
