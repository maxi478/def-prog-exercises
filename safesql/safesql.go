package safesql

import "internal/raw"

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string(text)}
}

func init() {
	raw.TrustedSQLCtor =
		func(unsafe string) TrustedSQL {
			return TrustedSQL{unsafe}
		}
}
