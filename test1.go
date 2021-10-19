package main
/*
import "fmt"

func a()  {
	x:=0
	y:=1
	for i:=0;i<3;i++{
		//x,y = x+y x,y同时赋值，赋值时只按照上次赋值变量结果赋值，y只取上次y的赋值结果赋值，不能取本次x,或者y的赋值的结果再进行赋值
		x,y = y,x+y
		fmt.Println("x---",x)
		fmt.Println("y-----",y)
	}
}

func b()  {
	x:=0
	y:=1
	for i:=0;i<3;i++{
		//x,y分别赋值时，y在赋值时会取本次x赋值的结果。
		x = y
		y = x+y
		fmt.Println("x---",x)
		fmt.Println("y-----",y)
	}
}



func main()  {
	a()
	fmt.Println("---------------------------------------")
	b()
	fmt.Println("---------------------------------------")*/
	/*通过测试可知道a()和b()的赋值过程是完全不同的
	分别赋值
	b()
	x = 0
	y = 1

	for i in range(3)
	i = 0

	x = y = 1
	x取本次赋值结果1
	y = x+y =1+1 = 2

	i = 1

	x = y = 2
	x取本次赋值结果2
	y = x+y = 2+2 = 4


	i = 2
	x = y = 4
	x取本次赋值结果4
	y = x+y = 4+4 = 8
	-------------------------------
		同时赋值
	a()

	x = 0
	y = 1

	for i in range(3)
	i = 0
	x,y = y,x+y

	x = y = 1
	x取上次赋值结果0
	y = x+y = 0+1 = 1


	i = 1
	x = y = 1
	x取上次赋值结果1
	y = x+y = 1+1 = 2



	i = 2
	x = y = 2
	x取上次赋值结果1
	y = x+y = 1+2=3

}*/