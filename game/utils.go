package game

func GetKeys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func ConvertLetterToRow(letter string) int {
	var row int

	switch letter[0] {
	case 'A':
		row = 0
	case 'B':
		row = 1
	case 'C':
		row = 2
	case 'D':
		row = 3
	case 'E':
		row = 4
	case 'F':
		row = 5
	case 'G':
		row = 6
	case 'H':
		row = 7
	case 'I':
		row = 8
	case 'J':
		row = 9
	default:
		return -1
	}

	return row
}
