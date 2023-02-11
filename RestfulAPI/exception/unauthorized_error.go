package exception

type UnauthorizedError struct {
	Error string `json:"error,omitempty"`
}

func NewUnAutohorizedError(error string) UnauthorizedError {
	return UnauthorizedError{Error: error}
}
