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

package v1alpha1

type AggFunction string

const (
	AggFunction_avg           AggFunction = "avg"
	AggFunction_countDistinct AggFunction = "countDistinct"
	AggFunction_count         AggFunction = "count"
	AggFunction_first         AggFunction = "first"
	AggFunction_last          AggFunction = "last"
	AggFunction_kurtosis      AggFunction = "kurtosis"
	AggFunction_max           AggFunction = "max"
	AggFunction_min           AggFunction = "min"
	AggFunction_skewness      AggFunction = "skewness"
	AggFunction_stddev_samp   AggFunction = "stddev_samp"
	AggFunction_stddev_pop    AggFunction = "stddev_pop"
	AggFunction_sum           AggFunction = "sum"
	AggFunction_sumDistinct   AggFunction = "sumDistinct"
	AggFunction_var_samp      AggFunction = "var_samp"
	AggFunction_var_pop       AggFunction = "var_pop"
)

type BackfillErrorCode string

const (
	BackfillErrorCode_ENCRYPTED_PARTITION_ERROR             BackfillErrorCode = "ENCRYPTED_PARTITION_ERROR"
	BackfillErrorCode_INTERNAL_ERROR                        BackfillErrorCode = "INTERNAL_ERROR"
	BackfillErrorCode_INVALID_PARTITION_TYPE_DATA_ERROR     BackfillErrorCode = "INVALID_PARTITION_TYPE_DATA_ERROR"
	BackfillErrorCode_MISSING_PARTITION_VALUE_ERROR         BackfillErrorCode = "MISSING_PARTITION_VALUE_ERROR"
	BackfillErrorCode_UNSUPPORTED_PARTITION_CHARACTER_ERROR BackfillErrorCode = "UNSUPPORTED_PARTITION_CHARACTER_ERROR"
)

type BlueprintRunState string

const (
	BlueprintRunState_RUNNING      BlueprintRunState = "RUNNING"
	BlueprintRunState_SUCCEEDED    BlueprintRunState = "SUCCEEDED"
	BlueprintRunState_FAILED       BlueprintRunState = "FAILED"
	BlueprintRunState_ROLLING_BACK BlueprintRunState = "ROLLING_BACK"
)

type BlueprintStatus string

const (
	BlueprintStatus_CREATING BlueprintStatus = "CREATING"
	BlueprintStatus_ACTIVE   BlueprintStatus = "ACTIVE"
	BlueprintStatus_UPDATING BlueprintStatus = "UPDATING"
	BlueprintStatus_FAILED   BlueprintStatus = "FAILED"
)

type CatalogEncryptionMode string

const (
	CatalogEncryptionMode_DISABLED CatalogEncryptionMode = "DISABLED"
	CatalogEncryptionMode_SSE_KMS  CatalogEncryptionMode = "SSE-KMS"
)

type CloudWatchEncryptionMode string

const (
	CloudWatchEncryptionMode_DISABLED CloudWatchEncryptionMode = "DISABLED"
	CloudWatchEncryptionMode_SSE_KMS  CloudWatchEncryptionMode = "SSE-KMS"
)

type ColumnStatisticsType string

const (
	ColumnStatisticsType_BOOLEAN ColumnStatisticsType = "BOOLEAN"
	ColumnStatisticsType_DATE    ColumnStatisticsType = "DATE"
	ColumnStatisticsType_DECIMAL ColumnStatisticsType = "DECIMAL"
	ColumnStatisticsType_DOUBLE  ColumnStatisticsType = "DOUBLE"
	ColumnStatisticsType_LONG    ColumnStatisticsType = "LONG"
	ColumnStatisticsType_STRING  ColumnStatisticsType = "STRING"
	ColumnStatisticsType_BINARY  ColumnStatisticsType = "BINARY"
)

type Comparator string

const (
	Comparator_EQUALS              Comparator = "EQUALS"
	Comparator_GREATER_THAN        Comparator = "GREATER_THAN"
	Comparator_LESS_THAN           Comparator = "LESS_THAN"
	Comparator_GREATER_THAN_EQUALS Comparator = "GREATER_THAN_EQUALS"
	Comparator_LESS_THAN_EQUALS    Comparator = "LESS_THAN_EQUALS"
)

type Compatibility string

