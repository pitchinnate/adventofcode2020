package day4

import (
	"adventofcode2020/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Byr string
	Iyr string
	Eyr string
	Hgt string
	Hcl string
	Ecl string
	Pid string
	Cid string
}

func validYear(year string, min int, max int) bool {
	date, err := strconv.Atoi(year)
	if err != nil {
		return false
	}
	return date >= min && date <= max
}

func validHeight(height string) bool {
	re := regexp.MustCompile(`(?P<Number>[0-9]+)(?P<Unit>cm|in)`)
	matches := re.FindStringSubmatch(height)
	if len(matches) == 0 {
		return false
	}
	h, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	if matches[2] == "in" && h >= 59 && h <= 76 {
		return true
	}
	if matches[2] == "cm" && h >= 150 && h <= 193 {
		return true
	}
	return false
}

func validHair(hair string) bool {
	hair = strings.TrimSpace(hair)
	if len(hair) != 7 {
		//log.Print("invalid hair color length", hair)
		return false
	}
	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	return re.MatchString(hair)
}

func validEye(eye string) bool {
	eye = strings.TrimSpace(eye)
	if len(eye) != 3 {
		//log.Print("invalid eye color length", eye)
		return false
	}
	return eye == "amb" || eye == "blu" || eye == "brn" || eye == "gry" || eye == "grn" || eye == "hzl" || eye == "oth"
}

func validPassportId(id string) bool {
	id = strings.TrimSpace(id)
	if len(id) != 9 {
		return false
	}
	re := regexp.MustCompile(`[0-9]{9}`)
	return re.MatchString(id)
}

func (passport *Passport) valid() bool {
	if passport.Byr == "" || passport.Iyr == "" || passport.Eyr == "" || passport.Hgt == "" || passport.Hcl == "" || passport.Ecl == "" || passport.Pid == "" {
		return false
	}
	return true
}

func (passport *Passport) betterValid() bool {
	if !passport.valid() {
		return false
	}
	if !validYear(passport.Byr, 1920, 2002) || !validYear(passport.Iyr, 2010, 2020) || !validYear(passport.Eyr, 2020, 2030) {
		return false
	}
	if !validHeight(passport.Hgt) || !validHair(passport.Hcl) || !validEye(passport.Ecl) || !validPassportId(passport.Pid) {
		return false
	}
	return true
}

func Run() {
	log.Println("Running Day 4 code...")
	allData := utils.SplitByLine(utils.ReadInputFile("day4/input.txt"))
	var passports []Passport
	var newPass Passport

	for _, row := range allData {
		//log.Print(row)
		if row == "" {
			passports = append(passports, newPass)
			newPass = Passport{}
			continue
		}
		datas := strings.Split(row, " ")
		//log.Print(datas)
		for _, data := range datas {
			if strings.TrimSpace(data) == "" {
				continue
			}
			pieces := strings.Split(data, ":")
			//log.Print(pieces)
			if len(pieces) == 2 {
				//log.Print(pieces[0], pieces[1])
				err := utils.SetField(&newPass, strings.Title(pieces[0]), pieces[1])
				if err != nil {
					log.Print(err)
				}
			}
		}
		//log.Printf("passport: %#v", newPass)
	}

	part1(passports)
	part2(passports)
}

func part1(passports []Passport) {
	valid := 0
	for _, passport := range passports {
		if passport.valid() {
			valid += 1
		}
	}
	log.Print("valid passports: ", valid)
}

func part2(passports []Passport) {
	valid := 0
	for _, passport := range passports {
		if passport.betterValid() {
			valid += 1
		}
	}
	log.Print("valid passports: ", valid)
}
