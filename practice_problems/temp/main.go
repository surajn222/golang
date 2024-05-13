package main

func main() {

}

func Add(slice1 []int) int {

	sum1 := 0
	for i := range slice1 {
		sum1 = sum1 + slice1[i]
	}
	return sum1
}
