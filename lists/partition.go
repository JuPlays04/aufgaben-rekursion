package lists

// Liefert zwei Listen:
// - Eine, die alle Elemente aus list enthält, die kleiner oder gleich key sind.
// - Eine, die alle übrigen Elemente aus list enthält.
func Partition(list []int, key int) ([]int, []int) {
	// Verwende Kopien von list, damit die ursprüngliche Liste nicht verändert wird.
	l1 := append([]int{}, list...)
	l2 := append([]int{}, list...)

	x := FilterLess(l1, key)
	y := FilterGreater(l2, key)

	return x, y
}
