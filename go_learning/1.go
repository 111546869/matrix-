// package main

// import "fmt"

// type Animal interface {
// 	Speak() string
// 	Eat() string
// }

// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "woof!"
// }

// func (d Dog) Eat() string{
// 	return "dog eat!"
// }


// type Cat struct{}
// func (c Cat) Speak() string {
// 	return "meow!"
// }

// func (c Cat) Eat() string {
// 	return "cat eat!"
// }

// func main(){
// 	var a Animal
// 	a=Dog{}
// 	fmt.Println(a.Speak())
// 	fmt.Println(a.Eat())
// 	a=Cat{}
// 	fmt.Println(a.Speak())
// 	fmt.Println(a.Eat())
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	for i:=0;i<10;i++ {
// 		defer fmt.Println(i)
// 		for j:=i+1;j<10;j++ {
// 			defer fmt.Println(j)
// 		}
// 	}
// 	fmt.Println("done")
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch :=make(chan int)

// 	go func() {
// 		ch<- 42
// 	}()

// 	val:=<-ch
// 	fmt.Println(val)
// }

// package main

// import(
// 	"fmt"
// )

// func squareWorker(nums []int,ch chan int) {
// 	sum:=0
// 	for _,num :=range nums {
// 		sum+=num*num
// 	}
// 	ch<- sum
// }

// func main(){
// 	nums:=[]int{1,2,3,4,5}
// 	ch :=make(chan int)

// 	go squareWorker(nums[:len(nums)/2],ch)
// 	go squareWorker(nums[len(nums)/2:],ch)

// 	sum1:=<-ch
// 	sum2:=<-ch
// 	total:=sum1+sum2

// 	fmt.Printf("total sum = %d\n",total)
// }

// package main

// import (
//     "fmt"
// )

// func sayHello() {
//     fmt.Println("Hello!")
// }

// func main() {
//     go sayHello() // 使用 go 关键字启动一个新的 goroutine
// 	go sayHello()
// 	go sayHello()
// 	go sayHello()
// 	go sayHello()
// 	go sayHello()
// 	go sayHello()
// 	fmt.Println("hello!")
//     // time.Sleep(1 * time.Second) // 等待 1 秒钟，以便 goroutine 有机会执行
// }


package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from ch1"
    }()

    go func() {
        time.Sleep(1 * time.Second)
        ch2 <- "Message from ch2"
    }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
