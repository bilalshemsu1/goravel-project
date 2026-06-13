package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"

	"goravel/app/facades"
	"goravel/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl")
	})

	facades.Route().Prefix("user").Group(func(router route.Router) {
		// userController = new UserController() -> PHP
		userController := &controllers.UserController{}

		// Route::get("/user", [UserController::Class, "index"]) -> laravel
		router.Get("/", userController.Index)
		router.Get("/:id", userController.Show)
	})
}
