package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func main() {
	filename := "Frankenstein.txt"
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(filename) // Error handling elided for brevity.
	io.Copy(buf, f)           // Error handling elided for brevity.
	f.Close()
	text := string(buf.Bytes())
	fmt.Println("Hello, playground")
	validRegex := regexp.MustCompile(`([\w'-]+|[.,!?;&])`)
	splitText := validRegex.FindAllString(text, -1)
	finalKeypairs := make(map[[2]string][]string)
	previousItems := [2]string{"#STARTING... 2#", "#STARTING... 1#"}
	for _, item := range splitText {
		finalKeypairs[previousItems] = append(finalKeypairs[previousItems], item)
		previousItems[0] = previousItems[1]
		previousItems[1] = item
	}
	endingItems := [2]string{"#ENDING... 1#", "#ENDING... 2#"}
	for i := 0; i < 2; i++ {
		finalKeypairs[previousItems] = append(finalKeypairs[previousItems], endingItems[i])
		previousItems[0] = previousItems[1]
		previousItems[1] = endingItems[i]
	}
	fmt.Println(finalKeypairs)
	fmt.Println("")
	lastSaid := [2]string{"#STARTING... 2#", "#STARTING... 1#"}
	count := 0
	for (lastSaid[1] != endingItems[1]) && (count < 1000) {
		randMax := len(finalKeypairs[lastSaid]) - 1
		var itemIndex int
		if randMax > 0 {
			randItem := rand.New(rand.NewSource(time.Now().UnixNano()))
			itemIndex = randItem.Intn(randMax)
		} else {
			itemIndex = randMax
		}
		item := finalKeypairs[lastSaid][itemIndex]
		fmt.Print(item + " ")
		lastSaid[0] = lastSaid[1]
		lastSaid[1] = item
		count++
	}
}
