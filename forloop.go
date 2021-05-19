package main

import "fmt"

func PrivFuncPrint(x int) {
	fmt.Println(x, "yazdirildi")
}

func main() {
	fmt.Println("test")

	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// 	// 1-9
	// }

	// x := 1
	// for {
	// 	fmt.Println(x)
	// 	x++
	// 	break
	// }

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "2 ye tam bolundu")
		} else {
			fmt.Println(i, "3'e tam bolundu")
		}

		switch i {
		case 5:
			fmt.Println("5 yazdirildi")
			PrivFuncPrint(i)
		}
	}

}
