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
