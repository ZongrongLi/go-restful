package dl

import "github.com/tiancai110a/test_user/pkg/errno"

type Base struct {
	Err *errno.Errno
}

func CheckStatus(b *Base) bool {
	if b == nil {
		b.Err = errno.ErrEmpty
		return false
	}
	if b.Err.Code != 0 {
		return false
	}
	return true
}
