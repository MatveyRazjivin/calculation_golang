package calculation

import "errors"

var (
	ErrDivisionByZero = errors.New("На ноль делить нельзя")
	ErrInvalidToken   = errors.New("Недопустимый токен")
	ErrIndexOfRange   = errors.New("Недостаточный размер массива")
)
