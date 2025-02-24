package task

import "time"

type Dto struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
