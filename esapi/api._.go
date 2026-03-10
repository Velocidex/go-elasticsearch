// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 9.3.0 (6ce0741): DO NOT EDIT

package esapi

// API contains the Elasticsearch APIs
type API struct {
	Bulk Bulk
}

// Cat contains the Cat APIs
type Cat struct {
	Aliases              CatAliases
	Allocation           CatAllocation
	CircuitBreaker       CatCircuitBreaker
	ComponentTemplates   CatComponentTemplates
	Count                CatCount
	Fielddata            CatFielddata
	Health               CatHealth
	Help                 CatHelp
	Indices              CatIndices
	MLDataFrameAnalytics CatMLDataFrameAnalytics
	MLDatafeeds          CatMLDatafeeds
	MLJobs               CatMLJobs
	MLTrainedModels      CatMLTrainedModels
	Master               CatMaster
	Nodeattrs            CatNodeattrs
	Nodes                CatNodes
	PendingTasks         CatPendingTasks
	Plugins              CatPlugins
	Recovery             CatRecovery
	Repositories         CatRepositories
	Segments             CatSegments
	Shards               CatShards
	Snapshots            CatSnapshots
	Tasks                CatTasks
	Templates            CatTemplates
	ThreadPool           CatThreadPool
	Transforms           CatTransforms
}

// Cluster contains the Cluster APIs
type Cluster struct {
	AllocationExplain            ClusterAllocationExplain
	DeleteComponentTemplate      ClusterDeleteComponentTemplate
	DeleteVotingConfigExclusions ClusterDeleteVotingConfigExclusions
	ExistsComponentTemplate      ClusterExistsComponentTemplate
	GetComponentTemplate         ClusterGetComponentTemplate
	GetSettings                  ClusterGetSettings
	Health                       ClusterHealth
	Info                         ClusterInfo
	PendingTasks                 ClusterPendingTasks
	PostVotingConfigExclusions   ClusterPostVotingConfigExclusions
	PutComponentTemplate         ClusterPutComponentTemplate
	PutSettings                  ClusterPutSettings
	RemoteInfo                   ClusterRemoteInfo
	Reroute                      ClusterReroute
	State                        ClusterState
	Stats                        ClusterStats
}

