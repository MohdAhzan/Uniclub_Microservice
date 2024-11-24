package utils

import "time"

func StringToTime(timeStr string) (time.Time, error) {
	layout := time.RFC3339
	parsedTime, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func  TimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}




