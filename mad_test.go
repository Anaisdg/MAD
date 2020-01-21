package mad

import (
	"reflect"
	"testing"
)

func TestMad(t *testing.T) {
	dataset := Dataset{Values: []float64{55.0, 51.0, 12.0}}

	t.Run("raw", func(t *testing.T) {
		got := dataset.Raw()
		want := []float64{55.0, 51.0, 12.0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("sort", func(t *testing.T) {
		got := dataset.Sort()
		want := []float64{12.0, 51.0, 55.0}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

	t.Run("median", func(t *testing.T) {
		got := dataset.Median()
		want := 51.0
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	})

}

// 	t.Run("difference", func(t *testing.T) {
// 		got := dataset.Diff()
// 		want := []float64{-39.0, 0.0, 4.0}
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("got %v want %v given", got, want)
// 		}
// 	})
// }

// test if we can return a dataset CHECK
// test if we can order the dataset CHECK
// test if we can find the median CHECK
// test if we can find the differences between median and each value
// test if we can find the sort that difference
// test if we can find the median value of that
