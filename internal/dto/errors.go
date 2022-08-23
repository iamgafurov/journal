package dto

import "errors"

var (
	ErrNoRows         = errors.New("no rows in result set")
	ErrNoRowsAffected = errors.New("no rows are affected")
)
