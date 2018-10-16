package utils

type Err struct {
	s string
}

func (e *Err) Error() string {
	return e.s
}
