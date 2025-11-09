package tiny

// An Operation is the source of a calculation chain, beginning with a path to the tiny.Context.
// Most operations yield a Formula, which can be used to evaluate the operation into a result.
//
// see.BaselessCalculation
type Operation[T any] std.Path
