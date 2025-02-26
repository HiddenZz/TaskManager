package task

import (
	"encoding/json"
	"net/http"
	domain "taskmanager.com/internal/domain/tasks"
)

func (hd Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto CreateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	task, err := hd.repository.Create(r.Context(), func() (*domain.Task, error) {
		return domain.CreateTask(dto.Name, dto.Desc)
	})
	if err != nil {
		http.Error(w, "create task competed error", http.StatusBadRequest)
		return
	}

	j, err := json.Marshal(ResponseDto{
		Id:        task.Id(),
		Desc:      task.Desc(),
		CreatedAt: task.CreateDate(),
		Name:      task.Name(),
	})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
