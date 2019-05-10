package db

func NewMemory() *Memory {
	return new(Memory)
}

type Memory struct{}
