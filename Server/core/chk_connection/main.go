package chk_connection

import (
	"fmt"
	"net"
)

func Check(connection net.Conn) (err error) {
	nbytes, err := connection.Write([]byte("try"))
	if err == nil {
		fmt.Println("[+] Alive ", nbytes)
	}
	return
}
