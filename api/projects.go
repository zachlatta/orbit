package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/zachlatta/orbit"
)

func serveProject(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["ID"])
	if err != nil {
		return err
	}

	project, err := store.Projects.Get(id)
	if err != nil {
		return err
	}

	return writeJSON(w, project)
}

func serveCreateProject(w http.ResponseWriter, r *http.Request) error {
	var project orbit.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		return err
	}

	// TODO create logic

	if err := store.Projects.Create(&project); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return writeJSON(w, project)
}