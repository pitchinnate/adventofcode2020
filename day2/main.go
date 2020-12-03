package day2

import (
	"adventofcode2020/utils"
	"log"
	"strconv"
	"strings"
)

type Entry struct {
	min int
	max int
	letter int32
	password string
}

func (entry *Entry) Valid() bool {
	counter := 0
	for _, val := range entry.password {
		if val == entry.letter {
			counter+=1
		}
	}
	return (counter >= entry.min && counter <= entry.max)
}

func (entry *Entry) ValidPass2() bool {
	counter := 0
	length := len(entry.password)
	if entry.min > length || entry.max > length {
		log.Print("Invalid one", entry)
		return false
	}
	if int32(entry.password[entry.min - 1]) == entry.letter {
		counter += 1
	}
	if int32(entry.password[entry.max - 1]) == entry.letter {
		counter += 1
	}
	return counter == 1
}

func Run() {
	log.Println("Running Day 2 code...")
	allData := utils.SplitByLine(utils.ReadInputFile("day2/input.txt"))
	var entries []Entry
	for _, data := range allData {
		trimmed := strings.TrimSpace(data)
		if trimmed == "" {
			continue
		}
 		pieces := strings.Split(trimmed, " ")
		pieces2 := strings.Split(pieces[0], "-")
		min, _ := strconv.Atoi(pieces2[0])
		max, _ := strconv.Atoi(pieces2[1])
		entries = append(entries, Entry{
			min: min,
			max: max,
			letter: int32(pieces[1][0]),
			password: pieces[2],
		})
	}
	part1(entries)
	part2(entries)
}

func part1(entries []Entry) {
	counter := 0
	for _, entry := range entries {
		if entry.Valid() {
			counter += 1
		}
	}
	log.Printf("found %d valid passwords", counter)
}

func part2(entries []Entry) {
	counter := 0
	for _, entry := range entries {
		if entry.ValidPass2() {
			counter += 1
		}
	}
	log.Printf("found %d valid passwords", counter)
}
