package main

/*包引入
包声明
函数
变量
表达式
*/
//
//import (
//	"fmt"
//)
//
//var e int
//
//func main() {
//	var a,b = 1,2
//
//
//	var f bool
//
//	var g string
//	fmt.Printf("%v %q\n",f,g)
//	fmt.Println("----------------------")
//	c:= "applenumber = %d&name = %s"
//	//这是我的第一个go 代码
//	d:= fmt.Sprintf(c,a,b)
//	fmt.Println(d)
//
//}
//----------------------------------------------------------
//
//func main()  {
//	const WIDTH = 10
//	const HEIGHT = 5
//	var area  int
//	const a,b,c = 1,false,"ttt"
//
//	area = WIDTH * HEIGHT
//	fmt.Printf("面积是%d", area)
//	println()
//	println(a,b,c)
//
//}

//import "unsafe"
//
//const (
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//)
//
//	func main(){
//		println(a,b,c)
//	}
//
//
//------------------------------

//import "fmt"
//const (
//	i=1<<iota
//	j=3<<iota
//	k
//	l
//)
//
//func main() {
//	fmt.Println("i=",i)
//	fmt.Println("j=",j)
//	fmt.Println("k=",k)
//	fmt.Println("l=",l)
//}

//---------------------------------

//import "fmt"
//
//func main()  {
//	a := 5
//	b := 10
//	c := a + b
//	d := a * b
//	e := a - b
//	f := a / b
//	g := a % b
//	if(c > 15) {
//		fmt.Println(a+b)
//	} else{
//		fmt.Println(a+b)
//	}
//	if(d >50){
//		fmt.Println(a*b)
//	}else{
//		fmt.Println(a*b)
//	}
//	if(e < 5){
//		fmt.Println(a-b)
//	}else{
//		fmt.Println(a-b)
//	}
//	if(f!=1){
//		fmt.Println(a/b)
//	}else{
//		fmt.Println(a/b)
//	}
//	if(g>=1){
//		fmt.Println(a%b)
//	}else{
//		fmt.Println(a%b)
//	}
//
//	println(a,b,c,d,e,f,g)
//
//
//}

//--------------------------------
//import "fmt"
//
//func main() {
//	var a int = 21
//	var b int = 10
//
//	if( a == b ) {
//		fmt.Printf("第一行 - a 等于 b\n" )
//	} else {
//		fmt.Printf("第一行 - a 不等于 b\n" )
//	}
//	if ( a < b ) {
//		fmt.Printf("第二行 - a 小于 b\n" )
//	} else {
//		fmt.Printf("第二行 - a 不小于 b\n" )
//	}
//
//	if ( a > b ) {
//		fmt.Printf("第三行 - a 大于 b\n" )
//	} else {
//		fmt.Printf("第三行 - a 不大于 b\n" )
//	}
//	/* Lets change value of a and b */
//	a = 5
//	b = 20
//	if ( a <= b ) {
//		fmt.Printf("第四行 - a 小于等于 b\n" )
//	}
//	if ( b >= a ) {
//		fmt.Printf("第五行 - b 大于等于 a\n" )
//	}
//}
//----------------------------------------
//
//import "fmt"
//
//func main()  {
//	var a bool = true
//	var b bool = false
//
//	if( a && b ){
//		fmt.Println("a和b都是true")
//	}else{
//		fmt.Println("a或b至少有一个不是true")
//	}
//	if( a || b ){
//		fmt.Println("a或者b是true")
//	}
//	if( !(a&&b) ){
//		fmt.Println("a和b不都是true")
//	}
//
//}
//--------------------------------------------
//import "fmt"
//
//func main() {
//	sum:=0
//	for i :=1;i<=10;i++{
//		sum+=i
//	}
//
//	fmt.Println(sum)
//}
//--------------------------------------------------

//

//func print1(a, b int)int {
//	var c int
//	c = a + b
//	return c
//}
//
//
//
//func main()  {
//	fmt.Println(print1(100,6))
//
//}

/*
func main(){
	var MyArraylist [10] int
	var i,j int

	for i=0;i<10;i++ {
		MyArraylist[i] = i+100
	}

	for j =0;j<10;j++ {
		fmt.Printf("%d time is %d\n",j,MyArraylist[j])
	}
}
*/

//------------------------
/*
import "fmt"

func main()  {
	var j int
	n:= [5] float32{1.0,2.4,2.5,4.5,1.3}
	fmt.Println(n)
	fmt.Println(n[3])
	for j= 0;j<5;j++{
		fmt.Printf("这个列表中的第%d是%f\n",j,n[j])
	}
}
*/
//-----------------------------------------
/*import "fmt"

func main()  {
	a:= [][] int{}
	row1 := []int {1,3,4}
	row2 := []int {2,5,6}
	row3 := []int{45,23,52,523 }
	a = append(a,row1)
	a = append(a,row2)
	a = append(a,row3)
	fmt.Println(a)
	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[2][2])
}*/

//------------------------------------------
/*
								指针
	变量是一个存放数据的容器，声明变量就是在计算机内存中申请一个地址的过程，或者叫申请内存空间的过程
	给变量赋值就是在 某个内存地址中存放数据的过程
	那么指针作为一种特殊的变量，它的本质就是指向某个变量的内存地址。 指针实际上就是对某个变量的代理，你要访问某个变量，可以先访问这个指向这个变量地址的指针。
	当一个指针没有指向任何变量的内存地址时，它被称为空指针。 它虽然是代理，但是它没有代理任何变量

	某个类型的指针变量可以指向任何类型的变量 ，但是在使用指针时，指针和指针所指向的变量类型要保持一致，这样变量和指针指向的内存地址和变量所在的内存位置才是一致的。
*/

