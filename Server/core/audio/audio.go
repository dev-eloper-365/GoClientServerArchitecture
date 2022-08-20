package audio

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Play(connection net.Conn) (err error) {
	fmt.Print("\baudio_name> ")
	CommandReader := bufio.NewReader(os.Stdin)
	user_command_raw, err := CommandReader.ReadString('\n')
	command := strings.TrimSpace(user_command_raw)
	if command == "stop" {
		nbytes, error := connection.Write([]byte(user_command_raw))
		if error != nil {
			panic(error)
		}
		fmt.Println("[+]", nbytes, "bytes written\n")
		return
	}

	nbytes, err := connection.Write([]byte(user_command_raw))
	if err != nil {
		panic(err)
	}
	fmt.Println("[+]", nbytes, "bytes written\n")

	reader := bufio.NewReader(connection)
	raw_audio_status, err := reader.ReadString('\n')
	audio_status := strings.TrimSpace(raw_audio_status)

	if audio_status == "incorrect" {
		fmt.Println("[-] File doesn't exist.")
		return
	} else {
		fmt.Println("[#] Press Enter to stop playing...")
		response_reader := bufio.NewReader(os.Stdin)
		response, error := response_reader.ReadString('\n')
		if error != nil {
			panic(error)
		}
		nes, err := connection.Write([]byte(response))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", nes, "bytes written\n")

	}
	return
}
