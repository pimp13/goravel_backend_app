package rules

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type Unique struct {
}

// Signature The name of the rule.
func (receiver *Unique) Signature() string {
	return "unique"
}

// Passes Determine if the validation rule passes.
func (receiver *Unique) Passes(data validation.Data, val any, options ...any) bool {
	tableName := options[0].(string)
	columnName := options[1].(string)
	isExists, err := facades.Orm().Query().Table(tableName).Where(columnName, val).Exists()
	if err != nil {
		facades.Log().Errorf(
			"failed to find from table #%s get column #%s is exists for unique: %v",
			tableName,
			columnName,
			err,
		)
		return false
	}
	// if isExists {
	// 	return true
	// }
	// return false
	return isExists
}

// Message Get the validation error message.
func (receiver *Unique) Message() string {
	return ":attribute has been taken."
}
