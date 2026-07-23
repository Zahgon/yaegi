package main

import "time"

type TimeValue time.Time

func (v *TimeValue) decode() { _ = "STUB: not implemented"; return }

func main() {
	var tv TimeValue
	tv.decode()
}
