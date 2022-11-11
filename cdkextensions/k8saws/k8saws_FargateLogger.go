package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws/internal"
)

// Creates a ConfigMap that configures logging for containers running in EKS on Fargate.
type FargateLogger interface {
	awscdk.Resource
	// The EKS cluster where Fargate logging is being configured.
	Cluster() awseks.ICluster
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// Collection of Fluent Bit filter plugins being configured for logging.
	Filters() *[]IFluentBitFilterPlugin
	// The Kubernetes manifest that creates the ConfigMap that Fargate uses to configure logging.
	Manifest() awseks.KubernetesManifest
	// The tree node.
	Node() constructs.Node
	// Collection of Fluent Bit output plugins being configured for logging.
	Outputs() *[]IFluentBitOutputPlugin
	// Collection of Fluent Bit parser plugins being configured for logging.
	Parsers() *[]IFluentBitParserPlugin
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	AddFargateProfile(profile awseks.FargateProfile) FargateLogger
	AddFilter(filter IFluentBitFilterPlugin) FargateLogger
	AddOutput(output IFluentBitOutputPlugin) FargateLogger
	AddParser(parser IFluentBitParserPlugin) FargateLogger
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

// The jsii proxy struct for FargateLogger
type jsiiProxy_FargateLogger struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_FargateLogger) Cluster() awseks.ICluster {
	var returns awseks.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Filters() *[]IFluentBitFilterPlugin {
	var returns *[]IFluentBitFilterPlugin
	_jsii_.Get(
		j,
		"filters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Manifest() awseks.KubernetesManifest {
	var returns awseks.KubernetesManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Outputs() *[]IFluentBitOutputPlugin {
	var returns *[]IFluentBitOutputPlugin
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Parsers() *[]IFluentBitParserPlugin {
	var returns *[]IFluentBitParserPlugin
	_jsii_.Get(
		j,
		"parsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateLogger) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates a new instance of the FargateLogger class.
func NewFargateLogger(scope constructs.Construct, id *string, props *FargateLoggerProps) FargateLogger {
	_init_.Initialize()

	if err := validateNewFargateLoggerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FargateLogger{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FargateLogger",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the FargateLogger class.
func NewFargateLogger_Override(f FargateLogger, scope constructs.Construct, id *string, props *FargateLoggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.FargateLogger",
		[]interface{}{scope, id, props},
		f,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func FargateLogger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFargateLogger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FargateLogger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func FargateLogger_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFargateLogger_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FargateLogger",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func FargateLogger_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateFargateLogger_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.FargateLogger",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) AddFargateProfile(profile awseks.FargateProfile) FargateLogger {
	if err := f.validateAddFargateProfileParameters(profile); err != nil {
		panic(err)
	}
	var returns FargateLogger

	_jsii_.Invoke(
		f,
		"addFargateProfile",
		[]interface{}{profile},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) AddFilter(filter IFluentBitFilterPlugin) FargateLogger {
	if err := f.validateAddFilterParameters(filter); err != nil {
		panic(err)
	}
	var returns FargateLogger

	_jsii_.Invoke(
		f,
		"addFilter",
		[]interface{}{filter},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) AddOutput(output IFluentBitOutputPlugin) FargateLogger {
	if err := f.validateAddOutputParameters(output); err != nil {
		panic(err)
	}
	var returns FargateLogger

	_jsii_.Invoke(
		f,
		"addOutput",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) AddParser(parser IFluentBitParserPlugin) FargateLogger {
	if err := f.validateAddParserParameters(parser); err != nil {
		panic(err)
	}
	var returns FargateLogger

	_jsii_.Invoke(
		f,
		"addParser",
		[]interface{}{parser},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := f.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (f *jsiiProxy_FargateLogger) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := f.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) GetResourceNameAttribute(nameAttr *string) *string {
	if err := f.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateLogger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

