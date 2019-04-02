package groutine

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {

	for i := 1; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
