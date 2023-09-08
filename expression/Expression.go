package expression

func Ternary[T any](boolean bool, trueValue T, falseValue T) T {
	if boolean {
		return trueValue
	} else {
		return falseValue
	}
}
