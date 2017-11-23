@echo off
git clone https://github.com/chatton/GoAskEliza.git
cd GoAskEliza/src
go build main.go
START main.exe
start "" http://localhost:8080
exit
