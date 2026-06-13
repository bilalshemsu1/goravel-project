package controllers

import (
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

// UserController extend Controller
type UserController struct{}

// public function index()
func (r *UserController) Index(ctx http.Context) http.Response {
	return ctx.Response().View().Make("user/index.tmpl", map[string]any{
		"List": "User List",
		"name": []string{"bilal", "kamil", "buddy"},
	})
}

// public function show(){...}
func (r *UserController) Show(ctx http.Context) http.Response {
	id := ctx.Request().Route("id")
	_, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Response().String(400, "Number only allowed")
	}

	return ctx.Response().String(200, "hellp world "+id)

}
