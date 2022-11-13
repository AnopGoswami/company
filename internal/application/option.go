package application

import (
	"time"
)

// An Option configures Application.
type Option interface {
	apply(*Application)
}

func Name(name string) Option {
	return optionFunc(func(app *Application) {
		if name != `` {
			app.name = name
		}
	})
}

func ShutdownTimeout(d time.Duration) Option {
	return optionFunc(func(app *Application) {
		if d > 0 {
			app.shutdownTimeout = d
		}
	})
}

// optionFunc wraps a function, so it satisfies the Option interface.
type optionFunc func(*Application)

func (f optionFunc) apply(app *Application) { f(app) }
