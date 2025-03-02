package search

import "fmt"

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {
	left, right := 0, len(list)-1
	result := -1

	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == x {
			result = mid
			right = mid - 1 // Suche weiter links nach der ersten Instanz
		} else if list[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	list := []int{1, 2, 2, 2, 3, 4, 5}
	x := 2
	fmt.Println(FindSorted(list, x)) // Erwartete Ausgabe: 1 (Index der ersten 2)
}
