package alert

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Dialog(connection net.Conn) (err error) {

	fmt.Print("\bAlert_msg> ")
	CommandReader := bufio.NewReader(os.Stdin)
	user_command_raw, err := CommandReader.ReadString('\n')
	if err != nil {
		fmt.Println("[+] Unable to read command ")
	}
	if user_command_raw == "stop" {
		return
	}
	nbytes, err := connection.Write([]byte(user_command_raw))
	fmt.Println("\n[+]", nbytes, "bytes written\n")
	return
}
