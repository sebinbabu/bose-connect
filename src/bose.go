package main

import (
    "fmt"
    "bufio"
    "strings"
    "os"
    b64 "encoding/base64"
)

func main() {
    dir := os.Getenv("HOME") + "/.bose"
    config := dir + "/connect_config"

    args := os.Args[1:]
    if(len(args) != 1) {
        fmt.Println("bose : incorrect arguments supplied")
    } else if(args[0] == "connect") {
        if(DoesExist(config) == false) {
            BoseLogin(dir, config)
        }
        if(BoseConnect(dir, config) == false) {
            fmt.Println("bose : connect failed")
        }
    } else if(args[0] == "disconnect") {
        BoseDisconnect()
    } else if(args[0] == "login") {
        BoseLogin(dir, config)
    } else {
        BoseHelp()
    }
}
    
func BoseConnect(dir, config string) bool {
    token := ReadFromFile(config)

    url := "http://10.7.0.1:8888/"
    request := "actionType=umlogin&authorization=" + token + "&language=0&userIpMac="
    ClientPost(url, request)
    return ClientCheckConnection()
}

func BoseDisconnect() {
    url := "http://10.7.0.1:8888/"
    request := "actionType=umlogout&language=0"
    ClientPost(url, request)
}

func BoseLogin(dir, config string) {
    user, pass := GetUserInput()
    text := user + "|" + pass;
    text = b64.StdEncoding.EncodeToString([]byte(text))

    CreateDirIfNotExist(dir)
    WriteToFile(config, text)
}

func BoseHelp() {
    fmt.Println("\nUsage:\n bose [option]\nA tool to connect to BOSE.\n\nOptions:\n login\t\tstore user credentials for login\n connect\tsignin to BOSE\n disconnect\tsignout from BOSE\n help\t\tdisplay this help")
}

func GetUserInput() (string, string) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("username: ")
    user, _ := reader.ReadString('\n')
    user = strings.Replace(user, "\n", "", -1)

    fmt.Print("password: ")
    pass, _ := reader.ReadString('\n')
    pass = strings.Replace(pass, "\n", "", -1)

    return user, pass
}