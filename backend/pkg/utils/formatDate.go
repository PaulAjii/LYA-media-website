package utils

import "time"

func FormatDate(dateStr string) (string, error) {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return "", err
	}
	return t.Format("January 02, 2006"), nil
}
