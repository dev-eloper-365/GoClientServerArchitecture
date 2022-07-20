package wall

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/reujab/wallpaper"
)

func Url(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	url_raw, err := reader.ReadString('\n')
	url := strings.TrimSpace(url_raw)
	if url == "stop" {
		fmt.Println("[+] Cancelled wallpaper change")
		return
	}
	background, err := wallpaper.Get()

	fmt.Println("\nCurrent wallpaper : \n", background)

	err = wallpaper.SetFromURL(url)
	if err != nil {
		panic(err)
	}
	return
}
