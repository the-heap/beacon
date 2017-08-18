/*
The clock interface allows us to use time.Time without using the concrete
type time.Time. In testing, we can mock clock so our tests return a
consistent time.
*/

package main

import "time"

type clock interface {
	Now() time.Time
}

type realClock struct{}

func (realClock) Now() time.Time { return time.Now() }
