package tiny

type Formula[T any] struct {
	Operation[T]
}

func (f *Formula[TOut]) Equals() TOut {

}
