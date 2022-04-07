package dt

import "time"

func GetTime(t string) (time.Time, error) {
	now := time.Now()
	d, err := time.ParseDuration(t)
	if err != nil {
		return now, err
	}
	return now.Add(d), nil
}
