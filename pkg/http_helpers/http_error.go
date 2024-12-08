package http_helpers

import (
	"fmt"
	"net/http"
	"phamduytien1805/pkg/validator"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	AppCode string `json:"code,omitempty"`   // application-specific error code
	Reason  string `json:"reason,omitempty"` // user-level status message
	Errors  any    `json:"errors,omitempty"` // user-level status message

}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// func (app *application) logError(r *http.Request, err error) {
// 	var (
// 		method = r.Method
// 		uri    = r.URL.RequestURI()
// 	)

// 	app.logger.Error(err.Error(), "method", method, "uri", uri)
// }

func errorResponse(w http.ResponseWriter, r *http.Request, status int, message string, err any) {
	render.Render(w, r, &ErrResponse{
		HTTPStatusCode: status,
		Reason:         message,
		Errors:         err,
	})
}
func errorResponseDefault(w http.ResponseWriter, r *http.Request, status int, message string) {
	errorResponse(w, r, status, message, nil)
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "the server encountered a problem and could not process your request"
	errorResponseDefault(w, r, http.StatusInternalServerError, message)
}

func NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	errorResponseDefault(w, r, http.StatusNotFound, message)
}

func MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	errorResponseDefault(w, r, http.StatusMethodNotAllowed, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	errorResponseDefault(w, r, http.StatusBadRequest, err.Error())
}

func FailedValidationResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := fmt.Sprintf("Request body is not valid")
	errorResponse(w, r, http.StatusUnprocessableEntity, message, validator.ValidatorErrors(err))
}

func EditConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "unable to update the record due to an edit conflict, please try again"
	if err != nil {
		message = err.Error()
	}
	errorResponseDefault(w, r, http.StatusConflict, message)
}

func InvalidAuthenticateResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "Fail to authenticate user"
	if err != nil {
		message = err.Error()
	}
	errorResponseDefault(w, r, http.StatusUnauthorized, message)
}