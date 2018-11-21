package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type TimetampList []int64

func (t TimetampList) Len() int {
	return len(t)
}

func (t TimetampList) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t TimetampList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

const filePath = "/Users/jemy/Downloads/part-00000-29"

func main() {
	fp, openErr := os.Open(filePath)
	if openErr != nil {
		fmt.Println(openErr)
		return
	}

	defer fp.Close()

	timestampList := make(TimetampList, 0, 24*12)
	timestampFlux := make(map[int64]int64)
	var maxBand int64 = 0
	var maxBandTime int64
	var total int64

	bScanner := bufio.NewScanner(fp)
	for bScanner.Scan() {
		line := bScanner.Text()
		line = strings.TrimSuffix(line, ")")
		line = strings.TrimPrefix(line, "(")

		items := strings.Split(line, ",")
		streamTs := items[0]
		fluxData := items[1]

		streamItems := strings.Split(streamTs, "_")
		timestamp, _ := strconv.ParseInt(streamItems[1], 10, 64)
		flux, _ := strconv.ParseInt(fluxData, 10, 64)
		if flux > maxBand {
			maxBand = flux
			maxBandTime = timestamp
		}

		total += flux
		if _, ok := timestampFlux[timestamp]; ok {
			timestampFlux[timestamp] += flux
		} else {
			timestampList = append(timestampList, timestamp)
			timestampFlux[timestamp] = flux
		}
	}

	sort.Sort(timestampList)

	/*
		for _, timestamp := range timestampList {
			t := time.Unix(timestamp, 0).String()
			fmt.Printf("%s\t%d\n", t, int64(float64(timestampFlux[timestamp]/1024/1024/1024)))
			fmt.Printf("%s\t%d\n", t, int64(float64(timestampFlux[timestamp]/1000/1000/1000)))
		}
	*/

	fmt.Println(maxBand / 1000 / 1000 / 1000)
	fmt.Println(maxBand * 8 / 1000 / 1000 / 1000 / 300)
	fmt.Println(time.Unix(maxBandTime, 0).String())
	fmt.Println(total / 1024 / 1024 / 1024)

}
