package RandGenerator

import (
	"sort"
)

type Dataset struct {
	data []float64
}

func NewDataset(gen Generator, size int) *Dataset {
	set := &Dataset{
		data: make([]float64, size),
	}
	for i := 0; i < size; i++ {
		set.data[i] = gen.Float64()
	}
	return set
}

func (ds *Dataset) Get(index int) float64 {
	return ds.data[index]
}

func (ds *Dataset) Len() int {
	return len(ds.data)
}

func (ds *Dataset) Mean() float64 {
	sum := 0.0
	for _, val := range ds.data {
		sum += val
	}
	return sum / float64(ds.Len())
}

func (ds *Dataset) Sort() *Dataset {
	result := make([]float64, len(ds.data))
	for i, elem := range ds.data {
		result[i] = elem
	}
	sort.Float64s(result)
	return &Dataset{data: result}
}

func (ds *Dataset) Subset(start, end int) *Dataset {
	result := make([]float64, end-start)
	for i := start; i < end; i++ {
		result[i] = ds.Get(i)
	}
	return &Dataset{data: result}
}
