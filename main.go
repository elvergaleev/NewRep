Package main

func main () {
import "fmt"

first:=2
second:=45

fmt.Println (calc(first, second, sumFunc)

sumFunc:=func (x, y int) int {return x+y)
}

func calc (x, y, action func (x, y int) int {return action (x, y)}

//user2 changing
