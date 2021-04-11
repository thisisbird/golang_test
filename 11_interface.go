package main
import "github.com/headfirstgo/gadget"

func playList(device gadget.TapePlayer,songs []string){
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	// player := gadget.TapeRecorder{} //只能使用對應的interface
	mixtape := []string {"Jessie's Girl","Whip It","9 to 5"}
	playList(player,mixtape)
}