/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package domainname

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"

	svcapitypes "github.com/crossplane/provider-aws/apis/apigatewayv2/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.
// TODO(muvaf): We can generate one-time boilerplate for these hooks but currently
// ACK doesn't support not generating if file exists.

// GenerateGetDomainNameInput returns input for read
// operation.
func GenerateGetDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.GetDomainNameInput {
	res := preGenerateGetDomainNameInput(cr, &svcsdk.GetDomainNameInput{})

	if cr.Status.AtProvider.DomainName != nil {
		res.SetDomainName(*cr.Status.AtProvider.DomainName)
	}

	return postGenerateGetDomainNameInput(cr, res)
}

// GenerateDomainName returns the current state in the form of *svcapitypes.DomainName.
func GenerateDomainName(resp *svcsdk.GetDomainNameOutput) *svcapitypes.DomainName {
	cr := &svcapitypes.DomainName{}

	if resp.ApiMappingSelectionExpression != nil {
		cr.Status.AtProvider.APIMappingSelectionExpression = resp.ApiMappingSelectionExpression
	}
	if resp.DomainName != nil {
		cr.Status.AtProvider.DomainName = resp.DomainName
	}

	return cr
}

// GenerateCreateDomainNameInput returns a create input.
func GenerateCreateDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.CreateDomainNameInput {
	res := preGenerateCreateDomainNameInput(cr, &svcsdk.CreateDomainNameInput{})

	if cr.Spec.ForProvider.DomainNameConfigurations != nil {
		f0 := []*svcsdk.DomainNameConfiguration{}
		for _, f0iter := range cr.Spec.ForProvider.DomainNameConfigurations {
			f0elem := &svcsdk.DomainNameConfiguration{}
			if f0iter.APIGatewayDomainName != nil {
				f0elem.SetApiGatewayDomainName(*f0iter.APIGatewayDomainName)
			}
			if f0iter.CertificateARN != nil {
				f0elem.SetCertificateArn(*f0iter.CertificateARN)
			}
			if f0iter.CertificateName != nil {
				f0elem.SetCertificateName(*f0iter.CertificateName)
			}
			if f0iter.CertificateUploadDate != nil {
				f0elem.SetCertificateUploadDate(f0iter.CertificateUploadDate.Time)
			}
			if f0iter.DomainNameStatus != nil {
				f0elem.SetDomainNameStatus(*f0iter.DomainNameStatus)
			}
			if f0iter.DomainNameStatusMessage != nil {
				f0elem.SetDomainNameStatusMessage(*f0iter.DomainNameStatusMessage)
			}
			if f0iter.EndpointType != nil {
				f0elem.SetEndpointType(*f0iter.EndpointType)
			}
			if f0iter.HostedZoneID != nil {
				f0elem.SetHostedZoneId(*f0iter.HostedZoneID)
			}
			if f0iter.SecurityPolicy != nil {
				f0elem.SetSecurityPolicy(*f0iter.SecurityPolicy)
			}
			f0 = append(f0, f0elem)
		}
		res.SetDomainNameConfigurations(f0)
	}
	if cr.Spec.ForProvider.MutualTLSAuthentication != nil {
		f1 := &svcsdk.MutualTlsAuthenticationInput{}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI != nil {
			f1.SetTruststoreUri(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreURI)
		}
		if cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion != nil {
			f1.SetTruststoreVersion(*cr.Spec.ForProvider.MutualTLSAuthentication.TruststoreVersion)
		}
		res.SetMutualTlsAuthentication(f1)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f2 := map[string]*string{}
		for f2key, f2valiter := range cr.Spec.ForProvider.Tags {
			var f2val string
			f2val = *f2valiter
			f2[f2key] = &f2val
		}
		res.SetTags(f2)
	}

	return postGenerateCreateDomainNameInput(cr, res)
}

// GenerateDeleteDomainNameInput returns a deletion input.
func GenerateDeleteDomainNameInput(cr *svcapitypes.DomainName) *svcsdk.DeleteDomainNameInput {
	res := preGenerateDeleteDomainNameInput(cr, &svcsdk.DeleteDomainNameInput{})

	return postGenerateDeleteDomainNameInput(cr, res)
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "NotFoundException"
}