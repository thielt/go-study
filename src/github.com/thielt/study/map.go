package main

import "fmt"

func main() {
	// hashmap 
	//alternate way to initialize
	//statePopulations = make(map[string]int)

	statePopulations := map[string]int {
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illinois": 12801539,
		"Ohio": 11614373,
	}
	// m := map[[3]int]string{} //converting to an array
	// map[California:39250017 Florida:20612439 Illinois:12801539 New York:19745289 Ohio:11614373 Pennsylvania:12802503 Texas:27862596] map[]
	//cannot use slice as a valid key type
	//fmt.Println(statePopulations, m)

	// fmt.Println(statePopulations["Ohio"])

	//return order of a map is never guaranteed

	//adding an entry:
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Illinois")
	fmt.Println(statePopulations["Illinois"]) //will print out 0

	pop, ok := statePopulations["Illinois"]
	//, ok syntax is used to retrieve boolean value if the item is still in the map
	fmt.Println(pop, ok)

	// manipulating copy of the map in one place will change it everywhere else
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
	//both will not have ohio anymore

}