package ternary

// If is the replacement ternary operator in Go
//
// Usage:
//
//	ternary.If(condition bool, elementReturnedIfTrue, elementReturnedIfFalse)
//
// Example:
//
//	ternary.If(true, "foo", "bar") // returns "foo"
//	ternary.If(false, "foo", "bar") // returns "bar"
func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func IfFunc[T any, R any](condition bool, a, b T, f func(T) R) R {
	if condition {
		return f(a)
	}
	return f(b)
}

func Iff[T any](condition bool, a, b func() T) T {
	if condition {
		return a()
	}
	return b()
}
