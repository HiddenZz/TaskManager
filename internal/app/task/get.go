package task

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (hd Handler) GetById(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "expected id - got empty string", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error when parsing id", http.StatusBadRequest)
		return
	}

	task, err := hd.repository.GetById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(Dto{
		Id:        task.Id(),
		Name:      task.Name(),
		Desc:      task.Desk(),
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
