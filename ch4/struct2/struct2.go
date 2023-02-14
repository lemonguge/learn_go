package main

import (
	"fmt"
	"time"

	"lemonguge.cn/learn_go/ch4/treesort"
)

// 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员。
type Employee struct {
	// 如果结构体成员名字是以大写字母开头的，那么该成员就是导出的，一个结构体可能同时包含导出和未导出的成员。
	ID      int
	Name    string
	Address string
	// 如果相邻的成员类型相同，则可以被合并到一行
	// Name, Address string
	DoB               time.Time
	Position          string
	Salary, ManagerID int
	// 可以使用顺序赋值来初始化结构体的成员，因此结构体成员的顺序调整可能导致编译不通过。

	// 一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。
	// director Employee // compile error: illegal cycle in declaration of Employee
	// 但是 S 类型的结构体可以包含 *S 指针类型的成员
	director *Employee
}

type point struct{ X, Y int }

// 结构体对应的零值是每个成员对应该类型的零值。
var dilbert Employee

func main() {
	fmt.Println(dilbert.ID)
	// dilbert 结构体变量的成员可以通过点操作符访问
	dilbert.Salary = 2000
	fmt.Println(dilbert.Salary)
	dilbert.Salary += 5000
	fmt.Println(dilbert.Salary)

	dilbert.Position = "Engineer"
	// 对成员取地址，然后通过指针访问
	pp := &dilbert.Position
	*pp = "Senior " + *pp
	fmt.Println(dilbert.Position)

	// 本月最佳雇员
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 语法糖，等价于
	// (*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(dilbert.Position)

	fmt.Println(employeeOfTheWeek().Position)
	// 如果将 employeeOfTheWeek 函数的返回值从 *Employee 指针类型改为 Employee 值类型，那么更新语句将不能编译通过。
	// 因为在赋值语句的左边要求是一个变量（调用函数返回的是值，并不是一个可寻址的变量）。
	employeeOfTheWeek2().Salary += 5000
	// employeeOfTheWeek().Salary += 5000 // compile error: cannot assign

	dilbert.director = employeeOfTheWeek2()

	is := []int{4, 1, 5, 2, 3}
	treesort.Sort(is)
	fmt.Println(is)

	// 以结构体成员定义的顺序为每个结构体成员指定一个字面值
	p := point{1, 2}
	fmt.Println(p)

	// 结构体可以作为函数的参数和返回值，如果考虑效率的话，较大的结构体通常会用指针的方式传入和返回。

	// 创建并初始化一个结构体变量，并返回结构体的地址
	p2 := &point{2, 1}
	// 等价于
	// p2 := new(point)
	// *p2 = point{1, 2}
	fmt.Println(*p2)

	// 如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用 == 或 != 运算符进行比较。
	// 下面两个比较的表达式是等价的
	fmt.Println(p.X == p2.X && p.Y == p2.Y) // false
	fmt.Println(p == *p2)                   // false
}

func employeeOfTheWeek() Employee {
	// 以成员名字和相应的值来初始化，可以包含部分或全部的成员
	return Employee{ID: 999, Name: "Bob", Position: "Project Manager"}
}

func employeeOfTheWeek2() *Employee {
	return &Employee{ID: 999, Name: "Bob", Position: "Project Manager"}
}

func AwardAnnualRaise(e *Employee) {
	// 如果要在函数内部修改结构体成员的话，用指针传入是必须的；
	// 因为在 Go 语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
	e.Salary = e.Salary * 105 / 100
}
