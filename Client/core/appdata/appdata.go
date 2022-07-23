package appdata

import (
	"net"
	"os"
	"path/filepath"
)

func Dir(conn net.Conn) (err error) {
	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]

	home, _ := os.UserHomeDir()
	com := os.Chdir(filepath.Join(home, "AppData\\Roaming", without_ext))
	if com != nil {
		panic(err)
	}
	return
}
