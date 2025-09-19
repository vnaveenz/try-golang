package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filRecords []Record
	for _, r := range in {
		if (predicate(r)) {
			filRecords = append(filRecords, r)
		}
	}
	return filRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	From := p.From
	To := p.To
	return func(r Record) bool {
		if (r.Day >= From && r.Day <= To) {
			return true
		} else {
			return false
		}
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		} else {
			return false
		}
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	expense := 0.0
	wthnRange := ByDaysPeriod(p)
	for _,r := range in {
		if wthnRange(r) {
			expense += r.Amount
		}
	}
	return expense
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryFound := false
	expensePerCategory := 0.0
	for _,r := range in {
		if (r.Category == c) {
			categoryFound = true
			break
		}
	}
	if !categoryFound {
		return 0, errors.New("unknown category")
	} else {
		filtrdRecs := Filter(in, ByCategory(c))
		expensePerCategory = TotalByPeriod(filtrdRecs, p)
		return expensePerCategory, nil
	}
}
