//go:build darwin

package main

/*
#cgo darwin LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

void SetWindowIgnoresMouseEvents(void *window) {
    [(NSWindow *)window setIgnoresMouseEvents:YES];
}
*/
import "C"

import (
	"unsafe"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func createBorderWindow(app *application.App) *application.WebviewWindow {
	borderWindow := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "border-window",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTransparent,
			TitleBar:                application.MacTitleBarHidden,
			DisableShadow:           true,
		},
		BackgroundType:          application.BackgroundTypeTransparent,
		BackgroundColour:        application.NewRGBA(0, 0, 0, 0),
		Frameless:               true,
		StartState:              application.WindowStateFullscreen,
		AlwaysOnTop:             true,
		DisableResize:           true,
		URL:                     "/borderpage.html",
		Centered:                true,
		FullscreenButtonEnabled: true,
	})

	handle, err := borderWindow.NativeWindowHandle()
	if err != nil {
		panic(err.Error())
	}
	C.SetWindowIgnoresMouseEvents(unsafe.Pointer(handle))

	return borderWindow
}
