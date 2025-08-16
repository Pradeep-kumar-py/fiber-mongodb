package routes

import (
	"github.com/Pradeep-kumar-py/fiber-mongodb/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	// Create user group
	user := app.Group("/api/user")

	// Define routes (similar to Express router.route())
	user.Post("/register", controllers.RegisterUser)
	user.Post("/login", controllers.LoginUser)
	user.Post("/logout", controllers.LogoutUser)
	user.Get("/getalluser", controllers.GetAllUsers)

	//     app.Route("/api/users", func(api fiber.Router) {
	//     api.Post("/register", controllers.RegisterUser)
	//     api.Post("/login", controllers.LoginUser)
	//     api.Post("/logout", controllers.LogoutUser)
	//     api.Get("/", controllers.GetAllUsers)
	//     api.Get("/:id", controllers.GetUserById)
	//     api.Put("/:id", controllers.UpdateUser)
	//     api.Delete("/:id", controllers.DeleteUser)
	// })
}
