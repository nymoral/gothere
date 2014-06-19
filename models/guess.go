package models

import (
	"time"
)

type Guess struct {
	Pk int

	Gamepk int
	Userpk int

	Result1 int
	Result2 int

	Points int
	Total  int

	Given time.Time
}
