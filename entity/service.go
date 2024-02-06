package entity

import "time"

type HoursService struct {
	Hour map[time.Time]time.Time
}
type Service struct {
	Name         string            `json:"name"`
	Hours        HoursService      `json:"hours"`
	Description  string            `json:"description"`
	Images       map[string]string `json:"images"`
	Professional string            `json:"professional"`
	CreatedTime  time.Time         `json:"created_time"`
}
