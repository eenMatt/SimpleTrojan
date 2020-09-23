
**Assignment 2c**
- Negar Farshchi
- Matthew Waiariki


**Building Binary**
```
rsrc -manifest PingPlotterTrial.exe.manifest -ico setup.ico -o PingPlotterTrial.syso
go build -ldflags -H=windowsgui -o PingPlotterTrial.exe
```