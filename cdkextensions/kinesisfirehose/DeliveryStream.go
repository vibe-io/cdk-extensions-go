package kinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/kinesisfirehose/internal"
)

type DeliveryStream interface {
	awscdk.Resource
	IDeliveryStream
	// The network connections associated with this resource.
	Connections() awsec2.Connections
	DeliveryStreamArn() *string
	DeliveryStreamName() *string
	Destination() DeliveryStreamDestination
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The principal to grant permissions to.
	GrantPrincipal() awsiam.IPrincipal
	Name() *string
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
	Resource() awskinesisfirehose.CfnDeliveryStream
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	StreamType() DeliveryStreamType
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
	Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant
	GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for DeliveryStream
type jsiiProxy_DeliveryStream struct {
	internal.Type__awscdkResource
	jsiiProxy_IDeliveryStream
}

func (j *jsiiProxy_DeliveryStream) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Destination() DeliveryStreamDestination {
	var returns DeliveryStreamDestination
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Resource() awskinesisfirehose.CfnDeliveryStream {
	var returns awskinesisfirehose.CfnDeliveryStream
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DeliveryStream) StreamType() DeliveryStreamType {
	var returns DeliveryStreamType
	_jsii_.Get(
		j,
		"streamType",
		&returns,
	)
	return returns
}


func NewDeliveryStream(scope constructs.Construct, id *string, props *DeliveryStreamProps) DeliveryStream {
	_init_.Initialize()

	if err := validateNewDeliveryStreamParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DeliveryStream{}

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDeliveryStream_Override(d DeliveryStream, scope constructs.Construct, id *string, props *DeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		[]interface{}{scope, id, props},
		d,
	)
}

func DeliveryStream_FromDeliveryStreamArn(scope constructs.Construct, id *string, deliveryStreamArn *string) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamArnParameters(scope, id, deliveryStreamArn); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"fromDeliveryStreamArn",
		[]interface{}{scope, id, deliveryStreamArn},
		&returns,
	)

	return returns
}

func DeliveryStream_FromDeliveryStreamAttributes(scope constructs.Construct, id *string, attrs *DeliveryStreamAttributes) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"fromDeliveryStreamAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

func DeliveryStream_FromDeliveryStreamName(scope constructs.Construct, id *string, deliveryStreamName *string) IDeliveryStream {
	_init_.Initialize()

	if err := validateDeliveryStream_FromDeliveryStreamNameParameters(scope, id, deliveryStreamName); err != nil {
		panic(err)
	}
	var returns IDeliveryStream

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"fromDeliveryStreamName",
		[]interface{}{scope, id, deliveryStreamName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func DeliveryStream_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDeliveryStream_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func DeliveryStream_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateDeliveryStream_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.kinesis_firehose.DeliveryStream",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := d.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (d *jsiiProxy_DeliveryStream) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := d.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GetResourceNameAttribute(nameAttr *string) *string {
	if err := d.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Grant(grantee awsiam.IGrantable, actions ...*string) awsiam.Grant {
	if err := d.validateGrantParameters(grantee); err != nil {
		panic(err)
	}
	args := []interface{}{grantee}
	for _, a := range actions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grant",
		args,
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) GrantPutRecords(grantee awsiam.IGrantable) awsiam.Grant {
	if err := d.validateGrantPutRecordsParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		d,
		"grantPutRecords",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Bytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3BytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Bytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3DataFreshness(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3DataFreshnessParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3DataFreshness",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricBackupToS3Records(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricBackupToS3RecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricBackupToS3Records",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricIncomingBytes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricIncomingBytesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingBytes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) MetricIncomingRecords(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricIncomingRecordsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricIncomingRecords",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
