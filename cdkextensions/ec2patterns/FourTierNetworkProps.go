package ec2patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2"
)

type FourTierNetworkProps struct {
	AddressManager IpAddressManager `field:"optional" json:"addressManager" yaml:"addressManager"`
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	Cidr ec2.IIpv4CidrAssignment `field:"optional" json:"cidr" yaml:"cidr"`
	DefaultInstanceTenancy awsec2.DefaultInstanceTenancy `field:"optional" json:"defaultInstanceTenancy" yaml:"defaultInstanceTenancy"`
	EnableDnsHostnames *bool `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	EnableDnsSupport *bool `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	FlowLogs *map[string]*FlowLogOptions `field:"optional" json:"flowLogs" yaml:"flowLogs"`
	GatewayEndpoints *map[string]*awsec2.GatewayVpcEndpointOptions `field:"optional" json:"gatewayEndpoints" yaml:"gatewayEndpoints"`
	MaxAzs *float64 `field:"optional" json:"maxAzs" yaml:"maxAzs"`
	NatGatewayProvider awsec2.NatProvider `field:"optional" json:"natGatewayProvider" yaml:"natGatewayProvider"`
	NatGateways *float64 `field:"optional" json:"natGateways" yaml:"natGateways"`
	NatGatewaySubnets *awsec2.SubnetSelection `field:"optional" json:"natGatewaySubnets" yaml:"natGatewaySubnets"`
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
	VpnConnections *map[string]*awsec2.VpnConnectionOptions `field:"optional" json:"vpnConnections" yaml:"vpnConnections"`
	VpnGateway *bool `field:"optional" json:"vpnGateway" yaml:"vpnGateway"`
	VpnGatewayAsn *float64 `field:"optional" json:"vpnGatewayAsn" yaml:"vpnGatewayAsn"`
	VpnRoutePropagation *[]*awsec2.SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
}

