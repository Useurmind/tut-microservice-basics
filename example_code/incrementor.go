package main

import "fmt"

// FIELDS
type Incrementor struct {
	// PRIVATE because its lowercase
	value int

	// PUBLIC because its uppercase
	UpperLimit int
}

// CONSTRUCTOR
func NewIncrementor(startValue int) *Incrementor {
	return &Incrementor{
		value: startValue,
		UpperLimit: 10,
	}
}

// MEMBER
func (i *Incrementor) Increment() (int, error) {
	if i.value >= i.UpperLimit {
		return i.value, fmt.Errorf("Value has reached the upper limit")
	}

	i.value++

	return i.value, nil
}

func main() {
	inc := NewIncrementor(0)
	inc.UpperLimit = 3

	v, err := inc.Increment()  
	fmt.Printf("v = %d, err = %v\n", v, err)
	// v = 1, err = <nil>
	v, err = inc.Increment()  
	fmt.Printf("v = %d, err = %v\n", v, err)
	// v = 2, err = <nil>
	v, err = inc.Increment()  
	fmt.Printf("v = %d, err = %v\n", v, err)
	// v = 3, err = <nil>
	v, err = inc.Increment()  
	fmt.Printf("v = %d, err = %v\n", v, err)
	// v = 3, err = "Value has reached the upper limit"	
}