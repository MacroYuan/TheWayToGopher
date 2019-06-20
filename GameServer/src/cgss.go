package main

import "GameServer/src/cg"

var centerClient *cg.CenterClient

func startCenterService() error {

}

func Help(args []string) int {

}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {

}

func Login(args []string) int {
	return 0
}

func ListPlayer(args []string) int {

}

func Send(args []string) int {

}

//将命令与处理函数对齐
func GetCommandHandlers() map[string]func(args []string) int {

}

func main() {

}
