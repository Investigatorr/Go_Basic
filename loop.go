// 자바와 비슷하게 코딩 / ' 와 " 구분해서 사용해야 함

package main

import "fmt"

// loop는 for 하나밖에 없음
func superAdd(numbers ...int) int {
	total := 0
	// _를 붙여야 numbers의 원소 갯수 index를 _로 보내 무시하고,
	// 원소의 값을 number로 받아서 우리가 원하는 덧셈 가능
	for _, number := range numbers {
		fmt.Println(number)
		total += number
	}
	return total
}

func main() {
	result := superAdd(2, 4, 6, 8, 10, 12)
	fmt.Println(result)
}