// Indices contains the Indices APIs
type Indices struct {
	AddBlock                  IndicesAddBlock
	Analyze                   IndicesAnalyze
	CancelMigrateReindex      IndicesCancelMigrateReindex
	ClearCache                IndicesClearCache
	Clone                     IndicesClone
	Close                     IndicesClose
	CreateDataStream          IndicesCreateDataStream
	CreateFrom                IndicesCreateFrom
	Create                    IndicesCreate
	DataStreamsStats          IndicesDataStreamsStats
	DeleteAlias               IndicesDeleteAlias
	DeleteDataLifecycle       IndicesDeleteDataLifecycle
	DeleteDataStreamOptions   IndicesDeleteDataStreamOptions
	DeleteDataStream          IndicesDeleteDataStream
	DeleteIndexTemplate       IndicesDeleteIndexTemplate
	Delete                    IndicesDelete
	DeleteSampleConfiguration IndicesDeleteSampleConfiguration
	DeleteTemplate            IndicesDeleteTemplate
	DiskUsage                 IndicesDiskUsage
	Downsample                IndicesDownsample
	ExistsAlias               IndicesExistsAlias
	ExistsIndexTemplate       IndicesExistsIndexTemplate
	Exists                    IndicesExists
	ExistsTemplate            IndicesExistsTemplate
	ExplainDataLifecycle      IndicesExplainDataLifecycle
	FieldUsageStats           IndicesFieldUsageStats
	Flush                     IndicesFlush
	Forcemerge                IndicesForcemerge
	GetAlias                  IndicesGetAlias
	GetAllSampleConfiguration IndicesGetAllSampleConfiguration
	GetDataLifecycle          IndicesGetDataLifecycle
	GetDataLifecycleStats     IndicesGetDataLifecycleStats
	GetDataStreamMappings     IndicesGetDataStreamMappings
	GetDataStreamOptions      IndicesGetDataStreamOptions
	GetDataStream             IndicesGetDataStream
	GetDataStreamSettings     IndicesGetDataStreamSettings
	GetFieldMapping           IndicesGetFieldMapping
	GetIndexTemplate          IndicesGetIndexTemplate
	GetMapping                IndicesGetMapping
	GetMigrateReindexStatus   IndicesGetMigrateReindexStatus
	Get                       IndicesGet
	GetSampleConfiguration    IndicesGetSampleConfiguration
	GetSample                 IndicesGetSample
	GetSampleStats            IndicesGetSampleStats
	GetSettings               IndicesGetSettings
	GetTemplate               IndicesGetTemplate
	MigrateReindex            IndicesMigrateReindex
	MigrateToDataStream       IndicesMigrateToDataStream
	ModifyDataStream          IndicesModifyDataStream
	Open                      IndicesOpen
	PromoteDataStream         IndicesPromoteDataStream
	PutAlias                  IndicesPutAlias
	PutDataLifecycle          IndicesPutDataLifecycle
	PutDataStreamMappings     IndicesPutDataStreamMappings
	PutDataStreamOptions      IndicesPutDataStreamOptions
	PutDataStreamSettings     IndicesPutDataStreamSettings
	PutIndexTemplate          IndicesPutIndexTemplate
	PutMapping                IndicesPutMapping
	PutSampleConfiguration    IndicesPutSampleConfiguration
	PutSettings               IndicesPutSettings
	PutTemplate               IndicesPutTemplate
	Recovery                  IndicesRecovery
	Refresh                   IndicesRefresh
	ReloadSearchAnalyzers     IndicesReloadSearchAnalyzers
	RemoveBlock               IndicesRemoveBlock
	ResolveCluster            IndicesResolveCluster
	ResolveIndex              IndicesResolveIndex
	Rollover                  IndicesRollover
	Segments                  IndicesSegments
	ShardStores               IndicesShardStores
	Shrink                    IndicesShrink
	SimulateIndexTemplate     IndicesSimulateIndexTemplate
	SimulateTemplate          IndicesSimulateTemplate
	Split                     IndicesSplit
	Stats                     IndicesStats
	UpdateAliases             IndicesUpdateAliases
	ValidateQuery             IndicesValidateQuery
}

// Ingest contains the Ingest APIs
type Ingest struct {
	DeleteGeoipDatabase      IngestDeleteGeoipDatabase
	DeleteIPLocationDatabase IngestDeleteIPLocationDatabase
	DeletePipeline           IngestDeletePipeline
	GeoIPStats               IngestGeoIPStats
	GetGeoipDatabase         IngestGetGeoipDatabase
	GetIPLocationDatabase    IngestGetIPLocationDatabase
	GetPipeline              IngestGetPipeline
	ProcessorGrok            IngestProcessorGrok
	PutGeoipDatabase         IngestPutGeoipDatabase
	PutIPLocationDatabase    IngestPutIPLocationDatabase
	PutPipeline              IngestPutPipeline
	Simulate                 IngestSimulate
}

// Nodes contains the Nodes APIs
type Nodes struct {
	ClearRepositoriesMeteringArchive NodesClearRepositoriesMeteringArchive
	GetRepositoriesMeteringInfo      NodesGetRepositoriesMeteringInfo
	HotThreads                       NodesHotThreads
	Info                             NodesInfo
	ReloadSecureSettings             NodesReloadSecureSettings
	Stats                            NodesStats
	Usage                            NodesUsage
}

// Remote contains the Remote APIs
type Remote struct {
}

