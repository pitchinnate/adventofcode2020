package day3

import (
	"adventofcode2020/utils"
	"log"
	"strings"
)

type Row struct {
	Data string
	Index int
}

const inputs = "#."
var tree = inputs[0]

func (row *Row) CheckForTree(position int) bool {
	moddedX := position
	if position >= 31 {
		moddedX = position % 31
	}
	result := row.Data[moddedX] == tree
	//log.Printf("check for tree on row at position %d - %d - %v\n%s", position, moddedX, result, row.Data)
	return result
}

func Run() {
	log.Println("Running Day 3 code...")
	allData := utils.SplitByLine(utils.ReadInputFile("day3/input.txt"))
	var rows []Row
	for index, data := range allData {
		if strings.TrimSpace(data) != "" {
			rows = append(rows, Row{data, index})
		}
	}
	part1(rows)
	part2(rows)
}

func part1(rows []Row) {
	findTrees(3, 1, rows)
}

func part2(rows []Row) {
	treesHit := findTrees(1, 1, rows)
	treesHit2 := findTrees(3, 1, rows)
	treesHit3 := findTrees(5, 1, rows)
	treesHit4 := findTrees(7, 1, rows)
	treesHit5 := findTrees(1, 2, rows)
	log.Printf("Answer: %d", treesHit * treesHit2 * treesHit3 * treesHit4 * treesHit5)
}

func findTrees(moveRight int, moveDown int, rows []Row) int {
	numRows := len(rows)
	treesHit := 0
	currentX := 0
	currentY := 0
	for {
		if rows[currentY].CheckForTree(currentX) {
			treesHit += 1
		}
		currentY += moveDown
		currentX += moveRight
		if currentY >= numRows {
			break
		}
	}
	log.Printf("Found this many trees: %d", treesHit)
	return treesHit
}
