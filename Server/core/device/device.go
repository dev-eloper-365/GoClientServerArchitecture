package device

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func Info(connection net.Conn) (err error) {

	CommandReader := bufio.NewReader(connection)
	deviceinfo_raw, err := CommandReader.ReadString('%')
	device_info := strings.TrimSpace(deviceinfo_raw)

	fmt.Println(device_info)
	return
}
