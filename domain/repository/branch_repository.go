package repository

import "github.com/tomocy/ritty-for-branches/domain/model"

type BranchRepository interface {
	GetBranches() []*model.Branch
	FindBranch(id string) (*model.Branch, error)
	SaveBranch(branch *model.Branch) error
}
