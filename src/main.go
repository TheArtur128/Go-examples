package main

import ("fmt")

func getNumericByte() int {
   var first int = 100
   var second int
   third := 0

   second = 50
   second += 6

   third = 100

   return first + second + third
}

func printGo() {
   fmt.Println(string(byte(71)) + string(byte(111)))
}

func main() {
}