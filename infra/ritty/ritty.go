package ritty

import derr "github.com/tomocy/ritty-for-branches/domain/error"

func New() *Ritty {
	return new(Ritty)
}

type Ritty struct {
	*branchAuth
}

func devErrorf(did string, msg interface{}) *derr.DevError {
	return derr.DevErrorf("failed to %s: %v", did, msg)
}
