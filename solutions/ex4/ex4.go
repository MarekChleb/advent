package ex4

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type Passport struct {
	BirthYear      string `json:"byr"`
	IssueYear      string `json:"iyr"`
	ExpirationYear string `json:"eyr"`
	Height         string `json:"hgt"`
	HairColor      string `json:"hcl"`
	EyeColor       string `json:"ecl"`
	PassportID     string `json:"pid"`
	CountryID      string `json:"cid"`
}

func PassportFromMap(data map[string]string) (Passport, error) {
	if !mapIsValidWithFields(data) {
		return Passport{}, fmt.Errorf("Map is invalid")
	}
	dataJSON, err := json.Marshal(data)
	if err != nil {
		return Passport{}, err
	}
	var p Passport
	err = json.Unmarshal(dataJSON, &p)
	return p, err
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

var validator = map[string]func(string) bool{
	"byr": BYRValidator,
	"iyr": IYRValidator,
	"eyr": EYRValidator,
	"hgt": HGTValidator,
	"hcl": HCLValidator,
	"ecl": ECLValidator,
	"pid": PIDValidator,
}

const Path = "./solutions/ex4"

func A(lines []string) {
	count := 0
	m := map[string]string{}

	for _, l := range lines {
		if l == "" {
			if mapIsValid(m) {
				count++
			}
			m = map[string]string{}
			continue
		}
		lSplitted := strings.Split(l, " ")
		for _, phr := range lSplitted {
			pS := strings.Split(phr, ":")
			m[pS[0]] = pS[1]
		}
	}

	fmt.Println(count)

}

func B(lines []string) {
	count := 0
	m := map[string]string{}

	for _, l := range lines {
		if l == "" {
			if mapIsValidWithFields(m) {
				count++
			}
			m = map[string]string{}
			continue
		}
		lSplitted := strings.Split(l, " ")
		for _, phr := range lSplitted {
			pS := strings.Split(phr, ":")
			m[pS[0]] = pS[1]
		}
	}

	fmt.Println(count)
}

func mapIsValid(m map[string]string) bool {
	for _, f := range requiredFields {
		if _, ok := m[f]; !ok {
			return false
		}
	}

	return true
}

func mapIsValidWithFields(m map[string]string) bool {
	for _, f := range requiredFields {
		if v, ok := m[f]; !ok {
			return false
		} else {
			if !validator[f](v) {
				return false
			}
		}
	}

	return true
}

func BYRValidator(s string) bool {
	// v, err := strconv.ParseInt(s, 10, 64)
	// if err != nil {
	// 	return false
	// }
	// return v >= 1920 && v <= 2002
	re := regexp.MustCompile("(19[2-9][0-9]|200[0-2])")
	return re.MatchString(s)
}

func IYRValidator(s string) bool {
	// v, err := strconv.ParseInt(s, 10, 64)
	// if err != nil {
	// 	return false
	// }
	// return v >= 2010 && v <= 2020
	re := regexp.MustCompile("20(1[0-9]|20)")
	return re.MatchString(s)
}

func EYRValidator(s string) bool {
	// v, err := strconv.ParseInt(s, 10, 64)
	// if err != nil {
	// 	return false
	// }
	// return v >= 2020 && v <= 2030
	re := regexp.MustCompile("20(2[0-9]|30)")
	return re.MatchString(s)
}

func HGTValidator(s string) bool {
	// rNum := regexp.MustCompile("^[0-9]*")
	// sNum := rNum.FindString(s)
	// val, err := strconv.ParseInt(sNum, 10, 64)
	// if err != nil {
	// 	return false
	// }

	// rS := regexp.MustCompile("[a-z]*$")
	// measure := rS.FindString(rNum.ReplaceAllString(s, ""))

	// if measure == "cm" {
	// 	return val >= 150 && val <= 193
	// } else if measure == "in" {
	// 	return val >= 59 && val <= 76
	// }

	// return false
	re := regexp.MustCompile("^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$")

	return re.MatchString(s)
}

func HCLValidator(s string) bool {
	re := regexp.MustCompile("^#[a-f0-9]{6}$")
	return re.MatchString(s)
}

func ECLValidator(s string) bool {
	re := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	return re.MatchString(s)
}

func PIDValidator(s string) bool {
	re := regexp.MustCompile("^[0-9]{9}$")
	return re.MatchString(s)
}
