package legacyconversions

import (
	"github.com/empijei/def-prog-exercises/safesql"
	"github.com/empijei/def-prog-exercises/safesql/internal/raw"
)

var TrustedSQLCtor = raw.TrustedSQLCtor.(func(string) safesql.TrustedSQL)

func RiskilyAssumeTrustedSQL(trusted string) safesql.TrustedSQL {
	return TrustedSQLCtor(trusted)
}
