//go:build !no_runtime_type_checking

package ssm

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_IDocumentContent) validateBindParameters(scope constructs.IConstruct) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	return nil
}