// Snapshot contains the Snapshot APIs
type Snapshot struct {
	CleanupRepository         SnapshotCleanupRepository
	Clone                     SnapshotClone
	CreateRepository          SnapshotCreateRepository
	Create                    SnapshotCreate
	DeleteRepository          SnapshotDeleteRepository
	Delete                    SnapshotDelete
	GetRepository             SnapshotGetRepository
	Get                       SnapshotGet
	RepositoryAnalyze         SnapshotRepositoryAnalyze
	RepositoryVerifyIntegrity SnapshotRepositoryVerifyIntegrity
	Restore                   SnapshotRestore
	Status                    SnapshotStatus
	VerifyRepository          SnapshotVerifyRepository
}

// Tasks contains the Tasks APIs
type Tasks struct {
	Cancel TasksCancel
	Get    TasksGet
	List   TasksList
}

// AsyncSearch contains the AsyncSearch APIs
type AsyncSearch struct {
	Delete AsyncSearchDelete
	Get    AsyncSearchGet
	Status AsyncSearchStatus
	Submit AsyncSearchSubmit
}

// CCR contains the CCR APIs
type CCR struct {
	DeleteAutoFollowPattern CCRDeleteAutoFollowPattern
	FollowInfo              CCRFollowInfo
	Follow                  CCRFollow
	FollowStats             CCRFollowStats
	ForgetFollower          CCRForgetFollower
	GetAutoFollowPattern    CCRGetAutoFollowPattern
	PauseAutoFollowPattern  CCRPauseAutoFollowPattern
	PauseFollow             CCRPauseFollow
	PutAutoFollowPattern    CCRPutAutoFollowPattern
	ResumeAutoFollowPattern CCRResumeAutoFollowPattern
	ResumeFollow            CCRResumeFollow
	Stats                   CCRStats
	Unfollow                CCRUnfollow
}

// ILM contains the ILM APIs
type ILM struct {
	DeleteLifecycle    ILMDeleteLifecycle
	ExplainLifecycle   ILMExplainLifecycle
	GetLifecycle       ILMGetLifecycle
	GetStatus          ILMGetStatus
	MigrateToDataTiers ILMMigrateToDataTiers
	MoveToStep         ILMMoveToStep
	PutLifecycle       ILMPutLifecycle
	RemovePolicy       ILMRemovePolicy
	Retry              ILMRetry
	Start              ILMStart
	Stop               ILMStop
}

// License contains the License APIs
type License struct {
	Delete         LicenseDelete
	GetBasicStatus LicenseGetBasicStatus
	Get            LicenseGet
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations            MigrationDeprecations
	GetFeatureUpgradeStatus MigrationGetFeatureUpgradeStatus
	PostFeatureUpgrade      MigrationPostFeatureUpgrade
}

