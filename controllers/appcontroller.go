package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// BaseController Base controller.
type BaseController struct {
	beego.Controller
}

// serveError Serves an error response.
// @param status HTTP response status.
// @param message Error message.
func (c *BaseController) serveError(status int, message string) {
	c.Data["json"] = APIError{StatusCode: status, ErrorMessage: message}
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
	c.StopRun()
}

// APIError Api error message.
type APIError struct {
	ErrorCode    string `json:"error_code"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

// Error Return the string error.
func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage)
}
