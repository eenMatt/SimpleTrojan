**Assignment 2c for ICT Information Systems Security @ RMIT University**
S3636862 -- Negar Farshchi
S3751243 -- Matthew Waiariki

**Target is for Windows 10 machine.**
Ability to handle if user is using a OneDrive Desktop or Normal Desktop.
Set custom SFTP server, use flag " -sftp "sftp server address"

**Building Binary**
1. Download required packages
```
go get  "github.com/elastic/go-sysinfo"
go get  "github.com/pkg/sftp"
go get  "golang.org/x/crypto/ssh"

```

2. build
```
go build -ldflags -H=windowsgui -o "\bin\PingPlotterTrial.exe" 

```
3. New Binary will be found in bin\