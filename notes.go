

//Slice = list
//map = dict

## Data Types and Interfaces/Contracts

//You will see some simalarities to C here, we declare a variable along with its type and contents (1)
// we then point to the location of our var in memory (2)
//next we must derefence the address from whichever registry is being used so that we can call a function
//now we reassign ptr to "100"
var count = int(42) 
ptr := &count
fmt.Println(*ptr) x*ptr = 100
fmt.Println(count)


//here we create a new data type for further use down the track, we say that the name is Person, with two fields(name and age)
//next we define a method attached to the type called SayHello which will print a defined string aswell as whatever is in the persons Name field

type Person struct { Name string
	Age int }
	func (p *Person) SayHello() { fmt.Println("Hello,", p.Name)
	   }
	   func main() {
	var guy = new(Person) {guy.Name = "Dave"
	| guy.SayHello()
	}


//an interface is a way of a type fulfilling a contract to be considered another type EG:
type Friend interface { SayHello()
}
//as above if say, a newly created "person" or "dog" type was able to say hello (does call the SayHello method), the would then be considered a friend and can be referenced as such in future code, let's look at this further.
//first we create a function that says if the type calling greet meets being a being a friend then go ahead with saying hello
func Greet(f Friend) { 
	f.SayHello()
}
//remember that new data type we create with struct and the method we implemented in it earlier? well now let's create a variable called guy which will be a new TYPE of PERSON with his name field as dave before we GREET him.

func main() {
	var guy = new(Person) guy.Name = "Dave" Greet(guy)
	}

//now notice that we are able to greet, as the datatype person is able to SayHello because we coded this as a method when we declared the type

//but we don't JUST want to say hello to people, we like dogs too!, let's create another datatype and implement a SayHello method:

type Dog struct {}
func (d *Dog) SayHello() {
	fmt.Println("Woof woof") 
}
func main() {
	var guy = new(Person) guy.Name = "Dave"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}

//closing  off interfaces and structs for now I will write a program utilising types and lists aswell as user input to demonstrate knowledge.package golang

## Loops, conditionals

//loop example, using boolean true, looks cleaner than bash ans JS, similar to python
if x == 1 {
	fmt.Println("X is equal to 1")
	} else {
	fmt.Println("X is not equal to 1")
	}

//here is what i would call an elif conditional equivalent, again, looks clean, the check is initiated by "switch" and "elif" appears to be replaced with "case" and finally "else" with default (side note, go infers x to be a string) - it was worth noting that this is handled differently to elif, as switch is actually taking these cases and COMPARING them against x.

switch x{
	case "foo":
		fmt.Println("Found foo") 
	case "bar":
		fmt.Println("Found bar") 
	default:
		fmt.Println("Default case")
}

//a variant on our friend mr switch here is a "type switch", in whch we check for the type of case, not the contents. The below example is attempting to retrieve the type of interface. The switch syntax is v := i.(type), meaning we declare v as being equal to i and it's type. the cases then get compared to switch and handle accordingly on a match or no match

func foo(i interface{}) { 
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!") 
	case string:
		fmt.Println("I'm a string!") 
	default:
		fmt.Println("Unknown type!") 
	}
}

//for loop, same logica as seen in bash and js, explained [HERE](<link to explaination)

for i := 0; i < 10; i++ { 
	fmt.Println(i)
}

//a more complex forloop example looping over a list and dict, or as go calls them, a slice and a map.
//a list called nums is declared, with the contents being 4 integers
//we then use range to iterate over our list, for each index and it's corresponding value, they will be printed.
//so, the range of our list is 4, so the output would be as follows: 1:2, 3:4, 6:7, 8:9

nums := []int{2,4,6,8}
for idx, val:= range nums {
fmt.Println(idx, val) 
}

## Goroutines (pseudo threads)

//one of the things that makes go powerful - concurrency through "goroutines".
//"goroutines" can be functions or methods able to run simultaneously (sometimes described as lightweight threads as the cost of creating them is minimal compared to actual threads)
//example below, when we jump into main notice we prface our dummy function with go, meaning that main() will CONTINUE WITHOUT WAITING FOR F TO COMPLETE

func f() {
	fmt.Println("f function")
}

func main() { 
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function") 
}

//go contains a data type called channels that provide a mechanism through which goroutines can synchronise their execution and communicate with one another.
//below we define a var c of type chan int. We choose int as the type for channel because we will be passing through the lengths of various strings between goroutines.
//we declare whether data is flowing to or from a channel with <-
//our function accepts a string as an argument aswell as a channel, it then places the length of the string into the channel.
//the main function pieces everything together, first a call to make is made the create the integer channel. Then multiple concurrent calls to to strlen() are made via go which will spin up multiple goroutines. strlen() is passed two string values as well as the channel to which we will send the results.
//last, we read data from the channel via <-, with data flowing from our channel. This takes data out of the channel and assigning them to variables. At this point execution is blocked until good data can be read from the channel.
//as each line completes the length of each string and their sum will be sent to stdout.

