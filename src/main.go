package main

import (
   "fmt";
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
   const NUMBER_OF_RUNS_OF_WALKING_ROUTINE int = 9
   const NUMBER_OF_STARTS_OF_STOP_ROUTINE int = 5

   exit := make(chan string)

   defer close(exit)

   routine := func (name string, message string) {
      fmt.Println(message)
      exit <- fmt.Sprintf("from %s", name)
   }

   for i := 0; i < NUMBER_OF_RUNS_OF_WALKING_ROUTINE; i++ {
      go routine("first", "Go")
   }

   for i := 0; i < NUMBER_OF_STARTS_OF_STOP_ROUTINE; i++ {
      go routine("second", "Don't Go")
   }

   for i := 0; i < NUMBER_OF_RUNS_OF_WALKING_ROUTINE + NUMBER_OF_STARTS_OF_STOP_ROUTINE; i++ {
      fmt.Println("Get", <- exit)
   }

   fmt.Println("Stop")
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

   for y := 0; y < columnLength; y++ {
      table[y] = make([]int, columnLength)

      for x := 0; x < columnLength; x++ {
         table[y][x] = (x + start) * (y + start)
      }
   }

   return table
}

func IndexOf[objectT comparable](object objectT, set []objectT) (int, bool) {
   for itemIndex, item := range set {
      if object == item {
         return itemIndex, true
      }
   }

   return -1, false
}

func Lower(line string) string {
   newLine := make([]rune, len(line))

   for letterIndex, letter := range line {
      if 65 <= letter && letter < 91 {
         newLine[letterIndex] = letter + 32
      } else {
         newLine[letterIndex] = letter
      }
   }

   return string(newLine)
}

func PositionReportIn[objectT comparable](location []objectT, object objectT) string {
   index, exists := IndexOf(object, location)

   if exists {
      return fmt.Sprintf("Object %v is in array %v at index %b", object, location, index)
   } else {
      return fmt.Sprintf("Object %v isn't in array %v", object, location)
   }
}

func IsNumber(resource any) bool {
   switch resource.(type) {
   case int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
      return true
   default:
      return false
   }
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
   defer func() {
      if err := recover(); err != nil {
         fmt.Println("There was a panic that", err)
      }
   }()

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

   for _, column := range CreateMultiplicationTable(1, 11) {
      fmt.Println(column)
   }

   fmt.Println(Lower("THIS MESSAGE MUST NOT BE IN UPPERCASE"))

   fmt.Println(PositionReportIn(numbers, 4))
   fmt.Println(PositionReportIn(numbers, 5))

   fmt.Printf(
      "Are numbers?\n\t%d - %t\n\t%f - %t\n\t%s - %t\n",
      256, IsNumber(256),
      1.9, IsNumber(1.9),
      "this is a string", IsNumber("this is a string"),
   )
}