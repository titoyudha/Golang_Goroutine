package goroutine

import (
	"fmt"
	"testing"
)

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

}