func strlen(s string, c chan int) {
	c <- len(s)
}

func main() {
	c := make(chan int)
	go strlen("Salutations", c)
	go strlen("World", c) 
	x, y := <-c, <-c
	fmt.Println(x, y, x+y) 
}

//^^^ while the above may seem like black-magic, concurrency and parallelism can get quite complicated. The rest of the book Will introduce buffered channels, wait groups, mutexes in addition to others. If the book does not dive deep enough, I will create a seperate section to explain these after some research.

## Error Handling

//Go does not include syntax for try/catch/finally
//Go instead uses a minimalistic approach that encourages the coder to check for errors as they occur rather than poison the code chain.
//below is go's built in error type


type error interface {
	Error() string
}

//what this means is that you can use any data type that implements a method named Error(), which returns a string value as an error, below is a custom error

type MyError string
func (e MyError) Error() string {
	return string(e)
}

//we implement this by creating a user-defined string type named MyError and implement an Error() string method for the type.
//the below main function is calling our function, then checking if the error is equal to nil, if not this indicates something has indeed got wrong and thus we handle the error further.

func foo() error{
	return errors.New("Some Error Happen Here :O")
}
func main () {
	if err := foo();err != nill {
		//Handle the error
	}

}

//NOTE regarding errors in GO: most functions and methods will return at least one value for errors, which is generally nil, indicating there is actually no error.

## Serialization 

//the below code defines a struct type named foo which is intialised in our main() function after which a call is made to json.Marshal() with foo passed into it.
//Marshal() will encode the struct to JSON, returning a byte slice which is then printed to stdout.
//once the below completes, the JSON-encoded string will look as follows: 
//`{"Bar":"Joe Junior","Baz":"Hello Shabado"}`

type Foo struct {
	Bar string
	Bar string
}

func main() {
	f:= Foo{"Joe Junior", "Hello Shabado"}
	b, _ := json.Marshal(f)
	fmt.PrintIn(string(b))
	json.Unmarshal(b,&f)
}

### field tags

//the next part of the above code will take the byte slice and decode this time via a call to json.Unmarshal(b, &f) which will produce an instance of foo struct. We can do the same for XML in a similar manner, with the exception that we use field tags.
//notice that field tags are wrapped in backticks, begin with the tag name followed by a colon and the directive enclossed in double quotes
//directives define how the fields get handled. In the example below we declare that Bar be treated as ann attribute named id - whereas normally it would be treated as an element.
//we also declare that Baz will be found in a subelement of parent, named child.
//should we modify the previous JSON encoding program to encode the structure as xml the result would now look as follows:
//`<Foo id="Joe Junior"><parent><child>Hello Shabado</child></parent></Foo>`

type Foo struct {
	Bar string `xml:"id,attr"`
	Baz string `xml:"parent>child"`
}

//the XML encodere determines the names of elements using the tag directives which means we have greater control over each field.
//we will be using field tags throughout this book for dealing with other data serialisation formats, including ASN.1 and MessagePack.
//We will eventually define our own custom tags, specifically when we learn to handle (server message block) SMB

# Chapter 2 - TCP stuff

//refresher on shaking hands
//if a port is open the threeway handshake takes place as we know. Client sends a SYN packet (extends their hand), the server responds with a SYN-ACK (extends their own hand) finally the client sends an ACK (grips the hand and shakes)
//if however the port i closed, the client sends its SYN but is however met instead with RST from the host
//if traffic is being filtered by a firewall, the client will typically receive no response after extending it's hand
//NOTE the further we go into making tools, the more we need to master our TCP knowledge

## Post Scanner
//a good way to go deeper into concurrency is by building a port scanner, as we abviously cannot scan all 65,535 ports using a single contiguous method.
//we will be connecting to scanme.nmap.org using Go's "net" package: "net.dial(network, address string)"
//the first arg is a string that identifies the kind of connection to initiate (which by the way, this can even be your own custom protocol)
//the second argument tells dial who the target is (notice that the arg takes a single string, not a string and an int)

//we call upon our knowledge of errors gained earlier in order to detremine of our connection has been succesfull.

package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.PrintIn("Connection successful")
	}
}

//now lets scan multiple ports using a for loop

for 1:=1; i <= 1024; i++ {
	
}
