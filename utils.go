package admingorm

import (
	"reflect"
	"strings"
)

// getGormColumnName retrieves the database column name for a given struct field,
// based on the GORM tags. If the 'gorm' tag specifies a 'column' option, that value is used.
// Otherwise, the struct field name is used as the column name.
func getGormColumnName(structField reflect.StructField) string {
	tag, ok := structField.Tag.Lookup("gorm")
	if !ok {
		return structField.Name
	}
	tagParts := strings.Split(tag, ";")
	for _, part := range tagParts {
		if strings.HasPrefix(part, "column:") {
			return strings.TrimPrefix(part, "column:")
		}
	}
	return structField.Name
}
