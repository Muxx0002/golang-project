package routes

import (
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/transport/handlers"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/transport/midleware"
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
	app.Use(midleware.AuthMiddleware)

	api := app.Group("/auth")
	api.Post("/sign-up", handlers.Registration)
	api.Post("/sign-in", handlers.Login)
	api.Post("/log-out", handlers.LogOut)

	articles := app.Group("/articles")
	articles.Get("/", handlers.ListArticlesHandler)
	articles.Get("/:id", handlers.ArticleByIDHandler)
	articles.Post("/:id/comments", handlers.CreateComment)

	admin := app.Group("/admin")
	admin.Use(midleware.IsAdmin)
	admin.Post("/articles", handlers.CreateArticle)
	admin.Put("/articles/:id", handlers.UpdateArticle)
	admin.Delete("/articles/:id", handlers.DeleteArticle)

	app.Listen(viper.GetString("server_port"))
}
