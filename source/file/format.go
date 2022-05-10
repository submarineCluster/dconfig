package file

import (
	"strings"

	"git.code.oa.com/tencent_abtest/go-common/dconfig/encoder"
)

func format(p string, e encoder.Encoder) string {
	parts := strings.Split(p, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}
	return e.String()
}
