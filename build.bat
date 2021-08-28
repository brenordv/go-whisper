echo off

echo Setting Architecture: amd64
set GOARCH=amd64

echo Setting OS: windows
set GOOS=windows
echo Building: WHISPER for Windows
go build -o .\.dist\whisper.exe ./cmd/whisper
7z a -tzip .\.dist\go-whisper--windows-amd64--%*.zip .\.dist\whisper.exe readme.md


echo Setting OS: linux
set GOOS=linux

echo Building: WHISPER for Linux
go build -o .\.dist\whisper ./cmd/whisper
7z a -tzip .\.dist\go-whisper--linux-amd64--%*.zip .\.dist\whisper readme.md

copy .\.dist\whisper.exe .\whisper.exe