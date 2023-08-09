//go:build !no_runtime_type_checking

package ssm

import (
	"fmt"
)

func (i *jsiiProxy_IAutomationDocument) validateArnForAutomationDefinitionVersionParameters(version *string) error {
	if version == nil {
		return fmt.Errorf("parameter version is required, but nil was provided")
	}

	return nil
}