const (
	Compatibility_NONE         Compatibility = "NONE"
	Compatibility_DISABLED     Compatibility = "DISABLED"
	Compatibility_BACKWARD     Compatibility = "BACKWARD"
	Compatibility_BACKWARD_ALL Compatibility = "BACKWARD_ALL"
	Compatibility_FORWARD      Compatibility = "FORWARD"
	Compatibility_FORWARD_ALL  Compatibility = "FORWARD_ALL"
	Compatibility_FULL         Compatibility = "FULL"
	Compatibility_FULL_ALL     Compatibility = "FULL_ALL"
)

type CompressionType string

const (
	CompressionType_gzip  CompressionType = "gzip"
	CompressionType_bzip2 CompressionType = "bzip2"
)

type ConnectionPropertyKey string

const (
	ConnectionPropertyKey_HOST                                     ConnectionPropertyKey = "HOST"
	ConnectionPropertyKey_PORT                                     ConnectionPropertyKey = "PORT"
	ConnectionPropertyKey_USERNAME                                 ConnectionPropertyKey = "USERNAME"
	ConnectionPropertyKey_PASSWORD                                 ConnectionPropertyKey = "PASSWORD"
	ConnectionPropertyKey_ENCRYPTED_PASSWORD                       ConnectionPropertyKey = "ENCRYPTED_PASSWORD"
	ConnectionPropertyKey_JDBC_DRIVER_JAR_URI                      ConnectionPropertyKey = "JDBC_DRIVER_JAR_URI"
	ConnectionPropertyKey_JDBC_DRIVER_CLASS_NAME                   ConnectionPropertyKey = "JDBC_DRIVER_CLASS_NAME"
	ConnectionPropertyKey_JDBC_ENGINE                              ConnectionPropertyKey = "JDBC_ENGINE"
	ConnectionPropertyKey_JDBC_ENGINE_VERSION                      ConnectionPropertyKey = "JDBC_ENGINE_VERSION"
	ConnectionPropertyKey_CONFIG_FILES                             ConnectionPropertyKey = "CONFIG_FILES"
	ConnectionPropertyKey_INSTANCE_ID                              ConnectionPropertyKey = "INSTANCE_ID"
	ConnectionPropertyKey_JDBC_CONNECTION_URL                      ConnectionPropertyKey = "JDBC_CONNECTION_URL"
	ConnectionPropertyKey_JDBC_ENFORCE_SSL                         ConnectionPropertyKey = "JDBC_ENFORCE_SSL"
	ConnectionPropertyKey_CUSTOM_JDBC_CERT                         ConnectionPropertyKey = "CUSTOM_JDBC_CERT"
	ConnectionPropertyKey_SKIP_CUSTOM_JDBC_CERT_VALIDATION         ConnectionPropertyKey = "SKIP_CUSTOM_JDBC_CERT_VALIDATION"
	ConnectionPropertyKey_CUSTOM_JDBC_CERT_STRING                  ConnectionPropertyKey = "CUSTOM_JDBC_CERT_STRING"
	ConnectionPropertyKey_CONNECTION_URL                           ConnectionPropertyKey = "CONNECTION_URL"
	ConnectionPropertyKey_KAFKA_BOOTSTRAP_SERVERS                  ConnectionPropertyKey = "KAFKA_BOOTSTRAP_SERVERS"
	ConnectionPropertyKey_KAFKA_SSL_ENABLED                        ConnectionPropertyKey = "KAFKA_SSL_ENABLED"
	ConnectionPropertyKey_KAFKA_CUSTOM_CERT                        ConnectionPropertyKey = "KAFKA_CUSTOM_CERT"
	ConnectionPropertyKey_KAFKA_SKIP_CUSTOM_CERT_VALIDATION        ConnectionPropertyKey = "KAFKA_SKIP_CUSTOM_CERT_VALIDATION"
	ConnectionPropertyKey_KAFKA_CLIENT_KEYSTORE                    ConnectionPropertyKey = "KAFKA_CLIENT_KEYSTORE"
	ConnectionPropertyKey_KAFKA_CLIENT_KEYSTORE_PASSWORD           ConnectionPropertyKey = "KAFKA_CLIENT_KEYSTORE_PASSWORD"
	ConnectionPropertyKey_KAFKA_CLIENT_KEY_PASSWORD                ConnectionPropertyKey = "KAFKA_CLIENT_KEY_PASSWORD"
	ConnectionPropertyKey_ENCRYPTED_KAFKA_CLIENT_KEYSTORE_PASSWORD ConnectionPropertyKey = "ENCRYPTED_KAFKA_CLIENT_KEYSTORE_PASSWORD"
	ConnectionPropertyKey_ENCRYPTED_KAFKA_CLIENT_KEY_PASSWORD      ConnectionPropertyKey = "ENCRYPTED_KAFKA_CLIENT_KEY_PASSWORD"
	ConnectionPropertyKey_SECRET_ID                                ConnectionPropertyKey = "SECRET_ID"
	ConnectionPropertyKey_CONNECTOR_URL                            ConnectionPropertyKey = "CONNECTOR_URL"
	ConnectionPropertyKey_CONNECTOR_TYPE                           ConnectionPropertyKey = "CONNECTOR_TYPE"
	ConnectionPropertyKey_CONNECTOR_CLASS_NAME                     ConnectionPropertyKey = "CONNECTOR_CLASS_NAME"
)

