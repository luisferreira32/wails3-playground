//go:build windows

package main

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/w32"
)

func getBoderWindowOptions() application.WebviewWindowOptions {
	return application.WebviewWindowOptions{
		Title:          "border-window",
		BackgroundType: application.BackgroundTypeTransparent,
		Frameless:      true,
		StartState:     application.WindowStateFullscreen,
		AlwaysOnTop:    true,
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
			ExStyle:                           w32.WS_EX_TOPMOST | w32.WS_EX_LAYERED | w32.WS_EX_TRANSPARENT,
		},
		URL: "/borderpage.html",
	}
}
