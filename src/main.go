package main

import ("fmt")

type User struct {
   name string
   age uint8
}

func (user *User) isAdult() bool {
   return user.age >= 18
}

func (user *User) createComingOfAgeMessage() string {
   if user.isAdult() {
      return "Adult"
   } else {
      return "Not adult"
   }
}

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