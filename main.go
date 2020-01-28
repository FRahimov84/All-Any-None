package main

type account struct {
	id      int64
	name    string
	balance int64
}
type accountError struct {
	message string
}

func (a accountError) Error() string {
	return a.message
}

func all(items []account, predicate func(item account) bool) bool {
	return index(items, predicate, true) == -1
}

func any(items []account, predicate func(item account) bool) bool {
	return index(items, predicate, false) != -1
}

func none(items []account, predicate func(item account) bool) bool {
	return index(items, predicate, false) == -1
}

func index(items []account, predicate func(item account) bool, reverse bool) int {
	for index, value := range items {
		if reverse {
			if !predicate(value) {
				return index
			}
		} else {
			if predicate(value) {
				return index
			}
		}
	}
	return -1
}

func find(items []account, predicate func(item account) bool) (account, error) {
	result := index(items, predicate, false)
	if result == -1 {
		return account{}, accountError{message: "not found"}
	}
	return items[result], nil
}

func main() {}
