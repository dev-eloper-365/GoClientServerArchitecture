package alert

import (
	"bufio"
	"net"
	"strings"

	"tawesoft.co.uk/go/dialog"
)

func Dialog(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	alert_raw, err := reader.ReadString('\n')
	alert_msg := strings.TrimSpace(alert_raw)
	if alert_msg == "stop" {
		return
	}
	dialog.Alert(alert_msg)
	return
}
