package middlewares

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danyducky/social/app"
	"github.com/danyducky/social/utils"
	"github.com/gin-gonic/gin"
)

// Response middleware for override future response.
// This middlewares need to be latest in global scope.
type ResponseMiddleware struct {
	ctx app.Context
}

// Creates response middleware instance.
func NewResponseMiddleware(ctx app.Context) ResponseMiddleware {
	return ResponseMiddleware{
		ctx: ctx,
	}
}

const (
	okResponse          = "Ok"
	serverErrorResponse = "Something went wrong. Please retry."
	badRequestResponse  = "Bad request"
)

var (
	OkStatuses = []int{http.StatusOK, http.StatusCreated}
)

// Setup this middleware.
func (m ResponseMiddleware) Setup() {
	middleware := func(ctx *gin.Context) {
		var wb *app.ResponseBuffer

		// Here we gonna rewrite to application response writer and manage it.
		if w, ok := ctx.Writer.(gin.ResponseWriter); ok {
			wb = app.NewResponseBuffer(w)
			ctx.Writer = wb
			ctx.Next()
		} else {
			ctx.Next()
			return
		}

		var response app.Response
		status := ctx.Writer.Status()
		var data interface{}
		json.Unmarshal(wb.Body.Bytes(), &data)
		wb.Body.Reset()

		if utils.Contains(OkStatuses, status) {
			response = app.OkResponse("Ok", data)
		} else if status == http.StatusInternalServerError {
			response = app.EmptyResponse(false, serverErrorResponse)
		} else {
			errors := app.MapContextErrors(ctx)
			response = app.BadResponse(badRequestResponse, errors)
		}

		body, err := json.Marshal(response)
		if err != nil {
			panic(err.Error())
		}

		wb.Body.Write(body)

		wb.Header().Set("Content-Type", "application/json")
		wb.Header().Set("Content-Length", strconv.Itoa(wb.Body.Len()))

		wb.Flush()
	}

	m.ctx.ApiGroup.Use(middleware)
}
