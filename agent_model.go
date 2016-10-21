package main

import "time"

type Agent struct {
	ReviewTime   time.Time
	Name         string
	IsSolicited  bool
	Tag          string
	NumOfWords   int
	Stars        float32
	Score        float32
	NumOfRecords int
}

type DataWriter interface {
	GetWritableData() Agent
}

func (a *Agent) GetWritableData() Agent {
	return *a
}
