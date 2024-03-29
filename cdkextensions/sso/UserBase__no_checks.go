//go:build no_runtime_type_checking

package sso

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_UserBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (u *jsiiProxy_UserBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateUserBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUserBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateUserBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewUserBaseParameters(scope constructs.IConstruct, id *string, props *awscdk.ResourceProps) error {
	return nil
}

