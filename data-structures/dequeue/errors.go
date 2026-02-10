package dequeue

import "errors"

var (
	ErrEmptyDequeue = errors.New("operation cannot be applied on empty dequeue")
)
