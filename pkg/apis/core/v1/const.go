/*
Copyright 2018 The Fission Authors.

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

package v1

const (
	EXECUTOR_INSTANCEID_LABEL string = "executorInstanceId"
	DEFAULT_FUNCTION_TIMEOUT  int    = 60
)

const (
	// ResourceVersionCount env variable is used for updating configmaps and secrets in pods
	ResourceVersionCount string = "RESOURCE_VERSION_COUNT"
)

const (
	ChecksumTypeSHA256 ChecksumType = "sha256"
)

const (
	// ArchiveTypeLiteral means the package contents are specified in the Literal field of
	// resource itself.
	ArchiveTypeLiteral ArchiveType = "literal"

	// ArchiveTypeUrl means the package contents are at the specified URL.
	ArchiveTypeUrl ArchiveType = "url"
)

const (
	BuildStatusPending   = "pending"
	BuildStatusRunning   = "running"
	BuildStatusSucceeded = "succeeded"
	BuildStatusFailed    = "failed"
	BuildStatusNone      = "none"
)

const (
	AllowedFunctionsPerContainerSingle   = "single"
	AllowedFunctionsPerContainerInfinite = "infinite"
)

const (
	ExecutorTypePoolmgr   ExecutorType = "poolmgr"
	ExecutorTypeNewdeploy ExecutorType = "newdeploy"
)

const (
	StrategyTypeExecution = "execution"
)

const (
	SharedVolumeUserfunc   = "userfunc"
	SharedVolumePackages   = "packages"
	SharedVolumeSecrets    = "secrets"
	SharedVolumeConfigmaps = "configmaps"
)

const (
	MessageQueueTypeNats  = "nats-streaming"
	MessageQueueTypeASQ   = "azure-storage-queue"
	MessageQueueTypeKafka = "kafka"
)

const (
	// FunctionReferenceFunctionName means that the function
	// reference is simply by function name.
	FunctionReferenceTypeFunctionName = "name"

	FunctionReferenceTypeFunctionWeights = "function-weights"

	// Other function reference types we'd like to support:
	//   Versioned function, latest version
	//   Versioned function. by semver "latest compatible"
	//   Set of function references (recursively), by percentage of traffic
)

const (
	// failure type currently supported is http status code. This could be extended
	// in the future.
	FailureTypeStatusCode FailureType = "status-code"

	// Status of canary config can be one of the following
	CanaryConfigStatusPending   = "pending"
	CanaryConfigStatusSucceeded = "succeeded"
	CanaryConfigStatusFailed    = "failed"
	CanaryConfigStatusAborted   = "aborted"

	// set a max number for iterations to prevent infinite processing of canary config
	MaxIterationsForCanaryConfig = 10
)

const (
	DefaultSpecializationTimeOut = 120
)

const (
	FETCH_SOURCE = iota
	FETCH_DEPLOYMENT
	FETCH_URL
)

// executor kubernetes object label key
const (
	ENVIRONMENT_NAMESPACE     = "environmentNamespace"
	ENVIRONMENT_NAME          = "environmentName"
	ENVIRONMENT_UID           = "environmentUid"
	FUNCTION_NAMESPACE        = "functionNamespace"
	FUNCTION_NAME             = "functionName"
	FUNCTION_UID              = "functionUid"
	FUNCTION_RESOURCE_VERSION = "functionResourceVersion"
	EXECUTOR_TYPE             = "executorType"
)

const (
	ANNOTATION_SVC_HOST = "svcHost"
)

const (
	ArchiveLiteralSizeLimit int64 = 256 * 1024
)

const (
	FissionBuilderSA = "fission-builder"
	FissionFetcherSA = "fission-fetcher"

	SecretConfigMapGetterCR = "secret-configmap-getter"
	SecretConfigMapGetterRB = "secret-configmap-getter-binding"

	PackageGetterCR = "package-getter"
	PackageGetterRB = "package-getter-binding"

	ClusterRole = "ClusterRole"
)

const (
	// jingtao add 全局配置所在的命名空间
	GlobalSecretConfigMapNS = "fission-secret-configmap"
	// jingtao add Cluster Role 在fission安装时添加的权限,复用fission自带的权限
	GlobalSecretConfigMapGetterCR = "secret-configmap-getter"
	// jingtao add 绑定服务帐号和角色之间的关系
	GlobalSecretConfigMapGetterRB = "global-secret-configmap-getter-binding"
	// jingtao add 日志输出配置的configmap名称
	LogCollectionConfigName = "fission-log-collection-config"
	// jingtao add 函数配置中的指明使用的日志输出的方式
	LogCollectionConfigKey = "log-collection-type"
	// jingtao add 只有函数使用的此命名规则的configMap可以携带日志的输出方式
	LogConfigMapName = "func-%s"
)
