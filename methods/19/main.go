package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("xxx")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	// couldn't convert number: strconv.Atoi: parsing "xxx": invalid syntax
	fmt.Println("Conterted integer:", i)
}
