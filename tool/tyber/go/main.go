package main

import (
	"fmt"
	"log"

	"berty.tech/berty/tool/tyber/go/v2/bind"
	"berty.tech/berty/tool/tyber/go/v2/bridge"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	bootstrap "github.com/asticode/go-astilectron-bootstrap"
)

// Window properties
const (
	center    = true
	height    = 700
	minHeight = 350
	width     = 1200
	minWidth  = 500
)

// App properties
const singleInstance = true

// Vars injected via ldflags by bundler
var (
	AppName            string
	BuiltAt            string
	VersionAstilectron string
	VersionElectron    string
)

func main() {
	// Create logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	// Init Go <-> JS bridge
	b := bridge.New(l)

	// Run bootstrap
	l.Printf("Running app built at %s\n", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    bind.Asset,
		AssetDir: bind.AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "bundler/resources/icons/icon.icns",
			AppIconDefaultPath: "bundler/resources/icons/icon.png",
			SingleInstance:     singleInstance,
			VersionAstilectron: VersionAstilectron,
			VersionElectron:    VersionElectron,
		},
		Logger: l,
		MenuOptions: []*astilectron.MenuItemOptions{{
			Label: astikit.StrPtr("File"),
			SubMenu: []*astilectron.MenuItemOptions{
				{
					Label:       astikit.StrPtr("Open File(s)"),
					OnClick:     b.OpenFiles,
					Accelerator: astilectron.NewAccelerator("CommandOrControl", "o"),
				},
				{Type: astilectron.MenuItemTypeSeparator},
				{
					Label:       astikit.StrPtr("Preferences..."),
					OnClick:     b.OpenPreferences,
					Accelerator: astilectron.NewAccelerator("CommandOrControl", ","),
				},
				{
					Label:       astikit.StrPtr("Developer Tools"),
					OnClick:     b.ToggleDevTools,
					Accelerator: astilectron.NewAccelerator("CommandOrControl", "Alt", "i"),
				},
				{Type: astilectron.MenuItemTypeSeparator},
				{
					Label: astikit.StrPtr(fmt.Sprintf("Quit %s", AppName)),
					Role:  astilectron.MenuItemRoleQuit,
				},
			},
		}},
		OnWait:        b.Init,
		RestoreAssets: bind.RestoreAssets,
		ResourcesPath: "bundler/resources",
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: b.HandleMessages,
			Options: &astilectron.WindowOptions{
				Center:    astikit.BoolPtr(center),
				MinHeight: astikit.IntPtr(minHeight),
				Height:    astikit.IntPtr(height),
				MinWidth:  astikit.IntPtr(minWidth),
				Width:     astikit.IntPtr(width),
			},
		}},
	}); err != nil {
		l.Fatal(fmt.Errorf("Running bootstrap failed: %w", err))
	}
}
