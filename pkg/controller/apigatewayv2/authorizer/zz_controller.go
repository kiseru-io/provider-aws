/*
Copyright 2021 The Crossplane Authors.

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

package authorizer

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdk "github.com/aws/aws-sdk-go/service/apigatewayv2"
	svcsdkapi "github.com/aws/aws-sdk-go/service/apigatewayv2/apigatewayv2iface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/apigatewayv2/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an Authorizer resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create Authorizer in AWS"
	errUpdate        = "cannot update Authorizer in AWS"
	errDescribe      = "failed to describe Authorizer"
	errDelete        = "failed to delete Authorizer"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.Authorizer)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := connectaws.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.Authorizer)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetAuthorizerInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetAuthorizerWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateAuthorizer(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
	if meta.WasDeleted(cr) { // There is no need to run isUpToDate if the resource is deleted
		return managed.ExternalObservation{
			ResourceExists:   true,
			ResourceUpToDate: true,
		}, nil
	}
	upToDate, diff, err := e.isUpToDate(ctx, cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		Diff:                    diff,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.Authorizer)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateAuthorizerInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateAuthorizerWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.AuthorizerCredentialsArn != nil {
		cr.Spec.ForProvider.AuthorizerCredentialsARN = resp.AuthorizerCredentialsArn
	} else {
		cr.Spec.ForProvider.AuthorizerCredentialsARN = nil
	}
	if resp.AuthorizerId != nil {
		cr.Status.AtProvider.AuthorizerID = resp.AuthorizerId
	} else {
		cr.Status.AtProvider.AuthorizerID = nil
	}
	if resp.AuthorizerPayloadFormatVersion != nil {
		cr.Spec.ForProvider.AuthorizerPayloadFormatVersion = resp.AuthorizerPayloadFormatVersion
	} else {
		cr.Spec.ForProvider.AuthorizerPayloadFormatVersion = nil
	}
	if resp.AuthorizerResultTtlInSeconds != nil {
		cr.Spec.ForProvider.AuthorizerResultTtlInSeconds = resp.AuthorizerResultTtlInSeconds
	} else {
		cr.Spec.ForProvider.AuthorizerResultTtlInSeconds = nil
	}
	if resp.AuthorizerType != nil {
		cr.Spec.ForProvider.AuthorizerType = resp.AuthorizerType
	} else {
		cr.Spec.ForProvider.AuthorizerType = nil
	}
	if resp.AuthorizerUri != nil {
		cr.Spec.ForProvider.AuthorizerURI = resp.AuthorizerUri
	} else {
		cr.Spec.ForProvider.AuthorizerURI = nil
	}
	if resp.EnableSimpleResponses != nil {
		cr.Spec.ForProvider.EnableSimpleResponses = resp.EnableSimpleResponses
	} else {
		cr.Spec.ForProvider.EnableSimpleResponses = nil
	}
	if resp.IdentitySource != nil {
		f7 := []*string{}
		for _, f7iter := range resp.IdentitySource {
			var f7elem string
			f7elem = *f7iter
			f7 = append(f7, &f7elem)
		}
		cr.Spec.ForProvider.IdentitySource = f7
	} else {
		cr.Spec.ForProvider.IdentitySource = nil
	}
	if resp.IdentityValidationExpression != nil {
		cr.Spec.ForProvider.IdentityValidationExpression = resp.IdentityValidationExpression
	} else {
		cr.Spec.ForProvider.IdentityValidationExpression = nil
	}
	if resp.JwtConfiguration != nil {
		f9 := &svcapitypes.JWTConfiguration{}
		if resp.JwtConfiguration.Audience != nil {
			f9f0 := []*string{}
			for _, f9f0iter := range resp.JwtConfiguration.Audience {
				var f9f0elem string
				f9f0elem = *f9f0iter
				f9f0 = append(f9f0, &f9f0elem)
			}
			f9.Audience = f9f0
		}
		if resp.JwtConfiguration.Issuer != nil {
			f9.Issuer = resp.JwtConfiguration.Issuer
		}
		cr.Spec.ForProvider.JWTConfiguration = f9
	} else {
		cr.Spec.ForProvider.JWTConfiguration = nil
	}
	if resp.Name != nil {
		cr.Spec.ForProvider.Name = resp.Name
	} else {
		cr.Spec.ForProvider.Name = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.Authorizer)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateAuthorizerInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateAuthorizerWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.Authorizer)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteAuthorizerInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteAuthorizerWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ApiGatewayV2API, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.ApiGatewayV2API
	preObserve     func(context.Context, *svcapitypes.Authorizer, *svcsdk.GetAuthorizerInput) error
	postObserve    func(context.Context, *svcapitypes.Authorizer, *svcsdk.GetAuthorizerOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.AuthorizerParameters, *svcsdk.GetAuthorizerOutput) error
	isUpToDate     func(context.Context, *svcapitypes.Authorizer, *svcsdk.GetAuthorizerOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.Authorizer, *svcsdk.CreateAuthorizerInput) error
	postCreate     func(context.Context, *svcapitypes.Authorizer, *svcsdk.CreateAuthorizerOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.Authorizer, *svcsdk.DeleteAuthorizerInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.Authorizer, *svcsdk.DeleteAuthorizerOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.Authorizer, *svcsdk.UpdateAuthorizerInput) error
	postUpdate     func(context.Context, *svcapitypes.Authorizer, *svcsdk.UpdateAuthorizerOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.Authorizer, *svcsdk.GetAuthorizerInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.GetAuthorizerOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.AuthorizerParameters, *svcsdk.GetAuthorizerOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.Authorizer, *svcsdk.GetAuthorizerOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.Authorizer, *svcsdk.CreateAuthorizerInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.CreateAuthorizerOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.Authorizer, *svcsdk.DeleteAuthorizerInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.DeleteAuthorizerOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.Authorizer, *svcsdk.UpdateAuthorizerInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.Authorizer, _ *svcsdk.UpdateAuthorizerOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
