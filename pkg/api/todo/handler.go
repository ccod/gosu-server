package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

	m "github.com/ccod/gosu-server/pkg/middleware"
	"github.com/ccod/gosu-server/pkg/models"
	re "github.com/ccod/gosu-server/pkg/response"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func list(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todos []models.Todo
	db.Find(&todos)
	re.RespondJSON(todos, w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todo models.Todo
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Could not convert id", 500)
		return
	}

	db.First(&todo, id)
	re.RespondJSON(todo, w, r)
}

func create(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Could not decode body into todo", 500)
		return
	}

	db.Create(&todo)
	re.RespondJSON(todo, w, r)
}

func update(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todo, updatedTodo models.Todo
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Could not convert id", 500)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		http.Error(w, "Could not decode body into todo", 500)
		return
	}

	db.First(&todo, id)
	if updatedTodo.Description != todo.Description {
		todo.Description = updatedTodo.Description
	}

	if updatedTodo.Completed != todo.Completed {
		todo.Completed = updatedTodo.Completed
	}

	db.Save(&todo)
	re.RespondJSON(todo, w, r)
}

func delete(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todo models.Todo
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Could not convert id", 500)
		return
	}

	db.First(&todo, id)
	db.Delete(&todo)
	re.RespondJSON(todo, w, r)
}

func completed(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(m.DBKey).(*gorm.DB)
	var todo models.Todo
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Could not convert id", 500)
		return
	}

	db.First(&todo, id)
	todo.Completed = true
	db.Save(&todo)
	re.RespondJSON(todo, w, r)
}
