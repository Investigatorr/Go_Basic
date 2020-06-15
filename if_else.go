package main

import "fmt"

// if절에 return 있으면 else 따로 필요 없음
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {  //이렇게 작성해야 조건문임을 알아보기 쉬움 (변수는 변수대로 계산, 조건문은 제약식에 걸리는 구조)
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(1))
}