// ML contains the ML APIs
type ML struct {
	ClearTrainedModelDeploymentCache MLClearTrainedModelDeploymentCache
	CloseJob                         MLCloseJob
	DeleteCalendarEvent              MLDeleteCalendarEvent
	DeleteCalendarJob                MLDeleteCalendarJob
	DeleteCalendar                   MLDeleteCalendar
	DeleteDataFrameAnalytics         MLDeleteDataFrameAnalytics
	DeleteDatafeed                   MLDeleteDatafeed
	DeleteExpiredData                MLDeleteExpiredData
	DeleteFilter                     MLDeleteFilter
	DeleteForecast                   MLDeleteForecast
	DeleteJob                        MLDeleteJob
	DeleteModelSnapshot              MLDeleteModelSnapshot
	DeleteTrainedModelAlias          MLDeleteTrainedModelAlias
	DeleteTrainedModel               MLDeleteTrainedModel
	EstimateModelMemory              MLEstimateModelMemory
	EvaluateDataFrame                MLEvaluateDataFrame
	ExplainDataFrameAnalytics        MLExplainDataFrameAnalytics
	FlushJob                         MLFlushJob
	Forecast                         MLForecast
	GetBuckets                       MLGetBuckets
	GetCalendarEvents                MLGetCalendarEvents
	GetCalendars                     MLGetCalendars
	GetCategories                    MLGetCategories
	GetDataFrameAnalytics            MLGetDataFrameAnalytics
	GetDataFrameAnalyticsStats       MLGetDataFrameAnalyticsStats
	GetDatafeedStats                 MLGetDatafeedStats
	GetDatafeeds                     MLGetDatafeeds
	GetFilters                       MLGetFilters
	GetInfluencers                   MLGetInfluencers
	GetJobStats                      MLGetJobStats
	GetJobs                          MLGetJobs
	GetMemoryStats                   MLGetMemoryStats
	GetModelSnapshotUpgradeStats     MLGetModelSnapshotUpgradeStats
	GetModelSnapshots                MLGetModelSnapshots
	GetOverallBuckets                MLGetOverallBuckets
	GetRecords                       MLGetRecords
	GetTrainedModels                 MLGetTrainedModels
	GetTrainedModelsStats            MLGetTrainedModelsStats
	InferTrainedModel                MLInferTrainedModel
	Info                             MLInfo
	OpenJob                          MLOpenJob
	PostCalendarEvents               MLPostCalendarEvents
	PostData                         MLPostData
	PreviewDataFrameAnalytics        MLPreviewDataFrameAnalytics
	PreviewDatafeed                  MLPreviewDatafeed
	PutCalendarJob                   MLPutCalendarJob
	PutCalendar                      MLPutCalendar
	PutDataFrameAnalytics            MLPutDataFrameAnalytics
	PutDatafeed                      MLPutDatafeed
	PutFilter                        MLPutFilter
	PutJob                           MLPutJob
	PutTrainedModelAlias             MLPutTrainedModelAlias
	PutTrainedModelDefinitionPart    MLPutTrainedModelDefinitionPart
	PutTrainedModel                  MLPutTrainedModel
	PutTrainedModelVocabulary        MLPutTrainedModelVocabulary
	ResetJob                         MLResetJob
	RevertModelSnapshot              MLRevertModelSnapshot
	SetUpgradeMode                   MLSetUpgradeMode
	StartDataFrameAnalytics          MLStartDataFrameAnalytics
	StartDatafeed                    MLStartDatafeed
	StartTrainedModelDeployment      MLStartTrainedModelDeployment
	StopDataFrameAnalytics           MLStopDataFrameAnalytics
	StopDatafeed                     MLStopDatafeed
	StopTrainedModelDeployment       MLStopTrainedModelDeployment
	UpdateDataFrameAnalytics         MLUpdateDataFrameAnalytics
	UpdateDatafeed                   MLUpdateDatafeed
	UpdateFilter                     MLUpdateFilter
	UpdateJob                        MLUpdateJob
	UpdateModelSnapshot              MLUpdateModelSnapshot
	UpdateTrainedModelDeployment     MLUpdateTrainedModelDeployment
	UpgradeJobSnapshot               MLUpgradeJobSnapshot
	ValidateDetector                 MLValidateDetector
	Validate                         MLValidate
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk MonitoringBulk
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob    RollupDeleteJob
	GetJobs      RollupGetJobs
	GetCaps      RollupGetRollupCaps
	GetIndexCaps RollupGetRollupIndexCaps
	PutJob       RollupPutJob
	Search       RollupRollupSearch
	StartJob     RollupStartJob
	StopJob      RollupStopJob
}