type ConnectionType string

const (
	ConnectionType_JDBC        ConnectionType = "JDBC"
	ConnectionType_SFTP        ConnectionType = "SFTP"
	ConnectionType_MONGODB     ConnectionType = "MONGODB"
	ConnectionType_KAFKA       ConnectionType = "KAFKA"
	ConnectionType_NETWORK     ConnectionType = "NETWORK"
	ConnectionType_MARKETPLACE ConnectionType = "MARKETPLACE"
	ConnectionType_CUSTOM      ConnectionType = "CUSTOM"
)

type CrawlState string

const (
	CrawlState_RUNNING    CrawlState = "RUNNING"
	CrawlState_CANCELLING CrawlState = "CANCELLING"
	CrawlState_CANCELLED  CrawlState = "CANCELLED"
	CrawlState_SUCCEEDED  CrawlState = "SUCCEEDED"
	CrawlState_FAILED     CrawlState = "FAILED"
	CrawlState_ERROR      CrawlState = "ERROR"
)

type CrawlerHistoryState string

const (
	CrawlerHistoryState_RUNNING   CrawlerHistoryState = "RUNNING"
	CrawlerHistoryState_COMPLETED CrawlerHistoryState = "COMPLETED"
	CrawlerHistoryState_FAILED    CrawlerHistoryState = "FAILED"
	CrawlerHistoryState_STOPPED   CrawlerHistoryState = "STOPPED"
)

type CrawlerLineageSettings string

const (
	CrawlerLineageSettings_ENABLE  CrawlerLineageSettings = "ENABLE"
	CrawlerLineageSettings_DISABLE CrawlerLineageSettings = "DISABLE"
)

type CrawlerState string

const (
	CrawlerState_READY    CrawlerState = "READY"
	CrawlerState_RUNNING  CrawlerState = "RUNNING"
	CrawlerState_STOPPING CrawlerState = "STOPPING"
)

type CsvHeaderOption string

const (
	CsvHeaderOption_UNKNOWN CsvHeaderOption = "UNKNOWN"
	CsvHeaderOption_PRESENT CsvHeaderOption = "PRESENT"
	CsvHeaderOption_ABSENT  CsvHeaderOption = "ABSENT"
)

type DQStopJobOnFailureTiming string

const (
	DQStopJobOnFailureTiming_Immediate     DQStopJobOnFailureTiming = "Immediate"
	DQStopJobOnFailureTiming_AfterDataLoad DQStopJobOnFailureTiming = "AfterDataLoad"
)

type DQTransformOutput string

const (
	DQTransformOutput_PrimaryInput      DQTransformOutput = "PrimaryInput"
	DQTransformOutput_EvaluationResults DQTransformOutput = "EvaluationResults"
)

type DataFormat string

const (
	DataFormat_AVRO     DataFormat = "AVRO"
	DataFormat_JSON     DataFormat = "JSON"
	DataFormat_PROTOBUF DataFormat = "PROTOBUF"
)

type DataQualityRuleResultStatus string

const (
	DataQualityRuleResultStatus_PASS  DataQualityRuleResultStatus = "PASS"
	DataQualityRuleResultStatus_FAIL  DataQualityRuleResultStatus = "FAIL"
	DataQualityRuleResultStatus_ERROR DataQualityRuleResultStatus = "ERROR"
)

type DeleteBehavior string

const (
	DeleteBehavior_LOG                   DeleteBehavior = "LOG"
	DeleteBehavior_DELETE_FROM_DATABASE  DeleteBehavior = "DELETE_FROM_DATABASE"
	DeleteBehavior_DEPRECATE_IN_DATABASE DeleteBehavior = "DEPRECATE_IN_DATABASE"
)

