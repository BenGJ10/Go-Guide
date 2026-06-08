package main

import "fmt"

type Player struct {
	Name   string
	Number int
	Nation string
}

func practiceMaps() {
	players := make(map[string]string) // players -> nation
	fmt.Println("Content of players map:", players)

	// Adding key-value pairs to the map
	players["Messi"] = "Argentina"
	players["Ronaldo"] = "Portugal"
	players["Neymar"] = "Brazil"

	fmt.Println("Content of players map:", players)

	// Deleting a key-value pair from the map
	delete(players, "Neymar")
	fmt.Println("\nDeleted Neymar from players map.")

	// Checking if a key exists in the map
	if _, exists := players["Neymar"]; exists {
		fmt.Println("Neymar is in the players map.")
	} else {
		fmt.Println("Neymar is not in the players map.")
	}

	// We can also create a map with struct values
	playerDetails := map[string]Player{
		"Messi":   {"Lionel Messi", 10, "Argentina"},
		"Ronaldo": {"Cristiano Ronaldo", 7, "Portugal"},
		"Lamine":  {"Lamine Yamal", 19, "Spain"},
	}
	fmt.Println("\nContent of playerDetails map:", playerDetails)

	// Making a map with slice values
	playerTeams := map[string][]string{
		"Messi":   {"Barcelona", "Paris Saint-Germain"},
		"Ronaldo": {"Manchester United", "Real Madrid", "Juventus"},
	}
	fmt.Println("\nContent of playerTeams map:", playerTeams)

	// We can append to the slice in the map
	playerTeams["Messi"] = append(playerTeams["Messi"], "Inter Miami")
	fmt.Println("\nUpdated playerTeams map:", playerTeams)

	// Get the teams of a player
	messiTeams := playerTeams["Messi"]
	fmt.Println("\nTeams of Messi:", messiTeams)

}
