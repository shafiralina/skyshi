package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"skyshi/models"
	"strconv"
)

var CreateActivities = func(w http.ResponseWriter, r *http.Request) {
	activity := &models.Activities{}
	db := models.GetDB()

	err := json.NewDecoder(r.Body).Decode(activity)
	if err != nil {
		models.RespondError(w, models.Message("failed", "error decode request body"), http.StatusBadRequest)
		return
	}

	resp := activity.Create(db)
	models.Respond(w, resp)
}

var UpdateActivities = func(w http.ResponseWriter, r *http.Request) {
	activity := &models.Activities{}
	db := models.GetDB()

	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	resp := activity.Update(id, db)
	models.Respond(w, resp)
}

var GetActivities = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	resp := models.GetActivities(int64(id), db)
	if resp["status"].(string) != "Success" {
		models.RespondError(w, models.Message(resp["status"].(string), resp["message"].(string)), http.StatusNotFound)
		return
	}
	models.Respond(w, resp)
}

var GetAllActivities = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	if r.Method != http.MethodGet {
		models.RespondError(w, models.Message("failed", "request method tidak valid"), http.StatusBadRequest)
		return
	}
	resp := models.GetAllActivities(db)
	models.Respond(w, resp)
}

var DeleteActivities = func(w http.ResponseWriter, r *http.Request) {
	db := models.GetDB()
	paramId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		models.RespondError(w, models.Message("failed", "bad request"), http.StatusBadRequest)
		return
	}

	resp := models.DeleteActivities(int64(id), db)
	if resp["status"].(string) != "Success" {
		models.RespondError(w, models.Message(resp["status"].(string), resp["message"].(string)), http.StatusNotFound)
		return
	}
	models.Respond(w, resp)
}
