package controllers

import (
	"app/structures"
	"app/middlewares"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

type Env struct {
	Some string
}

func Start() {
	r := mux.NewRouter()
	env := &Env{}
	r.Use(middlewares.Logging)
	r.HandleFunc("/get", env.GetWord).Methods("GET")
	r.HandleFunc("/add", env.AddWord).Methods("POST")
	http.ListenAndServe(":5080", r)
}

func (e *Env) Json(w http.ResponseWriter, data map[string]interface{}) {
	jsonOut, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonOut))
}

func (e *Env) ParseBody(r *http.Request, instance structures.BaseStruct) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(instance)
	if err != nil {
		panic(err)
	}
}