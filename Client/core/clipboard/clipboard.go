package clipboard

import (
	"bufio"
	"fmt"
	"github.com/d-tsuji/clipboard"
	"log"
	"net"
	"strings"
)

func Write(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	clip_command, err := reader.ReadString('\n')
	command := strings.TrimSpace(clip_command)
	if command == "stop" {
		return
	}
	if command == "read" {
		text, err := clipboard.Get()
		if err != nil {
			log.Fatal(err)
		}
		nbytes, err := connection.Write([]byte(text + "\n"))
		if err != nil {
			panic(err)
		}
		fmt.Println("[+]", nbytes, "bytes written\n")
	}
	if command == "write" {
		fmt.Println("entered")
		reader := bufio.NewReader(connection)
		clip_write, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		write := strings.TrimSpace(clip_write)
		if err := clipboard.Set(write); err != nil {
			log.Fatal(err)
		}
	}
	return
}
