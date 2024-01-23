package errors

import "errors"

func DbError() error {
	return errors.New("db error")
}

func NoTaskError() error {
	return errors.New("no task found")
}
