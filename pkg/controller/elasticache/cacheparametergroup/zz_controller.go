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

package cacheparametergroup

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/elasticache"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	svcsdkapi "github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/elasticache/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an CacheParameterGroup resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create CacheParameterGroup in AWS"
	errUpdate        = "cannot update CacheParameterGroup in AWS"
	errDescribe      = "failed to describe CacheParameterGroup"
	errDelete        = "failed to delete CacheParameterGroup"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.CacheParameterGroup)
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
	cr, ok := mg.(*svcapitypes.CacheParameterGroup)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeCacheParameterGroupsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeCacheParameterGroupsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.CacheParameterGroups) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateCacheParameterGroup(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.CacheParameterGroup)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateCacheParameterGroupInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateCacheParameterGroupWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.CacheParameterGroup.ARN != nil {
		cr.Status.AtProvider.ARN = resp.CacheParameterGroup.ARN
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.CacheParameterGroup.CacheParameterGroupFamily != nil {
		cr.Spec.ForProvider.CacheParameterGroupFamily = resp.CacheParameterGroup.CacheParameterGroupFamily
	} else {
		cr.Spec.ForProvider.CacheParameterGroupFamily = nil
	}
	if resp.CacheParameterGroup.CacheParameterGroupName != nil {
		cr.Status.AtProvider.CacheParameterGroupName = resp.CacheParameterGroup.CacheParameterGroupName
	} else {
		cr.Status.AtProvider.CacheParameterGroupName = nil
	}
	if resp.CacheParameterGroup.Description != nil {
		cr.Spec.ForProvider.Description = resp.CacheParameterGroup.Description
	} else {
		cr.Spec.ForProvider.Description = nil
	}
	if resp.CacheParameterGroup.IsGlobal != nil {
		cr.Status.AtProvider.IsGlobal = resp.CacheParameterGroup.IsGlobal
	} else {
		cr.Status.AtProvider.IsGlobal = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.CacheParameterGroup)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyCacheParameterGroupInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyCacheParameterGroupWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.CacheParameterGroup)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteCacheParameterGroupInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteCacheParameterGroupWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.ElastiCacheAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
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
	client         svcsdkapi.ElastiCacheAPI
	preObserve     func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsInput) error
	postObserve    func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsOutput) *svcsdk.DescribeCacheParameterGroupsOutput
	lateInitialize func(*svcapitypes.CacheParameterGroupParameters, *svcsdk.DescribeCacheParameterGroupsOutput) error
	isUpToDate     func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.CreateCacheParameterGroupInput) error
	postCreate     func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.CreateCacheParameterGroupOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DeleteCacheParameterGroupInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DeleteCacheParameterGroupOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.ModifyCacheParameterGroupInput) error
	postUpdate     func(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.CacheParameterGroupNameMessage, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.CacheParameterGroup, _ *svcsdk.DescribeCacheParameterGroupsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.CacheParameterGroup, list *svcsdk.DescribeCacheParameterGroupsOutput) *svcsdk.DescribeCacheParameterGroupsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.CacheParameterGroupParameters, *svcsdk.DescribeCacheParameterGroupsOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DescribeCacheParameterGroupsOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.CreateCacheParameterGroupInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.CacheParameterGroup, _ *svcsdk.CreateCacheParameterGroupOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.DeleteCacheParameterGroupInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.CacheParameterGroup, _ *svcsdk.DeleteCacheParameterGroupOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.CacheParameterGroup, *svcsdk.ModifyCacheParameterGroupInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.CacheParameterGroup, _ *svcsdk.CacheParameterGroupNameMessage, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
