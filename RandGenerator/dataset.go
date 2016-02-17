package RandGenerator

type Dataset struct {
	data []float64
}

func NewDataset(gen Generator, size int) *Dataset {
	set := &Dataset{
		data: make([]float64, size),
	}
	for i := 0; i < size; i++ {
		set.data[i] = gen.Float()
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
