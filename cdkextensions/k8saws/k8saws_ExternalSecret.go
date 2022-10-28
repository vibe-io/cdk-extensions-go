package k8saws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8saws/internal"
)

// Represents a Kubernetes secret that is being synchronized from an external provider.
//
// On a technical level, provides the configuration for how the external
// secrets operator service should manage the synchronization of the Kubernetes
// secret.
type ExternalSecret interface {
	awscdk.Resource
	// The EKS cluster where the secret should be created.
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
	// The Kubernetes manifest defining the configuration of how to synchronize the Kubernetes secret from the provider secrets.
	Manifest() awseks.KubernetesManifest
	// The name to use for the Kubernetes secret resource when it is synchronized into the cluster.
	Name() *string
	// The name where the synchronized secret should be created.
	Namespace() *string
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//    cross-environment scenarios.
	PhysicalName() *string
	// The frequency at which synchronization should occur.
	RefreshInterval() awscdk.Duration
	// The name of the Kubernetes secret.
	SecretName() *string
	// The collection of referenced provider secrets that are referenced in the Kubernetes secret.
	Secrets() *[]ISecretReference
	// The Kubernetes secret store resource that provides details and permissions to use for importing secrets from the provider.
	SecretStore() ISecretStore
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// Adds a provider secret reference to the synchronized Kubernetes secret.
	//
	// For external secrets that reference multiple provider secrets the keys of
	// all provider secrets will be merged into the single Kubernetes secret.
	//
	// Returns: The external secret resoiurce where the reference was added.
	AddSecret(secret ISecretReference) ExternalSecret
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

// The jsii proxy struct for ExternalSecret
type jsiiProxy_ExternalSecret struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_ExternalSecret) Cluster() awseks.ICluster {
	var returns awseks.ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Manifest() awseks.KubernetesManifest {
	var returns awseks.KubernetesManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) RefreshInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"refreshInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) SecretName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Secrets() *[]ISecretReference {
	var returns *[]ISecretReference
	_jsii_.Get(
		j,
		"secrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) SecretStore() ISecretStore {
	var returns ISecretStore
	_jsii_.Get(
		j,
		"secretStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExternalSecret) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


// Creates a new instance of the ExternalSecret class.
func NewExternalSecret(scope constructs.Construct, id *string, props *ExternalSecretProps) ExternalSecret {
	_init_.Initialize()

	if err := validateNewExternalSecretParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExternalSecret{}

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalSecret",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the ExternalSecret class.
func NewExternalSecret_Override(e ExternalSecret, scope constructs.Construct, id *string, props *ExternalSecretProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_aws.ExternalSecret",
		[]interface{}{scope, id, props},
		e,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ExternalSecret_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExternalSecret_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecret",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ExternalSecret_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateExternalSecret_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecret",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ExternalSecret_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateExternalSecret_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_aws.ExternalSecret",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecret) AddSecret(secret ISecretReference) ExternalSecret {
	if err := e.validateAddSecretParameters(secret); err != nil {
		panic(err)
	}
	var returns ExternalSecret

	_jsii_.Invoke(
		e,
		"addSecret",
		[]interface{}{secret},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecret) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_ExternalSecret) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExternalSecret) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
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

func (e *jsiiProxy_ExternalSecret) GetResourceNameAttribute(nameAttr *string) *string {
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

func (e *jsiiProxy_ExternalSecret) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
