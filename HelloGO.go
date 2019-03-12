package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

var (
	v1                = dict{1, 2}
	v2                = dict{X: 1}
	v3                = dict{}
	plk               = &dict{1, 2}
	sectionState bool = false
)

type dict struct {
	Y, X int
}

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

func section(title string) {
	if sectionState == false {
		fmt.Printf("----------\\/%s\\/----------\n", title)
		sectionState = true
	} else if sectionState == true {
		fmt.Printf("----------/\\%s/\\----------\n", title)
		sectionState = false
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func learnVar() {
	var hg = "Hello Go@"
	fmt.Println(hg)
}

func learnMath() {
	var randX int = rand.Intn(100)
	var sqrtX float64 = math.Sqrt(100)
	var piX float32 = math.Pi
	fmt.Println("myNum is:", randX)
	fmt.Printf("Squrt:%g \n", sqrtX)
	fmt.Printf("Pi:%g\n", piX)
	fmt.Println("result:", plus(randX, randX))
}

func learnTime() {
	fmt.Println(time.Now())
}

func learnSwap() {
	a, b := swap("A", "B")
	fmt.Println(a, b)
	fmt.Println(split(190))
}

func learnSwith() {
	section("switch")
	fmt.Print("Curr OS ver is:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s \n", os)
	}
	section("switch")
}

func learnIfelse() {
	section("if 和 else、循环与函数")
	sum := 0
	var randX int = rand.Intn(100)
	for sum < plus(randX, randX) {
		sum += 1
		if sum < 100 {
			println("Curr val is under 100")
		} else if sum == 100 {
			println("Curr val is equal 100")
		} else {
			println("Curr val is over 100")
		}
		fmt.Println(sum)
		section("if 和 else、循环与函数")
	}
}

func learnDefer() {
	section("defer")
	defer fmt.Println("world")
	fmt.Println("Hellow")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
	section("defer")
}

func learnPointer() {
	section("指针")
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	*p = *p / 37
	fmt.Println(j)
	section("指针")
}

func learnStruct() {
	section("结构体")
	v := dict{12, 3}
	fmt.Println(v.X)
	v.X = 100
	fmt.Println(v.X)
	section("结构体")
}

func learnStructFild() {
	section("结构体字段")
	v := dict{1, 2}
	Xp := v
	Xp.X = 20
	fmt.Println(Xp)
	section("结构体字段")
}

func learnStructpointer() {
	v := dict{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func learnStructLiterals() {
	section("结构体文法")
	fmt.Println(v1, plk, v2, v3)
	section("结构体文法")
}

func learnArray() {
	section("数组")
	var list [2]string
	list[0] = "hello"
	list[1] = "world"
	fmt.Println(list[0], list[1])
	fmt.Println(list)
	section("数组")
}

func learnSlice() {
	section("切片")
	primes := [6]int{1, 2, 3, 4, 5, 5}
	fmt.Println(primes)

	var s []int = primes[1:5]
	fmt.Println(s)
	section("切片")
}

func learnSlicePointer() {
	section("切片就像数组的引用")
	names := [4]string{
		"Hara", "Kira", "Sakura", "Jin",
	}
	fmt.Println(names)
	sl := names[0:2]
	kl := names[1:3]
	fmt.Println(sl, kl)

	kl[0] = "XXXXX"

	fmt.Println(sl, kl)
	fmt.Println(names)

	section("切片就像数组的引用")
}

func learnSliceLiterals() {
	section("切片文法")
	q := []int{2, 3, 4, 5, 6, 7, 78, 8, 9}
	fmt.Println(q)

	r := []bool{false, true, false, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{3, false},
		{3, false},
		{3, false},
		{3, false},
		{3, false},
	}
	fmt.Println(s)

}

func learnSliceBounds() {
	section("切片的默认行为")

	s := []int{1, 2, 3, 4, 5, 5, 6, 7, 8}

	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
}

func learnSliceLenCap() {
	section("切片的长度与容量")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(s)
	s = s[:0]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s = s[2:]
	printSlice(s)
}

func learnNilslice() {
	section("nil 切片")
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}

func learnMakeSlice() {
	section("用 make 创建切片")
	a := make([]int, 5)
	printSlice2("a", a)
	b := make([]int, 0, 5)
	printSlice2("b", b)
	c := b[:2]
	printSlice2("c", c)
	d := c[2:5]
	printSlice2("d", d)
}

func main() {
	//learnVar()
	//learnMath()
	//learnTime()
	//learnSwap()
	//learnSwith()
	////learnIfelse()
	//learnDefer()
	//learnPointer()
	//learnStructpointer()
	//learnStructLiterals()
	//learnArray()
	//learnSlice()
	//learnSlicePointer()
	//learnSliceLiterals()
	//learnSliceBounds()
	learnSliceLenCap()
	learnNilslice()
	learnMakeSlice()
}
