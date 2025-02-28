package task

import (
	"fmt"
	"net/http"
	"taskmanager.com/helpers/parse"
	l "taskmanager.com/pkg/logger"
)

func (hd Handler) Delete(w http.ResponseWriter, r *http.Request) {

	id, err := parse.IdStr(r.PathValue("id"))

	if err != nil {
		l.E(err)
		http.Error(w, fmt.Sprintf("bad id: %v", err), http.StatusBadRequest)
		return
	}

	err = hd.repository.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when delete task with id %d", id), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
