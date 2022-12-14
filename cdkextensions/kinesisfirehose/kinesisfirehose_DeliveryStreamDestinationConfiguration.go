package kinesisfirehose

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
)

type DeliveryStreamDestinationConfiguration struct {
	AmazonopensearchserviceDestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_AmazonopensearchserviceDestinationConfigurationProperty `field:"optional" json:"amazonopensearchserviceDestinationConfiguration" yaml:"amazonopensearchserviceDestinationConfiguration"`
	ElasticsearchDestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_ElasticsearchDestinationConfigurationProperty `field:"optional" json:"elasticsearchDestinationConfiguration" yaml:"elasticsearchDestinationConfiguration"`
	ExtendedS3DestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_ExtendedS3DestinationConfigurationProperty `field:"optional" json:"extendedS3DestinationConfiguration" yaml:"extendedS3DestinationConfiguration"`
	HttpEndpointDestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_HttpEndpointDestinationConfigurationProperty `field:"optional" json:"httpEndpointDestinationConfiguration" yaml:"httpEndpointDestinationConfiguration"`
	RedshiftDestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_RedshiftDestinationConfigurationProperty `field:"optional" json:"redshiftDestinationConfiguration" yaml:"redshiftDestinationConfiguration"`
	S3DestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_S3DestinationConfigurationProperty `field:"optional" json:"s3DestinationConfiguration" yaml:"s3DestinationConfiguration"`
	SplunkDestinationConfiguration *awskinesisfirehose.CfnDeliveryStream_SplunkDestinationConfigurationProperty `field:"optional" json:"splunkDestinationConfiguration" yaml:"splunkDestinationConfiguration"`
}

