package application

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	ErrAppRunning    = errors.New(`application is already running`)
	ErrAppNotRunning = errors.New(`application is not running `)
)

const (
	DefaultShutdownTimeout = 15 * time.Second
)

type Application struct {

	// name is the name of application b default this value contains file name of binary
	// file of launched application.
	name string

	// state is the state of Application example start, shutdown, running.
	state uint64

	// shutdownTimeout is the timeout for graceful shutting down server
	shutdownTimeout time.Duration

	// main application context
	ctx    context.Context
	cancel context.CancelFunc

	// database connection instance for application
	DB *gorm.DB
}

// Init is a function for creating and initializing all Application modules.
func (app *Application) Init() (err error) {

	app.DbConnect()
	app.AutoMigrate()

	log.Println(`application is initialised`)

	return nil
}

// Run is function for launching Application.
func (app *Application) Run() (err error) {

	app.StartServer()

	log.Println(`application is running`)

	return nil
}

// New is create new Application with Option's.
func New(opts ...Option) *Application {

	app := &Application{
		name:            filepath.Base(os.Args[0]),
		shutdownTimeout: DefaultShutdownTimeout,
	}
	for _, opt := range opts {
		opt.apply(app)
	}

	return app
}
