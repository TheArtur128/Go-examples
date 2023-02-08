package main

import (
   "fmt";
   "time";
)

type Descriptive interface {
   CreateResume() string
}

type UserError struct {
   user *User
   err error
}

func (err UserError) Error() string {
   return err.err.Error()
}

func (err UserError) Unwrap() error {
   return err.err
}

type Age uint8

func (age *Age) IsAdult() bool {
   return *age >= 18
}

type User struct {
   name string
   age Age
   isAlive bool
}

func (user *User) Die() error {
   if user.isAlive {
      user.isAlive = false

      return nil
   } else {
      err := UserError{
         user,
         fmt.Errorf("an user at %p is already dead", user),
      }

      return err
   }
}

func (user *User) CreateResume() string {
   return fmt.Sprintf(
      "%s is %d years old. %s and %s",
      user.name,
      user.age,
      user.createComingOfAgeMessage(),
      user.createLivelinessMessage(),
   )
}

func (user *User) createComingOfAgeMessage() string {
   if user.age.IsAdult() {
      return "Adult"
   } else {
      return "Not adult"
   }
}

func (user *User) createLivelinessMessage() string {
   if user.isAlive {
      return "Alive"
   } else {
      return "Not alive"
   }
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

func UnderKey[keyT comparable, valueT any](key keyT, value valueT) map[keyT]valueT {
   return map[keyT]valueT{
      key: value,
   }
}

func Sum[elementT int | float32](elements ...elementT) elementT {
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

func Go() {
   defer fmt.Println("Stop")

   for i := 0; i < 9; i++ {
      go fmt.Println("Go")
   }

   for i := 0; i < 5; i++ {
      go fmt.Println("Don't Go")
   }

   fmt.Println("START OF SLEEP")
   time.Sleep(500 * time.Millisecond)
   fmt.Println("END OF SLEEP")
}

func ReferenceFor[resourceT any](resource resourceT) *resourceT {
   return &resource
}

func ChangeTo[valueT any](value valueT, reference *valueT) {
   *reference = value
}

func CreateMultiplicationTable(start int, end int) [][]int {
   columnLength := end - start

   table := make([][]int, columnLength)

   for yIndex := 0; yIndex < columnLength; yIndex++ {
      table[yIndex] = make([]int, columnLength)

      for xIndex := 0; xIndex < columnLength; xIndex++ {
         table[yIndex][xIndex] = (xIndex + columnLength) * (yIndex + columnLength)
      }
   }

   return table
}

var user User

var numbers []int
var numbersWithMultiplied []int

func init() {
   user = User{"Max", Age(22), true}

   numbers = []int{ 2, 4, 8 }
   numbersWithMultiplied = AddExponentiated(numbers)  
}

func main() {
   fmt.Println(user.CreateResume())

   fmt.Println(numbersWithMultiplied)

   fmt.Println(numbers)
   fmt.Println(numbersWithMultiplied[len(numbers):])

   fmt.Println(Sum(numbersWithMultiplied...))

   fmt.Println(UnderKey("alphabet", CreateAlphabet()))

   fmt.Println(Echo(42) + 22)

   Go()

   fmt.Println(ReferenceFor(256))

   number := 64

   fmt.Printf("Was %d, now ", number)

   ChangeTo(256, &number)

   fmt.Println(number)

   fmt.Println(*new(User))

   fmt.Println(CreateMultiplicationTable(1, 11))
}