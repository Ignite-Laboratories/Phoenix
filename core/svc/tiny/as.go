package tiny

import "math/big"

func ParseAs[TOut any](operand any) TOut {

}

// as is used to convert the yielded result from the underlying Context's type into a supported type.
//
// NOTE: This is equivalent to calling 'Equals' on the Formula, just with a different output type.
//
// see.BaselessCalculation
type as[TOut any] struct {
	fn func() Formula[TOut]
}

func (a as[TOut]) Byte() byte {
}
func (a as[TOut]) Rune() rune {

}
func (a as[TOut]) Int() int {

}
func (a as[TOut]) Int8() int8 {

}
func (a as[TOut]) Int16() int16 {

}
func (a as[TOut]) Int32() int32 {

}
func (a as[TOut]) Int64() int64 {

}
func (a as[TOut]) Float32() float32 {

}
func (a as[TOut]) Float64() float64 {

}
func (a as[TOut]) UInt() uint {

}
func (a as[TOut]) UInt8() uint8 {

}
func (a as[TOut]) UInt16() uint16 {

}
func (a as[TOut]) UInt32() uint32 {

}
func (a as[TOut]) UInt64() uint64 {

}
func (a as[TOut]) UIntPtr() uintptr {

}
func (a as[TOut]) Complex64() complex64 {

}
func (a as[TOut]) Complex128() complex128 {

}
func (a as[TOut]) String() string {

}
func (a as[TOut]) BigInt() *big.Int {

}
func (a as[TOut]) BigRat() *big.Rat {

}
func (a as[TOut]) BigFloat() *big.Float {

}
