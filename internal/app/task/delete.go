package task

import (
	"fmt"
	"net/http"
	"strconv"
)

func (hd Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "error when parsing id", http.StatusBadRequest)
		return
	}

	err = hd.repository.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error when delete task with id %d", id), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
