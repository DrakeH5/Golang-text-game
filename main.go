package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapString = [7]string{
	"   zzzz    ",
	"  zz  zzz  ",
	"  z  s   z ",
	"  z      z ",
	"  z f x  z ",
	"  zz    z  ",
	"   zzzzzz  ",
}

func reciveInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	if strings.Contains(input, "inventory") {
		fmt.Println(inventory)
		scanner.Scan()
		input = scanner.Text()
	} else if strings.Contains(input, "map") {
		for i := 0; i < 7; i++ {
			fmt.Println(mapString[i])
		}
		scanner.Scan()
		input = scanner.Text()
	}
	return input
}

var inventory = [5]string{"", "", "", "", ""}

var treeFallen = false

func main() {
	fmt.Println("You wake up and your boat is no where to be seen.")
	startingPrompt()
}

func startingPrompt() {
	fmt.Println("You are on the beach of an island. What do you do?")
	input := reciveInput()
	if strings.Contains(input, "swim") {
		startSwim()
	} else if strings.Contains(input, "walk") {
		bigTree()
	} else {
		fmt.Println("Could not understand command")
		startingPrompt()
	}
}

func startSwim() {
	fmt.Println("You enter the water. After swimming a short distance, you see a fin slice through the water. What do you do?")
	input := reciveInput()
	if strings.Contains(input, "swim") {
		fmt.Println("The fin races towards you faster than you can swim away. You feel a jolt of pain race through your left leg as the shark pulls you under. GAME OVER")
	} else if strings.Contains(input, "dead") {
		fmt.Println("Acting dead proves inefficent. The shark eats you anyways. GAME OVER")
	} else if strings.Contains(input, "punch") {
		fmt.Println("Punching the shark when it aproches scares it off. Do you wish to keep swimming out to sea?")
		input := reciveInput()
		if strings.Contains(input, "back") {
			startingPrompt()
		} else if strings.Contains(input, "ahead") {
			fmt.Println("While swimming out to sea, the shark returns and drags you under. GAME OVER")
		} else if strings.Contains(input, "keep") {
			fmt.Println("While swimming out to sea, the shark returns and drags you under. GAME OVER")
		} else {
			fmt.Println("Could not understand command")
			startSwim()
		}
	} else if strings.Contains(input, "kick") {
		fmt.Println("Kicking the shark when it aproches scares it off. Do you wish to keep swimming out to sea?")
		input := reciveInput()
		if strings.Contains(input, "back") {
			startingPrompt()
		} else if strings.Contains(input, "ahead") {
			fmt.Println("While swimming out to sea, the shark returns and drags you under. GAME OVER")
		} else if strings.Contains(input, "keep") {
			fmt.Println("While swimming out to sea, the shark returns and drags you under. GAME OVER")
		} else {
			fmt.Println("Could not understand command")
			startSwim()
		}
	} else if strings.Contains(input, "ride") {
		fmt.Println("Riding the shark proves oddly effective! After a short while, a boat passes by and rescues you. YOU WIN")
	} else {
		fmt.Println("Could not understand command")
		startSwim()
	}
}

func bigTree() {
	fmt.Println("You are at a big tree")
	input := reciveInput()
	if strings.Contains(input, "climb") {
		fmt.Println("While climbing the tree, you fall and break your neck. GAME OVER")
	} else if strings.Contains(input, "punch") {
		fmt.Println("This isn't minecraft...")
		bigTree()
	} else if strings.Contains(input, "cut") || strings.Contains(input, "chop") || strings.Contains(input, "swing") {
		for i := 0; i < 5; i++ {
			if inventory[i] == "axe" {
				fmt.Println("The tree falls over and lands on the beach.")
				i = 6
				treeFallen = true
				fallenTree()
			}
			if i == 4 {
				fmt.Println("You must have an axe to cut down the tree.")
				bigTree()
			}
		}
	} else if strings.Contains(input, "walk") {
		treeStump()
	} else {
		fmt.Println("Could not understand command")
		bigTree()
	}
}

func treeStump() {
	for i := 0; i < 5; i++ {
		if inventory[i] == "axe" {
			i = 6
		}
		if i == 4 {
			fmt.Println("You approach a tree stump with an axe in it.")
			input := reciveInput()
			if strings.Contains(input, "axe") || strings.Contains(input, "pick") || strings.Contains(input, "pull") || strings.Contains(input, "grab") {
				fmt.Println("You pull the axe out of the stump. The blade is somewhat dull and rusty, but it will still work.")
				for i := 0; i < 5; i++ {
					if inventory[i] == "" {
						inventory[i] = "axe"
						i = 6
					}
				}
				input := reciveInput()
				if strings.Contains(input, "walk") {
					bigRock()
				} else {
					fmt.Println("Could not understand command")
					treeStump()
				}
			} else if strings.Contains(input, "walk") {
				bigRock()
			} else {
				fmt.Println("Could not understand command")
				treeStump()
			}
		} else if i == 6 {
			fmt.Println("You approach a tree stump.")
			input := reciveInput()
			if strings.Contains(input, "walk") {
				bigRock()
			} else {
				fmt.Println("Could not understand command")
				treeStump()
			}
		}
	}
}

func fallenTree() {
	fmt.Println("You are at the fallen tree. What do you wish to do?")
	input := reciveInput()
	if strings.Contains(input, "light on fire") || strings.Contains(input, "burn") {
		for i := 0; i < 5; i++ {
			if inventory[i] == "fire starter" {
				fmt.Println("The tree burns. A nearby ship sees the smoke and comes to save you. YOU WIN")
				i = 6
			}
			if i == 4 {
				fmt.Println("You must have a fire starter to burn the tree.")
				fallenTree()
			}
		}
	} else if strings.Contains(input, "walk") {
		bigRock()
	} else {
		fmt.Println("Could not understand command")
		fallenTree()
	}
}

func bigRock() {
	for i := 0; i < 5; i++ {
		if inventory[i] == "fire starter" {
			i = 6
		}
		if i == 4 {
			fmt.Println("You are at a large rock. On top of the rock is a firestarting kit with flint and steel kit.")
			input := reciveInput()
			if strings.Contains(input, "fire starter") || strings.Contains(input, "flint and steel") || strings.Contains(input, "pick") || strings.Contains(input, "pull") || strings.Contains(input, "grab") {
				fmt.Println("You grab the flint and steel.")
				for i := 0; i < 5; i++ {
					if inventory[i] == "" {
						inventory[i] = "fire starter"
						i = 6
					}
				}
				input := reciveInput()
				if strings.Contains(input, "walk") {
					if treeFallen == true {
						fallenTree()
					} else {
						bigTree()
					}
				} else {
					fmt.Println("Could not understand command")
					bigRock()
				}
			} else if strings.Contains(input, "walk") {
				if treeFallen == true {
					fallenTree()
				} else {
					bigTree()
				}
			} else {
				fmt.Println("Could not understand command")
				bigRock()
			}
		} else if i == 6 {
			fmt.Println("You are at a large rock.")
			input := reciveInput()
			if strings.Contains(input, "walk") {
				if treeFallen == true {
					fallenTree()
				} else {
					bigTree()
				}
			} else {
				fmt.Println("Could not understand command")
				bigRock()
			}
		}
	}
}
