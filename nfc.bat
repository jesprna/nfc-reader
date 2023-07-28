@echo off
setlocal

:start
nfcuid.exe -device=1
if %ERRORLEVEL% NEQ 0 goto error
goto start

:error
echo Application encountered an error or exited unexpectedly.
echo Restarting...
timeout /t 2 > nul
goto start