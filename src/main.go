package main

import ("fmt")

type User struct {
   name string
   age uint8
}

func (user *User) IsAdult() bool {
   return user.age >= 18
}

func (user *User) CreateComingOfAgeMessage() string {
   if user.IsAdult() {
      return "Adult"
   } else {
      return "Not adult"
   }
}

func (user *User) CreateResume() string {
   return fmt.Sprintf(
      "%s is %d years old. %s",
      user.name,
      user.age,
      user.CreateComingOfAgeMessage(),
   )
}




}

func PrintGo() {
   fmt.Println(string(byte(71)) + string(byte(111)))
}

func main() {
   user := User{"Max", 22}

   fmt.Println(user.CreateResume())
}