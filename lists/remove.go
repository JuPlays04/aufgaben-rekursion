package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {

	// removeAt entfernt das Element an der gegebenen Position rekursiv.
	// Basisfall: Leere Liste oder ungültige Position → Original-Liste zurückgeben
	if Empty(list) || pos < 0 {
		return list
	}

	// Rekursiver Fall: Erstes Element entfernen, wenn pos == 0
	if pos == 0 {
		return list[1:]
	}

	// Rekursiv den Rest der Liste verarbeiten
	return append([]int{list[0]}, RemoveElement(list[1:], pos-1)...)
}
