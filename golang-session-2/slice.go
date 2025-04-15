package main

import (
	"fmt"
)

/*
	This is a Slice project.
*/

func main() {
	//Slice
	var agents = []string{"Sage", "Phoenix", "Sova", "Brimstone"}
	_ = agents

	//Slice (make function)
	var controllers = make([]string, 3)
	_ = controllers

	fmt.Printf("%#v\n", controllers) //will be empty strings

	//Slice (append function)
	//Assign or set data to controllers
	controllers[0] = "Brimstone"
	controllers[1] = "Viper"
	controllers[2] = "Omen"

	fmt.Printf("%#v\n", controllers)

	//this is where append process, from 3 become 6
	controllers = append(controllers, "Astra", "Harbor", "Clove")

	fmt.Printf("%#v\n", controllers)

	//Slice (append function with ellipsis)
	var duelists = []string{"Neon", "Jett", "Waylay"}
	var initiators = []string{"Sova", "Tejo", "Breach"}

	duelists = append(duelists, initiators...)
	fmt.Printf("%#v\n", duelists)

	//Slice (copy function)
	overrideData := copy(duelists, initiators)

	fmt.Println("Duelist:", duelists)
	fmt.Println("Initiator:", initiators)
	fmt.Println("Number copied of elements:", overrideData)

	//Slice (Slicing)
	var sentinels1 = []string{"Sage", "Cypher", "Killjoy", "Vyse", "Deadlock"}
	var sentinels2 = sentinels1[1:4] //start index 1, ends in 4
	fmt.Printf("%#v\n", sentinels2)

	var sentinels3 = sentinels1[:3] //0 until index 3
	fmt.Printf("%#v\n", sentinels3)

	var sentinels4 = sentinels1[1:] //1 until slice ends
	fmt.Printf("%#v\n", sentinels4)

	var sentinels5 = sentinels1[:] //all slice, same with sentinels1[:len(sentinels1)]
	fmt.Printf("%#v\n", sentinels5)

	//Slice (Combining slicing and append)
	sentinels1 = append(sentinels1[:3], "Chamber") //re-assign index no 3 from Vyse to Chamber
	fmt.Printf("%#v\n", sentinels1)

	//Slice (Backing array)
	sentinels2 = sentinels1[2:4]
	sentinels2[0] = "New Sentinel 1"

	fmt.Println("Sentinels 1: ", sentinels1)
	fmt.Println("Sentinels 2: ", sentinels2)

	//Slice (Cap function)
	var proPlayers1 = []string{"TenZ", "Demon1", "Jawgemo", "lerraf"}
	fmt.Println("Pro Player 1 cap:", cap(proPlayers1)) //4
	fmt.Println("Pro Player 1 len:", len(proPlayers1)) //4
	fmt.Println("")

	var proPlayers2 = proPlayers1[0:3]
	fmt.Println("Pro Player 2 cap:", cap(proPlayers2)) //4
	fmt.Println("Pro Player 2 len:", len(proPlayers2)) //3
	fmt.Println("")

	var proPlayers3 = proPlayers1[1:]
	fmt.Println("Pro Player 3 cap:", cap(proPlayers3)) //4
	fmt.Println("Pro Player 3 len:", len(proPlayers3)) //3
	fmt.Println("")

	//Slice (Creating a new backing array)
	favAgents := []string{"Astra", "Omen", "Sova", "Reyna"}
	learnAgents := []string{}

	learnAgents = append(learnAgents, favAgents[0:2]...)
	favAgents[0] = "Neon"
	fmt.Println("Favorite Agents: ", favAgents)
	fmt.Println("Learning Agents: ", learnAgents)
}

/*
	Note:

	Result:
	[]string{"", "", ""}
	[]string{"Brimstone", "Viper", "Omen"}
	[]string{"Brimstone", "Viper", "Omen", "Astra", "Harbor", "Clove"}
	[]string{"Neon", "Jett", "Waylay", "Sova", "Tejo", "Breach"}
	Duelist: [Sova Tejo Breach Sova Tejo Breach]
	Initiator: [Sova Tejo Breach]
	Number copied of elements: 3
	[]string{"Cypher", "Killjoy", "Vyse"}
	[]string{"Sage", "Cypher", "Killjoy"}
	[]string{"Cypher", "Killjoy", "Vyse", "Deadlock"}
	[]string{"Sage", "Cypher", "Killjoy", "Vyse", "Deadlock"}
	[]string{"Sage", "Cypher", "Killjoy", "Chamber"}
	Sentinels 1:  [Sage Cypher New Sentinel 1 Chamber]
	Sentinels 2:  [New Sentinel 1 Chamber]
	Pro Player 1 cap: 4
	Pro Player 1 len: 4

	Pro Player 2 cap: 4
	Pro Player 2 len: 3

	Pro Player 3 cap: 3
	Pro Player 3 len: 3

	Favorite Agents:  [Neon Omen Sova Reyna]
	Learning Agents:  [Astra Omen]
*/
