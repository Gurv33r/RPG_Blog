package routes

import "time"

func validate(input string) bool {
	const iso = "2006-01-02"
	_, err := time.Parse(iso, input)
	return err != nil
}
