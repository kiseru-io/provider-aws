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

package responseheaderspolicy

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/cloudfront"
	svcsdk "github.com/aws/aws-sdk-go/service/cloudfront"
	svcsdkapi "github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cloudfront/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an ResponseHeadersPolicy resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create ResponseHeadersPolicy in AWS"
	errUpdate        = "cannot update ResponseHeadersPolicy in AWS"
	errDescribe      = "failed to describe ResponseHeadersPolicy"
	errDelete        = "failed to delete ResponseHeadersPolicy"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.ResponseHeadersPolicy)
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
	cr, ok := mg.(*svcapitypes.ResponseHeadersPolicy)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateGetResponseHeadersPolicyInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.GetResponseHeadersPolicyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateResponseHeadersPolicy(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.ResponseHeadersPolicy)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateResponseHeadersPolicyInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateResponseHeadersPolicyWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.ETag != nil {
		cr.Status.AtProvider.ETag = resp.ETag
	} else {
		cr.Status.AtProvider.ETag = nil
	}
	if resp.Location != nil {
		cr.Status.AtProvider.Location = resp.Location
	} else {
		cr.Status.AtProvider.Location = nil
	}
	if resp.ResponseHeadersPolicy != nil {
		f2 := &svcapitypes.ResponseHeadersPolicy_SDK{}
		if resp.ResponseHeadersPolicy.Id != nil {
			f2.ID = resp.ResponseHeadersPolicy.Id
		}
		if resp.ResponseHeadersPolicy.LastModifiedTime != nil {
			f2.LastModifiedTime = &metav1.Time{*resp.ResponseHeadersPolicy.LastModifiedTime}
		}
		if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig != nil {
			f2f2 := &svcapitypes.ResponseHeadersPolicyConfig{}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Comment != nil {
				f2f2.Comment = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Comment
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig != nil {
				f2f2f1 := &svcapitypes.ResponseHeadersPolicyCORSConfig{}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowCredentials != nil {
					f2f2f1.AccessControlAllowCredentials = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowCredentials
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowHeaders != nil {
					f2f2f1f1 := &svcapitypes.ResponseHeadersPolicyAccessControlAllowHeaders{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowHeaders.Items != nil {
						f2f2f1f1f0 := []*string{}
						for _, f2f2f1f1f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowHeaders.Items {
							var f2f2f1f1f0elem string
							f2f2f1f1f0elem = *f2f2f1f1f0iter
							f2f2f1f1f0 = append(f2f2f1f1f0, &f2f2f1f1f0elem)
						}
						f2f2f1f1.Items = f2f2f1f1f0
					}
					f2f2f1.AccessControlAllowHeaders = f2f2f1f1
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowMethods != nil {
					f2f2f1f2 := &svcapitypes.ResponseHeadersPolicyAccessControlAllowMethods{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowMethods.Items != nil {
						f2f2f1f2f0 := []*string{}
						for _, f2f2f1f2f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowMethods.Items {
							var f2f2f1f2f0elem string
							f2f2f1f2f0elem = *f2f2f1f2f0iter
							f2f2f1f2f0 = append(f2f2f1f2f0, &f2f2f1f2f0elem)
						}
						f2f2f1f2.Items = f2f2f1f2f0
					}
					f2f2f1.AccessControlAllowMethods = f2f2f1f2
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowOrigins != nil {
					f2f2f1f3 := &svcapitypes.ResponseHeadersPolicyAccessControlAllowOrigins{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowOrigins.Items != nil {
						f2f2f1f3f0 := []*string{}
						for _, f2f2f1f3f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlAllowOrigins.Items {
							var f2f2f1f3f0elem string
							f2f2f1f3f0elem = *f2f2f1f3f0iter
							f2f2f1f3f0 = append(f2f2f1f3f0, &f2f2f1f3f0elem)
						}
						f2f2f1f3.Items = f2f2f1f3f0
					}
					f2f2f1.AccessControlAllowOrigins = f2f2f1f3
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlExposeHeaders != nil {
					f2f2f1f4 := &svcapitypes.ResponseHeadersPolicyAccessControlExposeHeaders{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlExposeHeaders.Items != nil {
						f2f2f1f4f0 := []*string{}
						for _, f2f2f1f4f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlExposeHeaders.Items {
							var f2f2f1f4f0elem string
							f2f2f1f4f0elem = *f2f2f1f4f0iter
							f2f2f1f4f0 = append(f2f2f1f4f0, &f2f2f1f4f0elem)
						}
						f2f2f1f4.Items = f2f2f1f4f0
					}
					f2f2f1.AccessControlExposeHeaders = f2f2f1f4
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlMaxAgeSec != nil {
					f2f2f1.AccessControlMaxAgeSec = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.AccessControlMaxAgeSec
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.OriginOverride != nil {
					f2f2f1.OriginOverride = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CorsConfig.OriginOverride
				}
				f2f2.CORSConfig = f2f2f1
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CustomHeadersConfig != nil {
				f2f2f2 := &svcapitypes.ResponseHeadersPolicyCustomHeadersConfig{}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CustomHeadersConfig.Items != nil {
					f2f2f2f0 := []*svcapitypes.ResponseHeadersPolicyCustomHeader{}
					for _, f2f2f2f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.CustomHeadersConfig.Items {
						f2f2f2f0elem := &svcapitypes.ResponseHeadersPolicyCustomHeader{}
						if f2f2f2f0iter.Header != nil {
							f2f2f2f0elem.Header = f2f2f2f0iter.Header
						}
						if f2f2f2f0iter.Override != nil {
							f2f2f2f0elem.Override = f2f2f2f0iter.Override
						}
						if f2f2f2f0iter.Value != nil {
							f2f2f2f0elem.Value = f2f2f2f0iter.Value
						}
						f2f2f2f0 = append(f2f2f2f0, f2f2f2f0elem)
					}
					f2f2f2.Items = f2f2f2f0
				}
				f2f2.CustomHeadersConfig = f2f2f2
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name != nil {
				f2f2.Name = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.Name
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.RemoveHeadersConfig != nil {
				f2f2f4 := &svcapitypes.ResponseHeadersPolicyRemoveHeadersConfig{}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.RemoveHeadersConfig.Items != nil {
					f2f2f4f0 := []*svcapitypes.ResponseHeadersPolicyRemoveHeader{}
					for _, f2f2f4f0iter := range resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.RemoveHeadersConfig.Items {
						f2f2f4f0elem := &svcapitypes.ResponseHeadersPolicyRemoveHeader{}
						if f2f2f4f0iter.Header != nil {
							f2f2f4f0elem.Header = f2f2f4f0iter.Header
						}
						f2f2f4f0 = append(f2f2f4f0, f2f2f4f0elem)
					}
					f2f2f4.Items = f2f2f4f0
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.RemoveHeadersConfig.Quantity != nil {
					f2f2f4.Quantity = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.RemoveHeadersConfig.Quantity
				}
				f2f2.RemoveHeadersConfig = f2f2f4
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig != nil {
				f2f2f5 := &svcapitypes.ResponseHeadersPolicySecurityHeadersConfig{}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentSecurityPolicy != nil {
					f2f2f5f0 := &svcapitypes.ResponseHeadersPolicyContentSecurityPolicy{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentSecurityPolicy.ContentSecurityPolicy != nil {
						f2f2f5f0.ContentSecurityPolicy = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentSecurityPolicy.ContentSecurityPolicy
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentSecurityPolicy.Override != nil {
						f2f2f5f0.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentSecurityPolicy.Override
					}
					f2f2f5.ContentSecurityPolicy = f2f2f5f0
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentTypeOptions != nil {
					f2f2f5f1 := &svcapitypes.ResponseHeadersPolicyContentTypeOptions{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentTypeOptions.Override != nil {
						f2f2f5f1.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ContentTypeOptions.Override
					}
					f2f2f5.ContentTypeOptions = f2f2f5f1
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.FrameOptions != nil {
					f2f2f5f2 := &svcapitypes.ResponseHeadersPolicyFrameOptions{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.FrameOptions.FrameOption != nil {
						f2f2f5f2.FrameOption = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.FrameOptions.FrameOption
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.FrameOptions.Override != nil {
						f2f2f5f2.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.FrameOptions.Override
					}
					f2f2f5.FrameOptions = f2f2f5f2
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ReferrerPolicy != nil {
					f2f2f5f3 := &svcapitypes.ResponseHeadersPolicyReferrerPolicy{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ReferrerPolicy.Override != nil {
						f2f2f5f3.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ReferrerPolicy.Override
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ReferrerPolicy.ReferrerPolicy != nil {
						f2f2f5f3.ReferrerPolicy = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.ReferrerPolicy.ReferrerPolicy
					}
					f2f2f5.ReferrerPolicy = f2f2f5f3
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity != nil {
					f2f2f5f4 := &svcapitypes.ResponseHeadersPolicyStrictTransportSecurity{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.AccessControlMaxAgeSec != nil {
						f2f2f5f4.AccessControlMaxAgeSec = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.AccessControlMaxAgeSec
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.IncludeSubdomains != nil {
						f2f2f5f4.IncludeSubdomains = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.IncludeSubdomains
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.Override != nil {
						f2f2f5f4.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.Override
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.Preload != nil {
						f2f2f5f4.Preload = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.StrictTransportSecurity.Preload
					}
					f2f2f5.StrictTransportSecurity = f2f2f5f4
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection != nil {
					f2f2f5f5 := &svcapitypes.ResponseHeadersPolicyXSSProtection{}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.ModeBlock != nil {
						f2f2f5f5.ModeBlock = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.ModeBlock
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.Override != nil {
						f2f2f5f5.Override = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.Override
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.Protection != nil {
						f2f2f5f5.Protection = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.Protection
					}
					if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.ReportUri != nil {
						f2f2f5f5.ReportURI = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.SecurityHeadersConfig.XSSProtection.ReportUri
					}
					f2f2f5.XSSProtection = f2f2f5f5
				}
				f2f2.SecurityHeadersConfig = f2f2f5
			}
			if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.ServerTimingHeadersConfig != nil {
				f2f2f6 := &svcapitypes.ResponseHeadersPolicyServerTimingHeadersConfig{}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.ServerTimingHeadersConfig.Enabled != nil {
					f2f2f6.Enabled = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.ServerTimingHeadersConfig.Enabled
				}
				if resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.ServerTimingHeadersConfig.SamplingRate != nil {
					f2f2f6.SamplingRate = resp.ResponseHeadersPolicy.ResponseHeadersPolicyConfig.ServerTimingHeadersConfig.SamplingRate
				}
				f2f2.ServerTimingHeadersConfig = f2f2f6
			}
			f2.ResponseHeadersPolicyConfig = f2f2
		}
		cr.Status.AtProvider.ResponseHeadersPolicy = f2
	} else {
		cr.Status.AtProvider.ResponseHeadersPolicy = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.ResponseHeadersPolicy)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateResponseHeadersPolicyInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateResponseHeadersPolicyWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.ResponseHeadersPolicy)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteResponseHeadersPolicyInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteResponseHeadersPolicyWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.CloudFrontAPI, opts []option) *external {
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
	client         svcsdkapi.CloudFrontAPI
	preObserve     func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.GetResponseHeadersPolicyInput) error
	postObserve    func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.GetResponseHeadersPolicyOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.ResponseHeadersPolicyParameters, *svcsdk.GetResponseHeadersPolicyOutput) error
	isUpToDate     func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.GetResponseHeadersPolicyOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.CreateResponseHeadersPolicyInput) error
	postCreate     func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.CreateResponseHeadersPolicyOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.DeleteResponseHeadersPolicyInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.DeleteResponseHeadersPolicyOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.UpdateResponseHeadersPolicyInput) error
	postUpdate     func(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.UpdateResponseHeadersPolicyOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.GetResponseHeadersPolicyInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.ResponseHeadersPolicy, _ *svcsdk.GetResponseHeadersPolicyOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.ResponseHeadersPolicyParameters, *svcsdk.GetResponseHeadersPolicyOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.GetResponseHeadersPolicyOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.CreateResponseHeadersPolicyInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.ResponseHeadersPolicy, _ *svcsdk.CreateResponseHeadersPolicyOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.DeleteResponseHeadersPolicyInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.ResponseHeadersPolicy, _ *svcsdk.DeleteResponseHeadersPolicyOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.ResponseHeadersPolicy, *svcsdk.UpdateResponseHeadersPolicyInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.ResponseHeadersPolicy, _ *svcsdk.UpdateResponseHeadersPolicyOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
