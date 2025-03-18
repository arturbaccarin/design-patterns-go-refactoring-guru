package main

import "fmt"

/*
You’re developing a media player that can switch between different states: "Playing," "Paused," and "Stopped."
Implement the State design pattern in Go to handle these transitions. Create a State interface with methods like Play(),
Pause(), and Stop(). Then, implement concrete states for "Playing," "Paused," and "Stopped"
that change the media player's behavior accordingly. Finally, create a MediaPlayer struct that
maintains the current state and delegates actions to it. Simulate switching states by calling the methods and printing messages for each action.
*/

type MediaPlayer struct {
	song  string
	state State
}

func (m *MediaPlayer) Play() {
	m.state.HandleSong(m)
}

func (m *MediaPlayer) ChangeState(state State) {
	m.state = state
}

type State interface {
	HandleSong(*MediaPlayer)
}

type PlayingState struct {
	MediaPlayer *MediaPlayer
}

func (s *PlayingState) HandleSong(m *MediaPlayer) {
	fmt.Printf("Playing song: %s\n", m.song)
}

type PausedState struct {
	MediaPlayer *MediaPlayer
}

func (s *PausedState) HandleSong(m *MediaPlayer) {
	fmt.Printf("Paused song: %s\n", m.song)
}

type StoppedState struct {
	MediaPlayer *MediaPlayer
}

func (s *StoppedState) HandleSong(m *MediaPlayer) {
	fmt.Printf("Stopped song: %s\n", m.song)
}

func main() {
	player := &MediaPlayer{}

	player.song = "fear of the dark"

	player.ChangeState(&PlayingState{player})
	player.Play()

	player.ChangeState(&PausedState{player})
	player.Play()

	player.ChangeState(&StoppedState{player})
	player.Play()
}
