//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func validateNewClassificationStringParameters(value *string) error {
	return nil
}

