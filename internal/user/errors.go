package user

import "errors"

var (
	ErrorUserResourceConflict = errors.New("username or email are used")
)
