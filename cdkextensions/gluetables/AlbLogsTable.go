package gluetables

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/athena"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/gluetables/internal"
)

type AlbLogsTable interface {
	glue.Table
	// {@link TableProps.compressed}.
	Compressed() *bool
	// Boolean indicating whether to create default Athena queries for the ALB Logs.
	// See: [`CfnNamedQueries`](https://docs.aws.amazon.com/cdk/api/v1/python/aws_cdk.aws_athena/CfnNamedQuery.html)
	//
	CreateQueries() *bool
	// {@link TableProps.database:}.
	Database() glue.Database
	// {@link TableProps.dataFormat}.
	DataFormat() glue.DataFormat
	// {@link TableProps.description}.
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
	// Boolean for adding "friendly names" for the created Athena queries.
	FriendlyQueryNames() *bool
	// {@link TableProps.location}.
	Location() *string
	// {@link TableProps.name}.
	Name() *string
	// The tree node.
	Node() constructs.Node
	// {@link TableProps.owner}.
	Owner() *string
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	Resource() awsglue.CfnTable
	// {@link TableProps.retention}.
	Retention() awscdk.Duration
	// {@link TableProps.serdeName}.
	SerdeName() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	Status5xxNamedQuery() athena.NamedQuery
	// {@link TableProps.storedAsSubDirectories}.
	StoredAsSubDirectories() *bool
	TableArn() *string
	TableName() *string
	// {@link TableProps.tableType}.
	TableType() glue.TableType
	// {@link TableProps.targetTable}.
	TargetTable() glue.Table
	TopIpsNamedQuery() athena.NamedQuery
	// {@link TableProps.viewExpandedText}.
	ViewExpandedText() *string
	// {@link TableProps.viewOriginalText}.
	ViewOriginalText() *string
	// The name of the workgroup where namedqueries should be created.
	// See: [Setting up workgroups](https://docs.aws.amazon.com/athena/latest/ug/workgroups-procedure.html)
	//
	WorkGroup() athena.IWorkGroup
	AddColumn(column glue.Column)
	AddParameter(key *string, value *string)
	AddPartitionKey(column glue.Column)
	AddSerdeParameter(key *string, value *string)
	AddStorageParameter(key *string, value *string)
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
	RenderStorageDescriptor() *awsglue.CfnTable_StorageDescriptorProperty
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for AlbLogsTable
type jsiiProxy_AlbLogsTable struct {
	internal.Type__glueTable
}

func (j *jsiiProxy_AlbLogsTable) Compressed() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"compressed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) CreateQueries() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Database() glue.Database {
	var returns glue.Database
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) DataFormat() glue.DataFormat {
	var returns glue.DataFormat
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) FriendlyQueryNames() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"friendlyQueryNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Resource() awsglue.CfnTable {
	var returns awsglue.CfnTable
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Retention() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) SerdeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serdeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) Status5xxNamedQuery() athena.NamedQuery {
	var returns athena.NamedQuery
	_jsii_.Get(
		j,
		"status5xxNamedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) StoredAsSubDirectories() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"storedAsSubDirectories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) TableType() glue.TableType {
	var returns glue.TableType
	_jsii_.Get(
		j,
		"tableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) TargetTable() glue.Table {
	var returns glue.Table
	_jsii_.Get(
		j,
		"targetTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) TopIpsNamedQuery() athena.NamedQuery {
	var returns athena.NamedQuery
	_jsii_.Get(
		j,
		"topIpsNamedQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlbLogsTable) WorkGroup() athena.IWorkGroup {
	var returns athena.IWorkGroup
	_jsii_.Get(
		j,
		"workGroup",
		&returns,
	)
	return returns
}


// Creates a new instance of the AlbLogsTable class.
func NewAlbLogsTable(scope constructs.Construct, id *string, props *AlbLogsTableProps) AlbLogsTable {
	_init_.Initialize()

	if err := validateNewAlbLogsTableParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbLogsTable{}

	_jsii_.Create(
		"cdk-extensions.glue_tables.AlbLogsTable",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the AlbLogsTable class.
func NewAlbLogsTable_Override(a AlbLogsTable, scope constructs.Construct, id *string, props *AlbLogsTableProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.glue_tables.AlbLogsTable",
		[]interface{}{scope, id, props},
		a,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func AlbLogsTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbLogsTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.glue_tables.AlbLogsTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func AlbLogsTable_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAlbLogsTable_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.glue_tables.AlbLogsTable",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func AlbLogsTable_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateAlbLogsTable_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.glue_tables.AlbLogsTable",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLogsTable) AddColumn(column glue.Column) {
	if err := a.validateAddColumnParameters(column); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addColumn",
		[]interface{}{column},
	)
}

func (a *jsiiProxy_AlbLogsTable) AddParameter(key *string, value *string) {
	if err := a.validateAddParameterParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addParameter",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_AlbLogsTable) AddPartitionKey(column glue.Column) {
	if err := a.validateAddPartitionKeyParameters(column); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPartitionKey",
		[]interface{}{column},
	)
}

func (a *jsiiProxy_AlbLogsTable) AddSerdeParameter(key *string, value *string) {
	if err := a.validateAddSerdeParameterParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addSerdeParameter",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_AlbLogsTable) AddStorageParameter(key *string, value *string) {
	if err := a.validateAddStorageParameterParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addStorageParameter",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_AlbLogsTable) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := a.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (a *jsiiProxy_AlbLogsTable) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLogsTable) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := a.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLogsTable) GetResourceNameAttribute(nameAttr *string) *string {
	if err := a.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLogsTable) RenderStorageDescriptor() *awsglue.CfnTable_StorageDescriptorProperty {
	var returns *awsglue.CfnTable_StorageDescriptorProperty

	_jsii_.Invoke(
		a,
		"renderStorageDescriptor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbLogsTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

