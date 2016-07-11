package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var folderType = flag.String("f", "ym", "Create folder to group calendar folders.")
	var startDays int
	flag.IntVar(&startDays, "s", 0, "The days start to create calendar folders.(Default 0)")
	var totalDays int
	flag.IntVar(&totalDays, "n", 100, "Total days to create.")
	var beginningiDate = flag.String("b", "9999-01-01", "The beginning date of calendar folders. The first folder's name.")
	var endDate = flag.String("e", "9999-01-01", "The end date of calendar folders. The last folder's name.")
	flag.Parse()
	fmt.Printf("Hello, world.\n folderType: %s \n beginningDate: %s \n endDate: %s\n", folderType, beginningiDate, endDate)
	var dateString string
	var startDate = time.Now().AddDate(0, 0, startDays)
	for i := 0; i < totalDays; i++ {
		//fmt.Println("Count number: %d ", i)
		date := startDate.AddDate(0, 0, i)
		dateString = fmt.Sprint(date.Format("2006-01-02"))
		err := os.MkdirAll(dateString, 0777)
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			dir, _ := os.Getwd()
			fmt.Printf("Create Directory %d OK! : %s/%s\n", i, dir, dateString)
		}
	}
}
