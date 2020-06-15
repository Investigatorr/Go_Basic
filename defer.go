// 자바와 비슷하게 코딩 / ' 와 " 구분해서 사용해야 함

package main

import (
	"fmt"
	"strings"
)

// !!!꿀팁 : arguments를 여러 개 받고 싶으면 함수선언할 때 ...을 붙이면 된다
// 5. defer메서드 : 함수 호출이 완료되면 특정 문구를 띄워줌 (편리한 기능이라 많이 사용함)
func lenAndUpper(name string) (length int, uppercase string) {
	//함수 끝난거 알려주는 문법 (아주 강력함)
	defer fmt.Println("Im Done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	//선언해놓고 안 쓰면 굉장히 화를 내는 친구
	totallength, up := lenAndUpper("hwang ji hoon") //이거 다음에 defer 바로 출력
	fmt.Println(totallength, up)
}
