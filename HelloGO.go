package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func plus(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

type dict struct {
	Y, X int
}

var (
	v1  = dict{1, 2}
	v2  = dict{X: 1}
	v3  = dict{}
	plk = &dict{1, 2}
)

func section(title, stype string) {
	if stype == "t" {
		fmt.Printf("----------\\/%s\\/----------\n", title)
	} else if stype == "b" {
		fmt.Printf("----------/\\%s/\\----------\n", title)
	}
}

func main() {
	var hg = "Hello Go@"
	var randX int = rand.Intn(100)
	var sqrtX float64 = math.Sqrt(100)
	var piX float32 = math.Pi
	fmt.Println(hg, time.Now())
	fmt.Println("myNum is:", randX)
	fmt.Printf("Squrt:%g \n", sqrtX)
	fmt.Printf("Pi:%g\n", piX)
	fmt.Println("result:", plus(randX, randX))
	a, b := swap("A", "B")
	fmt.Println(a, b)
	fmt.Println(split(190))
	//sum := 0
	//for sum < plus(randX,randX){
	//	sum += 1
	//	if sum < 100{
	//		println("Curr val is under 100")
	//	}else if sum == 100 {
	//		println("Curr val is equal 100")
	//	}else{
	//		println("Curr val is over 100")
	//	}
	//fmt.Println(sum)
	fmt.Print("Curr OS ver is:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}
	//defer fmt.Println("world")
	//fmt.Println("Hellow")
	//
	//fmt.Println("counting")
	//for i := 0;i<10;i++{
	//	defer fmt.Println(i)
	//}
	//fmt.Println("Done")

	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)

	v := dict{12, 3}
	fmt.Println(v.X)
	v.X = 100
	fmt.Println(v.X)

	Xp := v
	Xp.X = 20
	fmt.Println(Xp)

	fmt.Println(v1, plk, v2, v3)

	section("数组", "t")
	var list [2]string
	list[0] = "hello"
	list[1] = "world"
	fmt.Println(list[0], list[1])
	fmt.Println(list)

	primes := [6]int{1, 2, 3, 4, 5, 5}
	fmt.Println(primes)
	section("数组", "b")

}
