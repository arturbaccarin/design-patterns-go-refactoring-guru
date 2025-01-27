package main

/*
You are developing a media player application. Your application can only play .mp3 files,
but you need to extend it to also play .mp4 and .vlc files.

Instead of modifying the existing media player code (to adhere to the Open-Closed Principle),
you will use the Adapter pattern to make the new file formats compatible.
*/
type MediaPlayer interface {
	PlaySong(song string)
}

type MP3Player struct{}

func (m *MP3Player) PlaySong(song string) {
	println("Playing MP3 song:", song)
}

type NewMediaPlayer interface {
	PlayMP4(song string)
	PlayVLC(song string)
}

type NewPlayer struct {
}

func (n *NewPlayer) PlayMP4(song string) {
	println("Playing MP4 song:", song)
}

func (n *NewPlayer) PlayVLC(song string) {
	println("Playing VLC song:", song)
}

type MediaPlayerAdapter struct {
	player MediaPlayer
}

func NewMediaPlayerAdapter(player MediaPlayer) *MediaPlayerAdapter {
	return &MediaPlayerAdapter{
		player: player,
	}
}

func (a *MediaPlayerAdapter) PlayMP4(song string) {
	a.player.PlaySong(song)
}

func (a *MediaPlayerAdapter) PlayVLC(song string) {
	a.player.PlaySong(song)
}

func main() {
	song := "fear of the dark"

	var player NewMediaPlayer

	player = &NewPlayer{}

	player.PlayMP4(song)
	player.PlayVLC(song)

	mediaPlayer := &MP3Player{}

	adapter := NewMediaPlayerAdapter(mediaPlayer)

	adapter.PlayMP4(song)
	adapter.PlayVLC(song)
}
