package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Printf("Hello, world.\n")
	var dateString string
	for i := 0; i < 100; i++ {
		//fmt.Println("Count number: %d ", i)
		date := time.Now().AddDate(0, 0, i)
		dateString = fmt.Sprint(date.Format("2006-01-02"))
		err := os.MkdirAll(dateString, 0777)
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Print("Create Directory OK!/n")
		}
	}
}