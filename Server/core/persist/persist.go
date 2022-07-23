package persist

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Schedule(connection net.Conn) (err error) {

	fmt.Print("\bSchedule_reboot> ")
	CommandReader := bufio.NewReader(os.Stdin)
	user_command_raw, err := CommandReader.ReadString('\n')
	user_command := strings.TrimSpace(user_command_raw)

	command := user_command[:8]

	if command == "schedule" {
		nbytes, err2 := connection.Write([]byte(user_command[9:] + "\n"))
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("[+]", nbytes, "bytes written")

		CommandReader := bufio.NewReader(connection)
		msg_raw, err1 := CommandReader.ReadString('\n')
		if err1 != nil {
			panic(err1)
		}
		msg := strings.TrimSpace(msg_raw)
		fmt.Println(msg)
	} else {
		fmt.Println("[-] Invalid Selection")
		return
	}
	return
}

func Remove(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	raw_response, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	response := strings.TrimSpace(raw_response)
	if response == "" {
		fmt.Println("[#] Persistance Disarmed")
	} else {
		fmt.Println("[#] " + response)
	}
	return
}