type EnableHybridValues string

const (
	EnableHybridValues_TRUE  EnableHybridValues = "TRUE"
	EnableHybridValues_FALSE EnableHybridValues = "FALSE"
)

type ExecutionClass string

const (
	ExecutionClass_FLEX     ExecutionClass = "FLEX"
	ExecutionClass_STANDARD ExecutionClass = "STANDARD"
)

type ExistCondition string

const (
	ExistCondition_MUST_EXIST ExistCondition = "MUST_EXIST"
	ExistCondition_NOT_EXIST  ExistCondition = "NOT_EXIST"
	ExistCondition_NONE       ExistCondition = "NONE"
)

type FieldName string

const (
	FieldName_CRAWL_ID   FieldName = "CRAWL_ID"
	FieldName_STATE      FieldName = "STATE"
	FieldName_START_TIME FieldName = "START_TIME"
	FieldName_END_TIME   FieldName = "END_TIME"
	FieldName_DPU_HOUR   FieldName = "DPU_HOUR"
)

type FilterLogicalOperator string

const (
	FilterLogicalOperator_AND FilterLogicalOperator = "AND"
	FilterLogicalOperator_OR  FilterLogicalOperator = "OR"
)

type FilterOperation string

const (
	FilterOperation_EQ     FilterOperation = "EQ"
	FilterOperation_LT     FilterOperation = "LT"
	FilterOperation_GT     FilterOperation = "GT"
	FilterOperation_LTE    FilterOperation = "LTE"
	FilterOperation_GTE    FilterOperation = "GTE"
	FilterOperation_REGEX  FilterOperation = "REGEX"
	FilterOperation_ISNULL FilterOperation = "ISNULL"
)

type FilterOperator string

const (
	FilterOperator_GT FilterOperator = "GT"
	FilterOperator_GE FilterOperator = "GE"
	FilterOperator_LT FilterOperator = "LT"
	FilterOperator_LE FilterOperator = "LE"
	FilterOperator_EQ FilterOperator = "EQ"
	FilterOperator_NE FilterOperator = "NE"
)

type FilterValueType string

const (
	FilterValueType_COLUMNEXTRACTED FilterValueType = "COLUMNEXTRACTED"
	FilterValueType_CONSTANT        FilterValueType = "CONSTANT"
)

type GlueRecordType string

const (
	GlueRecordType_DATE       GlueRecordType = "DATE"
	GlueRecordType_STRING     GlueRecordType = "STRING"
	GlueRecordType_TIMESTAMP  GlueRecordType = "TIMESTAMP"
	GlueRecordType_INT        GlueRecordType = "INT"
	GlueRecordType_FLOAT      GlueRecordType = "FLOAT"
	GlueRecordType_LONG       GlueRecordType = "LONG"
	GlueRecordType_BIGDECIMAL GlueRecordType = "BIGDECIMAL"
	GlueRecordType_BYTE       GlueRecordType = "BYTE"
	GlueRecordType_SHORT      GlueRecordType = "SHORT"
	GlueRecordType_DOUBLE     GlueRecordType = "DOUBLE"
)

type JDBCDataType string

