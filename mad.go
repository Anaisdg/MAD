package mad

import (
	"sort"
)

type Dataset struct {
	Values []float64
}

func (d *Dataset) Raw() []float64 {
	return d.Values
}

func (d *Dataset) Sort() []float64 {
	sort.Slice(d.Values, func(i, j int) bool {
		return d.Values[i] < d.Values[j]
	})
	// //Print to check if slice was ordered
	// for _, v := range d.Values {
	// 	fmt.Println(v)
	// }
	return d.Values
}

func (d *Dataset) Median() float64 {
	d.Sort()                  // sort the numbers
	lDataset := len(d.Values) // find length of Dataset
	mNumber := lDataset / 2   // find the median

	if lDataset%2 == 0 {
		return (d.Values[mNumber-1] + d.Values[mNumber]) / 2
	} else {
		return d.Values[mNumber]
	}

}
