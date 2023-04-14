package utils

import "DataMiner/internal/models"

func GetDayInString(weekday models.Weekday) string {
	switch weekday {
	case models.Monday:
		return "Monday"
	case models.Tuesday:
		return "Tuesday"
	case models.Wednesday:
		return "Wednesday"
	case models.Thursday:
		return "Thursday"
	case models.Friday:
		return "Friday"
	case models.Saturday:
		return "Saturday"
	case models.Sunday:
		return "Sunday"
	default:
		return ""
	}
}
