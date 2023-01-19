package service

import (
    "time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int32 {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := int32(time.Until(d).Hours())

    return dur/24

}
