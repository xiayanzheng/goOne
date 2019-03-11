package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

func plus(x ,y int)  int{
	return x + y
}

func swap(x, y string) (string, string) {
	return y,x
}

func split(sum int)(x,y int){
	x = sum *4 /9
	y = sum - x
	return
}

type list struct {
	Y int
	X int
}

func main() {
	var hg  = "Hello Go@"
	var randX int = rand.Intn(100)
	var sqrtX float64 = math.Sqrt(100)
	var piX float32 = math.Pi
	fmt.Println( hg,time.Now())
	fmt.Println("myNum is:",randX)
	fmt.Printf("Squrt:%g \n",sqrtX)
	fmt.Printf("Pi:%g\n",piX)
	fmt.Println("result:",plus(randX,randX))
	a,b := swap("A","B")
	fmt.Println(a,b)
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
		fmt.Printf("%s \n",os)
	}
	//defer fmt.Println("world")
	//fmt.Println("Hellow")
	//
	//fmt.Println("counting")
	//for i := 0;i<10;i++{
	//	defer fmt.Println(i)
	//}
	//fmt.Println("Done")

	i,j := 42,2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p=&j
	*p = *p /37
	fmt.Println(j)

	v := list{12,3}
	fmt.Println(v.X)
	v.X = 100
	fmt.Println(v.X)

	Xp := v
	Xp.X = 20
	fmt.Println(Xp)








	}
