//go:build !no_runtime_type_checking

package ec2

import (
	"fmt"
)

func validateIpamPoolCidrConfiguration_CidrParameters(cidr *string) error {
	if cidr == nil {
		return fmt.Errorf("parameter cidr is required, but nil was provided")
	}

	return nil
}

func validateIpamPoolCidrConfiguration_NetmaskParameters(length *float64) error {
	if length == nil {
		return fmt.Errorf("parameter length is required, but nil was provided")
	}

	return nil
}

