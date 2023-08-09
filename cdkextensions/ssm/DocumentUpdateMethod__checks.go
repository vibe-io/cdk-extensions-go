//go:build !no_runtime_type_checking

package ssm

import (
	"fmt"
)

func validateDocumentUpdateMethod_OfParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

