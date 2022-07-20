package appdata

import (
	"net"
	"os"
	"path/filepath"
)

func Dir(conn net.Conn) (err error) {
	home, _ := os.UserHomeDir()
	com := os.Chdir(filepath.Join(home, "AppData\\Roaming"))
	if com != nil {
		panic(err)
	}
	return
}
