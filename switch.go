// 자바와 비슷하게 코딩 / ' 와 " 구분해서 사용해야 함

package main

import "fmt"

// 7. switch (1) : switch뒤에 변수나 exp가 없는 문법
// if else와 유사하지만 보기 더 간결함
// 한가지 변수에 여러 조건을 줄 때 유용함
func canIDrink1(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	return true //swith조건에 안 걸린건 true로 뜨도록 바깥에 return
}

// switch (2) : switch뒤에 변수나 exp가 있는 문법
// 범위는 안되고 일치/불일치 비교만 가능
// fallthrough는 case걸려도 break안되고 아래까지 실행 (리턴 밑에 삽입불가)
func canIDrink2(age int) bool {
	koreanAge := age + 2
	switch koreanAge {
	case 18:
		return false
	case 20:
		return true
	case 30, 31, 33:
		return false
	}
	return true //swith조건에 안 걸린건 true로 뜨도록 바깥에 return
}

func main() {
	fmt.Println(canIDrink(26))
}