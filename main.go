package main

import (
	"context"
	"embed"
	"log"
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed frontend/dist
var assets embed.FS

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "wails3-demo",
		Description: "A demo of using raw HTML & CSS",
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	createBorderWindow(app)

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "clock-window",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTransparent,
			TitleBar:                application.MacTitleBarHidden,
			DisableShadow:           true,
		},
		BackgroundType: application.BackgroundTypeTransparent,
		Frameless:      true,
		AlwaysOnTop:    true,
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
		},
		URL:    "/",
		Width:  300,
		Height: 100,
	})

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	ctx, cancel := context.WithCancel(context.Background())
	app.OnShutdown(func() {
		cancel()
	})
	go func() {
		t := time.NewTicker(time.Second)
		for {
			now := time.Now().Format(time.TimeOnly)
			app.Events.Emit(&application.WailsEvent{
				Name: "time",
				Data: now,
			})
			select {
			case <-ctx.Done():
				return
			case <-t.C:
			}
		}
	}()

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
