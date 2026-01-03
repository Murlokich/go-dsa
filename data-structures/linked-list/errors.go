package linked_list

import "errors"

var (
	ErrEmptyList   = errors.New("operation cannot be applied on empty list")
	ErrNoSuchValue = errors.New("requested value does not exist in the list")
)
