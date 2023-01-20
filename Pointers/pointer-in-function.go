package main

import "fmt"

/*
Using pointer as parameter in function to save memory usage.
it's recommend to use it.
*/

type Music struct {
	name, artist, genre string
}

func updateGenre(music *Music, genre string) {
	music.genre = genre // set value
}

func main() {
	myMusic := &Music{
		name:   "Jakarta Hari Ini",
		artist: "For Revenge",
		genre:  "Dangdut",
	}
	fmt.Println(myMusic)

	updateGenre(myMusic, "Pop Alternative")
	fmt.Println(myMusic)
}
