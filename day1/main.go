package day1

import (
	"adventofcode2020/utils"
	"log"
)

func Run() {
	log.Println("Running Day 1 code...")
	allData := utils.StringsToInts(utils.SplitByLine(utils.ReadInputFile("day1/input.txt")))
	part1(allData)
	part2(allData)
}

func part1(allData []int64) {
	found := false
	for index, val := range allData {
		remainingData := allData[index+1:]
		for _, val2 := range remainingData {
			if val + val2 == 2020 {
				found = true
				log.Printf("two vals are: %d and %d multiplied to: %d", val, val2, val * val2)
				break
			}
		}
		if found {
			break
		}
	}
}

func part2(allData []int64) {
	found := false
	for index, val := range allData {
		remainingData := allData[index+1:]
		for index2, val2 := range remainingData {
			remainingData2 := allData[index2+1:]
			for _, val3 := range remainingData2 {
				if val+val2+val3 == 2020 {
					found = true
					log.Printf("three vals are: %d and %d and %d multiplied to: %d", val, val2, val3, val*val2*val3)
					break
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
}
