package main

import (
	"reflect"
	"testing"
)

func TestConveer(t *testing.T) {
	expected := []float64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}
	t.Run("OK", func(t *testing.T) {
		ch1 := make(chan uint8)
		ch2 := make(chan float64)

		go func() {
			for i := 0; i < 10; i++ {
				ch1 <- uint8(i)
			}
			close(ch1)
		}()

		go conveer(ch1, ch2)
		actualSlice := []float64{}
		for val := range ch2 {
			if reflect.TypeOf(val).Kind() != reflect.Float64 {
				t.Errorf("Wrong type. Expected float64, got %v", reflect.TypeOf(val).Kind())
			}
			actualSlice = append(actualSlice, val)
		}
		if !reflect.DeepEqual(expected, actualSlice) {
			t.Errorf("Wrong result. Expected %v, got %v", expected, actualSlice)
		}
	})
}
