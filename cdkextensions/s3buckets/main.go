package s3buckets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.AlbLogsBucket",
		reflect.TypeOf((*AlbLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_AlbLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.AlbLogsBucketProps",
		reflect.TypeOf((*AlbLogsBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.CloudfrontLogsBucket",
		reflect.TypeOf((*CloudfrontLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_CloudfrontLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.CloudfrontLogsBucketProps",
		reflect.TypeOf((*CloudfrontLogsBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.CloudtrailBucket",
		reflect.TypeOf((*CloudtrailBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_CloudtrailBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.CloudtrailBucketProps",
		reflect.TypeOf((*CloudtrailBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.FlowLogsBucket",
		reflect.TypeOf((*FlowLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "crawler", GoGetter: "Crawler"},
			_jsii_.MemberProperty{JsiiProperty: "crawlerSchedule", GoGetter: "CrawlerSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_FlowLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.FlowLogsBucketProps",
		reflect.TypeOf((*FlowLogsBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.LoggingAspectOptions",
		reflect.TypeOf((*LoggingAspectOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.RawBucket",
		reflect.TypeOf((*RawBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
		},
		func() interface{} {
			j := jsiiProxy_RawBucket{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.Type__awss3IBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.RawBucketProps",
		reflect.TypeOf((*RawBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.S3AccessLogsBucket",
		reflect.TypeOf((*S3AccessLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addLoggingAspect", GoMethod: "AddLoggingAspect"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_S3AccessLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.S3AccessLogsBucketProps",
		reflect.TypeOf((*S3AccessLogsBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.SesLogsBucket",
		reflect.TypeOf((*SesLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_SesLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.SesLogsBucketProps",
		reflect.TypeOf((*SesLogsBucketProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-extensions.s3_buckets.WafLogsBucket",
		reflect.TypeOf((*WafLogsBucket)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addEventNotification", GoMethod: "AddEventNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectCreatedNotification", GoMethod: "AddObjectCreatedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addObjectRemovedNotification", GoMethod: "AddObjectRemovedNotification"},
			_jsii_.MemberMethod{JsiiMethod: "addToResourcePolicy", GoMethod: "AddToResourcePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "arnForObjects", GoMethod: "ArnForObjects"},
			_jsii_.MemberProperty{JsiiProperty: "bucketArn", GoGetter: "BucketArn"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDomainName", GoGetter: "BucketDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketDualStackDomainName", GoGetter: "BucketDualStackDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketName", GoGetter: "BucketName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketRegionalDomainName", GoGetter: "BucketRegionalDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteDomainName", GoGetter: "BucketWebsiteDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "bucketWebsiteUrl", GoGetter: "BucketWebsiteUrl"},
			_jsii_.MemberProperty{JsiiProperty: "createQueries", GoGetter: "CreateQueries"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberMethod{JsiiMethod: "enableEventBridgeNotification", GoMethod: "EnableEventBridgeNotification"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionKey", GoGetter: "EncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyQueryNames", GoGetter: "FriendlyQueryNames"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "grantDelete", GoMethod: "GrantDelete"},
			_jsii_.MemberMethod{JsiiMethod: "grantPublicAccess", GoMethod: "GrantPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "grantPut", GoMethod: "GrantPut"},
			_jsii_.MemberMethod{JsiiMethod: "grantPutAcl", GoMethod: "GrantPutAcl"},
			_jsii_.MemberMethod{JsiiMethod: "grantRead", GoMethod: "GrantRead"},
			_jsii_.MemberMethod{JsiiMethod: "grantReadWrite", GoMethod: "GrantReadWrite"},
			_jsii_.MemberMethod{JsiiMethod: "grantWrite", GoMethod: "GrantWrite"},
			_jsii_.MemberProperty{JsiiProperty: "isWebsite", GoGetter: "IsWebsite"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailEvent", GoMethod: "OnCloudTrailEvent"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailPutObject", GoMethod: "OnCloudTrailPutObject"},
			_jsii_.MemberMethod{JsiiMethod: "onCloudTrailWriteObject", GoMethod: "OnCloudTrailWriteObject"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "resource", GoGetter: "Resource"},
			_jsii_.MemberMethod{JsiiMethod: "s3UrlForObject", GoMethod: "S3UrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "transferAccelerationUrlForObject", GoMethod: "TransferAccelerationUrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "urlForObject", GoMethod: "UrlForObject"},
			_jsii_.MemberMethod{JsiiMethod: "virtualHostedUrlForObject", GoMethod: "VirtualHostedUrlForObject"},
			_jsii_.MemberProperty{JsiiProperty: "workGroup", GoGetter: "WorkGroup"},
		},
		func() interface{} {
			j := jsiiProxy_WafLogsBucket{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RawBucket)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-extensions.s3_buckets.WafLogsBucketProps",
		reflect.TypeOf((*WafLogsBucketProps)(nil)).Elem(),
	)
}
