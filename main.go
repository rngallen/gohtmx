package main

import (
	"embed"
	"html/template"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/template/html/v2"
	"github.com/rngallen/gohtmx/services/logs"
)

//go:embed views/*
var embedDirViews embed.FS

//go:embed static/*
var embedDirStatic embed.FS

func main() {

	// Activate logger
	logs.AppLogger()

	// Embed html files to compiled file
	engine := html.NewFileSystem(http.FS(embedDirViews), ".html")

	app := fiber.New(fiber.Config{
		ServerHeader:  "GOFIBER",
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       "Go Htmx",
		Views:         engine,
		ViewsLayout:   "views/base/main",
	})

	// Embed static files to compiled file
	app.Use("/static", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirStatic),
		PathPrefix: "static",
		// Seconds 3600 => 60*60 1Hour
		MaxAge: 3600,
	}))

	app.Get("", home)
	app.Post("/addFilm", addFilm)

	logs.ErrorLogger.Fatalln(app.Listen(":80"))
}

type Film struct {
	Title    string `json:"title" form:"title"`
	Director string `json:"director" form:"director"`
}

// will act as database
var movies = []Film{
	{Title: "The Godfather", Director: "Francis Ford Coppola"},
	{Title: "Blade Runner", Director: "Ridley Scott"},
	{Title: "The Thing", Director: "John Carpenter"},
}

func home(c *fiber.Ctx) error {
	return c.Render("views/index", fiber.Map{"Movies": movies})
}

func addFilm(c *fiber.Ctx) error {

	var input Film
	if err := c.BodyParser(&input); err != nil {
		return c.Render("views/index", fiber.Map{"Error": err.Error()})
	}

	// Simulate time taken to interact with database
	time.Sleep(time.Millisecond / 5)
	movies = append(movies, Film{input.Title, input.Director})
	tmpl := template.Must(template.ParseFiles("views/index.html"))
	tmpl.ExecuteTemplate(c.Response().BodyWriter(), "film-list-element", Film{input.Title, input.Director})

	return nil
}