const (
	JDBCDataType_ARRAY                   JDBCDataType = "ARRAY"
	JDBCDataType_BIGINT                  JDBCDataType = "BIGINT"
	JDBCDataType_BINARY                  JDBCDataType = "BINARY"
	JDBCDataType_BIT                     JDBCDataType = "BIT"
	JDBCDataType_BLOB                    JDBCDataType = "BLOB"
	JDBCDataType_BOOLEAN                 JDBCDataType = "BOOLEAN"
	JDBCDataType_CHAR                    JDBCDataType = "CHAR"
	JDBCDataType_CLOB                    JDBCDataType = "CLOB"
	JDBCDataType_DATALINK                JDBCDataType = "DATALINK"
	JDBCDataType_DATE                    JDBCDataType = "DATE"
	JDBCDataType_DECIMAL                 JDBCDataType = "DECIMAL"
	JDBCDataType_DISTINCT                JDBCDataType = "DISTINCT"
	JDBCDataType_DOUBLE                  JDBCDataType = "DOUBLE"
	JDBCDataType_FLOAT                   JDBCDataType = "FLOAT"
	JDBCDataType_INTEGER                 JDBCDataType = "INTEGER"
	JDBCDataType_JAVA_OBJECT             JDBCDataType = "JAVA_OBJECT"
	JDBCDataType_LONGNVARCHAR            JDBCDataType = "LONGNVARCHAR"
	JDBCDataType_LONGVARBINARY           JDBCDataType = "LONGVARBINARY"
	JDBCDataType_LONGVARCHAR             JDBCDataType = "LONGVARCHAR"
	JDBCDataType_NCHAR                   JDBCDataType = "NCHAR"
	JDBCDataType_NCLOB                   JDBCDataType = "NCLOB"
	JDBCDataType_NULL                    JDBCDataType = "NULL"
	JDBCDataType_NUMERIC                 JDBCDataType = "NUMERIC"
	JDBCDataType_NVARCHAR                JDBCDataType = "NVARCHAR"
	JDBCDataType_OTHER                   JDBCDataType = "OTHER"
	JDBCDataType_REAL                    JDBCDataType = "REAL"
	JDBCDataType_REF                     JDBCDataType = "REF"
	JDBCDataType_REF_CURSOR              JDBCDataType = "REF_CURSOR"
	JDBCDataType_ROWID                   JDBCDataType = "ROWID"
	JDBCDataType_SMALLINT                JDBCDataType = "SMALLINT"
	JDBCDataType_SQLXML                  JDBCDataType = "SQLXML"
	JDBCDataType_STRUCT                  JDBCDataType = "STRUCT"
	JDBCDataType_TIME                    JDBCDataType = "TIME"
	JDBCDataType_TIME_WITH_TIMEZONE      JDBCDataType = "TIME_WITH_TIMEZONE"
	JDBCDataType_TIMESTAMP               JDBCDataType = "TIMESTAMP"
	JDBCDataType_TIMESTAMP_WITH_TIMEZONE JDBCDataType = "TIMESTAMP_WITH_TIMEZONE"
	JDBCDataType_TINYINT                 JDBCDataType = "TINYINT"
	JDBCDataType_VARBINARY               JDBCDataType = "VARBINARY"
	JDBCDataType_VARCHAR                 JDBCDataType = "VARCHAR"
)

type JdbcMetadataEntry string

const (
	JdbcMetadataEntry_COMMENTS JdbcMetadataEntry = "COMMENTS"
	JdbcMetadataEntry_RAWTYPES JdbcMetadataEntry = "RAWTYPES"
)

type JobBookmarksEncryptionMode string

const (
	JobBookmarksEncryptionMode_DISABLED JobBookmarksEncryptionMode = "DISABLED"
	JobBookmarksEncryptionMode_CSE_KMS  JobBookmarksEncryptionMode = "CSE-KMS"
)

type JobRunState string

const (
	JobRunState_STARTING  JobRunState = "STARTING"
	JobRunState_RUNNING   JobRunState = "RUNNING"
	JobRunState_STOPPING  JobRunState = "STOPPING"
	JobRunState_STOPPED   JobRunState = "STOPPED"
	JobRunState_SUCCEEDED JobRunState = "SUCCEEDED"
	JobRunState_FAILED    JobRunState = "FAILED"
	JobRunState_TIMEOUT   JobRunState = "TIMEOUT"
	JobRunState_ERROR     JobRunState = "ERROR"
	JobRunState_WAITING   JobRunState = "WAITING"
)

type JoinType string

const (
	JoinType_equijoin JoinType = "equijoin"
	JoinType_left     JoinType = "left"
	JoinType_right    JoinType = "right"
	JoinType_outer    JoinType = "outer"
	JoinType_leftsemi JoinType = "leftsemi"
	JoinType_leftanti JoinType = "leftanti"
)

type Language string

const (
	Language_PYTHON Language = "PYTHON"
	Language_SCALA  Language = "SCALA"
)

type LastCrawlStatus string

const (
	LastCrawlStatus_SUCCEEDED LastCrawlStatus = "SUCCEEDED"
	LastCrawlStatus_CANCELLED LastCrawlStatus = "CANCELLED"
	LastCrawlStatus_FAILED    LastCrawlStatus = "FAILED"
)

type Logical string

const (
	Logical_AND Logical = "AND"
	Logical_ANY Logical = "ANY"
)

type LogicalOperator string

const (
	LogicalOperator_EQUALS LogicalOperator = "EQUALS"
)

type MLUserDataEncryptionModeString string

