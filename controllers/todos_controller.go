package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"skyshi/models"
	"strconv"
)

var CreateTodos = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todos{}
	db := models.GetDB()

	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		models.RespondError(w, models.Message("failed", "error decode request body"), http.StatusBadRequest)
		return
	}

	resp := todo.Create(db)
	models.Respond(w, resp)
}

var UpdateTodos = func(w http.ResponseWriter, r *http.Request) {
	todo := &models.Todos{}
	db := models.GetDB()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	resp := todo.Update(int64(id), db)
	models.Respond(w, resp)
}

var GetTodos = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	resp := models.GetTodos(int64(id), db)
	if resp["status"].(string) != "Success" {
		models.RespondError(w, models.Message(resp["status"].(string), resp["message"].(string)), http.StatusNotFound)
		return
	}
	models.Respond(w, resp)
}

var GetAllTodos = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	if r.Method != http.MethodGet {
		models.RespondError(w, models.Message("failed", "request method tidak valid"), http.StatusBadRequest)
		return
	}

	paramActivity := r.URL.Query().Get("activity_group_id")
	activity_group_id, _ := strconv.Atoi(paramActivity)

	resp := models.GetAllTodos(int64(activity_group_id), db)
	models.Respond(w, resp)
}

var DeleteTodos = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	paramId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		models.RespondError(w, models.Message("failed", "bad request"), http.StatusBadRequest)
		return
	}

	resp := models.DeleteTodos(int64(id), db)
	if resp["status"].(string) != "Success" {
		models.RespondError(w, models.Message(resp["status"].(string), resp["message"].(string)), http.StatusNotFound)
		return
	}
	models.Respond(w, resp)
}
