//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ram

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISharedResource) validateBindParameters(scope constructs.IConstruct) error {
	return nil
}

