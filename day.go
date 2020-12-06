package main

import "./solutions"

type Day interface {
	Handle(string) ([]string, error)
}

func getDays() map[int]Day {
	days := make(map[int]Day)
	days[1] = &solutions.Day1{}
	days[2] = &solutions.Day2{}

	return days
}
