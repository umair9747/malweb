package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(banner)
	log.Println("MalWeb is starting...")
	takeInput()
	log.Println("Looking for URLs from URLhaus...")
	loadURLhaus()
	log.Println("Loaded the URLs from URLhaus into the memory\n")
	log.Println("Starting the scan on target(s)...\n")
	scanTargets()
}
