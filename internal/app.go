package internal

import (
	"cmp"
	"embed"
	"net/http"
	"os"

	"infinite-scroll/internal/employees"
	"infinite-scroll/internal/system/assets"

	"go.leapkit.dev/core/db"
	"go.leapkit.dev/core/render"
	"go.leapkit.dev/core/server"
)

var (
	//go:embed **/*.html **/*.html
	tmpls embed.FS

	// DB is the database connection builder function
	// that will be used by the application based on the driver and
	// connection string.
	DB = db.ConnectionFn(
		cmp.Or(os.Getenv("DATABASE_URL"), "database.db?_timeout=5000&_sync=1"),
		db.WithDriver("sqlite3"),
	)
)

// Server interface exposes the methods
// needed to start the server in the cmd/app package
type Server interface {
	Addr() string
	Handler() http.Handler
}

func New() Server {
	// Creating a new server instance with the
	// default host and port values.
	r := server.New(
		server.WithHost(cmp.Or(os.Getenv("HOST"), "0.0.0.0")),
		server.WithPort(cmp.Or(os.Getenv("PORT"), "3000")),
		server.WithSession(
			cmp.Or(os.Getenv("SESSION_SECRET"), "d720c059-9664-4980-8169-1158e167ae57"),
			cmp.Or(os.Getenv("SESSION_NAME"), "leapkit_session"),
		),

		server.WithAssets(assets.Files, "/internal/system/assets"),
	)

	r.Use(render.Middleware(
		render.TemplateFS(tmpls, "internal"),
		render.WithDefaultLayout("system/layout.html"),
	))

	r.HandleFunc("GET /{$}", employees.Index)

	return r
}
