package RandGenerator

type Dataset struct {
	data []float64
}

func NewDataset(gen Generator, size int64) *Dataset {
	set := &Dataset{
		data: make([]float64, size),
	}
	for i := 0; i < size; i++ {
		set.data[i] = gen.Float()
	}
	return set
}

func (ds *Dataset) Get(index int64) {
	return ds.data[index]
}

func (ds *Dataset) Len() int64 {
	return len(ds.data)
}

func (ds *Dataset) Mean() int64 {
	sum := 0
	for _, val := range ds.data {
		sum += val
	}
	return sum / ds.Len()
}
