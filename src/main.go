package main

import ("fmt")

type Descriptive interface {
   CreateResume() string
}

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

func CreateAlphabet() [26]string {
   var alphabet [26]string

   for letterIndex := 0; letterIndex < 26; letterIndex++ {
      alphabet[letterIndex] = string(letterIndex + 65)
   }

   return alphabet
}

func AddExponentiated(numbers []int) []int {
   for _, number := range numbers {
      numbers = append(numbers, number * number)
   }

   return numbers
}

func underKey[keyT comparable, valueT any](key keyT, value valueT) map[keyT]valueT {
   return map[keyT]valueT{
      key: value,
   }
}

func sum[elementT int](elements ...elementT) elementT {
   var sum_of_elements elementT

   for _, element := range elements {
      sum_of_elements += element
   }

   return sum_of_elements
}

func PrintGo() {
   fmt.Println(string(byte(71)) + string(byte(111)))
}


func Echo[resourceT any](resource resourceT) resourceT {
   defer fmt.Println(resource)
   return resource
}


var user User

func init() {
   user = User{"Max", 22}
}

func main() {
   fmt.Println(user.CreateResume())

   fmt.Println(underKey("alphabet", CreateAlphabet()))
}