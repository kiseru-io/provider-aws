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

package dbcluster

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/neptune"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/neptune/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeDBClustersInput returns input for read
// operation.
func GenerateDescribeDBClustersInput(cr *svcapitypes.DBCluster) *svcsdk.DescribeDBClustersInput {
	res := &svcsdk.DescribeDBClustersInput{}

	if cr.Status.AtProvider.DBClusterIdentifier != nil {
		res.SetDBClusterIdentifier(*cr.Status.AtProvider.DBClusterIdentifier)
	}

	return res
}

// GenerateDBCluster returns the current state in the form of *svcapitypes.DBCluster.
func GenerateDBCluster(resp *svcsdk.DescribeDBClustersOutput) *svcapitypes.DBCluster {
	cr := &svcapitypes.DBCluster{}

	found := false
	for _, elem := range resp.DBClusters {
		if elem.AllocatedStorage != nil {
			cr.Status.AtProvider.AllocatedStorage = elem.AllocatedStorage
		} else {
			cr.Status.AtProvider.AllocatedStorage = nil
		}
		if elem.AssociatedRoles != nil {
			f1 := []*svcapitypes.DBClusterRole{}
			for _, f1iter := range elem.AssociatedRoles {
				f1elem := &svcapitypes.DBClusterRole{}
				if f1iter.FeatureName != nil {
					f1elem.FeatureName = f1iter.FeatureName
				}
				if f1iter.RoleArn != nil {
					f1elem.RoleARN = f1iter.RoleArn
				}
				if f1iter.Status != nil {
					f1elem.Status = f1iter.Status
				}
				f1 = append(f1, f1elem)
			}
			cr.Status.AtProvider.AssociatedRoles = f1
		} else {
			cr.Status.AtProvider.AssociatedRoles = nil
		}
		if elem.AutomaticRestartTime != nil {
			cr.Status.AtProvider.AutomaticRestartTime = &metav1.Time{*elem.AutomaticRestartTime}
		} else {
			cr.Status.AtProvider.AutomaticRestartTime = nil
		}
		if elem.AvailabilityZones != nil {
			f3 := []*string{}
			for _, f3iter := range elem.AvailabilityZones {
				var f3elem string
				f3elem = *f3iter
				f3 = append(f3, &f3elem)
			}
			cr.Spec.ForProvider.AvailabilityZones = f3
		} else {
			cr.Spec.ForProvider.AvailabilityZones = nil
		}
		if elem.BackupRetentionPeriod != nil {
			cr.Spec.ForProvider.BackupRetentionPeriod = elem.BackupRetentionPeriod
		} else {
			cr.Spec.ForProvider.BackupRetentionPeriod = nil
		}
		if elem.CharacterSetName != nil {
			cr.Spec.ForProvider.CharacterSetName = elem.CharacterSetName
		} else {
			cr.Spec.ForProvider.CharacterSetName = nil
		}
		if elem.CloneGroupId != nil {
			cr.Status.AtProvider.CloneGroupID = elem.CloneGroupId
		} else {
			cr.Status.AtProvider.CloneGroupID = nil
		}
		if elem.ClusterCreateTime != nil {
			cr.Status.AtProvider.ClusterCreateTime = &metav1.Time{*elem.ClusterCreateTime}
		} else {
			cr.Status.AtProvider.ClusterCreateTime = nil
		}
		if elem.CopyTagsToSnapshot != nil {
			cr.Spec.ForProvider.CopyTagsToSnapshot = elem.CopyTagsToSnapshot
		} else {
			cr.Spec.ForProvider.CopyTagsToSnapshot = nil
		}
		if elem.CrossAccountClone != nil {
			cr.Status.AtProvider.CrossAccountClone = elem.CrossAccountClone
		} else {
			cr.Status.AtProvider.CrossAccountClone = nil
		}
		if elem.DBClusterArn != nil {
			cr.Status.AtProvider.DBClusterARN = elem.DBClusterArn
		} else {
			cr.Status.AtProvider.DBClusterARN = nil
		}
		if elem.DBClusterIdentifier != nil {
			cr.Status.AtProvider.DBClusterIdentifier = elem.DBClusterIdentifier
		} else {
			cr.Status.AtProvider.DBClusterIdentifier = nil
		}
		if elem.DBClusterMembers != nil {
			f12 := []*svcapitypes.DBClusterMember{}
			for _, f12iter := range elem.DBClusterMembers {
				f12elem := &svcapitypes.DBClusterMember{}
				if f12iter.DBClusterParameterGroupStatus != nil {
					f12elem.DBClusterParameterGroupStatus = f12iter.DBClusterParameterGroupStatus
				}
				if f12iter.DBInstanceIdentifier != nil {
					f12elem.DBInstanceIdentifier = f12iter.DBInstanceIdentifier
				}
				if f12iter.IsClusterWriter != nil {
					f12elem.IsClusterWriter = f12iter.IsClusterWriter
				}
				if f12iter.PromotionTier != nil {
					f12elem.PromotionTier = f12iter.PromotionTier
				}
				f12 = append(f12, f12elem)
			}
			cr.Status.AtProvider.DBClusterMembers = f12
		} else {
			cr.Status.AtProvider.DBClusterMembers = nil
		}
		if elem.DBClusterOptionGroupMemberships != nil {
			f13 := []*svcapitypes.DBClusterOptionGroupStatus{}
			for _, f13iter := range elem.DBClusterOptionGroupMemberships {
				f13elem := &svcapitypes.DBClusterOptionGroupStatus{}
				if f13iter.DBClusterOptionGroupName != nil {
					f13elem.DBClusterOptionGroupName = f13iter.DBClusterOptionGroupName
				}
				if f13iter.Status != nil {
					f13elem.Status = f13iter.Status
				}
				f13 = append(f13, f13elem)
			}
			cr.Status.AtProvider.DBClusterOptionGroupMemberships = f13
		} else {
			cr.Status.AtProvider.DBClusterOptionGroupMemberships = nil
		}
		if elem.DBClusterParameterGroup != nil {
			cr.Status.AtProvider.DBClusterParameterGroup = elem.DBClusterParameterGroup
		} else {
			cr.Status.AtProvider.DBClusterParameterGroup = nil
		}
		if elem.DBSubnetGroup != nil {
			cr.Status.AtProvider.DBSubnetGroup = elem.DBSubnetGroup
		} else {
			cr.Status.AtProvider.DBSubnetGroup = nil
		}
		if elem.DatabaseName != nil {
			cr.Spec.ForProvider.DatabaseName = elem.DatabaseName
		} else {
			cr.Spec.ForProvider.DatabaseName = nil
		}
		if elem.DbClusterResourceId != nil {
			cr.Status.AtProvider.DBClusterResourceID = elem.DbClusterResourceId
		} else {
			cr.Status.AtProvider.DBClusterResourceID = nil
		}
		if elem.DeletionProtection != nil {
			cr.Spec.ForProvider.DeletionProtection = elem.DeletionProtection
		} else {
			cr.Spec.ForProvider.DeletionProtection = nil
		}
		if elem.EarliestRestorableTime != nil {
			cr.Status.AtProvider.EarliestRestorableTime = &metav1.Time{*elem.EarliestRestorableTime}
		} else {
			cr.Status.AtProvider.EarliestRestorableTime = nil
		}
		if elem.EnabledCloudwatchLogsExports != nil {
			f20 := []*string{}
			for _, f20iter := range elem.EnabledCloudwatchLogsExports {
				var f20elem string
				f20elem = *f20iter
				f20 = append(f20, &f20elem)
			}
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = f20
		} else {
			cr.Status.AtProvider.EnabledCloudwatchLogsExports = nil
		}
		if elem.Endpoint != nil {
			cr.Status.AtProvider.Endpoint = elem.Endpoint
		} else {
			cr.Status.AtProvider.Endpoint = nil
		}
		if elem.Engine != nil {
			cr.Spec.ForProvider.Engine = elem.Engine
		} else {
			cr.Spec.ForProvider.Engine = nil
		}
		if elem.EngineVersion != nil {
			cr.Spec.ForProvider.EngineVersion = elem.EngineVersion
		} else {
			cr.Spec.ForProvider.EngineVersion = nil
		}
		if elem.HostedZoneId != nil {
			cr.Status.AtProvider.HostedZoneID = elem.HostedZoneId
		} else {
			cr.Status.AtProvider.HostedZoneID = nil
		}
		if elem.IAMDatabaseAuthenticationEnabled != nil {
			cr.Status.AtProvider.IAMDatabaseAuthenticationEnabled = elem.IAMDatabaseAuthenticationEnabled
		} else {
			cr.Status.AtProvider.IAMDatabaseAuthenticationEnabled = nil
		}
		if elem.KmsKeyId != nil {
			cr.Spec.ForProvider.KMSKeyID = elem.KmsKeyId
		} else {
			cr.Spec.ForProvider.KMSKeyID = nil
		}
		if elem.LatestRestorableTime != nil {
			cr.Status.AtProvider.LatestRestorableTime = &metav1.Time{*elem.LatestRestorableTime}
		} else {
			cr.Status.AtProvider.LatestRestorableTime = nil
		}
		if elem.MasterUsername != nil {
			cr.Spec.ForProvider.MasterUsername = elem.MasterUsername
		} else {
			cr.Spec.ForProvider.MasterUsername = nil
		}
		if elem.MultiAZ != nil {
			cr.Status.AtProvider.MultiAZ = elem.MultiAZ
		} else {
			cr.Status.AtProvider.MultiAZ = nil
		}
		if elem.PercentProgress != nil {
			cr.Status.AtProvider.PercentProgress = elem.PercentProgress
		} else {
			cr.Status.AtProvider.PercentProgress = nil
		}
		if elem.Port != nil {
			cr.Spec.ForProvider.Port = elem.Port
		} else {
			cr.Spec.ForProvider.Port = nil
		}
		if elem.PreferredBackupWindow != nil {
			cr.Spec.ForProvider.PreferredBackupWindow = elem.PreferredBackupWindow
		} else {
			cr.Spec.ForProvider.PreferredBackupWindow = nil
		}
		if elem.PreferredMaintenanceWindow != nil {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = elem.PreferredMaintenanceWindow
		} else {
			cr.Spec.ForProvider.PreferredMaintenanceWindow = nil
		}
		if elem.ReadReplicaIdentifiers != nil {
			f34 := []*string{}
			for _, f34iter := range elem.ReadReplicaIdentifiers {
				var f34elem string
				f34elem = *f34iter
				f34 = append(f34, &f34elem)
			}
			cr.Status.AtProvider.ReadReplicaIdentifiers = f34
		} else {
			cr.Status.AtProvider.ReadReplicaIdentifiers = nil
		}
		if elem.ReaderEndpoint != nil {
			cr.Status.AtProvider.ReaderEndpoint = elem.ReaderEndpoint
		} else {
			cr.Status.AtProvider.ReaderEndpoint = nil
		}
		if elem.ReplicationSourceIdentifier != nil {
			cr.Spec.ForProvider.ReplicationSourceIdentifier = elem.ReplicationSourceIdentifier
		} else {
			cr.Spec.ForProvider.ReplicationSourceIdentifier = nil
		}
		if elem.ServerlessV2ScalingConfiguration != nil {
			f37 := &svcapitypes.ServerlessV2ScalingConfiguration{}
			if elem.ServerlessV2ScalingConfiguration.MaxCapacity != nil {
				f37.MaxCapacity = elem.ServerlessV2ScalingConfiguration.MaxCapacity
			}
			if elem.ServerlessV2ScalingConfiguration.MinCapacity != nil {
				f37.MinCapacity = elem.ServerlessV2ScalingConfiguration.MinCapacity
			}
			cr.Spec.ForProvider.ServerlessV2ScalingConfiguration = f37
		} else {
			cr.Spec.ForProvider.ServerlessV2ScalingConfiguration = nil
		}
		if elem.Status != nil {
			cr.Status.AtProvider.Status = elem.Status
		} else {
			cr.Status.AtProvider.Status = nil
		}
		if elem.StorageEncrypted != nil {
			cr.Spec.ForProvider.StorageEncrypted = elem.StorageEncrypted
		} else {
			cr.Spec.ForProvider.StorageEncrypted = nil
		}
		if elem.VpcSecurityGroups != nil {
			f40 := []*svcapitypes.VPCSecurityGroupMembership{}
			for _, f40iter := range elem.VpcSecurityGroups {
				f40elem := &svcapitypes.VPCSecurityGroupMembership{}
				if f40iter.Status != nil {
					f40elem.Status = f40iter.Status
				}
				if f40iter.VpcSecurityGroupId != nil {
					f40elem.VPCSecurityGroupID = f40iter.VpcSecurityGroupId
				}
				f40 = append(f40, f40elem)
			}
			cr.Status.AtProvider.VPCSecurityGroups = f40
		} else {
			cr.Status.AtProvider.VPCSecurityGroups = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateDBClusterInput returns a create input.
func GenerateCreateDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.CreateDBClusterInput {
	res := &svcsdk.CreateDBClusterInput{}

	if cr.Spec.ForProvider.AvailabilityZones != nil {
		f0 := []*string{}
		for _, f0iter := range cr.Spec.ForProvider.AvailabilityZones {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetAvailabilityZones(f0)
	}
	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.CharacterSetName != nil {
		res.SetCharacterSetName(*cr.Spec.ForProvider.CharacterSetName)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DBSubnetGroupName != nil {
		res.SetDBSubnetGroupName(*cr.Spec.ForProvider.DBSubnetGroupName)
	}
	if cr.Spec.ForProvider.DatabaseName != nil {
		res.SetDatabaseName(*cr.Spec.ForProvider.DatabaseName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.DestinationRegion != nil {
		res.SetDestinationRegion(*cr.Spec.ForProvider.DestinationRegion)
	}
	if cr.Spec.ForProvider.EnableCloudwatchLogsExports != nil {
		f9 := []*string{}
		for _, f9iter := range cr.Spec.ForProvider.EnableCloudwatchLogsExports {
			var f9elem string
			f9elem = *f9iter
			f9 = append(f9, &f9elem)
		}
		res.SetEnableCloudwatchLogsExports(f9)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.Engine != nil {
		res.SetEngine(*cr.Spec.ForProvider.Engine)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*cr.Spec.ForProvider.GlobalClusterIdentifier)
	}
	if cr.Spec.ForProvider.KMSKeyID != nil {
		res.SetKmsKeyId(*cr.Spec.ForProvider.KMSKeyID)
	}
	if cr.Spec.ForProvider.MasterUserPassword != nil {
		res.SetMasterUserPassword(*cr.Spec.ForProvider.MasterUserPassword)
	}
	if cr.Spec.ForProvider.MasterUsername != nil {
		res.SetMasterUsername(*cr.Spec.ForProvider.MasterUsername)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreSignedURL != nil {
		res.SetPreSignedUrl(*cr.Spec.ForProvider.PreSignedURL)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ReplicationSourceIdentifier != nil {
		res.SetReplicationSourceIdentifier(*cr.Spec.ForProvider.ReplicationSourceIdentifier)
	}
	if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration != nil {
		f23 := &svcsdk.ServerlessV2ScalingConfiguration{}
		if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MaxCapacity != nil {
			f23.SetMaxCapacity(*cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MaxCapacity)
		}
		if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MinCapacity != nil {
			f23.SetMinCapacity(*cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MinCapacity)
		}
		res.SetServerlessV2ScalingConfiguration(f23)
	}
	if cr.Spec.ForProvider.SourceRegion != nil {
		res.SetSourceRegion(*cr.Spec.ForProvider.SourceRegion)
	}
	if cr.Spec.ForProvider.StorageEncrypted != nil {
		res.SetStorageEncrypted(*cr.Spec.ForProvider.StorageEncrypted)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f26 := []*svcsdk.Tag{}
		for _, f26iter := range cr.Spec.ForProvider.Tags {
			f26elem := &svcsdk.Tag{}
			if f26iter.Key != nil {
				f26elem.SetKey(*f26iter.Key)
			}
			if f26iter.Value != nil {
				f26elem.SetValue(*f26iter.Value)
			}
			f26 = append(f26, f26elem)
		}
		res.SetTags(f26)
	}
	if cr.Spec.ForProvider.VPCSecurityGroupIDs != nil {
		f27 := []*string{}
		for _, f27iter := range cr.Spec.ForProvider.VPCSecurityGroupIDs {
			var f27elem string
			f27elem = *f27iter
			f27 = append(f27, &f27elem)
		}
		res.SetVpcSecurityGroupIds(f27)
	}

	return res
}

// GenerateModifyDBClusterInput returns an update input.
func GenerateModifyDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.ModifyDBClusterInput {
	res := &svcsdk.ModifyDBClusterInput{}

	if cr.Spec.ForProvider.BackupRetentionPeriod != nil {
		res.SetBackupRetentionPeriod(*cr.Spec.ForProvider.BackupRetentionPeriod)
	}
	if cr.Spec.ForProvider.CopyTagsToSnapshot != nil {
		res.SetCopyTagsToSnapshot(*cr.Spec.ForProvider.CopyTagsToSnapshot)
	}
	if cr.Spec.ForProvider.DBClusterParameterGroupName != nil {
		res.SetDBClusterParameterGroupName(*cr.Spec.ForProvider.DBClusterParameterGroupName)
	}
	if cr.Spec.ForProvider.DeletionProtection != nil {
		res.SetDeletionProtection(*cr.Spec.ForProvider.DeletionProtection)
	}
	if cr.Spec.ForProvider.EnableIAMDatabaseAuthentication != nil {
		res.SetEnableIAMDatabaseAuthentication(*cr.Spec.ForProvider.EnableIAMDatabaseAuthentication)
	}
	if cr.Spec.ForProvider.EngineVersion != nil {
		res.SetEngineVersion(*cr.Spec.ForProvider.EngineVersion)
	}
	if cr.Spec.ForProvider.MasterUserPassword != nil {
		res.SetMasterUserPassword(*cr.Spec.ForProvider.MasterUserPassword)
	}
	if cr.Spec.ForProvider.OptionGroupName != nil {
		res.SetOptionGroupName(*cr.Spec.ForProvider.OptionGroupName)
	}
	if cr.Spec.ForProvider.Port != nil {
		res.SetPort(*cr.Spec.ForProvider.Port)
	}
	if cr.Spec.ForProvider.PreferredBackupWindow != nil {
		res.SetPreferredBackupWindow(*cr.Spec.ForProvider.PreferredBackupWindow)
	}
	if cr.Spec.ForProvider.PreferredMaintenanceWindow != nil {
		res.SetPreferredMaintenanceWindow(*cr.Spec.ForProvider.PreferredMaintenanceWindow)
	}
	if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration != nil {
		f16 := &svcsdk.ServerlessV2ScalingConfiguration{}
		if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MaxCapacity != nil {
			f16.SetMaxCapacity(*cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MaxCapacity)
		}
		if cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MinCapacity != nil {
			f16.SetMinCapacity(*cr.Spec.ForProvider.ServerlessV2ScalingConfiguration.MinCapacity)
		}
		res.SetServerlessV2ScalingConfiguration(f16)
	}
	if cr.Spec.ForProvider.VPCSecurityGroupIDs != nil {
		f17 := []*string{}
		for _, f17iter := range cr.Spec.ForProvider.VPCSecurityGroupIDs {
			var f17elem string
			f17elem = *f17iter
			f17 = append(f17, &f17elem)
		}
		res.SetVpcSecurityGroupIds(f17)
	}

	return res
}

// GenerateDeleteDBClusterInput returns a deletion input.
func GenerateDeleteDBClusterInput(cr *svcapitypes.DBCluster) *svcsdk.DeleteDBClusterInput {
	res := &svcsdk.DeleteDBClusterInput{}

	if cr.Status.AtProvider.DBClusterIdentifier != nil {
		res.SetDBClusterIdentifier(*cr.Status.AtProvider.DBClusterIdentifier)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "DBClusterNotFoundFault"
}
