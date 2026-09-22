package models

import "time"

type Investment struct {
	ID             int64
	Name           string
	FinishDate     time.Time
	OfficeLocation string
	Phone          string
}
