package service

import (
	"os"
	"os/exec"
	"plan4u/config"
	"runtime"

	"fyne.io/systray"
	"fyne.io/systray/example/icon"
	"github.com/labstack/gommon/log"
)

type SystemTray struct {
	Cfg *config.Config
}

func (s *SystemTray) OnReady() {
	systray.SetIcon(icon.Data)
	app_title := systray.AddMenuItem("plan4u app", "plan4u")
	app_title.Disable()
	launch_button := systray.AddMenuItem("Open", "Open in browser")
	quit_button := systray.AddMenuItem("Quit", "Quit the app")

	go openBrowser(launch_button, "http://"+s.Cfg.URL+":"+s.Cfg.PORT)

	go func() {
		<-quit_button.ClickedCh
		systray.Quit()
	}()
}

func (s *SystemTray) OnExit() {
	os.Exit(0)
}

func openBrowser(event *systray.MenuItem, url string) {
	for {
		<-event.ClickedCh
		var err error
		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			log.Errorf("unsupported platform")
		}
		if err != nil {
			panic(err)
		}
	}
}
