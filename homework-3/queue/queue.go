package queue

import "errors"

// CreateQueue - as factory with closure
func CreateQueue() (func(value interface{}), func() (interface{}, error)) {
	q := []interface{}{}

	push := func(value interface{}) {
		q = append(q, value)
	}

	shift := func() (interface{}, error) {
		if len(q) == 0 {
			return nil, errors.New("Queue has no elements")
		}

		val := q[0]
		q = q[1:]

		return val, nil
	}

	return push, shift
}

// Queue - as structure
type Queue struct {
	slice []interface{}
}

// Shift -
func (q *Queue) Shift() (interface{}, error) {
	if len(q.slice) == 0 {
		return nil, errors.New("Queue has no elements")
	}

	val := q.slice[0]
	q.slice = q.slice[1:]

	return val, nil
}

// Push -
func (q *Queue) Push(val interface{}) {
	q.slice = append(q.slice, val)
}
