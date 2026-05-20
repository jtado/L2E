package student

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		if size == 0 {
			fmt.Println()
		}
		return
	}
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}
	var chunked [][]int
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunked = append(chunked, slice[i:end])
	}
	fmt.Println(chunked)
}
