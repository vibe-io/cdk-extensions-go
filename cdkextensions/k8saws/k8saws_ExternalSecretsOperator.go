package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws/internal"
)

type ExternalSecretsOperator interface {
	awscdk.Resource
	Cluster() awseks.Cluster
	CreateNamespace() *bool
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	HelmChart() awseks.HelmChart
	Name() *string
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	OperatorName() *string
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
	RegisterSecretsManagerSecret(id *string, secret awssecretsmanager.ISecret, options *NamespacedExternalSecretOptions) ExternalSecret
	RegisterSsmParameterSecret(id *string, parameter awsssm.IParameter, options *NamespacedExternalSecretOptions) ExternalSecret
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for ExternalSecretsOperator
type jsiiProxy_ExternalSecretsOperator struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ExternalSecretsOperator) Cluster() awseks.Cluster {
	var returns awseks.Cluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) CreateNamespace() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) HelmChart() awseks.HelmChart {
	var returns awseks.HelmChart
	_jsii_.Get(
		j,
		"helmChart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) OperatorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecretsOperator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


func NewExternalSecretsOperator(scope constructs.Construct, id *string, props *ExternalSecretsOperatorProps) ExternalSecretsOperator {
	_init_.Initialize()

	if err := validateNewExternalSecretsOperatorParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalSecretsOperator{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewExternalSecretsOperator_Override(e ExternalSecretsOperator, scope constructs.Construct, id *string, props *ExternalSecretsOperatorProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ExternalSecretsOperator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalSecretsOperator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ExternalSecretsOperator_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateExternalSecretsOperator_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ExternalSecretsOperator_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateExternalSecretsOperator_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func ExternalSecretsOperator_DEFAULT_NAMESPACE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_aws.ExternalSecretsOperator",
		"DEFAULT_NAMESPACE",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_ExternalSecretsOperator) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := e.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) GetResourceNameAttribute(nameAttr *string) *string {
	if err := e.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) RegisterSecretsManagerSecret(id *string, secret awssecretsmanager.ISecret, options *NamespacedExternalSecretOptions) ExternalSecret {
	if err := e.validateRegisterSecretsManagerSecretParameters(id, secret, options); err != nil {
		panic(err)
	}
	var returns ExternalSecret

	_jsii_.Invoke(
		e,
		"registerSecretsManagerSecret",
		[]interface{}{id, secret, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) RegisterSsmParameterSecret(id *string, parameter awsssm.IParameter, options *NamespacedExternalSecretOptions) ExternalSecret {
	if err := e.validateRegisterSsmParameterSecretParameters(id, parameter, options); err != nil {
		panic(err)
	}
	var returns ExternalSecret

	_jsii_.Invoke(
		e,
		"registerSsmParameterSecret",
		[]interface{}{id, parameter, options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecretsOperator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

