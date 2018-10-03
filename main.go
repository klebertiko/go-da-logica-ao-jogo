// Packages are used to organise go source code for better reusability and readability
package main

// import blocks or import lines
import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"time"
)

// Structure
// A structure is a user defined type which represents a collection of fields.
// It can be used in places where it makes sense to group the data into a
// single unit rather than maintaining each of them as separate types

// If a struct type starts with a capital letter,
// then it is a exported type and it can be accessed from other packages
type Message struct {
	Text string
}

type EmailMessage struct {
	Email   string
	Message Message
}

// Method
// A method is just a function with a special receiver type
// that is written between the func keyword and the method name.
// The receiver can be either struct type or non struct type.
// The receiver is available for access inside the method.
func (m *Message) GetText() string {
	return m.Text
}

// Functions
// A function is a block of code that performs a specific task.
// A function takes a input, performs some calculations on the
// input and generates a output
func myPrint() {
	var value = "My custom print"
	fmt.Printf("==> %v", value)
}

// Exported names
// Any variable or function which starts with a capital
// letter are exported names in golang

// Exported function
func MyPrintExported() {
	var value = "My exported custom print"
	fmt.Printf("==> %v", value)
}

// Named return
func myPrintNamedReturn() (myString string) {
	var value = "My custom print"
	fmt.Printf("==> %v", value)
	return value
}

// Simple return
func mySum() int {
	var number1, number2 = 1, 2
	sum := number1 + number2
	fmt.Printf("SUM ==> %d", sum)
	return sum
}

// Multiple return
func myDiv() (int, int) {
	var number1, number2 = 1, 2
	div := number1 / number2
	mod := number1 % number2
	fmt.Printf("DIV ==> %d", div)
	return div, mod
}

// BEGIN
func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	n := make([]int, len(numbers))
	copy(n, numbers)

	pivotIndex := len(n) / 2
	pivot := n[pivotIndex]
	n = append(n[:pivotIndex], n[pivotIndex+1:]...)

	smaller, larger := partition(n, pivot)

	return append(append(quicksort(smaller), pivot), quicksort(larger)...)
}

func partition(numbers []int, pivot int) (smaller []int, larger []int) {
	for _, n := range numbers {
		if n <= pivot {
			smaller = append(smaller, n)
		} else {
			larger = append(larger, n)
		}
	}
	return smaller, larger
}

// END

