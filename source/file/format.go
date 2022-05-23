package file

import (
	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
)

func format(p string, e encoder.Encoder) string {
	return e.String()
}
