package main

import (
	"context"
	"errors"
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

	// create a new scraper, and load vendors
	scraper := wfmplatefficiency.NewScraper()
	err := scraper.LoadVendors()
	if err != nil {
		runtime.LogError(a.ctx, fmt.Sprintf("error: %v", err))
	}

	runtime.LogInfo(a.ctx, "Loaded vendors\n")

	// save scraper pointer to context
	a.ctx = context.WithValue(a.ctx, "scraper", scraper)
}

func (a *App) getScraperFromCtx() (*wfmplatefficiency.Scraper, error) {
	scraper, ok := a.ctx.Value("scraper").(*wfmplatefficiency.Scraper)
	if !ok {
		err := errors.New("could not get scraper from context")
		runtime.LogErrorf(a.ctx, "%v", err)
		return nil, err
	}

	return scraper, nil
}

func (a *App) FetchVendor(name string) string {
	scraper, err := a.getScraperFromCtx()
	if err != nil {
		return "error"
	}

	vendor := scraper.GetVendors()[name]

	runtime.LogInfof(a.ctx, "Fetching items for: %s\n", vendor.Name)
	err = vendor.GetVendorStats()
	// TODO: get live updates via channel
	if err != nil {
		runtime.LogErrorf(a.ctx, "failed to get vendor statistics %v", err)
	}

	return vendor.String()
}

func (a *App) GetVendor(name string) *wfmplatefficiency.Vendor {
	scraper, err := a.getScraperFromCtx()
	if err != nil {
		return nil
	}

	return scraper.GetVendors()[name]
}

func (a *App) GetVendorNames() []string {
	names := make([]string, 0)

	// TODO: turn this into a helper function
	scraper, err := a.getScraperFromCtx()
	if err != nil {
		return names
	}

	for _, v := range scraper.GetVendors() {
		names = append(names, v.Name)
	}

	return names
}
