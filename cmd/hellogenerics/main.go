package main

import "fmt"

type Stringer interface {
	String() string
}

type intS int

func (i intS) String() string {
	return fmt.Sprintf("%d", i)
}

func main() {
	fmt.Println("hello generics")
	prnt([]int{1, 2, 3})
	s := []string{"A", "B", "C"}
	prnt[string](s)
	prnt(stringify([]intS{4, 5, 6}))
}

func prnt[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func stringify[T Stringer](s []T) []string {
	ret := []string{}
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}
