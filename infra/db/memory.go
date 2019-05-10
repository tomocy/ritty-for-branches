package db

import (
	derr "github.com/tomocy/ritty-for-branches/domain/error"
	"github.com/tomocy/ritty-for-branches/domain/model"
)

func NewMemory() *Memory {
	return new(Memory)
}

type Memory struct {
	branches []*model.Branch
}

func (m *Memory) GetBranches() []*model.Branch {
	return m.branches
}

func (m *Memory) FindBranch(id string) (*model.Branch, error) {
	for _, stored := range m.branches {
		if stored.ID == id {
			return stored, nil
		}
	}

	return nil, derr.ValidationErrorf("no such branch")
}

func (m *Memory) SaveBranch(branch *model.Branch) error {
	for i, stored := range m.branches {
		if stored.ID == branch.ID {
			m.branches[i] = branch
			return nil
		}
	}

	m.branches = append(m.branches, branch)

	return nil
}
