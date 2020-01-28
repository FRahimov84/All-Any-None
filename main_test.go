package main

import "fmt"

var data = []account{
	{id: 1, name: "Vasya", balance: 1_000},
	{id: 2, name: "Petya", balance: 2_000},
	{id: 3, name: "Kostya", balance: 3_000},
	{id: 4, name: "Oleg", balance: 10_000},
	{id: 5, name: "Pavel", balance: 12_000},
	{id: 6, name: "Katya", balance: 3_000},
}

func Example_AllForTrue() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance >= 1_000
	}))
	// Output: true
}
func Example_AllForFalse() {
	fmt.Println(all(data, func(item account) bool {
		return item.balance >= 1_500
	}))
	// Output: false
}

func Example_AnyForTrue() {
	fmt.Println(any(data, func(item account) bool {
		return item.balance >= 10_000
	}))
	// Output: true
}
func Example_AnyForFalse() {
	fmt.Println(any(data, func(item account) bool {
		return item.balance > 12_000
	}))
	// Output: false
}

func Example_NoneForTrue() {
	fmt.Println(none(data, func(item account) bool {
		return item.balance == 3_001
	}))
	// Output: true
}
func Example_NoneForFalse() {
	fmt.Println(none(data, func(item account) bool {
		return item.balance != 3_000
	}))
	// Output: false
}

func Example_IndexWithResult() {
	fmt.Println(index(data, func(item account) bool {
		return item.balance > 3_000
	},false))
	// Output: 3
}
func Example_IndexWithoutResult() {
	fmt.Println(index(data, func(item account) bool {
		return item.balance <=300
	},false))
	// Output: -1
}

func Example_FindWithResult()  {
	result, ok := find(data, func(item account) bool {
		return item.balance == 3_000
	})
	if ok !=nil {
		fmt.Errorf(ok.Error())
	}else {
		fmt.Println(result)
	}
	// Output: {3 Kostya 3000}
}
func Example_FindWithoutResult()  {
	result, ok := find(data, func(item account) bool {
		return item.name == "Oleg 1"
	})
	if ok !=nil {
		fmt.Println(ok.Error())
	}else {
		fmt.Println(result)
	}
	// Output: not found
}