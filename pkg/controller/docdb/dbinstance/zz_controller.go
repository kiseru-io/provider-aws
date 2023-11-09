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

package dbinstance

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/docdb"
	svcsdk "github.com/aws/aws-sdk-go/service/docdb"
	svcsdkapi "github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/docdb/v1alpha1"
	connectaws "github.com/crossplane-contrib/provider-aws/pkg/utils/connect/aws"
	errorutils "github.com/crossplane-contrib/provider-aws/pkg/utils/errors"
)

const (
	errUnexpectedObject = "managed resource is not an DBInstance resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create DBInstance in AWS"
	errUpdate        = "cannot update DBInstance in AWS"
	errDescribe      = "failed to describe DBInstance"
	errDelete        = "failed to delete DBInstance"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.DBInstance)
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
	cr, ok := mg.(*svcapitypes.DBInstance)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeDBInstancesInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeDBInstancesWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.DBInstances) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateDBInstance(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)
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
	cr, ok := mg.(*svcapitypes.DBInstance)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateDBInstanceInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateDBInstanceWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, errorutils.Wrap(err, errCreate)
	}

	if resp.DBInstance.AutoMinorVersionUpgrade != nil {
		cr.Spec.ForProvider.AutoMinorVersionUpgrade = resp.DBInstance.AutoMinorVersionUpgrade
	} else {
		cr.Spec.ForProvider.AutoMinorVersionUpgrade = nil
	}
	if resp.DBInstance.AvailabilityZone != nil {
		cr.Spec.ForProvider.AvailabilityZone = resp.DBInstance.AvailabilityZone
	} else {
		cr.Spec.ForProvider.AvailabilityZone = nil
	}
	if resp.DBInstance.BackupRetentionPeriod != nil {
		cr.Status.AtProvider.BackupRetentionPeriod = resp.DBInstance.BackupRetentionPeriod
	} else {
		cr.Status.AtProvider.BackupRetentionPeriod = nil
	}
	if resp.DBInstance.CACertificateIdentifier != nil {
		cr.Status.AtProvider.CACertificateIdentifier = resp.DBInstance.CACertificateIdentifier
	} else {
		cr.Status.AtProvider.CACertificateIdentifier = nil
	}
	if resp.DBInstance.CopyTagsToSnapshot != nil {
		cr.Spec.ForProvider.CopyTagsToSnapshot = resp.DBInstance.CopyTagsToSnapshot
	} else {
		cr.Spec.ForProvider.CopyTagsToSnapshot = nil
	}
	if resp.DBInstance.DBClusterIdentifier != nil {
		cr.Status.AtProvider.DBClusterIdentifier = resp.DBInstance.DBClusterIdentifier
	} else {
		cr.Status.AtProvider.DBClusterIdentifier = nil
	}
	if resp.DBInstance.DBInstanceArn != nil {
		cr.Status.AtProvider.DBInstanceARN = resp.DBInstance.DBInstanceArn
	} else {
		cr.Status.AtProvider.DBInstanceARN = nil
	}
	if resp.DBInstance.DBInstanceClass != nil {
		cr.Spec.ForProvider.DBInstanceClass = resp.DBInstance.DBInstanceClass
	} else {
		cr.Spec.ForProvider.DBInstanceClass = nil
	}
	if resp.DBInstance.DBInstanceIdentifier != nil {
		cr.Status.AtProvider.DBInstanceIdentifier = resp.DBInstance.DBInstanceIdentifier
	} else {
		cr.Status.AtProvider.DBInstanceIdentifier = nil
	}
	if resp.DBInstance.DBInstanceStatus != nil {
		cr.Status.AtProvider.DBInstanceStatus = resp.DBInstance.DBInstanceStatus
	} else {
		cr.Status.AtProvider.DBInstanceStatus = nil
	}
	if resp.DBInstance.DBSubnetGroup != nil {
		f10 := &svcapitypes.DBSubnetGroup_SDK{}
		if resp.DBInstance.DBSubnetGroup.DBSubnetGroupArn != nil {
			f10.DBSubnetGroupARN = resp.DBInstance.DBSubnetGroup.DBSubnetGroupArn
		}
		if resp.DBInstance.DBSubnetGroup.DBSubnetGroupDescription != nil {
			f10.DBSubnetGroupDescription = resp.DBInstance.DBSubnetGroup.DBSubnetGroupDescription
		}
		if resp.DBInstance.DBSubnetGroup.DBSubnetGroupName != nil {
			f10.DBSubnetGroupName = resp.DBInstance.DBSubnetGroup.DBSubnetGroupName
		}
		if resp.DBInstance.DBSubnetGroup.SubnetGroupStatus != nil {
			f10.SubnetGroupStatus = resp.DBInstance.DBSubnetGroup.SubnetGroupStatus
		}
		if resp.DBInstance.DBSubnetGroup.Subnets != nil {
			f10f4 := []*svcapitypes.Subnet{}
			for _, f10f4iter := range resp.DBInstance.DBSubnetGroup.Subnets {
				f10f4elem := &svcapitypes.Subnet{}
				if f10f4iter.SubnetAvailabilityZone != nil {
					f10f4elemf0 := &svcapitypes.AvailabilityZone{}
					if f10f4iter.SubnetAvailabilityZone.Name != nil {
						f10f4elemf0.Name = f10f4iter.SubnetAvailabilityZone.Name
					}
					f10f4elem.SubnetAvailabilityZone = f10f4elemf0
				}
				if f10f4iter.SubnetIdentifier != nil {
					f10f4elem.SubnetIdentifier = f10f4iter.SubnetIdentifier
				}
				if f10f4iter.SubnetStatus != nil {
					f10f4elem.SubnetStatus = f10f4iter.SubnetStatus
				}
				f10f4 = append(f10f4, f10f4elem)
			}
			f10.Subnets = f10f4
		}
		if resp.DBInstance.DBSubnetGroup.VpcId != nil {
			f10.VPCID = resp.DBInstance.DBSubnetGroup.VpcId
		}
		cr.Status.AtProvider.DBSubnetGroup = f10
	} else {
		cr.Status.AtProvider.DBSubnetGroup = nil
	}
	if resp.DBInstance.DbiResourceId != nil {
		cr.Status.AtProvider.DBIResourceID = resp.DBInstance.DbiResourceId
	} else {
		cr.Status.AtProvider.DBIResourceID = nil
	}
	if resp.DBInstance.EnabledCloudwatchLogsExports != nil {
		f12 := []*string{}
		for _, f12iter := range resp.DBInstance.EnabledCloudwatchLogsExports {
			var f12elem string
			f12elem = *f12iter
			f12 = append(f12, &f12elem)
		}
		cr.Status.AtProvider.EnabledCloudwatchLogsExports = f12
	} else {
		cr.Status.AtProvider.EnabledCloudwatchLogsExports = nil
	}
	if resp.DBInstance.Endpoint != nil {
		f13 := &svcapitypes.Endpoint{}
		if resp.DBInstance.Endpoint.Address != nil {
			f13.Address = resp.DBInstance.Endpoint.Address
		}
		if resp.DBInstance.Endpoint.HostedZoneId != nil {
			f13.HostedZoneID = resp.DBInstance.Endpoint.HostedZoneId
		}
		if resp.DBInstance.Endpoint.Port != nil {
			f13.Port = resp.DBInstance.Endpoint.Port
		}
		cr.Status.AtProvider.Endpoint = f13
	} else {
		cr.Status.AtProvider.Endpoint = nil
	}
	if resp.DBInstance.Engine != nil {
		cr.Spec.ForProvider.Engine = resp.DBInstance.Engine
	} else {
		cr.Spec.ForProvider.Engine = nil
	}
	if resp.DBInstance.EngineVersion != nil {
		cr.Status.AtProvider.EngineVersion = resp.DBInstance.EngineVersion
	} else {
		cr.Status.AtProvider.EngineVersion = nil
	}
	if resp.DBInstance.InstanceCreateTime != nil {
		cr.Status.AtProvider.InstanceCreateTime = &metav1.Time{*resp.DBInstance.InstanceCreateTime}
	} else {
		cr.Status.AtProvider.InstanceCreateTime = nil
	}
	if resp.DBInstance.KmsKeyId != nil {
		cr.Status.AtProvider.KMSKeyID = resp.DBInstance.KmsKeyId
	} else {
		cr.Status.AtProvider.KMSKeyID = nil
	}
	if resp.DBInstance.LatestRestorableTime != nil {
		cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*resp.DBInstance.LatestRestorableTime}
	} else {
		cr.Status.AtProvider.LatestRestorableTime = nil
	}
	if resp.DBInstance.PendingModifiedValues != nil {
		f19 := &svcapitypes.PendingModifiedValues{}
		if resp.DBInstance.PendingModifiedValues.AllocatedStorage != nil {
			f19.AllocatedStorage = resp.DBInstance.PendingModifiedValues.AllocatedStorage
		}
		if resp.DBInstance.PendingModifiedValues.BackupRetentionPeriod != nil {
			f19.BackupRetentionPeriod = resp.DBInstance.PendingModifiedValues.BackupRetentionPeriod
		}
		if resp.DBInstance.PendingModifiedValues.CACertificateIdentifier != nil {
			f19.CACertificateIdentifier = resp.DBInstance.PendingModifiedValues.CACertificateIdentifier
		}
		if resp.DBInstance.PendingModifiedValues.DBInstanceClass != nil {
			f19.DBInstanceClass = resp.DBInstance.PendingModifiedValues.DBInstanceClass
		}
		if resp.DBInstance.PendingModifiedValues.DBInstanceIdentifier != nil {
			f19.DBInstanceIdentifier = resp.DBInstance.PendingModifiedValues.DBInstanceIdentifier
		}
		if resp.DBInstance.PendingModifiedValues.DBSubnetGroupName != nil {
			f19.DBSubnetGroupName = resp.DBInstance.PendingModifiedValues.DBSubnetGroupName
		}
		if resp.DBInstance.PendingModifiedValues.EngineVersion != nil {
			f19.EngineVersion = resp.DBInstance.PendingModifiedValues.EngineVersion
		}
		if resp.DBInstance.PendingModifiedValues.Iops != nil {
			f19.IOPS = resp.DBInstance.PendingModifiedValues.Iops
		}
		if resp.DBInstance.PendingModifiedValues.LicenseModel != nil {
			f19.LicenseModel = resp.DBInstance.PendingModifiedValues.LicenseModel
		}
		if resp.DBInstance.PendingModifiedValues.MasterUserPassword != nil {
			f19.MasterUserPassword = resp.DBInstance.PendingModifiedValues.MasterUserPassword
		}
		if resp.DBInstance.PendingModifiedValues.MultiAZ != nil {
			f19.MultiAZ = resp.DBInstance.PendingModifiedValues.MultiAZ
		}
		if resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports != nil {
			f19f11 := &svcapitypes.PendingCloudwatchLogsExports{}
			if resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable != nil {
				f19f11f0 := []*string{}
				for _, f19f11f0iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToDisable {
					var f19f11f0elem string
					f19f11f0elem = *f19f11f0iter
					f19f11f0 = append(f19f11f0, &f19f11f0elem)
				}
				f19f11.LogTypesToDisable = f19f11f0
			}
			if resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable != nil {
				f19f11f1 := []*string{}
				for _, f19f11f1iter := range resp.DBInstance.PendingModifiedValues.PendingCloudwatchLogsExports.LogTypesToEnable {
					var f19f11f1elem string
					f19f11f1elem = *f19f11f1iter
					f19f11f1 = append(f19f11f1, &f19f11f1elem)
				}
				f19f11.LogTypesToEnable = f19f11f1
			}
			f19.PendingCloudwatchLogsExports = f19f11
		}
		if resp.DBInstance.PendingModifiedValues.Port != nil {
			f19.Port = resp.DBInstance.PendingModifiedValues.Port
		}
		if resp.DBInstance.PendingModifiedValues.StorageType != nil {
			f19.StorageType = resp.DBInstance.PendingModifiedValues.StorageType
		}
		cr.Status.AtProvider.PendingModifiedValues = f19
	} else {
		cr.Status.AtProvider.PendingModifiedValues = nil
	}
	if resp.DBInstance.PreferredBackupWindow != nil {
		cr.Status.AtProvider.PreferredBackupWindow = resp.DBInstance.PreferredBackupWindow
	} else {
		cr.Status.AtProvider.PreferredBackupWindow = nil
	}
	if resp.DBInstance.PreferredMaintenanceWindow != nil {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = resp.DBInstance.PreferredMaintenanceWindow
	} else {
		cr.Spec.ForProvider.PreferredMaintenanceWindow = nil
	}
	if resp.DBInstance.PromotionTier != nil {
		cr.Spec.ForProvider.PromotionTier = resp.DBInstance.PromotionTier
	} else {
		cr.Spec.ForProvider.PromotionTier = nil
	}
	if resp.DBInstance.PubliclyAccessible != nil {
		cr.Status.AtProvider.PubliclyAccessible = resp.DBInstance.PubliclyAccessible
	} else {
		cr.Status.AtProvider.PubliclyAccessible = nil
	}
	if resp.DBInstance.StatusInfos != nil {
		f24 := []*svcapitypes.DBInstanceStatusInfo{}
		for _, f24iter := range resp.DBInstance.StatusInfos {
			f24elem := &svcapitypes.DBInstanceStatusInfo{}
			if f24iter.Message != nil {
				f24elem.Message = f24iter.Message
			}
			if f24iter.Normal != nil {
				f24elem.Normal = f24iter.Normal
			}
			if f24iter.Status != nil {
				f24elem.Status = f24iter.Status
			}
			if f24iter.StatusType != nil {
				f24elem.StatusType = f24iter.StatusType
			}
			f24 = append(f24, f24elem)
		}
		cr.Status.AtProvider.StatusInfos = f24
	} else {
		cr.Status.AtProvider.StatusInfos = nil
	}
	if resp.DBInstance.StorageEncrypted != nil {
		cr.Status.AtProvider.StorageEncrypted = resp.DBInstance.StorageEncrypted
	} else {
		cr.Status.AtProvider.StorageEncrypted = nil
	}
	if resp.DBInstance.VpcSecurityGroups != nil {
		f26 := []*svcapitypes.VPCSecurityGroupMembership{}
		for _, f26iter := range resp.DBInstance.VpcSecurityGroups {
			f26elem := &svcapitypes.VPCSecurityGroupMembership{}
			if f26iter.Status != nil {
				f26elem.Status = f26iter.Status
			}
			if f26iter.VpcSecurityGroupId != nil {
				f26elem.VPCSecurityGroupID = f26iter.VpcSecurityGroupId
			}
			f26 = append(f26, f26elem)
		}
		cr.Status.AtProvider.VPCSecurityGroups = f26
	} else {
		cr.Status.AtProvider.VPCSecurityGroups = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.DBInstance)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateModifyDBInstanceInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.ModifyDBInstanceWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, errorutils.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.DBInstance)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteDBInstanceInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteDBInstanceWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, errorutils.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.DocDBAPI, opts []option) *external {
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
	client         svcsdkapi.DocDBAPI
	preObserve     func(context.Context, *svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesInput) error
	postObserve    func(context.Context, *svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesOutput) *svcsdk.DescribeDBInstancesOutput
	lateInitialize func(*svcapitypes.DBInstanceParameters, *svcsdk.DescribeDBInstancesOutput) error
	isUpToDate     func(context.Context, *svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesOutput) (bool, string, error)
	preCreate      func(context.Context, *svcapitypes.DBInstance, *svcsdk.CreateDBInstanceInput) error
	postCreate     func(context.Context, *svcapitypes.DBInstance, *svcsdk.CreateDBInstanceOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.DBInstance, *svcsdk.DeleteDBInstanceInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.DBInstance, *svcsdk.DeleteDBInstanceOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.DBInstance, *svcsdk.ModifyDBInstanceInput) error
	postUpdate     func(context.Context, *svcapitypes.DBInstance, *svcsdk.ModifyDBInstanceOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.DBInstance, _ *svcsdk.DescribeDBInstancesOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.DBInstance, list *svcsdk.DescribeDBInstancesOutput) *svcsdk.DescribeDBInstancesOutput {
	return list
}

func nopLateInitialize(*svcapitypes.DBInstanceParameters, *svcsdk.DescribeDBInstancesOutput) error {
	return nil
}
func alwaysUpToDate(context.Context, *svcapitypes.DBInstance, *svcsdk.DescribeDBInstancesOutput) (bool, string, error) {
	return true, "", nil
}

func nopPreCreate(context.Context, *svcapitypes.DBInstance, *svcsdk.CreateDBInstanceInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.DBInstance, _ *svcsdk.CreateDBInstanceOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.DBInstance, *svcsdk.DeleteDBInstanceInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.DBInstance, _ *svcsdk.DeleteDBInstanceOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.DBInstance, *svcsdk.ModifyDBInstanceInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.DBInstance, _ *svcsdk.ModifyDBInstanceOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
