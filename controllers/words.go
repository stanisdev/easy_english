package controllers

import (
	"net/http"
	"app/structures"
	_ "fmt"
)

func (e *Env) GetWord(w http.ResponseWriter, r *http.Request) {
	e.Json(w, map[string]interface{} {
		"word": "Language",
	})
}

func (e *Env) AddWord(w http.ResponseWriter, r *http.Request) {
	word := &structures.Word{}
	e.ParseBody(r, word)
	e.Json(w, map[string]interface{} {
		"value": word.Value,
	})
}