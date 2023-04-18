# Golang Tutorial from Udemy

Most important things:

- Pointers - `&` and `*`, see more in structures/main.go
  - Turn address into value with \*address
  - Turn value into address with &value
- is not OOP
- Structures

## Pointers

- Whenever you pass an integer, float, string, or struct into a function, it creates a copy of that value and these copies are used inside of the function.
- &variable - gives you the memory address of the value this variable is pointing at
- \*pointer - gives you the value this memory address is pointing at
- After the line func (lo *location) someFunction(), the *location specifies the type of the receiver that the function expects.
- If put name:="BILL" and print(\*&name), it will print BILL
- Reference vs Value Types. Reference means that the value is stored somewhere else in memory and the variable is just a pointer to that value aka dont need to worry about pointers. Value means that the value is stored in the variable itself, need to worry about pointers.
  - Value Types: int, float, string, bool, structs
  - Reference Types: slices, maps, channels, pointers, functions
- When we create a slice, go will auto create 2 data structure, an array and a struct that records the length and capacity of the slice plus a reference to the underlying array
- Golang everything is pass by value.

## Maps

Maps are similar to dict (python) or object (JS). Map key and value must be of the same type.

Maps Vs Structures

Maps (Reference Type):

- All keys must be of the same type
- All values must be of the same type
- Keys are indexed
- Use to represent a collection of related properties
- Don't need to know all the keys at compile time

Structures (Value Type):

- Values can be of different types
- Keys are not indexed
- Need to know all the keys at compile time
- Use to represent a thing with alot of different properties

## Interfaces

Interfaces makes it easier for us to reuse code between different parts of our code base. E.g:
Shuffling a deck of cards, can only take in type deck: func (d deck) shuffle() but if we use interface, we can use any type that has shuffle method. Read more inside the interfaces folder.

- Concrete Type

  - map/struct/int/string/customType

- Interface Type
  - type interface

Interface are not generic types, need to be implicit, a contract to help us manage type.

- Interfaces can have nested interfaces

Another example:
Reading HTTP Request body have a function printHTTP([]customDataType), reading file on hard drive printFile([]string) etc but they all have different data type BUT SAME FUNCTION. The solution is a reader interface. No matter what source of input, put a reader interface infront of that and output a data that anyone can work with like a byte slice []byte.

## Concurrency

Basically enable concurrency by using goroutines and channels. Go routine that exist in our running program, this can be used to run code concurrently. Many go routine to one go scheduler to one cpu core. Scheduler runs ONE routine until it finishes or makes a blocking call then it will run another routine. By default Go scheduler only use one CPU core but you can specify how many cores you want to use.

Concurrency - We can have multiple threads executing code. If one thread blocks another one is picked up and worked on.
Parallelism - We can do multiple things at the exact same nanosecond
Child go routine - Child routines created by the go keyword

Channels

- Channel will be used to communicate between go routines. Channels input type need to be of the same type.
- Sending Data with Channels
  - channel <- 42 // send data to channel
  - myNumber <- channel // wait for data to be sent to channel and then assign it to myNumber
  - fmt.Println(<-channel) // receive data from channel and print
