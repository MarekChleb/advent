package ex4

import (
	"fmt"
	"testing"
)

func TestPassportFromMap(t *testing.T) {
	m := map[string]string{
		"ecl": "gry", "pid": "860033327", "eyr": "2020", "hcl": "#fffffd",
		"byr": "1937", "iyr": "2017", "cid": "147", "hgt": "183cm",
	}

	p, err := PassportFromMap(m)
	fmt.Println(p, err)
}
