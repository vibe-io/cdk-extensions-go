//go:build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (v *jsiiProxy_VpnConnection) validateAddTunnelConfigurationParameters(options *TunnelOptions) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateMetricTunnelDataInParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateMetricTunnelDataOutParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func (v *jsiiProxy_VpnConnection) validateMetricTunnelStateParameters(props *awscloudwatch.MetricOptions) error {
	return nil
}

func validateVpnConnection_IsConstructParameters(x interface{}) error {
	return nil
}

func validateVpnConnection_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateVpnConnection_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewVpnConnectionParameters(scope constructs.Construct, id *string, props *VpnConnectionProps) error {
	return nil
}

