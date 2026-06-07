package main

import (
	"fmt"
)

/* Stack
A stack is a LIFO(last in first out). Think of a stack of dinner plates.
the last plate you put on top on the pile is the first plate you must take off.
*/

type Stack[T any] struct{
	Elements []T
}
func NewStack[T any]() *Stack[T]{
	return &Stack[T]{Elements: make([]T,0)}
}
func (s *Stack[T])Push(val T){
	s.Elements=append(s.Elements, val)
}
func (s *Stack[T])Pop()(T,bool){
	var zero T
	if s.IsEmpty(){
		return zero,false
	}
	val:=s.Elements[len(s.Elements)-1]
	s.Elements=s.Elements[:len(s.Elements)-1]
	return val,true

}
func (s *Stack[T])Peek()(T,bool){
	var zero T
	if s.IsEmpty(){
		return zero,false
	}
	return s.Elements[len(s.Elements)-1],true
}
func (s *Stack[T])IsEmpty()bool{
	if len(s.Elements)==0{
		return true
	}
	return false
}
func (s *Stack[T])Size()int{
	return  len(s.Elements)
}
func (s *Stack[T])Display(){
	fmt.Println(s.Elements)
}
func main(){
	fmt.Println("--- Integer Stack ---")
	intStack:=NewStack[int]()
	intStack.Push(10)
	intStack.Push(20)
	intStack.Display()
	fmt.Println("Stack Size:",intStack.Size())


	fmt.Println("---String Stack---")
	stringStack:=NewStack[string]()
	stringStack.Push("abc")
	stringStack.Push("xyz")
	stringStack.Display()
	fmt.Println("Stack Size:",stringStack.Size())
}