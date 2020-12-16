package ex16

import (
	"fmt"
	"testing"
)

func TestIsRange(t *testing.T) {
	fmt.Println(isRange("class: 1-3 or 5-7"))
}

func TestGetRange(t *testing.T) {
	name, validator := getRange("class: 1-3 or 5-7")
	fmt.Println(name)
	fmt.Println(validator(1), validator(2), validator(4))
}