/*import "fmt"

func main()  {
	a:=10
	var ip *int2
	ip = &a
	var ptr *float32
	fmt.Printf("这个变量的地址是%x\n",&a)
	fmt.Printf("*p的地址是%x\n",ip)
	fmt.Printf("*p的值是%d\n",*ip)
	if (ptr!=nil){
		fmt.Printf("ptr不是空指针,它的值是%x\n",ptr)
	}else{
		fmt.Printf("ptr是空指针,它的值是%x\n",ptr)
	}

}*/
/*
import "fmt"

type Player struct{
	name string
	age int
	sex string
	occupation string
}

func main()  {
	var farmer Player
	var Programer Player

	farmer.name = "alen"
	farmer.age = 33
	farmer.sex = "男"
	farmer.occupation = "农民"

	Programer.name = "Ulibrchi"
	Programer.age = 25
	Programer.sex = "男"
	Programer.occupation = "程序员"


	printPlayer(&farmer)
	printPlayer(&Programer)

}

func printPlayer(a *Player)  {
	fmt.Println(a.name,"\n")
	fmt.Println(a.age,"\n")
	fmt.Println(a.sex,"\n")
	fmt.Println(a.occupation,"\n")

}*/

//---------------------------------
/*import  "fmt"

func main()  {
	var a int = 30
	var b *int

	b = &a

	fmt.Println("\n",a)
	fmt.Println("\n",b)
	fmt.Println("\n",*b)
}*/
//--------------------------------
/*import "fmt"


func main()  {
	a1 :=[3]int{34,23,12}
	fmt.Println(a1)
	var c int
	var b [3]*int
	for c =0;c<3;c++{
		b[c] = &a1[c]
	}
	for c=0;c<3;c++{
		fmt.Printf("%d\n",*b[c])
	}

}
*/
//-------------------------
/*import "fmt"

func main()  {
	var n = make([]int,7)
	fmt.Println(n)
	printSlice(n)
}

func printSlice(x []int){
	fmt.Printf("长度%d , 大小%d, 切片数组%v",len(x),cap(x),x)
}*/
//------------------------------------
/*
import "fmt"

func main()  {
	var numbers []int
	numbers = append(numbers,2)
	fmt.Println(numbers)

	numbers = append(numbers,4)
	fmt.Println(numbers)

	numbers = append(numbers,6,5,3)
	fmt.Printf("numbers的长度是%d,容量是%d,numbers是%v\n",len(numbers),cap(numbers),numbers)

	var numbers1 = make([]int,len(numbers),(cap(numbers))*2)

	copy(numbers1,numbers)

	appendSlice(numbers1)
}

func appendSlice(x []int){
	fmt.Printf("numbers1的长度是：%d 容量是：%d,numbers1是：%v",len(x),cap(x),x)

}*/

/*import "fmt"

func main()  {
	nums := [4]int{13,32,41,53}
	sum:= 0
	for _,num :=range nums{
		sum +=num
		fmt.Printf("num是：%d,sum是：%v\n",num,sum)
	}
}*/
//----------------------------------------

/*import "fmt"

func main()  {
	//var countryCapitalMap map[string] string
	//countryCapitalMap = make(map[string]string)


	countryCapitalMap := make(map[string]string)

	countryCapitalMap ["France"] = "巴黎"
	countryCapitalMap ["China"] = "北京"
	countryCapitalMap ["Russia"] = "莫斯科"
	countryCapitalMap ["Cuba"] = "哈瓦那"
	for country :=range countryCapitalMap{
		fmt.Println(country,"首都是:",countryCapitalMap[country])

	}

	capital,ok := countryCapitalMap ["American"]

	if(ok){
	fmt.Printf("American 的capital是%s",capital)
	}else {
	fmt.Printf("American 的capital不存在\n")
	}

	delete(countryCapitalMap,"France")
	for country:= range countryCapitalMap{
		fmt.Println("删除后的country\n",country)
	}
}*/

//--------------------------------
/*import "fmt"

func jc(n uint64)(re uint64){
	if (n>0)||(n<13){
		re = n*jc(n-1)
		return re
	}
	return 1
}


func main()  {
	i:=12
	fmt.Printf("%d的阶乘是%d",i,jc(uint64(i)))

}*/
//go 语言中接口的实现方法是隐式的。
/*import "fmt"

type Yifu interface {
	wear()
}


type Clothes struct {

}

func (tshirt Clothes)wear()  {
	fmt.Println("穿衣服")

}

type Shortes struct {

}

func (xtep Shortes)wear()  {
	fmt.Println("穿袜子")


}

func main()  {
	var chuanyifu Yifu
	chuanyifu = new(Clothes)
	chuanyifu.wear()
	chuanyifu = new(Shortes)
	chuanyifu.wear()

}*/

//--------------------
//错误处理---------------------------------------------------不理解

/*import (
	"errors"
	"fmt"
)

func Sqrt(n int)(int,error)  {
	if(n<0){
		return 0,errors.New("squrt root is not vaild number!")
	}
	return n + 1, nil
}

func main()  {
	var i1 int = 5
	_, err := Sqrt(i1)
	if err != nil {
		return
	}
	fmt.Println("开平方")

}*/
//---------------------------------------------------
/*import "fmt"

func summary(s []int,c chan int){
	sum:= 0
	//_是匿名变量
	for _,v := range s{
		sum+=v
	}
	c <-sum
}

func main()  {
	s:=[]int{7,2,8,-9,4,0}

	c:=make(chan int)
	go summary(s[:len(s)/2],c)
	go summary(s[len(s)/2:],c)
	x,y :=<-c,<-c
	fmt.Println(x,y,x+y)


}*/
