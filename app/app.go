package app

import "context"

type App struct {
	version string
	ctx     context.Context
}

func NewApp(v string) *App {
	// force inappropriate key
	ctx := context.WithValue(context.Background(), "version", v)
	return &App{
		version: v,
		ctx:     ctx,
	}
}

func (a *App) Version() string {
	return a.version
}
