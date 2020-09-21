rsrc -manifest WinDirStat-Install.exe.manifest -ico setup.ico -o WinDirStat-Install.syso
go build -ldflags -H=windowsgui -o WinDirStat-Install.exe