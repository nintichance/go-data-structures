package main

import (
	"errors"
)

type Stack struct {
	Storage map[uint]string
	Size    uint
}

func (s *Stack) IsEmpty() bool {
	return s.Size == 0
}

func (s *Stack) Push(element string) {
	if s.Size == 0 {
		s.Storage = make(map[uint]string)
	}
	s.Size += 1

	s.Storage[s.Size] = element
}

func (s *Stack) Pop() (string, error) {
	s.Size -= 1
	if s.Size == 0 {
		return "", errors.New("cannot pop an empty *stack")
	}
	popped := s.Storage[s.Size]
	delete(s.Storage, s.Size)
	return popped, nil
}

func (s *Stack) Peek() string {
	return s.Storage[s.Size]
}

func balanceParens(str string) bool {
	s := &Stack{}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" || string(str[i]) == "{" || string(str[i]) == "[" {
			s.Push(string(str[i]))
		} else if s.IsEmpty() {
			return false
		} else {
			if string(str[i]) == ")" && s.Peek() != "(" ||
				string(str[i]) == "}" && s.Peek() != "{" ||
				string(str[i]) == "]" && s.Peek() != "[" {
				return false
			} else {
				s.Pop()
			}
		}

	}

	return s.IsEmpty()
}
