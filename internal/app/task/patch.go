package task

import (
	"encoding/json"
	"net/http"
	domain "taskmanager.com/internal/domain/tasks"
)

func (hd Handler) Patch(w http.ResponseWriter, r *http.Request) {
	var dto UpdateRequestDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	task, err := hd.repository.Update(r.Context(), func() (*domain.Task, error) {
		return domain.TemplateTask(dto.Id, dto.Name, dto.Desc), nil
	})
	if err != nil {
		http.Error(w, "update task competed error", http.StatusBadRequest)
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