const (
	MLUserDataEncryptionModeString_DISABLED MLUserDataEncryptionModeString = "DISABLED"
	MLUserDataEncryptionModeString_SSE_KMS  MLUserDataEncryptionModeString = "SSE-KMS"
)

type NodeType string

const (
	NodeType_CRAWLER NodeType = "CRAWLER"
	NodeType_JOB     NodeType = "JOB"
	NodeType_TRIGGER NodeType = "TRIGGER"
)

type ParamType string

const (
	ParamType_str     ParamType = "str"
	ParamType_int     ParamType = "int"
	ParamType_float   ParamType = "float"
	ParamType_complex ParamType = "complex"
	ParamType_bool    ParamType = "bool"
	ParamType_list    ParamType = "list"
	ParamType_null    ParamType = "null"
)

type ParquetCompressionType string

const (
	ParquetCompressionType_snappy       ParquetCompressionType = "snappy"
	ParquetCompressionType_lzo          ParquetCompressionType = "lzo"
	ParquetCompressionType_gzip         ParquetCompressionType = "gzip"
	ParquetCompressionType_uncompressed ParquetCompressionType = "uncompressed"
	ParquetCompressionType_none         ParquetCompressionType = "none"
)

type PartitionIndexStatus string

const (
	PartitionIndexStatus_CREATING PartitionIndexStatus = "CREATING"
	PartitionIndexStatus_ACTIVE   PartitionIndexStatus = "ACTIVE"
	PartitionIndexStatus_DELETING PartitionIndexStatus = "DELETING"
	PartitionIndexStatus_FAILED   PartitionIndexStatus = "FAILED"
)

type Permission string

const (
	Permission_ALL                  Permission = "ALL"
	Permission_SELECT               Permission = "SELECT"
	Permission_ALTER                Permission = "ALTER"
	Permission_DROP                 Permission = "DROP"
	Permission_DELETE               Permission = "DELETE"
	Permission_INSERT               Permission = "INSERT"
	Permission_CREATE_DATABASE      Permission = "CREATE_DATABASE"
	Permission_CREATE_TABLE         Permission = "CREATE_TABLE"
	Permission_DATA_LOCATION_ACCESS Permission = "DATA_LOCATION_ACCESS"
)

type PermissionType string

const (
	PermissionType_COLUMN_PERMISSION      PermissionType = "COLUMN_PERMISSION"
	PermissionType_CELL_FILTER_PERMISSION PermissionType = "CELL_FILTER_PERMISSION"
)

type PiiType string

const (
	PiiType_RowAudit      PiiType = "RowAudit"
	PiiType_RowMasking    PiiType = "RowMasking"
	PiiType_ColumnAudit   PiiType = "ColumnAudit"
	PiiType_ColumnMasking PiiType = "ColumnMasking"
)

type PrincipalType string

const (
	PrincipalType_USER  PrincipalType = "USER"
	PrincipalType_ROLE  PrincipalType = "ROLE"
	PrincipalType_GROUP PrincipalType = "GROUP"
)

type QuoteChar string

const (
	QuoteChar_quote        QuoteChar = "quote"
	QuoteChar_quillemet    QuoteChar = "quillemet"
	QuoteChar_single_quote QuoteChar = "single_quote"
	QuoteChar_disabled     QuoteChar = "disabled"
)

type RecrawlBehavior string

const (
	RecrawlBehavior_CRAWL_EVERYTHING       RecrawlBehavior = "CRAWL_EVERYTHING"
	RecrawlBehavior_CRAWL_NEW_FOLDERS_ONLY RecrawlBehavior = "CRAWL_NEW_FOLDERS_ONLY"
	RecrawlBehavior_CRAWL_EVENT_MODE       RecrawlBehavior = "CRAWL_EVENT_MODE"
)

type RegistryStatus string

const (
	RegistryStatus_AVAILABLE RegistryStatus = "AVAILABLE"
	RegistryStatus_DELETING  RegistryStatus = "DELETING"
)

type ResourceShareType string

const (
	ResourceShareType_FOREIGN ResourceShareType = "FOREIGN"
	ResourceShareType_ALL     ResourceShareType = "ALL"
)

type ResourceType string

const (
	ResourceType_JAR     ResourceType = "JAR"
	ResourceType_FILE    ResourceType = "FILE"
	ResourceType_ARCHIVE ResourceType = "ARCHIVE"
)

