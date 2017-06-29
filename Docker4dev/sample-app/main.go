package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	var (
		baseValue = 50.00
		done      = make(chan bool)
		update    = make(chan float64, 1)
	)
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	go vCoinCorrector(done, update, baseValue)
	go vCoinTracker(update)

	<-done
}

func vCoinCorrector(done chan<- bool, update chan<- float64, baseValue float64) {
	r := rand.New(rand.NewSource(99))
	for {
		correction := r.Float64()
		log.Println("Correction :  ", correction)
		update <- baseValue + correction
		time.Sleep(500 * time.Millisecond)
	}
	done <- true
}

func vCoinTracker(update <-chan float64) {
	tm.Clear()
	tm.MoveCursor(0, tm.Width()/2)
	tm.Println(tm.Background(tm.Color(tm.Bold("vCoin Tracker"), tm.RED), tm.BLACK))
	tm.MoveCursor(2, tm.Width()/2-15)
	tm.Println("===========================================")

	for {
		tm.MoveCursor(5, tm.Width()/2)
		tm.Printf("%.2f\n", <-update)
		tm.Flush()
	}
}
