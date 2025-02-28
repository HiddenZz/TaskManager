package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"taskmanager.com/helpers/parse"
	l "taskmanager.com/pkg/logger"
)

func (hd Handler) GetById(w http.ResponseWriter, r *http.Request) {

	id, err := parse.IdStr(r.PathValue("id"))

	if err != nil {
		l.E(err)
		http.Error(w, fmt.Sprintf("bad id: %v", err), http.StatusBadRequest)
		return
	}

	task, err := hd.repository.GetById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		l.E(err)
		return
	}

	j, err := json.Marshal(ResponseDto{
		Id:        task.Id(),
		Name:      task.Name(),
		Desc:      task.Desc(),
		CreatedAt: task.CreateDate(),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
