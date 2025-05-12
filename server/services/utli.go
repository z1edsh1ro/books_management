package services

import "time"

func ConvertDate(data string) (time.Time, error) {
	date, err := time.Parse(time.RFC3339, data)

	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}
