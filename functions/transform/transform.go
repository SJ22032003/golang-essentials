package transform

type TransformArgs struct {
	Numbers   *[]int
	Transform func(int) int
}

func TransformNumbers(args TransformArgs) []int {
	transformNumbers := make([]int, len(*args.Numbers))
	for index, val := range *args.Numbers {
		transformNumbers[index] = args.Transform(val)
	}
	return transformNumbers
}

func Triple(num int) int {
	return 3 * num
}

func Double(num int) int {
	return 2 * num
}
