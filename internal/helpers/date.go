package helpers

import "time"

func InTime(end, check time.Time) bool {
	return check.Before(end)
}
