package device

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"net"
	"strconv"
)

type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	RAM      uint64 `bson:ram`
	Disk     uint64 `bson:disk`
}

func Info(connection net.Conn) (err error) {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("\\") // If you're in Unix change this "\\" for "/"

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.Disk = diskStat.Total / 1024 / 1024
	v, _ := mem.VirtualMemory()

	string := "\nHostname \t: " + info.Hostname + "\nPlatform \t: " + info.Platform + "\nCPU\t \t: " + info.CPU + "\nDisk Space \t: " + strconv.FormatUint(info.Disk, 10) + " MB" + "\nRAM       \t: " + strconv.FormatUint(info.RAM, 10) + " MB" + "\nUsed RAM       \t: " + fmt.Sprint(v.UsedPercent) + "%"
	nbytes, err := connection.Write([]byte(string))
	fmt.Println("[+]", nbytes, "bytes written")
	return
}
