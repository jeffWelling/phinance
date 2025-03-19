package internal

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	sqlite3Migrate "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

var (
	// Version is the current version of the application
	Version          = "0.0.1"
	ProgramName      = "phinances"
	DatabaseFilename = "phinances.db"
	DBMigrationsPath = "db/migrations"
)

type Accounty struct {
	accountID   int
	accountName string
	accountType string
	createdAt   string
	currency    string
}

type Transaction struct {
	fromAccount     int
	toAccount       int
	amount          int
	transactionDate string
}

type AccountRelationship struct {
	parentAccountID int
	childAccountID  int
}

type PhinanceContext struct {
	logger *slog.Logger
	db     *gorm.DB
}

func Start() int {
	var context PhinanceContext
	context.logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	context.logger.Info("Starting up " + ProgramName + " version " + Version)

	// Open or create 'app.db' as your sqlite file
	db, err := sql.Open("sqlite3", "phinances.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Create a migrate driver using the existing DB connection
	driver, err := sqlite3Migrate.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("Failed to create migrate driver: %v", err)
	}

	// Point to our migrations folder; the second argument must match the DB name
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"ql",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	// Apply all up migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to run migration up: %v", err)
	}

	// If there's nothing to migrate, ErrNoChange is returned
	fmt.Println("Migrations applied successfully (or no changes required).")

	router := gin.Default()

	// Load the templates
	router.LoadHTMLGlob("templates/*")

	// The first WebUI will exist with no pefix, just "/"
	router.GET("/", func(c *gin.Context) { homepage(c, context.db) })
	router.GET("/ui/v1/", func(c *gin.Context) { homepage(c, context.db) })
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
