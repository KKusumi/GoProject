package main

type S string

func (s S) String() string {
	return string(s)
}
