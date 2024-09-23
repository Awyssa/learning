package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	fmt.Println("Lesson 8.5: Channels extended")

	var chickenChannel = make(chan string)
	var websites = []string{"Waitrose.com", "Tesco.com", "Sainsbury.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20

		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}
