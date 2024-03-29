package ec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/ec2/internal"
)

type IpamResourceDiscovery interface {
	awscdk.Resource
	IIpamResourceDiscovery
	Description() *string
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The resource discovery ARN.
	IpamResourceDiscoveryArn() *string
	// The resource discovery ID.
	IpamResourceDiscoveryId() *string
	// Defines if the resource discovery is the default.
	//
	// The default resource
	// discovery is the resource discovery automatically created when you create
	// an IPAM.
	IpamResourceDiscoveryIsDefault() awscdk.IResolvable
	// The owner ID.
	IpamResourceDiscoveryOwnerId() *string
	// The resource discovery Region.
	IpamResourceDiscoveryRegion() *string
	// The resource discovery's state.
	//
	// - create-in-progress - Resource discovery is being created.
	// - create-complete - Resource discovery creation is complete.
	// - create-failed - Resource discovery creation has failed.
	// - modify-in-progress - Resource discovery is being modified.
	// - modify-complete - Resource discovery modification is complete.
	// - modify-failed - Resource discovery modification has failed.
	// - delete-in-progress - Resource discovery is being deleted.
	// - delete-complete - Resource discovery deletion is complete.
	// - delete-failed - Resource discovery deletion has failed.
	// - isolate-in-progress - AWS account that created the resource discovery
	// has been removed and the resource discovery is being isolated.
	// - isolate-complete - Resource discovery isolation is complete.
	// - restore-in-progress - AWS account that created the resource discovery
	// and was isolated has been restored.
	IpamResourceDiscoveryState() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	Regions() *[]*string
	Resource() awsec2.CfnIPAMResourceDiscovery
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	AddIpam(id *string, options *IpamProps) IIpam
	AddRegion(region *string)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	AssociateIpam(ipam IIpam) IIpamResourceDiscoveryAssociation
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for IpamResourceDiscovery
type jsiiProxy_IpamResourceDiscovery struct {
	internal.Type__awscdkResource
	jsiiProxy_IIpamResourceDiscovery
}

func (j *jsiiProxy_IpamResourceDiscovery) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryIsDefault() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryIsDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryOwnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) IpamResourceDiscoveryState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) Resource() awsec2.CfnIPAMResourceDiscovery {
	var returns awsec2.CfnIPAMResourceDiscovery
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IpamResourceDiscovery) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewIpamResourceDiscovery(scope constructs.IConstruct, id *string, props *IpamResourceDiscoveryProps) IpamResourceDiscovery {
	_init_.Initialize()

	if err := validateNewIpamResourceDiscoveryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpamResourceDiscovery{}

	_jsii_.Create(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewIpamResourceDiscovery_Override(i IpamResourceDiscovery, scope constructs.IConstruct, id *string, props *IpamResourceDiscoveryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		[]interface{}{scope, id, props},
		i,
	)
}

// Imports an existing IPAM resource discovery by specifying its Amazon Resource Name (ARN).
//
// Returns: An object representing the imported IPAM resource discovery.
func IpamResourceDiscovery_FromIpamResourceDiscoveryArn(scope constructs.IConstruct, id *string, ipamResourceDiscoveryArn *string) IIpamResourceDiscovery {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_FromIpamResourceDiscoveryArnParameters(scope, id, ipamResourceDiscoveryArn); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscovery

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"fromIpamResourceDiscoveryArn",
		[]interface{}{scope, id, ipamResourceDiscoveryArn},
		&returns,
	)

	return returns
}

// Imports an existing IPAM resource discovery by explicitly specifying its attributes.
//
// Returns: An object representing the imported IPAM resource discovery.
func IpamResourceDiscovery_FromIpamResourceDiscoveryAttributes(scope constructs.IConstruct, id *string, attrs *IpamResourceDiscoveryAttributes) IIpamResourceDiscovery {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_FromIpamResourceDiscoveryAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscovery

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"fromIpamResourceDiscoveryAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Imports an existing IPAM resource discovery by explicitly specifying its AWS generated ID.
//
// Returns: An object representing the imported IPAM resource discovery.
func IpamResourceDiscovery_FromIpamResourceDiscoveryId(scope constructs.IConstruct, id *string, ipamResourceDiscoveryId *string) IIpamResourceDiscovery {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_FromIpamResourceDiscoveryIdParameters(scope, id, ipamResourceDiscoveryId); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscovery

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"fromIpamResourceDiscoveryId",
		[]interface{}{scope, id, ipamResourceDiscoveryId},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func IpamResourceDiscovery_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func IpamResourceDiscovery_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func IpamResourceDiscovery_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateIpamResourceDiscovery_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func IpamResourceDiscovery_ARN_FORMAT() awscdk.ArnFormat {
	_init_.Initialize()
	var returns awscdk.ArnFormat
	_jsii_.StaticGet(
		"cdk-extensions.ec2.IpamResourceDiscovery",
		"ARN_FORMAT",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) AddIpam(id *string, options *IpamProps) IIpam {
	if err := i.validateAddIpamParameters(id, options); err != nil {
		panic(err)
	}
	var returns IIpam

	_jsii_.Invoke(
		i,
		"addIpam",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) AddRegion(region *string) {
	if err := i.validateAddRegionParameters(region); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addRegion",
		[]interface{}{region},
	)
}

func (i *jsiiProxy_IpamResourceDiscovery) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IpamResourceDiscovery) AssociateIpam(ipam IIpam) IIpamResourceDiscoveryAssociation {
	if err := i.validateAssociateIpamParameters(ipam); err != nil {
		panic(err)
	}
	var returns IIpamResourceDiscoveryAssociation

	_jsii_.Invoke(
		i,
		"associateIpam",
		[]interface{}{ipam},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IpamResourceDiscovery) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