// Security contains the Security APIs
type Security struct {
	ActivateUserProfile         SecurityActivateUserProfile
	Authenticate                SecurityAuthenticate
	BulkDeleteRole              SecurityBulkDeleteRole
	BulkPutRole                 SecurityBulkPutRole
	BulkUpdateAPIKeys           SecurityBulkUpdateAPIKeys
	ChangePassword              SecurityChangePassword
	ClearAPIKeyCache            SecurityClearAPIKeyCache
	ClearCachedPrivileges       SecurityClearCachedPrivileges
	ClearCachedRealms           SecurityClearCachedRealms
	ClearCachedRoles            SecurityClearCachedRoles
	ClearCachedServiceTokens    SecurityClearCachedServiceTokens
	CreateAPIKey                SecurityCreateAPIKey
	CreateCrossClusterAPIKey    SecurityCreateCrossClusterAPIKey
	CreateServiceToken          SecurityCreateServiceToken
	DelegatePki                 SecurityDelegatePki
	DeletePrivileges            SecurityDeletePrivileges
	DeleteRoleMapping           SecurityDeleteRoleMapping
	DeleteRole                  SecurityDeleteRole
	DeleteServiceToken          SecurityDeleteServiceToken
	DeleteUser                  SecurityDeleteUser
	DisableUserProfile          SecurityDisableUserProfile
	DisableUser                 SecurityDisableUser
	EnableUserProfile           SecurityEnableUserProfile
	EnableUser                  SecurityEnableUser
	EnrollKibana                SecurityEnrollKibana
	EnrollNode                  SecurityEnrollNode
	GetAPIKey                   SecurityGetAPIKey
	GetBuiltinPrivileges        SecurityGetBuiltinPrivileges
	GetPrivileges               SecurityGetPrivileges
	GetRoleMapping              SecurityGetRoleMapping
	GetRole                     SecurityGetRole
	GetServiceAccounts          SecurityGetServiceAccounts
	GetServiceCredentials       SecurityGetServiceCredentials
	GetSettings                 SecurityGetSettings
	GetStats                    SecurityGetStats
	GetToken                    SecurityGetToken
	GetUserPrivileges           SecurityGetUserPrivileges
	GetUserProfile              SecurityGetUserProfile
	GetUser                     SecurityGetUser
	GrantAPIKey                 SecurityGrantAPIKey
	HasPrivileges               SecurityHasPrivileges
	HasPrivilegesUserProfile    SecurityHasPrivilegesUserProfile
	InvalidateAPIKey            SecurityInvalidateAPIKey
	InvalidateToken             SecurityInvalidateToken
	OidcAuthenticate            SecurityOidcAuthenticate
	OidcLogout                  SecurityOidcLogout
	OidcPrepareAuthentication   SecurityOidcPrepareAuthentication
	PutPrivileges               SecurityPutPrivileges
	PutRoleMapping              SecurityPutRoleMapping
	PutRole                     SecurityPutRole
	PutUser                     SecurityPutUser
	QueryAPIKeys                SecurityQueryAPIKeys
	QueryRole                   SecurityQueryRole
	QueryUser                   SecurityQueryUser
	SamlAuthenticate            SecuritySamlAuthenticate
	SamlCompleteLogout          SecuritySamlCompleteLogout
	SamlInvalidate              SecuritySamlInvalidate
	SamlLogout                  SecuritySamlLogout
	SamlPrepareAuthentication   SecuritySamlPrepareAuthentication
	SamlServiceProviderMetadata SecuritySamlServiceProviderMetadata
	SuggestUserProfiles         SecuritySuggestUserProfiles
	UpdateAPIKey                SecurityUpdateAPIKey
	UpdateCrossClusterAPIKey    SecurityUpdateCrossClusterAPIKey
	UpdateSettings              SecurityUpdateSettings
	UpdateUserProfileData       SecurityUpdateUserProfileData
}

// SQL contains the SQL APIs
type SQL struct {
	ClearCursor    SQLClearCursor
	DeleteAsync    SQLDeleteAsync
	GetAsync       SQLGetAsync
	GetAsyncStatus SQLGetAsyncStatus
	Query          SQLQuery
	Translate      SQLTranslate
}

// SSL contains the SSL APIs
type SSL struct {
	Certificates SSLCertificates
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetSettings     WatcherGetSettings
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	QueryWatches    WatcherQueryWatches
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
	UpdateSettings  WatcherUpdateSettings
}

// XPack contains the XPack APIs
type XPack struct {
	Info  XPackInfo
	Usage XPackUsage
}

// New creates new API
func New(t Transport) *API {
	return &API{
		Bulk: newBulkFunc(t),
	}
}
