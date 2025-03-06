package common

type GenericData interface {
	Data() any
}

type Data[T DataType] struct {
	data any
	GenericData
}

func (d *Data[T]) Data() any {
	return d.data
}
