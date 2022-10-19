package s3buckets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/vibe-io/cdk-extensions-go/cdkextensions/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/glue"
	"github.com/vibe-io/cdk-extensions-go/cdkextensions/gluetables"
)

type S3AccessLogsBucket interface {
	RawBucket
	// The ARN of the bucket.
	BucketArn() *string
	// The IPv4 DNS name of the specified bucket.
	BucketDomainName() *string
	// The IPv6 DNS name of the specified bucket.
	BucketDualStackDomainName() *string
	// The name of the bucket.
	BucketName() *string
	// The regional domain name of the specified bucket.
	BucketRegionalDomainName() *string
	// The Domain name of the static website.
	BucketWebsiteDomainName() *string
	// The URL of the static website.
	BucketWebsiteUrl() *string
	CreateQueries() *bool
	Database() glue.Database
	// Optional KMS encryption key associated with this bucket.
	EncryptionKey() awskms.IKey
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	FriendlyQueryNames() *bool
	// If this bucket has been configured for static website hosting.
	IsWebsite() *bool
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
	// The resource policy associated with this bucket.
	//
	// If `autoCreatePolicy` is true, a `BucketPolicy` will be created upon the
	// first call to addToResourcePolicy(s).
	Policy() awss3.BucketPolicy
	SetPolicy(val awss3.BucketPolicy)
	Resource() awss3.CfnBucket
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	Table() gluetables.S3AccessLogsTable
	// Adds a bucket notification event destination.
	AddEventNotification(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter)
	AddLoggingAspect(scope constructs.IConstruct, options *LoggingAspectOptions)
	// Subscribes a destination to receive notifications when an object is created in the bucket.
	//
	// This is identical to calling
	// `onEvent(s3.EventType.OBJECT_CREATED)`.
	AddObjectCreatedNotification(_dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter)
	// Subscribes a destination to receive notifications when an object is removed from the bucket.
	//
	// This is identical to calling
	// `onEvent(EventType.OBJECT_REMOVED)`.
	AddObjectRemovedNotification(_dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter)
	// Adds a statement to the resource policy for a principal (i.e. account/role/service) to perform actions on this bucket and/or its contents. Use `bucketArn` and `arnForObjects(keys)` to obtain ARNs for this bucket or objects.
	//
	// Note that the policy statement may or may not be added to the policy.
	// For example, when an `IBucket` is created from an existing bucket,
	// it's not possible to tell whether the bucket already has a policy
	// attached, let alone to re-use that policy to add more statements to it.
	// So it's safest to do nothing in these cases.
	AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult
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
	// Returns an ARN that represents all objects within the bucket that match the key pattern specified.
	//
	// To represent all keys, specify ``"*"``.
	ArnForObjects(_keyPattern *string) *string
	// Enables event bridge notification, causing all events below to be sent to EventBridge:.
	//
	// - Object Deleted (DeleteObject)
	// - Object Deleted (Lifecycle expiration)
	// - Object Restore Initiated
	// - Object Restore Completed
	// - Object Restore Expired
	// - Object Storage Class Changed
	// - Object Access Tier Changed
	// - Object ACL Updated
	// - Object Tags Added
	// - Object Tags Deleted.
	EnableEventBridgeNotification()
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
	// Grants s3:DeleteObject* permission to an IAM principal for objects in this bucket.
	GrantDelete(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant
	// Allows unrestricted access to objects from this bucket.
	//
	// IMPORTANT: This permission allows anyone to perform actions on S3 objects
	// in this bucket, which is useful for when you configure your bucket as a
	// website and want everyone to be able to read objects in the bucket without
	// needing to authenticate.
	//
	// Without arguments, this method will grant read ("s3:GetObject") access to
	// all objects ("*") in the bucket.
	//
	// The method returns the `iam.Grant` object, which can then be modified
	// as needed. For example, you can add a condition that will restrict access only
	// to an IPv4 range like this:
	//
	//      const grant = bucket.grantPublicAccess();
	//      grant.resourceStatement!.addCondition(‘IpAddress’, { “aws:SourceIp”: “54.240.143.0/24” });
	GrantPublicAccess(_keyPrefix *string, _allowedActions ...*string) awsiam.Grant
	// Grants s3:PutObject* and s3:Abort* permissions for this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	GrantPut(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant
	// Grant the given IAM identity permissions to modify the ACLs of objects in the given Bucket.
	//
	// If your application has the '@aws-cdk/aws-s3:grantWriteWithoutAcl' feature flag set,
	// calling {@link grantWrite} or {@link grantReadWrite} no longer grants permissions to modify the ACLs of the objects;
	// in this case, if you need to modify object ACLs, call this method explicitly.
	GrantPutAcl(_identity awsiam.IGrantable, _objectsKeyPattern *string) awsiam.Grant
	// Grant read permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If encryption is used, permission to use the key to decrypt the contents
	// of the bucket will also be granted to the same principal.
	GrantRead(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant
	// Grants read/write permissions for this bucket and it's contents to an IAM principal (Role/Group/User).
	//
	// If an encryption key is used, permission to use the key for
	// encrypt/decrypt will also be granted.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantReadWrite(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant
	// Grant write permissions to this bucket to an IAM principal.
	//
	// If encryption is used, permission to use the key to encrypt the contents
	// of written files will also be granted to the same principal.
	//
	// Before CDK version 1.85.0, this method granted the `s3:PutObject*` permission that included `s3:PutObjectAcl`,
	// which could be used to grant read/write object access to IAM principals in other accounts.
	// If you want to get rid of that behavior, update your CDK version to 1.85.0 or later,
	// and make sure the `@aws-cdk/aws-s3:grantWriteWithoutAcl` feature flag is set to `true`
	// in the `context` key of your cdk.json file.
	// If you've already updated, but still need the principal to have permissions to modify the ACLs,
	// use the {@link grantPutAcl} method.
	GrantWrite(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant
	// Defines a CloudWatch event that triggers when something happens to this bucket.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailEvent(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object is uploaded to the specified paths (keys) in this bucket using the PutObject API call.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using `onCloudTrailWriteObject` may be preferable.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailPutObject(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule
	// Defines an AWS CloudWatch event that triggers when an object at the specified paths (keys) in this bucket are written to.
	//
	// This includes
	// the events PutObject, CopyObject, and CompleteMultipartUpload.
	//
	// Note that some tools like `aws s3 cp` will automatically use either
	// PutObject or the multipart upload API depending on the file size,
	// so using this method may be preferable to `onCloudTrailPutObject`.
	//
	// Requires that there exists at least one CloudTrail Trail in your account
	// that captures the event. This method will not create the Trail.
	OnCloudTrailWriteObject(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule
	// The S3 URL of an S3 object.
	//
	// For example:
	// - `s3://onlybucket`
	// - `s3://bucket/key`.
	S3UrlForObject(_key *string) *string
	// Returns a string representation of this construct.
	ToString() *string
	// The https Transfer Acceleration URL of an S3 object.
	//
	// Specify `dualStack: true` at the options
	// for dual-stack endpoint (connect to the bucket over IPv6). For example:
	//
	// - `https://bucket.s3-accelerate.amazonaws.com`
	// - `https://bucket.s3-accelerate.amazonaws.com/key`
	TransferAccelerationUrlForObject(_key *string, _options *awss3.TransferAccelerationUrlOptions) *string
	// The https URL of an S3 object. For example:.
	//
	// - `https://s3.us-west-1.amazonaws.com/onlybucket`
	// - `https://s3.us-west-1.amazonaws.com/bucket/key`
	// - `https://s3.cn-north-1.amazonaws.com.cn/china-bucket/mykey`
	UrlForObject(_key *string) *string
	// The virtual hosted-style URL of an S3 object. Specify `regional: false` at the options for non-regional URL. For example:.
	//
	// - `https://only-bucket.s3.us-west-1.amazonaws.com`
	// - `https://bucket.s3.us-west-1.amazonaws.com/key`
	// - `https://bucket.s3.amazonaws.com/key`
	// - `https://china-bucket.s3.cn-north-1.amazonaws.com.cn/mykey`
	VirtualHostedUrlForObject(_key *string, _options *awss3.VirtualHostedStyleUrlOptions) *string
}

// The jsii proxy struct for S3AccessLogsBucket
type jsiiProxy_S3AccessLogsBucket struct {
	jsiiProxy_RawBucket
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketDualStackDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketDualStackDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketRegionalDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionalDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketWebsiteDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) BucketWebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketWebsiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) CreateQueries() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"createQueries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Database() glue.Database {
	var returns glue.Database
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) EncryptionKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) FriendlyQueryNames() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"friendlyQueryNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) IsWebsite() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isWebsite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Policy() awss3.BucketPolicy {
	var returns awss3.BucketPolicy
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Resource() awss3.CfnBucket {
	var returns awss3.CfnBucket
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3AccessLogsBucket) Table() gluetables.S3AccessLogsTable {
	var returns gluetables.S3AccessLogsTable
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}


// Creates a new instance of the S3AccessLogsBucket class.
func NewS3AccessLogsBucket(scope constructs.Construct, id *string, props *S3AccessLogsBucketProps) S3AccessLogsBucket {
	_init_.Initialize()

	if err := validateNewS3AccessLogsBucketParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3AccessLogsBucket{}

	_jsii_.Create(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Creates a new instance of the S3AccessLogsBucket class.
func NewS3AccessLogsBucket_Override(s S3AccessLogsBucket, scope constructs.Construct, id *string, props *S3AccessLogsBucketProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_S3AccessLogsBucket)SetPolicy(val awss3.BucketPolicy) {
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func S3AccessLogsBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateS3AccessLogsBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func S3AccessLogsBucket_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateS3AccessLogsBucket_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func S3AccessLogsBucket_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateS3AccessLogsBucket_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) AddEventNotification(_event awss3.EventType, _dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter) {
	if err := s.validateAddEventNotificationParameters(_event, _dest, &_filters); err != nil {
		panic(err)
	}
	args := []interface{}{_event, _dest}
	for _, a := range _filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addEventNotification",
		args,
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) AddLoggingAspect(scope constructs.IConstruct, options *LoggingAspectOptions) {
	if err := s.validateAddLoggingAspectParameters(scope, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addLoggingAspect",
		[]interface{}{scope, options},
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) AddObjectCreatedNotification(_dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter) {
	if err := s.validateAddObjectCreatedNotificationParameters(_dest, &_filters); err != nil {
		panic(err)
	}
	args := []interface{}{_dest}
	for _, a := range _filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addObjectCreatedNotification",
		args,
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) AddObjectRemovedNotification(_dest awss3.IBucketNotificationDestination, _filters ...*awss3.NotificationKeyFilter) {
	if err := s.validateAddObjectRemovedNotificationParameters(_dest, &_filters); err != nil {
		panic(err)
	}
	args := []interface{}{_dest}
	for _, a := range _filters {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"addObjectRemovedNotification",
		args,
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) AddToResourcePolicy(permission awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := s.validateAddToResourcePolicyParameters(permission); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		s,
		"addToResourcePolicy",
		[]interface{}{permission},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) ArnForObjects(_keyPattern *string) *string {
	if err := s.validateArnForObjectsParameters(_keyPattern); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"arnForObjects",
		[]interface{}{_keyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) EnableEventBridgeNotification() {
	_jsii_.InvokeVoid(
		s,
		"enableEventBridgeNotification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3AccessLogsBucket) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantDelete(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant {
	if err := s.validateGrantDeleteParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantDelete",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantPublicAccess(_keyPrefix *string, _allowedActions ...*string) awsiam.Grant {
	args := []interface{}{_keyPrefix}
	for _, a := range _allowedActions {
		args = append(args, a)
	}

	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantPublicAccess",
		args,
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantPut(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant {
	if err := s.validateGrantPutParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantPut",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantPutAcl(_identity awsiam.IGrantable, _objectsKeyPattern *string) awsiam.Grant {
	if err := s.validateGrantPutAclParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantPutAcl",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantRead(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant {
	if err := s.validateGrantReadParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantRead",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantReadWrite(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant {
	if err := s.validateGrantReadWriteParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantReadWrite",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) GrantWrite(_identity awsiam.IGrantable, _objectsKeyPattern interface{}) awsiam.Grant {
	if err := s.validateGrantWriteParameters(_identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		s,
		"grantWrite",
		[]interface{}{_identity, _objectsKeyPattern},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) OnCloudTrailEvent(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := s.validateOnCloudTrailEventParameters(_id, _options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onCloudTrailEvent",
		[]interface{}{_id, _options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) OnCloudTrailPutObject(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := s.validateOnCloudTrailPutObjectParameters(_id, _options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onCloudTrailPutObject",
		[]interface{}{_id, _options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) OnCloudTrailWriteObject(_id *string, _options *awss3.OnCloudTrailBucketEventOptions) awsevents.Rule {
	if err := s.validateOnCloudTrailWriteObjectParameters(_id, _options); err != nil {
		panic(err)
	}
	var returns awsevents.Rule

	_jsii_.Invoke(
		s,
		"onCloudTrailWriteObject",
		[]interface{}{_id, _options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) S3UrlForObject(_key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"s3UrlForObject",
		[]interface{}{_key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) TransferAccelerationUrlForObject(_key *string, _options *awss3.TransferAccelerationUrlOptions) *string {
	if err := s.validateTransferAccelerationUrlForObjectParameters(_options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"transferAccelerationUrlForObject",
		[]interface{}{_key, _options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) UrlForObject(_key *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"urlForObject",
		[]interface{}{_key},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3AccessLogsBucket) VirtualHostedUrlForObject(_key *string, _options *awss3.VirtualHostedStyleUrlOptions) *string {
	if err := s.validateVirtualHostedUrlForObjectParameters(_options); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"virtualHostedUrlForObject",
		[]interface{}{_key, _options},
		&returns,
	)

	return returns
}

