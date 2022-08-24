package queries

import (
	"fmt"
	"strings"
)

func (q EsQuery) Build() string {
	var sb strings.Builder

	for _, eq := range q.Equals {
		sb.WriteString(fmt.Sprintf(`%s:"%v" %s `, eq.Field, eq.Value, eq.Operator))
	}

	return strings.TrimSpace(sb.String())
}
