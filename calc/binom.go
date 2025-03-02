package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	if k == n {
		return 1
	}
	if n == 0 || k == 0 {
		return 1
	}
	return BinomialCoefficient(n-1, k-1) + BinomialCoefficient(n-1, k)
}
