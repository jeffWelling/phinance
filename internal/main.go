package internal

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	gormSqlite3 "gorm.io/driver/sqlite"
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

	databasePassword := os.Getenv("PHINANCES_DATABASE_PASSWORD")
	if databasePassword == "" {
		logger.Error("PHINANCES_DATABASE_PASSWORD not set")
		return 1
	}
	dbConnectionString := fmt.Sprintf("sqlite3://%s?_pragma=cipher_version(4)&key=%s", DatabaseFilename, databasePassword)

	db, err := sql.Open("sqlite3", DatabaseFilename)

	if err != nil {
		logger.Error("Failed to open database: " + err.Error())
		return 1
	}

	migrateDriver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		logger.Error("Failed to open database for migration: " + err.Error())
		return 1
	}

	// Apply database migrations before starting
	migrateInstance, err := migrate.NewWithDatabaseInstance(
		"file://"+DBMigrationsPath,
		"ql", migrateDriver,
	)
	if err != nil {
		logger.Error("Failed to open database for migration: " + err.Error())
		return 1
	}
	if err := migrateInstance.Up(); err != nil {
		logger.Error("Failed to apply database migrations: " + err.Error())
		return 1
	}
	logger.Info("Applied database migrations")
	db.Close()

	logger.Info("Opening database for gorm: " + dbConnectionString)
	gormDB, err := gorm.Open(gormSqlite3.Open(DatabaseFilename), &gorm.Config{})
	if err != nil {
		// Only throw error if the migrations failed due to 'no change'
		if !strings.Contains(err.Error(), "no change") {
			logger.Error("Failed to open database for gorm: " + err.Error())
			return 1
		}
	}

	router := gin.Default()

	// Load the templates
	router.LoadHTMLGlob("templates/*")

	// The first WebUI will exist with no pefix, just "/"
	router.GET("/", func(c *gin.Context) { homepage(c, gormDB) })
	router.GET("/ui/v1/", func(c *gin.Context) { homepage(c, gormDB) })
	router.GET("/ui/v1/accounts", accountsList)

	// The API for react will exist under "/api/v1/"

	router.Run(":8080")

	return 0
}

func homepage(c *gin.Context, db *gorm.DB) {
	c.HTML(200, "home.html", nil)
}

func accountsList(c *gin.Context) {
	c.HTML(200, "accounts.html", nil)
}
