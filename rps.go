package main

import (
"fmt"
"github.com/guptarohit/asciigraph"
"log"
"net/http"
"time"
)

const url = "http://localhost/"

func request() {
	_, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func testRpc(reqCount, threshold int) (int, float64) {
	start := time.Now()
	for i := 0; i < reqCount; i++ {
		request()
	}
	end := time.Now()
	duration := end.Sub(start).Milliseconds()
	passed := int(duration) <= threshold
	rps := 1000 / float64(duration) / float64(reqCount)
	log.Printf("[%d requests]\tpassed: %v\t|\ttime: %dms\t|\trps: %f", reqCount, passed, duration, rps)
	return int(duration), rps
}


func main() {
	step := 10
	limit := 100
	thresholdMs := 1000
	rpsHistory := make([]float64, 0)
	durationHistory := make([]float64, 0)
	passedCount := 0
	i := 0
	for reqCount := 25; reqCount <= limit; reqCount += step {
		duration, rps := testRpc(reqCount, thresholdMs)
		durationHistory = append(durationHistory, float64(duration))
		rpsHistory =  append(rpsHistory, rps * 0.01)
		if duration < thresholdMs {
			passedCount++
		}
		i++
	}
	log.Printf("[Total]\t%d/%d tests passed!\t|\tThreshold: <= %dms", passedCount, i, thresholdMs)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("\n[⇘] RPS")
	rpsGraph := asciigraph.Plot(rpsHistory, asciigraph.Width(30), asciigraph.Height(10))
	fmt.Println(rpsGraph)

	durGraph := asciigraph.Plot(durationHistory, asciigraph.Width(30), asciigraph.Height(10))
	fmt.Println("\n[⇘] Duration (sec)")
	fmt.Println(durGraph)
}