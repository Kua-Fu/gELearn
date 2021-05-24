package main

import "fmt"

type t struct {
}

// f1 f2
func f1() error {
	fmt.Println("f1")
	fmt.Println("ff")
	return nil
}

func main() {
	f1()
	fmt.Println("hello world!")
	fmt.Println("hello sunshine!")
}

// ESIndexDateInfo = {

// 	"wksp_1_keyevent": [
// 		{
// 			"wksp_1_keyevent-000001":
// 			[1611736990817, 1619343990838]
// 		},
// 		{
// 			"wksp_1_keyevent-000002":
// 			[1611736990817, 1619343990838]
// 		}
// 	],
// 	"wksp_2_keyevent": [
// 		{
// 			"wksp_2_keyevent-000001":
// 			[1611736990817, 1619343990838]
// 		},
// 		{
// 			"wksp_2_keyevent-000002":
// 			[1611736990817, 1619343990838]
// 		}
// 	],
// }
