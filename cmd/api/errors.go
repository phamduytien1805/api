package api

import (
	"fmt"
	"net/http"

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

func (app *application) logError(r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message string, err any) {
	render.Render(w, r, &ErrResponse{
		HTTPStatusCode: status,
		Reason:         message,
		Errors:         err,
	})
}
func (app *application) errorResponseDefault(w http.ResponseWriter, r *http.Request, status int, message string) {
	app.errorResponse(w, r, status, message, nil)
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponseDefault(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponseDefault(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponseDefault(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponseDefault(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := fmt.Sprintf("Request body is not valid")
	app.errorResponse(w, r, http.StatusUnprocessableEntity, message, app.validator.ValidatorErrors(err))
}

func (app *application) editConflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "unable to update the record due to an edit conflict, please try again"

	app.logError(r, err)

	if err != nil {
		message = err.Error()
	}
	app.errorResponseDefault(w, r, http.StatusConflict, message)
}

func (app *application) invalidAuthenticateResponse(w http.ResponseWriter, r *http.Request, err error) {
	message := "Fail to authenticate user"

	app.logError(r, err)

	if err != nil {
		message = err.Error()
	}
	app.errorResponseDefault(w, r, http.StatusUnauthorized, message)
}
