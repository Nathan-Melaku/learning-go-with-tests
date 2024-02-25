package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("collection of slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}

func TestSumTrails(t *testing.T) {
	checksum := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	}
	t.Run("collection of trail sums", func(t *testing.T) {
		got := SumTrails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		checksum(t, got, expected)
	})

	t.Run("collection of trail sums zero array", func(t *testing.T) {
		got := SumTrails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		checksum(t, got, expected)
	})
}
