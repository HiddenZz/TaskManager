package tasks

import (
	"fmt"
	"time"
)

type Task struct {
	id         int
	name       string
	desc       string
	createDate time.Time
}

func NewTask(id int, name string, desc string, createDate time.Time) (*Task, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}

	return &Task{id: id, name: name, desc: desc, createDate: createDate}, nil
}

func TemplateTask(id int, name string, desc string) *Task {
	return &Task{
		id:   id,
		name: name,
		desc: desc,
	}
}

func CreateTask(name string, desc string) (*Task, error) {
	return NewTask(0, name, desc, time.Now())
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("name is requeired")
	}
	return nil
}

func (t *Task) Id() int {
	return t.id
}
func (t *Task) Name() string {
	return t.name
}

func (t *Task) Desc() string {
	return t.desc
}
func (t *Task) CreateDate() time.Time {
	return t.createDate
}
