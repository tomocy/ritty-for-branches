package model

import derr "github.com/tomocy/ritty-for-branches/domain/error"

func validationErrorf(did string, msg interface{}) *derr.ValidationError {
	return derr.ValidationErrorf("failed to %s: %v", did, msg)
}
