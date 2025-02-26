package task

import "time"

type CreateRequestDto struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ResponseDto struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type UpdateRequestDto struct {
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`
	Id   int    `json:"id"`
}
