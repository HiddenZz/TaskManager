package parse

import (
	"errors"
	"strconv"
)

var (
	IdIsEmpty     = errors.New("id must not be empty")
	IdFailedParse = errors.New("failed parse id")
	IdNotValid    = errors.New("id not valid")
)

func IdStr(param string) (int, error) {

	if param == "" {
		return 0, IdIsEmpty
	}
	id, err := strconv.Atoi(param)

	if err != nil {
		return id, IdFailedParse
	}

	if id <= 0 {
		return id, IdNotValid
	}
	return id, nil
}
