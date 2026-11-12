package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syscall"
	"time"
)

// clearAppData deletes library cache, settings, and artwork, preserving the installation UUID
func (app *MiyooPod) clearAppData() {
	dc := app.DC

	dc.SetHexColor(app.CurrentTheme.BG)
	dc.Clear()

	dc.SetHexColor(app.CurrentTheme.HeaderBG)
	dc.DrawRectangle(0, 0, SCREEN_WIDTH, HEADER_HEIGHT)
	dc.Fill()

	dc.SetFontFace(app.FontHeader)
	dc.SetHexColor(app.CurrentTheme.HeaderTxt)
	dc.DrawStringAnchored("Clear App Data", SCREEN_WIDTH/2, HEADER_HEIGHT/2, 0.5, 0.5)

	dc.SetFontFace(app.FontMenu)
	dc.SetHexColor(app.CurrentTheme.ItemTxt)
	dc.DrawStringAnchored("Clear all app data?", SCREEN_WIDTH/2, SCREEN_HEIGHT/2-20, 0.5, 0.5)

	dc.SetFontFace(app.FontSmall)
	dc.SetHexColor(app.CurrentTheme.Dim)
	dc.DrawStringAnchored("This will reset settings and rescan your library", SCREEN_WIDTH/2, SCREEN_HEIGHT/2+16, 0.5, 0.5)

	barY := float64(SCREEN_HEIGHT - STATUS_BAR_HEIGHT)

	dc.SetHexColor(app.CurrentTheme.HeaderBG)
	dc.DrawRectangle(0, barY, SCREEN_WIDTH, STATUS_BAR_HEIGHT)
	dc.Fill()

	dc.SetHexColor(app.CurrentTheme.ProgBG)
	dc.SetLineWidth(1)
	dc.DrawLine(0, barY, SCREEN_WIDTH, barY)
	dc.Stroke()

	centerY := barY + float64(STATUS_BAR_HEIGHT)/2
	dc.SetFontFace(app.FontSmall)
	app.drawButtonLegend(12, centerY, "A", "Confirm")
	app.drawButtonLegend(150, centerY, "B", "Cancel")

	app.triggerRefresh()

	for app.Running {
		key := Key(C_GetKeyPress())
		if key == NONE {
			time.Sleep(33 * time.Millisecond)
			continue
		}

		switch key {
		case A:
			app.performClearAppData()
			return
		case B, MENU:
			app.LastKey = NONE
			app.drawCurrentScreen()
			return
		}
	}
}

// performClearAppData does the actual deletion and restarts the app
func (app *MiyooPod) performClearAppData() {
	installID := app.InstallationID

	dc := app.DC
	dc.SetHexColor(app.CurrentTheme.BG)
	dc.Clear()
	dc.SetFontFace(app.FontMenu)
	dc.SetHexColor(app.CurrentTheme.HeaderTxt)
	dc.DrawStringAnchored("Clearing app data...", SCREEN_WIDTH/2, SCREEN_HEIGHT/2, 0.5, 0.5)
	app.triggerRefresh()

	if err := os.Remove(LIBRARY_JSON_PATH); err != nil && !os.IsNotExist(err) {
		logMsg(fmt.Sprintf("WARNING: Failed to remove library cache: %v", err))
	}

	if err := os.Remove(SETTINGS_PATH); err != nil && !os.IsNotExist(err) {
		logMsg(fmt.Sprintf("WARNING: Failed to remove settings: %v", err))
	}

	if err := os.RemoveAll(ARTWORK_DIR); err != nil && !os.IsNotExist(err) {
		logMsg(fmt.Sprintf("WARNING: Failed to remove artwork: %v", err))
	}

	minimalSettings := Settings{
		InstallationID: installID,
	}
	data, err := json.MarshalIndent(minimalSettings, "", "  ")
	if err == nil {
		os.WriteFile(SETTINGS_PATH, data, 0644)
	}

	logMsg("INFO: App data cleared, restarting...")

	if app.Playing != nil && app.Playing.State == StatePlaying {
		audioStop()
	}

	close(app.RefreshChan)
	sdlCleanup()

	launchPath := "./launch.sh"
	syscall.Exec("/bin/sh", []string{"/bin/sh", launchPath}, os.Environ())
}
