package ex12

import (
	"fmt"
	"testing"
)

func TestParseMove(t *testing.T) {
	mType, value := parseMove("W100")
	fmt.Println(mType, value)
}
