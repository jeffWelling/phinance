package internal

import (
	"log/slog"
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	// Version is the current version of the application
	Version          = "0.0.1"
	ProgramName      = "phinances"
	DatabaseFilename = "phinances.db"
	DBMigrationsPath = "db/migrations"
)

func Start() int {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Starting up " + ProgramName + " version " + Version)

	// router := gin.Default()

	// // Load the templates
	// router.LoadHTMLGlob("templates/*")

	// // The first WebUI will exist with no pefix, just "/"
	// router.GET("/", func(c *gin.Context) { homepage(c, gormDB) })
	// router.GET("/ui/v1/", func(c *gin.Context) { homepage(c, gormDB) })
	// router.GET("/ui/v1/accounts", accountsList)

	// // The API for react will exist under "/api/v1/"

	// router.Run(":8080")

	return 0
}

func homepage(c *gin.Context, db *gorm.DB) {
	c.HTML(200, "home.html", nil)
}

func accountsList(c *gin.Context) {
	c.HTML(200, "accounts.html", nil)
}
