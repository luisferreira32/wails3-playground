//go:build darwin

package main

import "github.com/wailsapp/wails/v3/pkg/application"

func getBoderWindowOptions() application.WebviewWindowOptions {
	return application.WebviewWindowOptions{
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
	}
}
