package main

import "fmt"

func main() {
	fmt.Println(createUsers())
	fmt.Println(createCats())
	fmt.Println(createDogs())
}

// TODO: help me to fix me.
func createUsers() [5]string {
	return [5]string{
		"Chewbarka",
		"Jimmy Chew",
		"Bark Wahlberg",
		"Droolius Caesar",
		"Bark Twain",
	}
}

// TODO: help me to fix me.
func createCats() [2]string {
	var cats [2]string
	cats[0] = "Abbie"
	cats[1] = "Oward"
	return cats
}

// TODO: help me to fix me.
func createDogs() (res [2]string) {
	var dogs [2]string
	dogs[0] = "Amber"
	dogs[1] = "Stoot"
	return dogs
}
