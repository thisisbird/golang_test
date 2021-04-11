package main

import (
	"fmt"
	"github.com/headfirstgo/gadget"
)

type Player interface{
	Play(string)
	Stop()
}

func TryOut(player Player){
	player.Play("Test Track")
	player.Stop()
	// player.Record() 無法使用，請使用以下方式
	recorder ,ok := player.(gadget.TapeRecorder) //斷言原始型別是TapeRecorder
	if ok{
		recorder.Record()
	}else{
		fmt.Println("Player was not a TapeRecorder")
	}
}

func playList(device Player,songs []string){
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string {"Jessie's Girl","Whip It","9 to 5"}
	var player Player = gadget.TapePlayer{}
	playList(player,mixtape)
	player = gadget.TapeRecorder{}
	playList(player,mixtape)

	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}