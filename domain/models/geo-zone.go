package models

import "time"

type GeoZone struct {
	ID          int
	Name        string
	Description string
	UpdatedDate time.Time
	CreatedDate time.Time
}
