package main

import (
	"fmt"
)

const a int16 = 27

func main()  {
	const myConst int = 23
	fmt.Printf("%v, %T\n", myConst, myConst) // 23, int

	const a int = 14
 fmt.Printf("%v, %T\n", a, a) // 14, int
} 