package api

type ErrorResponse struct {
	Err string
}

type error interface {
	Error() string
}
