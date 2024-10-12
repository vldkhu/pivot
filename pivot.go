package pivot

// Record представляет собой структуру данных с категорией и значением.
type Record struct {
	Category string
	Value    int
}

// Pivot принимает срез записей и возвращает карту, сгруппированную по категориям.
func Pivot(records []Record) map[string][]int {
	pivoted := make(map[string][]int)

	for _, record := range records {
		pivoted[record.Category] = append(pivoted[record.Category], record.Value)
	}

	return pivoted
}
