// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ElasticsearchDomainSpec defines the desired state of ElasticsearchDomain.
type ElasticsearchDomainSpec struct {
	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Option to allow references to indices in an HTTP request body. Must be false
	// when configuring access to individual sub-resources. By default, the value
	// is true. See Configuration Advanced Options (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Specifies advanced security options.
	AdvancedSecurityOptions *AdvancedSecurityOptionsInput `json:"advancedSecurityOptions,omitempty"`
	// Specifies Auto-Tune options.
	AutoTuneOptions *AutoTuneOptionsInput `json:"autoTuneOptions,omitempty"`
	// Options to specify the Cognito user and identity pools for Kibana authentication.
	// For more information, see Amazon Cognito Authentication for Kibana (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`
	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// The name of the Elasticsearch domain that you are creating. Domain names
	// are unique across the domains owned by an account within an AWS region. Domain
	// names must start with a lowercase letter and can contain the following characters:
	// a-z (lowercase), 0-9, and - (hyphen).
	// +kubebuilder:validation:Required
	DomainName *string `json:"domainName"`
	// Options to enable, disable and specify the type and size of EBS storage volumes.
	EBSOptions *EBSOptions `json:"ebsOptions,omitempty"`
	// Configuration options for an Elasticsearch domain. Specifies the instance
	// type and number of instances in the domain cluster.
	ElasticsearchClusterConfig *ElasticsearchClusterConfig `json:"elasticsearchClusterConfig,omitempty"`
	// String of format X.Y to specify version for the Elasticsearch domain eg.
	// "1.5" or "2.3". For more information, see Creating Elasticsearch Domains
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains)
	// in the Amazon Elasticsearch Service Developer Guide.
	ElasticsearchVersion *string `json:"elasticsearchVersion,omitempty"`
	// Specifies the Encryption At Rest Options.
	EncryptionAtRestOptions *EncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty"`
	// Map of LogType and LogPublishingOption, each containing options to publish
	// a given type of Elasticsearch log.
	LogPublishingOptions map[string]*LogPublishingOption `json:"logPublishingOptions,omitempty"`
	// Specifies the NodeToNodeEncryptionOptions.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`
	// Option to set time, in UTC format, of the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *SnapshotOptions `json:"snapshotOptions,omitempty"`
	// A list of Tag added during domain creation.
	TagList []*Tag `json:"tagList,omitempty"`
	// Options to specify the subnets and security groups for VPC endpoint. For
	// more information, see Creating a VPC (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *VPCOptions `json:"vpcOptions,omitempty"`
}

// ElasticsearchDomainStatus defines the observed state of ElasticsearchDomain
type ElasticsearchDomainStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The domain creation status. True if the creation of an Elasticsearch domain
	// is complete. False if domain creation is still in progress.
	// +kubebuilder:validation:Optional
	Created *bool `json:"created,omitempty"`
	// The domain deletion status. True if a delete request has been received for
	// the domain but resource cleanup is still in progress. False if the domain
	// has not been deleted. Once domain deletion is complete, the status of the
	// domain is no longer returned.
	// +kubebuilder:validation:Optional
	Deleted *bool `json:"deleted,omitempty"`
	// The unique identifier for the specified Elasticsearch domain.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainID,omitempty"`
	// The Elasticsearch domain endpoint that you use to submit index and search
	// requests.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty"`
	// Map containing the Elasticsearch domain endpoints used to submit index and
	// search requests. Example key, value: 'vpc','vpc-endpoint-h2dsd34efgyghrtguk5gt6j2foh4.us-east-1.es.amazonaws.com'.
	// +kubebuilder:validation:Optional
	Endpoints map[string]*string `json:"endpoints,omitempty"`
	// The status of the Elasticsearch domain configuration. True if Amazon Elasticsearch
	// Service is processing configuration changes. False if the configuration is
	// active.
	// +kubebuilder:validation:Optional
	Processing *bool `json:"processing,omitempty"`
	// The current status of the Elasticsearch domain's service software.
	// +kubebuilder:validation:Optional
	ServiceSoftwareOptions *ServiceSoftwareOptions `json:"serviceSoftwareOptions,omitempty"`
	// The status of an Elasticsearch domain version upgrade. True if Amazon Elasticsearch
	// Service is undergoing a version upgrade. False if the configuration is active.
	// +kubebuilder:validation:Optional
	UpgradeProcessing *bool `json:"upgradeProcessing,omitempty"`
}

// ElasticsearchDomain is the Schema for the ElasticsearchDomains API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type ElasticsearchDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainSpec   `json:"spec,omitempty"`
	Status            ElasticsearchDomainStatus `json:"status,omitempty"`
}

// ElasticsearchDomainList contains a list of ElasticsearchDomain
// +kubebuilder:object:root=true
type ElasticsearchDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchDomain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElasticsearchDomain{}, &ElasticsearchDomainList{})
}
