package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadInputFile(path string) string {
	contents, _ := ioutil.ReadFile(fmt.Sprintf("c:/www/adventofcode2020/%s", path))
	return string(contents)
}

func SplitByLine(content string) []string {
	return strings.Split(content, "\n")
}

func StringsToInts(content []string) []int64 {
	var vals []int64
	for _, val := range content {
		trimmed := strings.TrimSpace(val)
		if trimmed != "" {
			intVal, _ := strconv.Atoi(trimmed)
			vals = append(vals, int64(intVal))
		}
	}
	return vals
}
