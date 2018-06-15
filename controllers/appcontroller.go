package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// serveError Serves an error response.
// @param status HTTP response status.
// @param message Error message.
func (c *BaseController) serveError(status int, message string) {
	c.Data["json"] = ApiError{StatusCode: status, ErrorMessage: message}
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
	c.StopRun()
}

type ApiError struct {
	ErrorCode    string `json:"error_code"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.ErrorMessage)
}
