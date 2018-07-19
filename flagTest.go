package main

import "flag"

var (
	roomConfigID string //房间ID
)

func main() {
	flag.StringVar(&roomConfigID, "room_id", "", "游戏房间ID")
	
}
