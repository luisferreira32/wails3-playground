//go:build darwin

package main

import "github.com/wailsapp/wails/v3/pkg/application"

func getBoderWindowOptions() application.WebviewWindowOptions {
	return application.WebviewWindowOptions{
		Title: "border-window",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundType: application.BackgroundTypeTransparent,
		Frameless:      true,
		StartState:     application.WindowStateFullscreen,
		AlwaysOnTop:    true,
		URL:            "/borderpage.html",
	}
}
