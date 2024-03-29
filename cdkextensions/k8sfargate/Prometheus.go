package k8sfargate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/aps"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/k8sfargate/internal"
)

// Deploys Prometheus into EKS.
//
// The service is run in Fargate and writes data to Amazon Managed Service for
// Prometheus which provides the backing data store.
// See: [Official Helm chart documentation](https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus#readme)
//
type Prometheus interface {
	awscdk.Resource
	// The Helm chart that was used to deploy Prometheus.
	Chart() awseks.HelmChart
	// The EKS cluster where Prometheus should be deployed.
	Cluster() awseks.Cluster
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The Fargate profile used for running the service in Fargate.
	FargateProfile() awseks.FargateProfile
	// The Kubernetes namespace where the service should be deployed.
	Namespace() *string
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
	// Configures the queue used to write to Amazon Managed Service for Prometheus.
	QueueConfiguration() *QueueConfiguration
	// The service account that Prometheus will use to gain permissions for Kubernetes and AWS.
	ServiceAccount() awseks.ServiceAccount
	// Name of the Kubernetes service account that should be created and used by Prometheus.
	ServiceAccountName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The Amazon Managed Service for Prometheus workspace where the Prometheus server should sned its data.
	Workspace() aps.IWorkspace
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

// The jsii proxy struct for Prometheus
type jsiiProxy_Prometheus struct {
	internal.Type__awscdkResource
}

func (j *jsiiProxy_Prometheus) Chart() awseks.HelmChart {
	var returns awseks.HelmChart
	_jsii_.Get(
		j,
		"chart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Cluster() awseks.Cluster {
	var returns awseks.Cluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) FargateProfile() awseks.FargateProfile {
	var returns awseks.FargateProfile
	_jsii_.Get(
		j,
		"fargateProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) QueueConfiguration() *QueueConfiguration {
	var returns *QueueConfiguration
	_jsii_.Get(
		j,
		"queueConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) ServiceAccount() awseks.ServiceAccount {
	var returns awseks.ServiceAccount
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Prometheus) Workspace() aps.IWorkspace {
	var returns aps.IWorkspace
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}


// Creates a new instance of the Prometheus class.
func NewPrometheus(scope constructs.IConstruct, id *string, props *PrometheusProps) Prometheus {
	_init_.Initialize()

	if err := validateNewPrometheusParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Prometheus{}

	_jsii_.Create(
		"cdk-extensions.k8s_fargate.Prometheus",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the Prometheus class.
func NewPrometheus_Override(p Prometheus, scope constructs.IConstruct, id *string, props *PrometheusProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.k8s_fargate.Prometheus",
		[]interface{}{scope, id, props},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func Prometheus_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePrometheus_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_fargate.Prometheus",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Prometheus_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePrometheus_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_fargate.Prometheus",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Prometheus_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validatePrometheus_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.k8s_fargate.Prometheus",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Prometheus_CHART_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_fargate.Prometheus",
		"CHART_NAME",
		&returns,
	)
	return returns
}

func Prometheus_CHART_REPOSITORY() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_fargate.Prometheus",
		"CHART_REPOSITORY",
		&returns,
	)
	return returns
}

func Prometheus_DEFAULT_NAMESPACE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-extensions.k8s_fargate.Prometheus",
		"DEFAULT_NAMESPACE",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_Prometheus) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := p.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (p *jsiiProxy_Prometheus) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheus) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := p.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheus) GetResourceNameAttribute(nameAttr *string) *string {
	if err := p.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Prometheus) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