func main() {
	// Variable is the name given to a memory location
	// to store a value of a specific type.
	// There are various syntaxes to declare variables in golang
	// variable declaration
	var number1 int

	// variable assignment
	number1 = 1

	// variable declaration with initial value
	var number2 int = 2 // can be inferred var number2 = 2

	fmt.Println(Message{fmt.Sprintf("Bem vindo programador nº%d", number1)})
	log.Println(Message{fmt.Sprintf("[LOG] Bem vindo programador nº%d", number2)})

	// The following are the basic types available in golang:
	// bool
	// A bool type represents a boolean and is either true or false
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	// - Numeric Types
	//   - Signed Integers
	//     - int8: represents 8 bit signed integers
	//     - size: 8 bits
	//     - range: -128 to 127

	//     - int16: represents 16 bit signed integers
	//     - size: 16 bits
	//     - range: -32768 to 32767

	//     - int32: represents 32 bit signed integers
	//     - size: 32 bits
	//     - range: -2147483648 to 2147483647

	//     - int64: represents 64 bit signed integers
	//     - size: 64 bits
	//     - range: -9223372036854775808 to 9223372036854775807

	//     - int: represents 32 or 64 bit integers depending on the underlying platform. You should generally be using int to represent integers unless there is a need to use a specific sized integer.
	//     - size: 32 bits in 32 bit systems and 64 bit in 64 bit systems.
	//     - range: -2147483648 to 2147483647 in 32 bit systems and -9223372036854775808 to 9223372036854775807 in 64 bit systems

	//   - Unsigned Inetegers
	//     - uint8: represents 8 bit unsigned integers
	//     - size: 8 bits
	//     - range: 0 to 255

	//     - uint16: represents 16 bit unsigned integers
	//     - size: 16 bits
	//     - range: 0 to 65535

	//     - uint32: represents 32 bit unsigned integers
	//     - size: 32 bits
	//     - range: 0 to 4294967295

	//     - uint64: represents 64 bit unsigned integers
	//     - size: 64 bits
	//     - range: 0 to 18446744073709551615

	//     - uint : represents 32 or 64 bit unsigned integers depending on the underlying platform.
	//     - size : 32 bits in 32 bit systems and 64 bits in 64 bit systems.
	//     - range : 0 to 4294967295 in 32 bit systems and 0 to 18446744073709551615 in 64 bit systems

	//   - Floating point
	//     - float32: 32 bit floating point numbers
	//     - float64: 64 bit floating point numbers

	//   - Complex types
	//     - complex64: complex numbers which have float32 real and imaginary parts
	//     - complex128: complex numbers with float64 real and imaginary parts

	//   - byte: is an alias of uint8

	//   - rune: is an alias of int32

	// string
	// Strings are a collection of bytes in golang
	text := "My Text"

	// Constants
	// The term constant is used in Go to denote fixed values
	const constant = 55 //allowed
	fmt.Printf("number constant value: %d, type: %T\n", constant, constant)
	const stringConstant = "Hello World" // string constant
	fmt.Printf("string constant value: %s, type: %T\n", stringConstant, stringConstant)
	// constant = 89       //reassignment not allowed

	// Blank identifier
	div, _ := myDiv()
	fmt.Println(div)

	// if
	// An if is a conditional statement, just true or false can be used with
	// an if instruction
	if true {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// loop
	// A loop statement is used to execute a block of code repeatedly
	// Golang has only one repeating structure, for,
	// which can be used in different ways according to the context
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i > 5 {
			break
		}
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// LIKE C WHILE
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	for i := range text {
		fmt.Println(i)
	}

	// defer
	// A defer statement defers the execution of a function until the surrounding function returns.
	defer fmt.Println("world")
	fmt.Println("hello")

	// switch
	// A switch is a conditional statement which evaluates an expression
	// and compares it against a list of possible matches and executes
	// blocks of code according to the match
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("No finger")
	}

	switch finger {
	case 1, 2, 3, 4, 5:
		fmt.Println("Finger")
	default:
		fmt.Println("No finger")
	}

	// Pointers
	// A pointer is a variable which stores the memory address of another variable.
	number3 := 3
	var number3Pointer *int = &number3
	fmt.Printf("Type of number3 is %T\n", number3Pointer)
	// & operator is used to get the address of a variable

	// Dereferencing a pointer means accessing the value of
	// the variable which the pointer points to
	number4 := &number3
	fmt.Println("The address of number3 is ", number4)
	fmt.Println("The value of number3 is ", *number4)

	// Arrays
	// An array is a collection of elements that belong to the same type.
	var intArray [3]int //int array with length 3
	intArray[0] = 12    // array index starts at 0
	intArray[1] = 78
	intArray[2] = 50
	fmt.Println(intArray)

	intArrayShort := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(intArrayShort)

	// The size of the array is a part of the type.
	// Hence [5]int and [25]int are distinct types.
	// Because of this, arrays cannot be resized.
	// Don't worry about this restriction since slices exist to overcome this.

	// Slices
	// A slice is a convenient, flexible and powerful wrapper on top of an array.
	// Slices do not own any data on their own. They are the just references to
	// existing arrays
	// The capacity of the slice is the number of elements in the underlying array
	// starting from the index from which the slice is created.
	array := [5]int{76, 77, 78, 79, 80}
	var slice []int = array[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(slice)

	// map
	type Vertex struct {
		Lat, Long float64
	}

	var m map[string]Vertex

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literals
	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	// map literals omitted
	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m)
}

// 	Random numbers in Go don’t seem random? Surprise! The rand package defaults to a seed of 1.
// That’s great if you need a bunch of random numbers at the start of your program. Not great if you expect a different outcome each time you run the program.
// A solution is to seed rand with Unix time. Try it in the init function:
func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	cmd := exec.Command("echo", "Go - Da logica ao jogo")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
