package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func askQuestion(question string) string { //Ask the user a question, do some basic cleaning up, and return the value
	fmt.Print(question)
	reader := bufio.NewReader(os.Stdin)
	id, errread := reader.ReadString('\n')
	id = strings.TrimSpace(id)
	if errread != nil {
		log.Fatal(errread)
	}
	fmt.Println("")
	return id
}

func looptheids(count int, prefix string, isSignalBadge bool, criticalblocks [3]string) { //Make prefix a two letter all caps hex code
	var typeofresp string
	if isSignalBadge {
		typeofresp = "7F"
	} else {
		typeofresp = "01"
	}
	for i := 0; i < count; i++ {
		fmt.Println("00" + criticalblocks[0] + "0000" + criticalblocks[1] + prefix + "00A" + strconv.Itoa(i) + criticalblocks[2] + "13000100" + typeofresp + "00")
	}
}

func main() {
	var criticalblocks [3]string
	starttoken := askQuestion("Input one of your request tokens: ")
	if len(starttoken) == 31 { //Handle a 31 bit badge. Credit to Simonomi for figuring out the true format https://github.com/d1str0/dc29-badge/issues/1#issuecomment-894018025
		fmt.Println("Looks like this is a 31 token. Correcting...")
		criticalblocks[0] = "0" + starttoken[2:3]                //Second hex set, also a component in the ID?
		criticalblocks[1] = starttoken[3:4] + starttoken[8:9]    //Fifth hex set
		criticalblocks[2] = starttoken[9:10] + starttoken[16:19] //9th and 10th hex set, the ID of the device
	} else {
		criticalblocks[0] = starttoken[2:4]   //Second hex set, also a component in the ID?
		criticalblocks[1] = starttoken[8:10]  //Fifth hex set
		criticalblocks[2] = starttoken[16:20] //9th and 10th hex set, the ID of the device
	}
	fmt.Println(`To get to the signal tier: `)
	looptheids(7, "AA", true, criticalblocks)
	fmt.Println(`To spread the signal: `)
	looptheids(10, "AB", false, criticalblocks)
	looptheids(10, "AC", false, criticalblocks)
}
