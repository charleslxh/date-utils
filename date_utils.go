package date_utils

import "time"

func GetBeginOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	return time.Date(y, 1, 1, 0, 0, 0, 0, t.Location())
}

func GetBeginOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
}

func GetBeginOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func GetBeginOfHour(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, t.Hour(), 0, 0, 0, t.Location())
}

func GetLastDayOfYear(t time.Time) time.Time {
	y, _, _ := t.Date()
	t = time.Date(y+1, 1, 1, 0, 0, 0, 0, t.Location())
	return t.AddDate(0, 0, -1)
}

func GetLastDayOfMonth(t time.Time) time.Time {
	return GetBeginOfMonth(t).AddDate(0, 1, -1)
}

func GetTomorrow() time.Time {
	return GetBeginOfDay(time.Now()).AddDate(0, 0, 1)
}

func GetYesterday() time.Time {
	return GetBeginOfDay(time.Now()).AddDate(0, 0, -1)
}

func GeDayAfterTomorrow() time.Time {
	return GetTomorrow().AddDate(0, 0, 1)
}

func GeDayBeforeYesterday() time.Time {
	return GetYesterday().AddDate(0, 0, -1)
}

func AddHour(t time.Time, hours int) time.Time {
	return t.Add(time.Duration(hours) * time.Hour)
}

func AddMinutes(t time.Time, minutes int) time.Time {
	return t.Add(time.Duration(minutes) * time.Minute)
}

func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func ParseExcelNumber(days int) time.Time {
	start := time.Date(1899, 12, 30, 0, 0, 0, 0, time.Now().Location())
	return start.Add(time.Duration(days * 24) * time.Hour)
}

func GetDaysBetween(t1 time.Time, t2 time.Time) int {
	duration := int64(0)
	timestamp1 := GetBeginOfDay(t1).Unix()
	timestamp2 := GetBeginOfDay(t2).Unix()

	if timestamp1 > timestamp2 {
		duration = timestamp1 - timestamp2
	} else {
		duration = timestamp2 - timestamp1
	}

	return int((duration / 86400) | 0)
}