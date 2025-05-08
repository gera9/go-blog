package shareddtos

import "net/http"

type ErrResponse struct {
	StatusCode  int    `json:"status_code,omitempty"`
	DevMessage  string `json:"dev_message,omitempty"`
	UserMessage string `json:"user_message,omitempty"`
}

func (err ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(err.StatusCode)
	return nil
}

func NewBadRequestErr(devErr error, userErr ...error) ErrResponse {
	return NewErrResponse(http.StatusBadRequest, devErr, userErr...)
}

func NewInternalServerErr(devErr error, userErr ...error) ErrResponse {
	return NewErrResponse(http.StatusInternalServerError, devErr, userErr...)
}

func NewErrResponse(statusCode int, devErr error, userErr ...error) ErrResponse {
	userErrAux := devErr
	if len(userErr) > 0 {
		userErrAux = userErr[0]
	}
	return ErrResponse{
		StatusCode:  statusCode,
		DevMessage:  devErr.Error(),
		UserMessage: userErrAux.Error(),
	}
}
