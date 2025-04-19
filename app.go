package main

import (
	"context"
	"fmt"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	scraper := wfmplatefficiency.NewScraper()
	err := scraper.LoadVendors()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("error: %v", err))
	}

	a.ctx = context.WithValue(a.ctx, "scraper", scraper)
}

func (a *App) GetVendorNames() []string {
	names := make([]string, 0)

	// TODO: turn this into a helper function
	scraper, ok := a.ctx.Value("scraper").(*wfmplatefficiency.Scraper)
	if !ok {
		return names
	}

	for _, v := range scraper.GetVendors() {
		names = append(names, v.Name)
	}

	return names
}
