package routes

import (
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/transport/handlers"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/transport/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func Routes() {
	front_path := viper.GetString("front_path")
	app := fiber.New(fiber.Config{
		ReadTimeout:  viper.GetDuration("read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration("write_timeout") * time.Second,
		IdleTimeout:  viper.GetDuration("idle_timeout") * time.Second,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.SendFile(front_path + "/index.html")
		},
	})
	app.Static("/assets/", front_path+"/assets/")
	app.Use(cors.New())
	app.Use(recover.New())
	api := app.Group("/auth")
	api.Use(middleware.AuthMiddleware)
	api.Post("/sign-up", handlers.Registration)
	api.Post("/sign-in", handlers.Login)
	api.Post("/log-out", handlers.LogOut)
	api.Get("/me", handlers.UserProfile)
	api.Put("/me", handlers.UpdateUserData)
	api.Post("/forgot-password") //запрос на сброс пароля
	app.Post("/reset-password")  //подтверждение сброса, установка нового пароля

	articles := app.Group("/articles")
	articles.Get("/", handlers.ListArticlesHandler)
	articles.Get("/:id", handlers.ArticleByIDHandler)
	articles.Post("/:id/comments", middleware.AuthMiddleware, handlers.CreateComment)
	articles.Delete("/comments/:id", handlers.DeleteCommentByIDHandler)

	admin := app.Group("/admin")
	admin.Use(middleware.IsAdmin)
	admin.Post("/articles", handlers.CreateArticle)
	admin.Get("/comments", handlers.ListCommentsHandler)
	admin.Get("/users", handlers.ListUsersHandler)
	admin.Put("/articles/:id", handlers.UpdateArticle)
	admin.Delete("/articles/:id", handlers.DeleteArticle)
	admin.Delete("/comments/:id", handlers.DeleteCommentByIDHandler)
	admin.Delete("/users/:id", handlers.DeleteUserHandler)

	app.Listen(viper.GetString("server_port"))
}
