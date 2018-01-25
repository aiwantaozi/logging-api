package logging

import (
	"github.com/rancher/norman/types"
)

const (
	nameField = "name"
)

func Validator(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}) error {
	//todo
	// value, ok := data[nameField]
	// if !ok {
	// 	return nil
	// }

	// if err := validate(value); err != nil {
	// 	return httperror.NewAPIError(httperror.InvalidFormat, fmt.Sprintf("invalid %s %s: %s", nameField, value,
	// 		strings.Join(errs, ",")))
	// }

	return nil
}
