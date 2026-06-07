package main

import "testing"

func TestPushandPop(t *testing.T){
	s:=NewStack[int]()

	// Stack should be empty initially
	if !s.IsEmpty(){
		 t.Errorf("Expected new stack should be empty but IsEmpty() return false")
	}
	// Push elements and check size
	s.Push(10)
	s.Push(20)
	if s.Size()!= 2{
		 t.Errorf("Expected size to be 2 after pushing twice, got %d",s.Size())
	}

	// Pop element and check boolean
	_,ok:=s.Pop()
	if !ok{
		t.Errorf("Expected Pop to be succeed, but got ok = false")
	}
	// Check is actual element is removed
	beforeSize:=s.Size()
	_,ok =s.Pop()
	if !ok{
		t.Errorf("Expected Pop to be succeed, but got ok = false")
	}
	if beforeSize-1 != s.Size(){
		t.Errorf("Expected size to reduce by 1, but got %d",s.Size())
	} 
}

func TestStack_EmptyEdgeCases(t *testing.T){
	// Define the test table structure
	tests:=[]struct{
		name string
		action func(s *Stack[string])(string,bool)
		wantVal  string
		wantSuccess bool
	}{
		{
			name: "Pop on empty stack",
			action: func(s *Stack[string])(string,bool){
				return s.Pop()
			},
			wantVal: "",
			wantSuccess: false,
		},
		{
		name: "Peek on empty stack",
		action: func (s *Stack[string])(string,bool)  {
			return s.Peek()
		},
		wantVal: "",
		wantSuccess: false,
		},
	}
	// Run the table
	for _,tt:=range tests{
		t.Run(tt.name, func(t *testing.T) {
			s:=NewStack[string]()
			// s.Push("Uday") // To check the failed cases
			gotVal,gotSuccess:=tt.action(s)
			if gotSuccess!=tt.wantSuccess{
				t.Errorf("Success status= %v, want %v",gotSuccess,tt.wantSuccess)
			}
			if gotVal !=tt.wantVal{
				t.Errorf("Return value %q, want %v",gotVal,tt.wantVal)
			}
		})
	}
}