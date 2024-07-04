package store

import "errors"

//Файл, содержащий специальные стандартные ошибки
var (
	ErrRecordNotFound = errors.New("record not found")
)
