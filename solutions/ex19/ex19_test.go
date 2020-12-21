package ex19

import (
	"fmt"
	"testing"
)

func TestParseLine(t *testing.T) {
	fmt.Println(parseLine("70: 90 92 | 53 50"))
	fmt.Println(parseLine("53: \"a\""))
}