type S3EncryptionMode string

const (
	S3EncryptionMode_DISABLED S3EncryptionMode = "DISABLED"
	S3EncryptionMode_SSE_KMS  S3EncryptionMode = "SSE-KMS"
	S3EncryptionMode_SSE_S3   S3EncryptionMode = "SSE-S3"
)

type ScheduleState string

const (
	ScheduleState_SCHEDULED     ScheduleState = "SCHEDULED"
	ScheduleState_NOT_SCHEDULED ScheduleState = "NOT_SCHEDULED"
	ScheduleState_TRANSITIONING ScheduleState = "TRANSITIONING"
)

type SchemaDiffType string

const (
	SchemaDiffType_SYNTAX_DIFF SchemaDiffType = "SYNTAX_DIFF"
)

type SchemaStatus string

const (
	SchemaStatus_AVAILABLE SchemaStatus = "AVAILABLE"
	SchemaStatus_PENDING   SchemaStatus = "PENDING"
	SchemaStatus_DELETING  SchemaStatus = "DELETING"
)

type SchemaVersionStatus string

const (
	SchemaVersionStatus_AVAILABLE SchemaVersionStatus = "AVAILABLE"
	SchemaVersionStatus_PENDING   SchemaVersionStatus = "PENDING"
	SchemaVersionStatus_FAILURE   SchemaVersionStatus = "FAILURE"
	SchemaVersionStatus_DELETING  SchemaVersionStatus = "DELETING"
)

type Separator string

const (
	Separator_comma     Separator = "comma"
	Separator_ctrla     Separator = "ctrla"
	Separator_pipe      Separator = "pipe"
	Separator_semicolon Separator = "semicolon"
	Separator_tab       Separator = "tab"
)

type SessionStatus string

const (
	SessionStatus_PROVISIONING SessionStatus = "PROVISIONING"
	SessionStatus_READY        SessionStatus = "READY"
	SessionStatus_FAILED       SessionStatus = "FAILED"
	SessionStatus_TIMEOUT      SessionStatus = "TIMEOUT"
	SessionStatus_STOPPING     SessionStatus = "STOPPING"
	SessionStatus_STOPPED      SessionStatus = "STOPPED"
)

type Sort string

const (
	Sort_ASC  Sort = "ASC"
	Sort_DESC Sort = "DESC"
)

type SortDirectionType string

const (
	SortDirectionType_DESCENDING SortDirectionType = "DESCENDING"
	SortDirectionType_ASCENDING  SortDirectionType = "ASCENDING"
)

type SourceControlAuthStrategy string

const (
	SourceControlAuthStrategy_PERSONAL_ACCESS_TOKEN SourceControlAuthStrategy = "PERSONAL_ACCESS_TOKEN"
	SourceControlAuthStrategy_AWS_SECRETS_MANAGER   SourceControlAuthStrategy = "AWS_SECRETS_MANAGER"
)

type SourceControlProvider string

const (
	SourceControlProvider_GITHUB          SourceControlProvider = "GITHUB"
	SourceControlProvider_AWS_CODE_COMMIT SourceControlProvider = "AWS_CODE_COMMIT"
)

type StartingPosition string

const (
	StartingPosition_latest       StartingPosition = "latest"
	StartingPosition_trim_horizon StartingPosition = "trim_horizon"
	StartingPosition_earliest     StartingPosition = "earliest"
)

type StatementState string

const (
	StatementState_WAITING    StatementState = "WAITING"
	StatementState_RUNNING    StatementState = "RUNNING"
	StatementState_AVAILABLE  StatementState = "AVAILABLE"
	StatementState_CANCELLING StatementState = "CANCELLING"
	StatementState_CANCELLED  StatementState = "CANCELLED"
	StatementState_ERROR      StatementState = "ERROR"
)

type TargetFormat string

const (
	TargetFormat_json    TargetFormat = "json"
	TargetFormat_csv     TargetFormat = "csv"
	TargetFormat_avro    TargetFormat = "avro"
	TargetFormat_orc     TargetFormat = "orc"
	TargetFormat_parquet TargetFormat = "parquet"
)

type TaskRunSortColumnType string

const (
	TaskRunSortColumnType_TASK_RUN_TYPE TaskRunSortColumnType = "TASK_RUN_TYPE"
	TaskRunSortColumnType_STATUS        TaskRunSortColumnType = "STATUS"
	TaskRunSortColumnType_STARTED       TaskRunSortColumnType = "STARTED"
)

type TaskStatusType string

const (
	TaskStatusType_STARTING  TaskStatusType = "STARTING"
	TaskStatusType_RUNNING   TaskStatusType = "RUNNING"
	TaskStatusType_STOPPING  TaskStatusType = "STOPPING"
	TaskStatusType_STOPPED   TaskStatusType = "STOPPED"
	TaskStatusType_SUCCEEDED TaskStatusType = "SUCCEEDED"
	TaskStatusType_FAILED    TaskStatusType = "FAILED"
	TaskStatusType_TIMEOUT   TaskStatusType = "TIMEOUT"
)

type TaskType string

const (
	TaskType_EVALUATION              TaskType = "EVALUATION"
	TaskType_LABELING_SET_GENERATION TaskType = "LABELING_SET_GENERATION"
	TaskType_IMPORT_LABELS           TaskType = "IMPORT_LABELS"
	TaskType_EXPORT_LABELS           TaskType = "EXPORT_LABELS"
	TaskType_FIND_MATCHES            TaskType = "FIND_MATCHES"
)

type TransformSortColumnType string

const (
	TransformSortColumnType_NAME           TransformSortColumnType = "NAME"
	TransformSortColumnType_TRANSFORM_TYPE TransformSortColumnType = "TRANSFORM_TYPE"
	TransformSortColumnType_STATUS         TransformSortColumnType = "STATUS"
	TransformSortColumnType_CREATED        TransformSortColumnType = "CREATED"
	TransformSortColumnType_LAST_MODIFIED  TransformSortColumnType = "LAST_MODIFIED"
)

type TransformStatusType string

const (
	TransformStatusType_NOT_READY TransformStatusType = "NOT_READY"
	TransformStatusType_READY     TransformStatusType = "READY"
	TransformStatusType_DELETING  TransformStatusType = "DELETING"
)

type TransformType string

const (
	TransformType_FIND_MATCHES TransformType = "FIND_MATCHES"
)

type TriggerState string

const (
	TriggerState_CREATING     TriggerState = "CREATING"
	TriggerState_CREATED      TriggerState = "CREATED"
	TriggerState_ACTIVATING   TriggerState = "ACTIVATING"
	TriggerState_ACTIVATED    TriggerState = "ACTIVATED"
	TriggerState_DEACTIVATING TriggerState = "DEACTIVATING"
	TriggerState_DEACTIVATED  TriggerState = "DEACTIVATED"
	TriggerState_DELETING     TriggerState = "DELETING"
	TriggerState_UPDATING     TriggerState = "UPDATING"
)

type TriggerType string

const (
	TriggerType_SCHEDULED   TriggerType = "SCHEDULED"
	TriggerType_CONDITIONAL TriggerType = "CONDITIONAL"
	TriggerType_ON_DEMAND   TriggerType = "ON_DEMAND"
	TriggerType_EVENT       TriggerType = "EVENT"
)

type UnionType string

const (
	UnionType_ALL      UnionType = "ALL"
	UnionType_DISTINCT UnionType = "DISTINCT"
)

type UpdateBehavior string

const (
	UpdateBehavior_LOG                UpdateBehavior = "LOG"
	UpdateBehavior_UPDATE_IN_DATABASE UpdateBehavior = "UPDATE_IN_DATABASE"
)

type UpdateCatalogBehavior string

const (
	UpdateCatalogBehavior_UPDATE_IN_DATABASE UpdateCatalogBehavior = "UPDATE_IN_DATABASE"
	UpdateCatalogBehavior_LOG                UpdateCatalogBehavior = "LOG"
)

type WorkerType string

const (
	WorkerType_Standard WorkerType = "Standard"
	WorkerType_G_1X     WorkerType = "G.1X"
	WorkerType_G_2X     WorkerType = "G.2X"
	WorkerType_G_025X   WorkerType = "G.025X"
)

type WorkflowRunStatus string

const (
	WorkflowRunStatus_RUNNING   WorkflowRunStatus = "RUNNING"
	WorkflowRunStatus_COMPLETED WorkflowRunStatus = "COMPLETED"
	WorkflowRunStatus_STOPPING  WorkflowRunStatus = "STOPPING"
	WorkflowRunStatus_STOPPED   WorkflowRunStatus = "STOPPED"
	WorkflowRunStatus_ERROR     WorkflowRunStatus = "ERROR"
)
