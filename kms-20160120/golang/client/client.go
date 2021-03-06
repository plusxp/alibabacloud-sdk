// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ListSecretsRequest struct {
	FetchTags  *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) SetFetchTags(v string) *ListSecretsRequest {
	s.FetchTags = &v
	return s
}

func (s *ListSecretsRequest) SetPageNumber(v int) *ListSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsRequest) SetPageSize(v int) *ListSecretsRequest {
	s.PageSize = &v
	return s
}

type ListSecretsResponse struct {
	RequestId  *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber *int                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                           `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	TotalCount *int                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	SecretList *ListSecretsResponseSecretList `json:"SecretList,omitempty" xml:"SecretList,omitempty" require:"true" type:"Struct"`
}

func (s ListSecretsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretsResponse) SetRequestId(v string) *ListSecretsResponse {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponse) SetPageNumber(v int) *ListSecretsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsResponse) SetPageSize(v int) *ListSecretsResponse {
	s.PageSize = &v
	return s
}

func (s *ListSecretsResponse) SetTotalCount(v int) *ListSecretsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListSecretsResponse) SetSecretList(v *ListSecretsResponseSecretList) *ListSecretsResponse {
	s.SecretList = v
	return s
}

type ListSecretsResponseSecretList struct {
	Secret []*ListSecretsResponseSecretListSecret `json:"Secret,omitempty" xml:"Secret,omitempty" require:"true" type:"Repeated"`
}

func (s ListSecretsResponseSecretList) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseSecretList) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseSecretList) SetSecret(v []*ListSecretsResponseSecretListSecret) *ListSecretsResponseSecretList {
	s.Secret = v
	return s
}

type ListSecretsResponseSecretListSecret struct {
	SecretName        *string                                  `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	CreateTime        *string                                  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime        *string                                  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	PlannedDeleteTime *string                                  `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty" require:"true"`
	Tags              *ListSecretsResponseSecretListSecretTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s ListSecretsResponseSecretListSecret) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseSecretListSecret) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseSecretListSecret) SetSecretName(v string) *ListSecretsResponseSecretListSecret {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseSecretListSecret) SetCreateTime(v string) *ListSecretsResponseSecretListSecret {
	s.CreateTime = &v
	return s
}

func (s *ListSecretsResponseSecretListSecret) SetUpdateTime(v string) *ListSecretsResponseSecretListSecret {
	s.UpdateTime = &v
	return s
}

func (s *ListSecretsResponseSecretListSecret) SetPlannedDeleteTime(v string) *ListSecretsResponseSecretListSecret {
	s.PlannedDeleteTime = &v
	return s
}

func (s *ListSecretsResponseSecretListSecret) SetTags(v *ListSecretsResponseSecretListSecretTags) *ListSecretsResponseSecretListSecret {
	s.Tags = v
	return s
}

type ListSecretsResponseSecretListSecretTags struct {
	Tag []*ListSecretsResponseSecretListSecretTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s ListSecretsResponseSecretListSecretTags) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseSecretListSecretTags) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseSecretListSecretTags) SetTag(v []*ListSecretsResponseSecretListSecretTagsTag) *ListSecretsResponseSecretListSecretTags {
	s.Tag = v
	return s
}

type ListSecretsResponseSecretListSecretTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListSecretsResponseSecretListSecretTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListSecretsResponseSecretListSecretTagsTag) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseSecretListSecretTagsTag) SetTagKey(v string) *ListSecretsResponseSecretListSecretTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListSecretsResponseSecretListSecretTagsTag) SetTagValue(v string) *ListSecretsResponseSecretListSecretTagsTag {
	s.TagValue = &v
	return s
}

type ListSecretVersionIdsRequest struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	IncludeDeprecated *string `json:"IncludeDeprecated,omitempty" xml:"IncludeDeprecated,omitempty"`
	PageNumber        *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize          *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecretVersionIdsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsRequest) SetSecretName(v string) *ListSecretVersionIdsRequest {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetIncludeDeprecated(v string) *ListSecretVersionIdsRequest {
	s.IncludeDeprecated = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageNumber(v int) *ListSecretVersionIdsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsRequest) SetPageSize(v int) *ListSecretVersionIdsRequest {
	s.PageSize = &v
	return s
}

type ListSecretVersionIdsResponse struct {
	PageNumber *int                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                                    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName *string                                 `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	TotalCount *int                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	VersionIds *ListSecretVersionIdsResponseVersionIds `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" require:"true" type:"Struct"`
}

func (s ListSecretVersionIdsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponse) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponse) SetPageNumber(v int) *ListSecretVersionIdsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetPageSize(v int) *ListSecretVersionIdsResponse {
	s.PageSize = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetRequestId(v string) *ListSecretVersionIdsResponse {
	s.RequestId = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetSecretName(v string) *ListSecretVersionIdsResponse {
	s.SecretName = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetTotalCount(v int) *ListSecretVersionIdsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListSecretVersionIdsResponse) SetVersionIds(v *ListSecretVersionIdsResponseVersionIds) *ListSecretVersionIdsResponse {
	s.VersionIds = v
	return s
}

type ListSecretVersionIdsResponseVersionIds struct {
	VersionId []*ListSecretVersionIdsResponseVersionIdsVersionId `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseVersionIds) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseVersionIds) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseVersionIds) SetVersionId(v []*ListSecretVersionIdsResponseVersionIdsVersionId) *ListSecretVersionIdsResponseVersionIds {
	s.VersionId = v
	return s
}

type ListSecretVersionIdsResponseVersionIdsVersionId struct {
	CreateTime    *string                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	VersionId     *string                                                       `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	VersionStages *ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" require:"true" type:"Struct"`
}

func (s ListSecretVersionIdsResponseVersionIdsVersionId) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseVersionIdsVersionId) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseVersionIdsVersionId) SetCreateTime(v string) *ListSecretVersionIdsResponseVersionIdsVersionId {
	s.CreateTime = &v
	return s
}

func (s *ListSecretVersionIdsResponseVersionIdsVersionId) SetVersionId(v string) *ListSecretVersionIdsResponseVersionIdsVersionId {
	s.VersionId = &v
	return s
}

func (s *ListSecretVersionIdsResponseVersionIdsVersionId) SetVersionStages(v *ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages) *ListSecretVersionIdsResponseVersionIdsVersionId {
	s.VersionStages = v
	return s
}

type ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" require:"true" type:"Repeated"`
}

func (s ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages) String() string {
	return tea.Prettify(s)
}

func (s ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages) GoString() string {
	return s.String()
}

func (s *ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages) SetVersionStage(v []*string) *ListSecretVersionIdsResponseVersionIdsVersionIdVersionStages {
	s.VersionStage = v
	return s
}

type DescribeSecretRequest struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	FetchTags  *string `json:"FetchTags,omitempty" xml:"FetchTags,omitempty"`
}

func (s DescribeSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecretRequest) SetSecretName(v string) *DescribeSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretRequest) SetFetchTags(v string) *DescribeSecretRequest {
	s.FetchTags = &v
	return s
}

type DescribeSecretResponse struct {
	RequestId         *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Arn               *string                     `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	SecretName        *string                     `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	EncryptionKeyId   *string                     `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty" require:"true"`
	Description       *string                     `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	CreateTime        *string                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	UpdateTime        *string                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty" require:"true"`
	PlannedDeleteTime *string                     `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty" require:"true"`
	Tags              *DescribeSecretResponseTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponse) SetRequestId(v string) *DescribeSecretResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretResponse) SetArn(v string) *DescribeSecretResponse {
	s.Arn = &v
	return s
}

func (s *DescribeSecretResponse) SetSecretName(v string) *DescribeSecretResponse {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretResponse) SetEncryptionKeyId(v string) *DescribeSecretResponse {
	s.EncryptionKeyId = &v
	return s
}

func (s *DescribeSecretResponse) SetDescription(v string) *DescribeSecretResponse {
	s.Description = &v
	return s
}

func (s *DescribeSecretResponse) SetCreateTime(v string) *DescribeSecretResponse {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecretResponse) SetUpdateTime(v string) *DescribeSecretResponse {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSecretResponse) SetPlannedDeleteTime(v string) *DescribeSecretResponse {
	s.PlannedDeleteTime = &v
	return s
}

func (s *DescribeSecretResponse) SetTags(v *DescribeSecretResponseTags) *DescribeSecretResponse {
	s.Tags = v
	return s
}

type DescribeSecretResponseTags struct {
	Tag []*DescribeSecretResponseTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSecretResponseTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseTags) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseTags) SetTag(v []*DescribeSecretResponseTagsTag) *DescribeSecretResponseTags {
	s.Tag = v
	return s
}

type DescribeSecretResponseTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s DescribeSecretResponseTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecretResponseTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseTagsTag) SetTagKey(v string) *DescribeSecretResponseTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSecretResponseTagsTag) SetTagValue(v string) *DescribeSecretResponseTagsTag {
	s.TagValue = &v
	return s
}

type UpdateSecretRequest struct {
	SecretName  *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretRequest) SetSecretName(v string) *UpdateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretRequest) SetDescription(v string) *UpdateSecretRequest {
	s.Description = &v
	return s
}

type UpdateSecretResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
}

func (s UpdateSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretResponse) SetRequestId(v string) *UpdateSecretResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSecretResponse) SetSecretName(v string) *UpdateSecretResponse {
	s.SecretName = &v
	return s
}

type GetSecretValueRequest struct {
	SecretName   *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	VersionStage *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueRequest) GoString() string {
	return s.String()
}

func (s *GetSecretValueRequest) SetSecretName(v string) *GetSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionStage(v string) *GetSecretValueRequest {
	s.VersionStage = &v
	return s
}

func (s *GetSecretValueRequest) SetVersionId(v string) *GetSecretValueRequest {
	s.VersionId = &v
	return s
}

type GetSecretValueResponse struct {
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName     *string                              `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	VersionId      *string                              `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	CreateTime     *string                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	SecretData     *string                              `json:"SecretData,omitempty" xml:"SecretData,omitempty" require:"true"`
	SecretDataType *string                              `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty" require:"true"`
	VersionStages  *GetSecretValueResponseVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" require:"true" type:"Struct"`
}

func (s GetSecretValueResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponse) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponse) SetRequestId(v string) *GetSecretValueResponse {
	s.RequestId = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretName(v string) *GetSecretValueResponse {
	s.SecretName = &v
	return s
}

func (s *GetSecretValueResponse) SetVersionId(v string) *GetSecretValueResponse {
	s.VersionId = &v
	return s
}

func (s *GetSecretValueResponse) SetCreateTime(v string) *GetSecretValueResponse {
	s.CreateTime = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretData(v string) *GetSecretValueResponse {
	s.SecretData = &v
	return s
}

func (s *GetSecretValueResponse) SetSecretDataType(v string) *GetSecretValueResponse {
	s.SecretDataType = &v
	return s
}

func (s *GetSecretValueResponse) SetVersionStages(v *GetSecretValueResponseVersionStages) *GetSecretValueResponse {
	s.VersionStages = v
	return s
}

type GetSecretValueResponseVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" require:"true" type:"Repeated"`
}

func (s GetSecretValueResponseVersionStages) String() string {
	return tea.Prettify(s)
}

func (s GetSecretValueResponseVersionStages) GoString() string {
	return s.String()
}

func (s *GetSecretValueResponseVersionStages) SetVersionStage(v []*string) *GetSecretValueResponseVersionStages {
	s.VersionStage = v
	return s
}

type GetRandomPasswordRequest struct {
	PasswordLength          *string `json:"PasswordLength,omitempty" xml:"PasswordLength,omitempty"`
	ExcludeCharacters       *string `json:"ExcludeCharacters,omitempty" xml:"ExcludeCharacters,omitempty"`
	ExcludeLowercase        *string `json:"ExcludeLowercase,omitempty" xml:"ExcludeLowercase,omitempty"`
	ExcludeUppercase        *string `json:"ExcludeUppercase,omitempty" xml:"ExcludeUppercase,omitempty"`
	ExcludeNumbers          *string `json:"ExcludeNumbers,omitempty" xml:"ExcludeNumbers,omitempty"`
	ExcludePunctuation      *string `json:"ExcludePunctuation,omitempty" xml:"ExcludePunctuation,omitempty"`
	RequireEachIncludedType *string `json:"RequireEachIncludedType,omitempty" xml:"RequireEachIncludedType,omitempty"`
}

func (s GetRandomPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordRequest) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordRequest) SetPasswordLength(v string) *GetRandomPasswordRequest {
	s.PasswordLength = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeCharacters(v string) *GetRandomPasswordRequest {
	s.ExcludeCharacters = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeLowercase(v string) *GetRandomPasswordRequest {
	s.ExcludeLowercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeUppercase(v string) *GetRandomPasswordRequest {
	s.ExcludeUppercase = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludeNumbers(v string) *GetRandomPasswordRequest {
	s.ExcludeNumbers = &v
	return s
}

func (s *GetRandomPasswordRequest) SetExcludePunctuation(v string) *GetRandomPasswordRequest {
	s.ExcludePunctuation = &v
	return s
}

func (s *GetRandomPasswordRequest) SetRequireEachIncludedType(v string) *GetRandomPasswordRequest {
	s.RequireEachIncludedType = &v
	return s
}

type GetRandomPasswordResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RandomPassword *string `json:"RandomPassword,omitempty" xml:"RandomPassword,omitempty" require:"true"`
}

func (s GetRandomPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRandomPasswordResponse) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordResponse) SetRequestId(v string) *GetRandomPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *GetRandomPasswordResponse) SetRandomPassword(v string) *GetRandomPasswordResponse {
	s.RandomPassword = &v
	return s
}

type RestoreSecretRequest struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
}

func (s RestoreSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretRequest) GoString() string {
	return s.String()
}

func (s *RestoreSecretRequest) SetSecretName(v string) *RestoreSecretRequest {
	s.SecretName = &v
	return s
}

type RestoreSecretResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
}

func (s RestoreSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreSecretResponse) GoString() string {
	return s.String()
}

func (s *RestoreSecretResponse) SetRequestId(v string) *RestoreSecretResponse {
	s.RequestId = &v
	return s
}

func (s *RestoreSecretResponse) SetSecretName(v string) *RestoreSecretResponse {
	s.SecretName = &v
	return s
}

type CreateSecretRequest struct {
	SecretName      *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	VersionId       *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	EncryptionKeyId *string `json:"EncryptionKeyId,omitempty" xml:"EncryptionKeyId,omitempty"`
	SecretData      *string `json:"SecretData,omitempty" xml:"SecretData,omitempty" require:"true"`
	SecretDataType  *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Tags            *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretRequest) GoString() string {
	return s.String()
}

func (s *CreateSecretRequest) SetSecretName(v string) *CreateSecretRequest {
	s.SecretName = &v
	return s
}

func (s *CreateSecretRequest) SetVersionId(v string) *CreateSecretRequest {
	s.VersionId = &v
	return s
}

func (s *CreateSecretRequest) SetEncryptionKeyId(v string) *CreateSecretRequest {
	s.EncryptionKeyId = &v
	return s
}

func (s *CreateSecretRequest) SetSecretData(v string) *CreateSecretRequest {
	s.SecretData = &v
	return s
}

func (s *CreateSecretRequest) SetSecretDataType(v string) *CreateSecretRequest {
	s.SecretDataType = &v
	return s
}

func (s *CreateSecretRequest) SetDescription(v string) *CreateSecretRequest {
	s.Description = &v
	return s
}

func (s *CreateSecretRequest) SetTags(v string) *CreateSecretRequest {
	s.Tags = &v
	return s
}

type CreateSecretResponse struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Arn        *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	VersionId  *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
}

func (s CreateSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretResponse) SetRequestId(v string) *CreateSecretResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSecretResponse) SetArn(v string) *CreateSecretResponse {
	s.Arn = &v
	return s
}

func (s *CreateSecretResponse) SetVersionId(v string) *CreateSecretResponse {
	s.VersionId = &v
	return s
}

func (s *CreateSecretResponse) SetSecretName(v string) *CreateSecretResponse {
	s.SecretName = &v
	return s
}

type PutSecretValueRequest struct {
	VersionId      *string `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	SecretName     *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	SecretData     *string `json:"SecretData,omitempty" xml:"SecretData,omitempty" require:"true"`
	SecretDataType *string `json:"SecretDataType,omitempty" xml:"SecretDataType,omitempty"`
	VersionStages  *string `json:"VersionStages,omitempty" xml:"VersionStages,omitempty"`
}

func (s PutSecretValueRequest) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueRequest) GoString() string {
	return s.String()
}

func (s *PutSecretValueRequest) SetVersionId(v string) *PutSecretValueRequest {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretName(v string) *PutSecretValueRequest {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretData(v string) *PutSecretValueRequest {
	s.SecretData = &v
	return s
}

func (s *PutSecretValueRequest) SetSecretDataType(v string) *PutSecretValueRequest {
	s.SecretDataType = &v
	return s
}

func (s *PutSecretValueRequest) SetVersionStages(v string) *PutSecretValueRequest {
	s.VersionStages = &v
	return s
}

type PutSecretValueResponse struct {
	RequestId     *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName    *string                              `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	VersionId     *string                              `json:"VersionId,omitempty" xml:"VersionId,omitempty" require:"true"`
	VersionStages *PutSecretValueResponseVersionStages `json:"VersionStages,omitempty" xml:"VersionStages,omitempty" require:"true" type:"Struct"`
}

func (s PutSecretValueResponse) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponse) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponse) SetRequestId(v string) *PutSecretValueResponse {
	s.RequestId = &v
	return s
}

func (s *PutSecretValueResponse) SetSecretName(v string) *PutSecretValueResponse {
	s.SecretName = &v
	return s
}

func (s *PutSecretValueResponse) SetVersionId(v string) *PutSecretValueResponse {
	s.VersionId = &v
	return s
}

func (s *PutSecretValueResponse) SetVersionStages(v *PutSecretValueResponseVersionStages) *PutSecretValueResponse {
	s.VersionStages = v
	return s
}

type PutSecretValueResponseVersionStages struct {
	VersionStage []*string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" require:"true" type:"Repeated"`
}

func (s PutSecretValueResponseVersionStages) String() string {
	return tea.Prettify(s)
}

func (s PutSecretValueResponseVersionStages) GoString() string {
	return s.String()
}

func (s *PutSecretValueResponseVersionStages) SetVersionStage(v []*string) *PutSecretValueResponseVersionStages {
	s.VersionStage = v
	return s
}

type DeleteSecretRequest struct {
	SecretName                 *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	ForceDeleteWithoutRecovery *string `json:"ForceDeleteWithoutRecovery,omitempty" xml:"ForceDeleteWithoutRecovery,omitempty"`
	RecoveryWindowInDays       *string `json:"RecoveryWindowInDays,omitempty" xml:"RecoveryWindowInDays,omitempty"`
}

func (s DeleteSecretRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretRequest) SetSecretName(v string) *DeleteSecretRequest {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretRequest) SetForceDeleteWithoutRecovery(v string) *DeleteSecretRequest {
	s.ForceDeleteWithoutRecovery = &v
	return s
}

func (s *DeleteSecretRequest) SetRecoveryWindowInDays(v string) *DeleteSecretRequest {
	s.RecoveryWindowInDays = &v
	return s
}

type DeleteSecretResponse struct {
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	PlannedDeleteTime *string `json:"PlannedDeleteTime,omitempty" xml:"PlannedDeleteTime,omitempty" require:"true"`
}

func (s DeleteSecretResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretResponse) SetRequestId(v string) *DeleteSecretResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteSecretResponse) SetSecretName(v string) *DeleteSecretResponse {
	s.SecretName = &v
	return s
}

func (s *DeleteSecretResponse) SetPlannedDeleteTime(v string) *DeleteSecretResponse {
	s.PlannedDeleteTime = &v
	return s
}

type UpdateSecretVersionStageRequest struct {
	SecretName        *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	VersionStage      *string `json:"VersionStage,omitempty" xml:"VersionStage,omitempty" require:"true"`
	RemoveFromVersion *string `json:"RemoveFromVersion,omitempty" xml:"RemoveFromVersion,omitempty"`
	MoveToVersion     *string `json:"MoveToVersion,omitempty" xml:"MoveToVersion,omitempty"`
}

func (s UpdateSecretVersionStageRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageRequest) SetSecretName(v string) *UpdateSecretVersionStageRequest {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetVersionStage(v string) *UpdateSecretVersionStageRequest {
	s.VersionStage = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetRemoveFromVersion(v string) *UpdateSecretVersionStageRequest {
	s.RemoveFromVersion = &v
	return s
}

func (s *UpdateSecretVersionStageRequest) SetMoveToVersion(v string) *UpdateSecretVersionStageRequest {
	s.MoveToVersion = &v
	return s
}

type UpdateSecretVersionStageResponse struct {
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty" require:"true"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateSecretVersionStageResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSecretVersionStageResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretVersionStageResponse) SetSecretName(v string) *UpdateSecretVersionStageResponse {
	s.SecretName = &v
	return s
}

func (s *UpdateSecretVersionStageResponse) SetRequestId(v string) *UpdateSecretVersionStageResponse {
	s.RequestId = &v
	return s
}

type AsymmetricDecryptRequest struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	Algorithm      *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty" require:"true"`
}

func (s AsymmetricDecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptRequest) SetCiphertextBlob(v string) *AsymmetricDecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyId(v string) *AsymmetricDecryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetKeyVersionId(v string) *AsymmetricDecryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricDecryptRequest) SetAlgorithm(v string) *AsymmetricDecryptRequest {
	s.Algorithm = &v
	return s
}

type AsymmetricDecryptResponse struct {
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s AsymmetricDecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricDecryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricDecryptResponse) SetPlaintext(v string) *AsymmetricDecryptResponse {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricDecryptResponse) SetKeyId(v string) *AsymmetricDecryptResponse {
	s.KeyId = &v
	return s
}

func (s *AsymmetricDecryptResponse) SetRequestId(v string) *AsymmetricDecryptResponse {
	s.RequestId = &v
	return s
}

func (s *AsymmetricDecryptResponse) SetKeyVersionId(v string) *AsymmetricDecryptResponse {
	s.KeyVersionId = &v
	return s
}

type AsymmetricVerifyRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty" require:"true"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty" require:"true"`
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s AsymmetricVerifyRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyRequest) SetKeyId(v string) *AsymmetricVerifyRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetKeyVersionId(v string) *AsymmetricVerifyRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetAlgorithm(v string) *AsymmetricVerifyRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetDigest(v string) *AsymmetricVerifyRequest {
	s.Digest = &v
	return s
}

func (s *AsymmetricVerifyRequest) SetValue(v string) *AsymmetricVerifyRequest {
	s.Value = &v
	return s
}

type AsymmetricVerifyResponse struct {
	Value        *bool   `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s AsymmetricVerifyResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricVerifyResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricVerifyResponse) SetValue(v bool) *AsymmetricVerifyResponse {
	s.Value = &v
	return s
}

func (s *AsymmetricVerifyResponse) SetKeyId(v string) *AsymmetricVerifyResponse {
	s.KeyId = &v
	return s
}

func (s *AsymmetricVerifyResponse) SetRequestId(v string) *AsymmetricVerifyResponse {
	s.RequestId = &v
	return s
}

func (s *AsymmetricVerifyResponse) SetKeyVersionId(v string) *AsymmetricVerifyResponse {
	s.KeyVersionId = &v
	return s
}

type AsymmetricSignRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty" require:"true"`
	Digest       *string `json:"Digest,omitempty" xml:"Digest,omitempty" require:"true"`
}

func (s AsymmetricSignRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricSignRequest) SetKeyId(v string) *AsymmetricSignRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignRequest) SetKeyVersionId(v string) *AsymmetricSignRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricSignRequest) SetAlgorithm(v string) *AsymmetricSignRequest {
	s.Algorithm = &v
	return s
}

func (s *AsymmetricSignRequest) SetDigest(v string) *AsymmetricSignRequest {
	s.Digest = &v
	return s
}

type AsymmetricSignResponse struct {
	Value        *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s AsymmetricSignResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricSignResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricSignResponse) SetValue(v string) *AsymmetricSignResponse {
	s.Value = &v
	return s
}

func (s *AsymmetricSignResponse) SetKeyId(v string) *AsymmetricSignResponse {
	s.KeyId = &v
	return s
}

func (s *AsymmetricSignResponse) SetRequestId(v string) *AsymmetricSignResponse {
	s.RequestId = &v
	return s
}

func (s *AsymmetricSignResponse) SetKeyVersionId(v string) *AsymmetricSignResponse {
	s.KeyVersionId = &v
	return s
}

type AsymmetricEncryptRequest struct {
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty" require:"true"`
}

func (s AsymmetricEncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptRequest) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptRequest) SetPlaintext(v string) *AsymmetricEncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyId(v string) *AsymmetricEncryptRequest {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetKeyVersionId(v string) *AsymmetricEncryptRequest {
	s.KeyVersionId = &v
	return s
}

func (s *AsymmetricEncryptRequest) SetAlgorithm(v string) *AsymmetricEncryptRequest {
	s.Algorithm = &v
	return s
}

type AsymmetricEncryptResponse struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s AsymmetricEncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s AsymmetricEncryptResponse) GoString() string {
	return s.String()
}

func (s *AsymmetricEncryptResponse) SetCiphertextBlob(v string) *AsymmetricEncryptResponse {
	s.CiphertextBlob = &v
	return s
}

func (s *AsymmetricEncryptResponse) SetKeyId(v string) *AsymmetricEncryptResponse {
	s.KeyId = &v
	return s
}

func (s *AsymmetricEncryptResponse) SetRequestId(v string) *AsymmetricEncryptResponse {
	s.RequestId = &v
	return s
}

func (s *AsymmetricEncryptResponse) SetKeyVersionId(v string) *AsymmetricEncryptResponse {
	s.KeyVersionId = &v
	return s
}

type GetPublicKeyRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s GetPublicKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyRequest) GoString() string {
	return s.String()
}

func (s *GetPublicKeyRequest) SetKeyId(v string) *GetPublicKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyRequest) SetKeyVersionId(v string) *GetPublicKeyRequest {
	s.KeyVersionId = &v
	return s
}

type GetPublicKeyResponse struct {
	PublicKey    *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s GetPublicKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponse) SetPublicKey(v string) *GetPublicKeyResponse {
	s.PublicKey = &v
	return s
}

func (s *GetPublicKeyResponse) SetKeyId(v string) *GetPublicKeyResponse {
	s.KeyId = &v
	return s
}

func (s *GetPublicKeyResponse) SetRequestId(v string) *GetPublicKeyResponse {
	s.RequestId = &v
	return s
}

func (s *GetPublicKeyResponse) SetKeyVersionId(v string) *GetPublicKeyResponse {
	s.KeyVersionId = &v
	return s
}

type GenerateDataKeyWithoutPlaintextRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeySpec           *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes     *int                   `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyWithoutPlaintextRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetKeySpec(v string) *GenerateDataKeyWithoutPlaintextRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetNumberOfBytes(v int) *GenerateDataKeyWithoutPlaintextRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyWithoutPlaintextRequest {
	s.EncryptionContext = v
	return s
}

type GenerateDataKeyWithoutPlaintextResponse struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s GenerateDataKeyWithoutPlaintextResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyWithoutPlaintextResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetCiphertextBlob(v string) *GenerateDataKeyWithoutPlaintextResponse {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetKeyId(v string) *GenerateDataKeyWithoutPlaintextResponse {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetRequestId(v string) *GenerateDataKeyWithoutPlaintextResponse {
	s.RequestId = &v
	return s
}

func (s *GenerateDataKeyWithoutPlaintextResponse) SetKeyVersionId(v string) *GenerateDataKeyWithoutPlaintextResponse {
	s.KeyVersionId = &v
	return s
}

type UpdateKeyDescriptionRequest struct {
	KeyId       *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
}

func (s UpdateKeyDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionRequest) SetKeyId(v string) *UpdateKeyDescriptionRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateKeyDescriptionRequest) SetDescription(v string) *UpdateKeyDescriptionRequest {
	s.Description = &v
	return s
}

type UpdateKeyDescriptionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateKeyDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateKeyDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeyDescriptionResponse) SetRequestId(v string) *UpdateKeyDescriptionResponse {
	s.RequestId = &v
	return s
}

type DescribeKeyVersionRequest struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s DescribeKeyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionRequest) SetKeyId(v string) *DescribeKeyVersionRequest {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionRequest) SetKeyVersionId(v string) *DescribeKeyVersionRequest {
	s.KeyVersionId = &v
	return s
}

type DescribeKeyVersionResponse struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersion *DescribeKeyVersionResponseKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" require:"true" type:"Struct"`
}

func (s DescribeKeyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponse) SetRequestId(v string) *DescribeKeyVersionResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyVersionResponse) SetKeyVersion(v *DescribeKeyVersionResponseKeyVersion) *DescribeKeyVersionResponse {
	s.KeyVersion = v
	return s
}

type DescribeKeyVersionResponseKeyVersion struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty" require:"true"`
}

func (s DescribeKeyVersionResponseKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyVersionResponseKeyVersion) GoString() string {
	return s.String()
}

func (s *DescribeKeyVersionResponseKeyVersion) SetKeyId(v string) *DescribeKeyVersionResponseKeyVersion {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyVersionResponseKeyVersion) SetKeyVersionId(v string) *DescribeKeyVersionResponseKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *DescribeKeyVersionResponseKeyVersion) SetCreationDate(v string) *DescribeKeyVersionResponseKeyVersion {
	s.CreationDate = &v
	return s
}

type UpdateRotationPolicyRequest struct {
	KeyId                   *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty" require:"true"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
}

func (s UpdateRotationPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyRequest) SetKeyId(v string) *UpdateRotationPolicyRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetEnableAutomaticRotation(v bool) *UpdateRotationPolicyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *UpdateRotationPolicyRequest) SetRotationInterval(v string) *UpdateRotationPolicyRequest {
	s.RotationInterval = &v
	return s
}

type UpdateRotationPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateRotationPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRotationPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateRotationPolicyResponse) SetRequestId(v string) *UpdateRotationPolicyResponse {
	s.RequestId = &v
	return s
}

type ListKeyVersionsRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKeyVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsRequest) SetKeyId(v string) *ListKeyVersionsRequest {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageNumber(v int) *ListKeyVersionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsRequest) SetPageSize(v int) *ListKeyVersionsRequest {
	s.PageSize = &v
	return s
}

type ListKeyVersionsResponse struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalCount  *int                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber  *int                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize    *int                                `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	KeyVersions *ListKeyVersionsResponseKeyVersions `json:"KeyVersions,omitempty" xml:"KeyVersions,omitempty" require:"true" type:"Struct"`
}

func (s ListKeyVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponse) SetRequestId(v string) *ListKeyVersionsResponse {
	s.RequestId = &v
	return s
}

func (s *ListKeyVersionsResponse) SetTotalCount(v int) *ListKeyVersionsResponse {
	s.TotalCount = &v
	return s
}

func (s *ListKeyVersionsResponse) SetPageNumber(v int) *ListKeyVersionsResponse {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsResponse) SetPageSize(v int) *ListKeyVersionsResponse {
	s.PageSize = &v
	return s
}

func (s *ListKeyVersionsResponse) SetKeyVersions(v *ListKeyVersionsResponseKeyVersions) *ListKeyVersionsResponse {
	s.KeyVersions = v
	return s
}

type ListKeyVersionsResponseKeyVersions struct {
	KeyVersion []*ListKeyVersionsResponseKeyVersionsKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" require:"true" type:"Repeated"`
}

func (s ListKeyVersionsResponseKeyVersions) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseKeyVersions) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseKeyVersions) SetKeyVersion(v []*ListKeyVersionsResponseKeyVersionsKeyVersion) *ListKeyVersionsResponseKeyVersions {
	s.KeyVersion = v
	return s
}

type ListKeyVersionsResponseKeyVersionsKeyVersion struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty" require:"true"`
}

func (s ListKeyVersionsResponseKeyVersionsKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s ListKeyVersionsResponseKeyVersionsKeyVersion) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseKeyVersionsKeyVersion) SetKeyId(v string) *ListKeyVersionsResponseKeyVersionsKeyVersion {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsResponseKeyVersionsKeyVersion) SetKeyVersionId(v string) *ListKeyVersionsResponseKeyVersionsKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *ListKeyVersionsResponseKeyVersionsKeyVersion) SetCreationDate(v string) *ListKeyVersionsResponseKeyVersionsKeyVersion {
	s.CreationDate = &v
	return s
}

type CreateKeyVersionRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s CreateKeyVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionRequest) SetKeyId(v string) *CreateKeyVersionRequest {
	s.KeyId = &v
	return s
}

type CreateKeyVersionResponse struct {
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersion *CreateKeyVersionResponseKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" require:"true" type:"Struct"`
}

func (s CreateKeyVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponse) SetRequestId(v string) *CreateKeyVersionResponse {
	s.RequestId = &v
	return s
}

func (s *CreateKeyVersionResponse) SetKeyVersion(v *CreateKeyVersionResponseKeyVersion) *CreateKeyVersionResponse {
	s.KeyVersion = v
	return s
}

type CreateKeyVersionResponseKeyVersion struct {
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty" require:"true"`
}

func (s CreateKeyVersionResponseKeyVersion) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyVersionResponseKeyVersion) GoString() string {
	return s.String()
}

func (s *CreateKeyVersionResponseKeyVersion) SetKeyId(v string) *CreateKeyVersionResponseKeyVersion {
	s.KeyId = &v
	return s
}

func (s *CreateKeyVersionResponseKeyVersion) SetKeyVersionId(v string) *CreateKeyVersionResponseKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *CreateKeyVersionResponseKeyVersion) SetCreationDate(v string) *CreateKeyVersionResponseKeyVersion {
	s.CreationDate = &v
	return s
}

type DescribeServiceRequest struct {
}

func (s DescribeServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceRequest) GoString() string {
	return s.String()
}

type DescribeServiceResponse struct {
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ProtectionLevels *DescribeServiceResponseProtectionLevels `json:"ProtectionLevels,omitempty" xml:"ProtectionLevels,omitempty" require:"true" type:"Struct"`
	KeySpecs         *DescribeServiceResponseKeySpecs         `json:"KeySpecs,omitempty" xml:"KeySpecs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) SetRequestId(v string) *DescribeServiceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceResponse) SetProtectionLevels(v *DescribeServiceResponseProtectionLevels) *DescribeServiceResponse {
	s.ProtectionLevels = v
	return s
}

func (s *DescribeServiceResponse) SetKeySpecs(v *DescribeServiceResponseKeySpecs) *DescribeServiceResponse {
	s.KeySpecs = v
	return s
}

type DescribeServiceResponseProtectionLevels struct {
	ProtectionLevel []*DescribeServiceResponseProtectionLevelsProtectionLevel `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeServiceResponseProtectionLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseProtectionLevels) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseProtectionLevels) SetProtectionLevel(v []*DescribeServiceResponseProtectionLevelsProtectionLevel) *DescribeServiceResponseProtectionLevels {
	s.ProtectionLevel = v
	return s
}

type DescribeServiceResponseProtectionLevelsProtectionLevel struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty" require:"true"`
}

func (s DescribeServiceResponseProtectionLevelsProtectionLevel) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseProtectionLevelsProtectionLevel) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseProtectionLevelsProtectionLevel) SetType(v string) *DescribeServiceResponseProtectionLevelsProtectionLevel {
	s.Type = &v
	return s
}

type DescribeServiceResponseKeySpecs struct {
	KeySpec []*DescribeServiceResponseKeySpecsKeySpec `json:"KeySpec,omitempty" xml:"KeySpec,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeServiceResponseKeySpecs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseKeySpecs) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseKeySpecs) SetKeySpec(v []*DescribeServiceResponseKeySpecsKeySpec) *DescribeServiceResponseKeySpecs {
	s.KeySpec = v
	return s
}

type DescribeServiceResponseKeySpecsKeySpec struct {
	Name                      *string                                                          `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	SupportedProtectionLevels *DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels `json:"SupportedProtectionLevels,omitempty" xml:"SupportedProtectionLevels,omitempty" require:"true" type:"Struct"`
	Usages                    *DescribeServiceResponseKeySpecsKeySpecUsages                    `json:"Usages,omitempty" xml:"Usages,omitempty" require:"true" type:"Struct"`
}

func (s DescribeServiceResponseKeySpecsKeySpec) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseKeySpecsKeySpec) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseKeySpecsKeySpec) SetName(v string) *DescribeServiceResponseKeySpecsKeySpec {
	s.Name = &v
	return s
}

func (s *DescribeServiceResponseKeySpecsKeySpec) SetSupportedProtectionLevels(v *DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels) *DescribeServiceResponseKeySpecsKeySpec {
	s.SupportedProtectionLevels = v
	return s
}

func (s *DescribeServiceResponseKeySpecsKeySpec) SetUsages(v *DescribeServiceResponseKeySpecsKeySpecUsages) *DescribeServiceResponseKeySpecsKeySpec {
	s.Usages = v
	return s
}

type DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels struct {
	SupportedProtectionLevel []*string `json:"SupportedProtectionLevel,omitempty" xml:"SupportedProtectionLevel,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels) SetSupportedProtectionLevel(v []*string) *DescribeServiceResponseKeySpecsKeySpecSupportedProtectionLevels {
	s.SupportedProtectionLevel = v
	return s
}

type DescribeServiceResponseKeySpecsKeySpecUsages struct {
	Usage []*string `json:"Usage,omitempty" xml:"Usage,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeServiceResponseKeySpecsKeySpecUsages) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponseKeySpecsKeySpecUsages) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponseKeySpecsKeySpecUsages) SetUsage(v []*string) *DescribeServiceResponseKeySpecsKeySpecUsages {
	s.Usage = v
	return s
}

type UpdateAliasRequest struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
}

func (s UpdateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasRequest) GoString() string {
	return s.String()
}

func (s *UpdateAliasRequest) SetKeyId(v string) *UpdateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *UpdateAliasRequest) SetAliasName(v string) *UpdateAliasRequest {
	s.AliasName = &v
	return s
}

type UpdateAliasResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UpdateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) SetRequestId(v string) *UpdateAliasResponse {
	s.RequestId = &v
	return s
}

type UntagResourceRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	TagKeys    *string `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" require:"true"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s UntagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceRequest) GoString() string {
	return s.String()
}

func (s *UntagResourceRequest) SetKeyId(v string) *UntagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *UntagResourceRequest) SetTagKeys(v string) *UntagResourceRequest {
	s.TagKeys = &v
	return s
}

func (s *UntagResourceRequest) SetSecretName(v string) *UntagResourceRequest {
	s.SecretName = &v
	return s
}

type UntagResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourceResponse) GoString() string {
	return s.String()
}

func (s *UntagResourceResponse) SetRequestId(v string) *UntagResourceResponse {
	s.RequestId = &v
	return s
}

type TagResourceRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	Tags       *string `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true"`
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s TagResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourceRequest) GoString() string {
	return s.String()
}

func (s *TagResourceRequest) SetKeyId(v string) *TagResourceRequest {
	s.KeyId = &v
	return s
}

func (s *TagResourceRequest) SetTags(v string) *TagResourceRequest {
	s.Tags = &v
	return s
}

func (s *TagResourceRequest) SetSecretName(v string) *TagResourceRequest {
	s.SecretName = &v
	return s
}

type TagResourceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourceResponse) GoString() string {
	return s.String()
}

func (s *TagResourceResponse) SetRequestId(v string) *TagResourceResponse {
	s.RequestId = &v
	return s
}

type ScheduleKeyDeletionRequest struct {
	KeyId               *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	PendingWindowInDays *int    `json:"PendingWindowInDays,omitempty" xml:"PendingWindowInDays,omitempty"`
}

func (s ScheduleKeyDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s ScheduleKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionRequest) SetKeyId(v string) *ScheduleKeyDeletionRequest {
	s.KeyId = &v
	return s
}

func (s *ScheduleKeyDeletionRequest) SetPendingWindowInDays(v int) *ScheduleKeyDeletionRequest {
	s.PendingWindowInDays = &v
	return s
}

type ScheduleKeyDeletionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ScheduleKeyDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s ScheduleKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionResponse) SetRequestId(v string) *ScheduleKeyDeletionResponse {
	s.RequestId = &v
	return s
}

type ListResourceTagsRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s ListResourceTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTagsRequest) SetKeyId(v string) *ListResourceTagsRequest {
	s.KeyId = &v
	return s
}

type ListResourceTagsResponse struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Tags      *ListResourceTagsResponseTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s ListResourceTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponse) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponse) SetRequestId(v string) *ListResourceTagsResponse {
	s.RequestId = &v
	return s
}

func (s *ListResourceTagsResponse) SetTags(v *ListResourceTagsResponseTags) *ListResourceTagsResponse {
	s.Tags = v
	return s
}

type ListResourceTagsResponseTags struct {
	Tag []*ListResourceTagsResponseTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s ListResourceTagsResponseTags) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponseTags) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseTags) SetTag(v []*ListResourceTagsResponseTagsTag) *ListResourceTagsResponseTags {
	s.Tag = v
	return s
}

type ListResourceTagsResponseTagsTag struct {
	KeyId    *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s ListResourceTagsResponseTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListResourceTagsResponseTagsTag) GoString() string {
	return s.String()
}

func (s *ListResourceTagsResponseTagsTag) SetKeyId(v string) *ListResourceTagsResponseTagsTag {
	s.KeyId = &v
	return s
}

func (s *ListResourceTagsResponseTagsTag) SetTagKey(v string) *ListResourceTagsResponseTagsTag {
	s.TagKey = &v
	return s
}

func (s *ListResourceTagsResponseTagsTag) SetTagValue(v string) *ListResourceTagsResponseTagsTag {
	s.TagValue = &v
	return s
}

type ListKeysRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeysRequest) GoString() string {
	return s.String()
}

func (s *ListKeysRequest) SetPageNumber(v int) *ListKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListKeysRequest) SetPageSize(v int) *ListKeysRequest {
	s.PageSize = &v
	return s
}

type ListKeysResponse struct {
	TotalCount *int    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ListKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeysResponse) GoString() string {
	return s.String()
}

func (s *ListKeysResponse) SetTotalCount(v int) *ListKeysResponse {
	s.TotalCount = &v
	return s
}

func (s *ListKeysResponse) SetPageNumber(v int) *ListKeysResponse {
	s.PageNumber = &v
	return s
}

func (s *ListKeysResponse) SetPageSize(v int) *ListKeysResponse {
	s.PageSize = &v
	return s
}

func (s *ListKeysResponse) SetRequestId(v string) *ListKeysResponse {
	s.RequestId = &v
	return s
}

type ListAliasesByKeyIdRequest struct {
	KeyId      *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	PageNumber *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesByKeyIdRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdRequest) SetKeyId(v string) *ListAliasesByKeyIdRequest {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageNumber(v int) *ListAliasesByKeyIdRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdRequest) SetPageSize(v int) *ListAliasesByKeyIdRequest {
	s.PageSize = &v
	return s
}

type ListAliasesByKeyIdResponse struct {
	TotalCount *int                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                               `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Aliases    *ListAliasesByKeyIdResponseAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" require:"true" type:"Struct"`
}

func (s ListAliasesByKeyIdResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponse) SetTotalCount(v int) *ListAliasesByKeyIdResponse {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetPageNumber(v int) *ListAliasesByKeyIdResponse {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetPageSize(v int) *ListAliasesByKeyIdResponse {
	s.PageSize = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetRequestId(v string) *ListAliasesByKeyIdResponse {
	s.RequestId = &v
	return s
}

func (s *ListAliasesByKeyIdResponse) SetAliases(v *ListAliasesByKeyIdResponseAliases) *ListAliasesByKeyIdResponse {
	s.Aliases = v
	return s
}

type ListAliasesByKeyIdResponseAliases struct {
	Alias []*ListAliasesByKeyIdResponseAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" require:"true" type:"Repeated"`
}

func (s ListAliasesByKeyIdResponseAliases) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseAliases) SetAlias(v []*ListAliasesByKeyIdResponseAliasesAlias) *ListAliasesByKeyIdResponseAliases {
	s.Alias = v
	return s
}

type ListAliasesByKeyIdResponseAliasesAlias struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
	AliasArn  *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty" require:"true"`
}

func (s ListAliasesByKeyIdResponseAliasesAlias) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesByKeyIdResponseAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesByKeyIdResponseAliasesAlias) SetKeyId(v string) *ListAliasesByKeyIdResponseAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesByKeyIdResponseAliasesAlias) SetAliasName(v string) *ListAliasesByKeyIdResponseAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesByKeyIdResponseAliasesAlias) SetAliasArn(v string) *ListAliasesByKeyIdResponseAliasesAlias {
	s.AliasArn = &v
	return s
}

type ListAliasesRequest struct {
	PageNumber *int `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAliasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListAliasesRequest) SetPageNumber(v int) *ListAliasesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesRequest) SetPageSize(v int) *ListAliasesRequest {
	s.PageSize = &v
	return s
}

type ListAliasesResponse struct {
	TotalCount *int                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty" require:"true"`
	PageNumber *int                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageSize   *int                        `json:"PageSize,omitempty" xml:"PageSize,omitempty" require:"true"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Aliases    *ListAliasesResponseAliases `json:"Aliases,omitempty" xml:"Aliases,omitempty" require:"true" type:"Struct"`
}

func (s ListAliasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesResponse) SetTotalCount(v int) *ListAliasesResponse {
	s.TotalCount = &v
	return s
}

func (s *ListAliasesResponse) SetPageNumber(v int) *ListAliasesResponse {
	s.PageNumber = &v
	return s
}

func (s *ListAliasesResponse) SetPageSize(v int) *ListAliasesResponse {
	s.PageSize = &v
	return s
}

func (s *ListAliasesResponse) SetRequestId(v string) *ListAliasesResponse {
	s.RequestId = &v
	return s
}

func (s *ListAliasesResponse) SetAliases(v *ListAliasesResponseAliases) *ListAliasesResponse {
	s.Aliases = v
	return s
}

type ListAliasesResponseAliases struct {
	Alias []*ListAliasesResponseAliasesAlias `json:"Alias,omitempty" xml:"Alias,omitempty" require:"true" type:"Repeated"`
}

func (s ListAliasesResponseAliases) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseAliases) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseAliases) SetAlias(v []*ListAliasesResponseAliasesAlias) *ListAliasesResponseAliases {
	s.Alias = v
	return s
}

type ListAliasesResponseAliasesAlias struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
	AliasArn  *string `json:"AliasArn,omitempty" xml:"AliasArn,omitempty" require:"true"`
}

func (s ListAliasesResponseAliasesAlias) String() string {
	return tea.Prettify(s)
}

func (s ListAliasesResponseAliasesAlias) GoString() string {
	return s.String()
}

func (s *ListAliasesResponseAliasesAlias) SetKeyId(v string) *ListAliasesResponseAliasesAlias {
	s.KeyId = &v
	return s
}

func (s *ListAliasesResponseAliasesAlias) SetAliasName(v string) *ListAliasesResponseAliasesAlias {
	s.AliasName = &v
	return s
}

func (s *ListAliasesResponseAliasesAlias) SetAliasArn(v string) *ListAliasesResponseAliasesAlias {
	s.AliasArn = &v
	return s
}

type ImportKeyMaterialRequest struct {
	KeyId                 *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	EncryptedKeyMaterial  *string `json:"EncryptedKeyMaterial,omitempty" xml:"EncryptedKeyMaterial,omitempty" require:"true"`
	ImportToken           *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty" require:"true"`
	KeyMaterialExpireUnix *int64  `json:"KeyMaterialExpireUnix,omitempty" xml:"KeyMaterialExpireUnix,omitempty" require:"true"`
}

func (s ImportKeyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialRequest) SetKeyId(v string) *ImportKeyMaterialRequest {
	s.KeyId = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetEncryptedKeyMaterial(v string) *ImportKeyMaterialRequest {
	s.EncryptedKeyMaterial = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetImportToken(v string) *ImportKeyMaterialRequest {
	s.ImportToken = &v
	return s
}

func (s *ImportKeyMaterialRequest) SetKeyMaterialExpireUnix(v int64) *ImportKeyMaterialRequest {
	s.KeyMaterialExpireUnix = &v
	return s
}

type ImportKeyMaterialResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ImportKeyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyMaterialResponse) SetRequestId(v string) *ImportKeyMaterialResponse {
	s.RequestId = &v
	return s
}

type GetParametersForImportRequest struct {
	KeyId             *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	WrappingAlgorithm *string `json:"WrappingAlgorithm,omitempty" xml:"WrappingAlgorithm,omitempty" require:"true"`
	WrappingKeySpec   *string `json:"WrappingKeySpec,omitempty" xml:"WrappingKeySpec,omitempty" require:"true"`
}

func (s GetParametersForImportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportRequest) GoString() string {
	return s.String()
}

func (s *GetParametersForImportRequest) SetKeyId(v string) *GetParametersForImportRequest {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingAlgorithm(v string) *GetParametersForImportRequest {
	s.WrappingAlgorithm = &v
	return s
}

func (s *GetParametersForImportRequest) SetWrappingKeySpec(v string) *GetParametersForImportRequest {
	s.WrappingKeySpec = &v
	return s
}

type GetParametersForImportResponse struct {
	KeyId           *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	ImportToken     *string `json:"ImportToken,omitempty" xml:"ImportToken,omitempty" require:"true"`
	PublicKey       *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty" require:"true"`
	TokenExpireTime *string `json:"TokenExpireTime,omitempty" xml:"TokenExpireTime,omitempty" require:"true"`
}

func (s GetParametersForImportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetParametersForImportResponse) GoString() string {
	return s.String()
}

func (s *GetParametersForImportResponse) SetKeyId(v string) *GetParametersForImportResponse {
	s.KeyId = &v
	return s
}

func (s *GetParametersForImportResponse) SetRequestId(v string) *GetParametersForImportResponse {
	s.RequestId = &v
	return s
}

func (s *GetParametersForImportResponse) SetImportToken(v string) *GetParametersForImportResponse {
	s.ImportToken = &v
	return s
}

func (s *GetParametersForImportResponse) SetPublicKey(v string) *GetParametersForImportResponse {
	s.PublicKey = &v
	return s
}

func (s *GetParametersForImportResponse) SetTokenExpireTime(v string) *GetParametersForImportResponse {
	s.TokenExpireTime = &v
	return s
}

type GenerateDataKeyRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeySpec           *string                `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
	NumberOfBytes     *int                   `json:"NumberOfBytes,omitempty" xml:"NumberOfBytes,omitempty"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s GenerateDataKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyRequest) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyRequest) SetKeyId(v string) *GenerateDataKeyRequest {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyRequest) SetKeySpec(v string) *GenerateDataKeyRequest {
	s.KeySpec = &v
	return s
}

func (s *GenerateDataKeyRequest) SetNumberOfBytes(v int) *GenerateDataKeyRequest {
	s.NumberOfBytes = &v
	return s
}

func (s *GenerateDataKeyRequest) SetEncryptionContext(v map[string]interface{}) *GenerateDataKeyRequest {
	s.EncryptionContext = v
	return s
}

type GenerateDataKeyResponse struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	Plaintext      *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s GenerateDataKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateDataKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateDataKeyResponse) SetCiphertextBlob(v string) *GenerateDataKeyResponse {
	s.CiphertextBlob = &v
	return s
}

func (s *GenerateDataKeyResponse) SetKeyId(v string) *GenerateDataKeyResponse {
	s.KeyId = &v
	return s
}

func (s *GenerateDataKeyResponse) SetPlaintext(v string) *GenerateDataKeyResponse {
	s.Plaintext = &v
	return s
}

func (s *GenerateDataKeyResponse) SetRequestId(v string) *GenerateDataKeyResponse {
	s.RequestId = &v
	return s
}

func (s *GenerateDataKeyResponse) SetKeyVersionId(v string) *GenerateDataKeyResponse {
	s.KeyVersionId = &v
	return s
}

type EncryptRequest struct {
	KeyId             *string                `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	Plaintext         *string                `json:"Plaintext,omitempty" xml:"Plaintext,omitempty" require:"true"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s EncryptRequest) String() string {
	return tea.Prettify(s)
}

func (s EncryptRequest) GoString() string {
	return s.String()
}

func (s *EncryptRequest) SetKeyId(v string) *EncryptRequest {
	s.KeyId = &v
	return s
}

func (s *EncryptRequest) SetPlaintext(v string) *EncryptRequest {
	s.Plaintext = &v
	return s
}

func (s *EncryptRequest) SetEncryptionContext(v map[string]interface{}) *EncryptRequest {
	s.EncryptionContext = v
	return s
}

type EncryptResponse struct {
	CiphertextBlob *string `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	KeyId          *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId   *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s EncryptResponse) String() string {
	return tea.Prettify(s)
}

func (s EncryptResponse) GoString() string {
	return s.String()
}

func (s *EncryptResponse) SetCiphertextBlob(v string) *EncryptResponse {
	s.CiphertextBlob = &v
	return s
}

func (s *EncryptResponse) SetKeyId(v string) *EncryptResponse {
	s.KeyId = &v
	return s
}

func (s *EncryptResponse) SetRequestId(v string) *EncryptResponse {
	s.RequestId = &v
	return s
}

func (s *EncryptResponse) SetKeyVersionId(v string) *EncryptResponse {
	s.KeyVersionId = &v
	return s
}

type EnableKeyRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s EnableKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableKeyRequest) GoString() string {
	return s.String()
}

func (s *EnableKeyRequest) SetKeyId(v string) *EnableKeyRequest {
	s.KeyId = &v
	return s
}

type EnableKeyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s EnableKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableKeyResponse) GoString() string {
	return s.String()
}

func (s *EnableKeyResponse) SetRequestId(v string) *EnableKeyResponse {
	s.RequestId = &v
	return s
}

type DisableKeyRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s DisableKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableKeyRequest) GoString() string {
	return s.String()
}

func (s *DisableKeyRequest) SetKeyId(v string) *DisableKeyRequest {
	s.KeyId = &v
	return s
}

type DisableKeyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DisableKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableKeyResponse) GoString() string {
	return s.String()
}

func (s *DisableKeyResponse) SetRequestId(v string) *DisableKeyResponse {
	s.RequestId = &v
	return s
}

type DescribeRegionsRequest struct {
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

type DescribeRegionsResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Regions   *DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Struct"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v *DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	Region []*DescribeRegionsResponseRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegion(v []*DescribeRegionsResponseRegionsRegion) *DescribeRegionsResponseRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseRegionsRegion struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseRegionsRegion {
	s.RegionId = &v
	return s
}

type DescribeKeyRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s DescribeKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeKeyRequest) SetKeyId(v string) *DescribeKeyRequest {
	s.KeyId = &v
	return s
}

type DescribeKeyResponse struct {
	RequestId   *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyMetadata *DescribeKeyResponseKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" require:"true" type:"Struct"`
}

func (s DescribeKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponse) SetRequestId(v string) *DescribeKeyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeKeyResponse) SetKeyMetadata(v *DescribeKeyResponseKeyMetadata) *DescribeKeyResponse {
	s.KeyMetadata = v
	return s
}

type DescribeKeyResponseKeyMetadata struct {
	CreationDate       *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty" require:"true"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	KeyId              *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyState           *string `json:"KeyState,omitempty" xml:"KeyState,omitempty" require:"true"`
	KeyUsage           *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" require:"true"`
	DeleteDate         *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty" require:"true"`
	Creator            *string `json:"Creator,omitempty" xml:"Creator,omitempty" require:"true"`
	Arn                *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	Origin             *string `json:"Origin,omitempty" xml:"Origin,omitempty" require:"true"`
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty" require:"true"`
	ProtectionLevel    *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty" require:"true"`
	PrimaryKeyVersion  *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty" require:"true"`
	LastRotationDate   *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty" require:"true"`
	AutomaticRotation  *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty" require:"true"`
	RotationInterval   *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty" require:"true"`
	NextRotationDate   *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty" require:"true"`
	KeySpec            *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty" require:"true"`
}

func (s DescribeKeyResponseKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s DescribeKeyResponseKeyMetadata) GoString() string {
	return s.String()
}

func (s *DescribeKeyResponseKeyMetadata) SetCreationDate(v string) *DescribeKeyResponseKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetDescription(v string) *DescribeKeyResponseKeyMetadata {
	s.Description = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetKeyId(v string) *DescribeKeyResponseKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetKeyState(v string) *DescribeKeyResponseKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetKeyUsage(v string) *DescribeKeyResponseKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetDeleteDate(v string) *DescribeKeyResponseKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetCreator(v string) *DescribeKeyResponseKeyMetadata {
	s.Creator = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetArn(v string) *DescribeKeyResponseKeyMetadata {
	s.Arn = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetOrigin(v string) *DescribeKeyResponseKeyMetadata {
	s.Origin = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetMaterialExpireTime(v string) *DescribeKeyResponseKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetProtectionLevel(v string) *DescribeKeyResponseKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetPrimaryKeyVersion(v string) *DescribeKeyResponseKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetLastRotationDate(v string) *DescribeKeyResponseKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetAutomaticRotation(v string) *DescribeKeyResponseKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetRotationInterval(v string) *DescribeKeyResponseKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetNextRotationDate(v string) *DescribeKeyResponseKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *DescribeKeyResponseKeyMetadata) SetKeySpec(v string) *DescribeKeyResponseKeyMetadata {
	s.KeySpec = &v
	return s
}

type DeleteKeyMaterialRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s DeleteKeyMaterialRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyMaterialRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialRequest) SetKeyId(v string) *DeleteKeyMaterialRequest {
	s.KeyId = &v
	return s
}

type DeleteKeyMaterialResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteKeyMaterialResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyMaterialResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyMaterialResponse) SetRequestId(v string) *DeleteKeyMaterialResponse {
	s.RequestId = &v
	return s
}

type DeleteAliasRequest struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
}

func (s DeleteAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasRequest) GoString() string {
	return s.String()
}

func (s *DeleteAliasRequest) SetAliasName(v string) *DeleteAliasRequest {
	s.AliasName = &v
	return s
}

type DeleteAliasResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAliasResponse) GoString() string {
	return s.String()
}

func (s *DeleteAliasResponse) SetRequestId(v string) *DeleteAliasResponse {
	s.RequestId = &v
	return s
}

type DecryptRequest struct {
	CiphertextBlob    *string                `json:"CiphertextBlob,omitempty" xml:"CiphertextBlob,omitempty" require:"true"`
	EncryptionContext map[string]interface{} `json:"EncryptionContext,omitempty" xml:"EncryptionContext,omitempty"`
}

func (s DecryptRequest) String() string {
	return tea.Prettify(s)
}

func (s DecryptRequest) GoString() string {
	return s.String()
}

func (s *DecryptRequest) SetCiphertextBlob(v string) *DecryptRequest {
	s.CiphertextBlob = &v
	return s
}

func (s *DecryptRequest) SetEncryptionContext(v map[string]interface{}) *DecryptRequest {
	s.EncryptionContext = v
	return s
}

type DecryptResponse struct {
	Plaintext    *string `json:"Plaintext,omitempty" xml:"Plaintext,omitempty" require:"true"`
	KeyId        *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty" require:"true"`
}

func (s DecryptResponse) String() string {
	return tea.Prettify(s)
}

func (s DecryptResponse) GoString() string {
	return s.String()
}

func (s *DecryptResponse) SetPlaintext(v string) *DecryptResponse {
	s.Plaintext = &v
	return s
}

func (s *DecryptResponse) SetKeyId(v string) *DecryptResponse {
	s.KeyId = &v
	return s
}

func (s *DecryptResponse) SetRequestId(v string) *DecryptResponse {
	s.RequestId = &v
	return s
}

func (s *DecryptResponse) SetKeyVersionId(v string) *DecryptResponse {
	s.KeyVersionId = &v
	return s
}

type CreateKeyRequest struct {
	Description             *string `json:"Description,omitempty" xml:"Description,omitempty"`
	KeyUsage                *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty"`
	Origin                  *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	ProtectionLevel         *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty"`
	EnableAutomaticRotation *bool   `json:"EnableAutomaticRotation,omitempty" xml:"EnableAutomaticRotation,omitempty"`
	RotationInterval        *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty"`
	KeySpec                 *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty"`
}

func (s CreateKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyRequest) SetDescription(v string) *CreateKeyRequest {
	s.Description = &v
	return s
}

func (s *CreateKeyRequest) SetKeyUsage(v string) *CreateKeyRequest {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyRequest) SetOrigin(v string) *CreateKeyRequest {
	s.Origin = &v
	return s
}

func (s *CreateKeyRequest) SetProtectionLevel(v string) *CreateKeyRequest {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyRequest) SetEnableAutomaticRotation(v bool) *CreateKeyRequest {
	s.EnableAutomaticRotation = &v
	return s
}

func (s *CreateKeyRequest) SetRotationInterval(v string) *CreateKeyRequest {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyRequest) SetKeySpec(v string) *CreateKeyRequest {
	s.KeySpec = &v
	return s
}

type CreateKeyResponse struct {
	RequestId   *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KeyMetadata *CreateKeyResponseKeyMetadata `json:"KeyMetadata,omitempty" xml:"KeyMetadata,omitempty" require:"true" type:"Struct"`
}

func (s CreateKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateKeyResponse) SetRequestId(v string) *CreateKeyResponse {
	s.RequestId = &v
	return s
}

func (s *CreateKeyResponse) SetKeyMetadata(v *CreateKeyResponseKeyMetadata) *CreateKeyResponse {
	s.KeyMetadata = v
	return s
}

type CreateKeyResponseKeyMetadata struct {
	CreationDate       *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty" require:"true"`
	Description        *string `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	KeyId              *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	KeyState           *string `json:"KeyState,omitempty" xml:"KeyState,omitempty" require:"true"`
	KeyUsage           *string `json:"KeyUsage,omitempty" xml:"KeyUsage,omitempty" require:"true"`
	DeleteDate         *string `json:"DeleteDate,omitempty" xml:"DeleteDate,omitempty" require:"true"`
	Creator            *string `json:"Creator,omitempty" xml:"Creator,omitempty" require:"true"`
	Arn                *string `json:"Arn,omitempty" xml:"Arn,omitempty" require:"true"`
	Origin             *string `json:"Origin,omitempty" xml:"Origin,omitempty" require:"true"`
	MaterialExpireTime *string `json:"MaterialExpireTime,omitempty" xml:"MaterialExpireTime,omitempty" require:"true"`
	ProtectionLevel    *string `json:"ProtectionLevel,omitempty" xml:"ProtectionLevel,omitempty" require:"true"`
	PrimaryKeyVersion  *string `json:"PrimaryKeyVersion,omitempty" xml:"PrimaryKeyVersion,omitempty" require:"true"`
	LastRotationDate   *string `json:"LastRotationDate,omitempty" xml:"LastRotationDate,omitempty" require:"true"`
	AutomaticRotation  *string `json:"AutomaticRotation,omitempty" xml:"AutomaticRotation,omitempty" require:"true"`
	RotationInterval   *string `json:"RotationInterval,omitempty" xml:"RotationInterval,omitempty" require:"true"`
	NextRotationDate   *string `json:"NextRotationDate,omitempty" xml:"NextRotationDate,omitempty" require:"true"`
	KeySpec            *string `json:"KeySpec,omitempty" xml:"KeySpec,omitempty" require:"true"`
}

func (s CreateKeyResponseKeyMetadata) String() string {
	return tea.Prettify(s)
}

func (s CreateKeyResponseKeyMetadata) GoString() string {
	return s.String()
}

func (s *CreateKeyResponseKeyMetadata) SetCreationDate(v string) *CreateKeyResponseKeyMetadata {
	s.CreationDate = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetDescription(v string) *CreateKeyResponseKeyMetadata {
	s.Description = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetKeyId(v string) *CreateKeyResponseKeyMetadata {
	s.KeyId = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetKeyState(v string) *CreateKeyResponseKeyMetadata {
	s.KeyState = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetKeyUsage(v string) *CreateKeyResponseKeyMetadata {
	s.KeyUsage = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetDeleteDate(v string) *CreateKeyResponseKeyMetadata {
	s.DeleteDate = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetCreator(v string) *CreateKeyResponseKeyMetadata {
	s.Creator = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetArn(v string) *CreateKeyResponseKeyMetadata {
	s.Arn = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetOrigin(v string) *CreateKeyResponseKeyMetadata {
	s.Origin = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetMaterialExpireTime(v string) *CreateKeyResponseKeyMetadata {
	s.MaterialExpireTime = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetProtectionLevel(v string) *CreateKeyResponseKeyMetadata {
	s.ProtectionLevel = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetPrimaryKeyVersion(v string) *CreateKeyResponseKeyMetadata {
	s.PrimaryKeyVersion = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetLastRotationDate(v string) *CreateKeyResponseKeyMetadata {
	s.LastRotationDate = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetAutomaticRotation(v string) *CreateKeyResponseKeyMetadata {
	s.AutomaticRotation = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetRotationInterval(v string) *CreateKeyResponseKeyMetadata {
	s.RotationInterval = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetNextRotationDate(v string) *CreateKeyResponseKeyMetadata {
	s.NextRotationDate = &v
	return s
}

func (s *CreateKeyResponseKeyMetadata) SetKeySpec(v string) *CreateKeyResponseKeyMetadata {
	s.KeySpec = &v
	return s
}

type CreateAliasRequest struct {
	KeyId     *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty" require:"true"`
}

func (s CreateAliasRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasRequest) GoString() string {
	return s.String()
}

func (s *CreateAliasRequest) SetKeyId(v string) *CreateAliasRequest {
	s.KeyId = &v
	return s
}

func (s *CreateAliasRequest) SetAliasName(v string) *CreateAliasRequest {
	s.AliasName = &v
	return s
}

type CreateAliasResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateAliasResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAliasResponse) GoString() string {
	return s.String()
}

func (s *CreateAliasResponse) SetRequestId(v string) *CreateAliasResponse {
	s.RequestId = &v
	return s
}

type CancelKeyDeletionRequest struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s CancelKeyDeletionRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelKeyDeletionRequest) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionRequest) SetKeyId(v string) *CancelKeyDeletionRequest {
	s.KeyId = &v
	return s
}

type CancelKeyDeletionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CancelKeyDeletionResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionResponse) SetRequestId(v string) *CancelKeyDeletionResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("kms"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) ListSecretsWithOptions(request *ListSecretsRequest, runtime *util.RuntimeOptions) (_result *ListSecretsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSecretsResponse{}
	_body, _err := client.DoRequest(tea.String("ListSecrets"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecrets(request *ListSecretsRequest) (_result *ListSecretsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretsResponse{}
	_body, _err := client.ListSecretsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecretVersionIdsWithOptions(request *ListSecretVersionIdsRequest, runtime *util.RuntimeOptions) (_result *ListSecretVersionIdsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.DoRequest(tea.String("ListSecretVersionIds"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecretVersionIds(request *ListSecretVersionIdsRequest) (_result *ListSecretVersionIdsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecretVersionIdsResponse{}
	_body, _err := client.ListSecretVersionIdsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecretWithOptions(request *DescribeSecretRequest, runtime *util.RuntimeOptions) (_result *DescribeSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSecretResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecret(request *DescribeSecretRequest) (_result *DescribeSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecretResponse{}
	_body, _err := client.DescribeSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSecretWithOptions(request *UpdateSecretRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateSecretResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSecret(request *UpdateSecretRequest) (_result *UpdateSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretResponse{}
	_body, _err := client.UpdateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecretValueWithOptions(request *GetSecretValueRequest, runtime *util.RuntimeOptions) (_result *GetSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetSecretValueResponse{}
	_body, _err := client.DoRequest(tea.String("GetSecretValue"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecretValue(request *GetSecretValueRequest) (_result *GetSecretValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSecretValueResponse{}
	_body, _err := client.GetSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRandomPasswordWithOptions(request *GetRandomPasswordRequest, runtime *util.RuntimeOptions) (_result *GetRandomPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.DoRequest(tea.String("GetRandomPassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRandomPassword(request *GetRandomPasswordRequest) (_result *GetRandomPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetRandomPasswordResponse{}
	_body, _err := client.GetRandomPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreSecretWithOptions(request *RestoreSecretRequest, runtime *util.RuntimeOptions) (_result *RestoreSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestoreSecretResponse{}
	_body, _err := client.DoRequest(tea.String("RestoreSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreSecret(request *RestoreSecretRequest) (_result *RestoreSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestoreSecretResponse{}
	_body, _err := client.RestoreSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecretWithOptions(request *CreateSecretRequest, runtime *util.RuntimeOptions) (_result *CreateSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateSecretResponse{}
	_body, _err := client.DoRequest(tea.String("CreateSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecret(request *CreateSecretRequest) (_result *CreateSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecretResponse{}
	_body, _err := client.CreateSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutSecretValueWithOptions(request *PutSecretValueRequest, runtime *util.RuntimeOptions) (_result *PutSecretValueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PutSecretValueResponse{}
	_body, _err := client.DoRequest(tea.String("PutSecretValue"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutSecretValue(request *PutSecretValueRequest) (_result *PutSecretValueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PutSecretValueResponse{}
	_body, _err := client.PutSecretValueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecretWithOptions(request *DeleteSecretRequest, runtime *util.RuntimeOptions) (_result *DeleteSecretResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteSecret"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecret(request *DeleteSecretRequest) (_result *DeleteSecretResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecretResponse{}
	_body, _err := client.DeleteSecretWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSecretVersionStageWithOptions(request *UpdateSecretVersionStageRequest, runtime *util.RuntimeOptions) (_result *UpdateSecretVersionStageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateSecretVersionStage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSecretVersionStage(request *UpdateSecretVersionStageRequest) (_result *UpdateSecretVersionStageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateSecretVersionStageResponse{}
	_body, _err := client.UpdateSecretVersionStageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsymmetricDecryptWithOptions(request *AsymmetricDecryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricDecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.DoRequest(tea.String("AsymmetricDecrypt"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsymmetricDecrypt(request *AsymmetricDecryptRequest) (_result *AsymmetricDecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricDecryptResponse{}
	_body, _err := client.AsymmetricDecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsymmetricVerifyWithOptions(request *AsymmetricVerifyRequest, runtime *util.RuntimeOptions) (_result *AsymmetricVerifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.DoRequest(tea.String("AsymmetricVerify"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsymmetricVerify(request *AsymmetricVerifyRequest) (_result *AsymmetricVerifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricVerifyResponse{}
	_body, _err := client.AsymmetricVerifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsymmetricSignWithOptions(request *AsymmetricSignRequest, runtime *util.RuntimeOptions) (_result *AsymmetricSignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AsymmetricSignResponse{}
	_body, _err := client.DoRequest(tea.String("AsymmetricSign"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsymmetricSign(request *AsymmetricSignRequest) (_result *AsymmetricSignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricSignResponse{}
	_body, _err := client.AsymmetricSignWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AsymmetricEncryptWithOptions(request *AsymmetricEncryptRequest, runtime *util.RuntimeOptions) (_result *AsymmetricEncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.DoRequest(tea.String("AsymmetricEncrypt"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AsymmetricEncrypt(request *AsymmetricEncryptRequest) (_result *AsymmetricEncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AsymmetricEncryptResponse{}
	_body, _err := client.AsymmetricEncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicKeyWithOptions(request *GetPublicKeyRequest, runtime *util.RuntimeOptions) (_result *GetPublicKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.DoRequest(tea.String("GetPublicKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublicKey(request *GetPublicKeyRequest) (_result *GetPublicKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicKeyResponse{}
	_body, _err := client.GetPublicKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDataKeyWithoutPlaintextWithOptions(request *GenerateDataKeyWithoutPlaintextRequest, runtime *util.RuntimeOptions) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.DoRequest(tea.String("GenerateDataKeyWithoutPlaintext"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDataKeyWithoutPlaintext(request *GenerateDataKeyWithoutPlaintextRequest) (_result *GenerateDataKeyWithoutPlaintextResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDataKeyWithoutPlaintextResponse{}
	_body, _err := client.GenerateDataKeyWithoutPlaintextWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateKeyDescriptionWithOptions(request *UpdateKeyDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateKeyDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateKeyDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateKeyDescription(request *UpdateKeyDescriptionRequest) (_result *UpdateKeyDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateKeyDescriptionResponse{}
	_body, _err := client.UpdateKeyDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeyVersionWithOptions(request *DescribeKeyVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeKeyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKeyVersion(request *DescribeKeyVersionRequest) (_result *DescribeKeyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyVersionResponse{}
	_body, _err := client.DescribeKeyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRotationPolicyWithOptions(request *UpdateRotationPolicyRequest, runtime *util.RuntimeOptions) (_result *UpdateRotationPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateRotationPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRotationPolicy(request *UpdateRotationPolicyRequest) (_result *UpdateRotationPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRotationPolicyResponse{}
	_body, _err := client.UpdateRotationPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListKeyVersionsWithOptions(request *ListKeyVersionsRequest, runtime *util.RuntimeOptions) (_result *ListKeyVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.DoRequest(tea.String("ListKeyVersions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListKeyVersions(request *ListKeyVersionsRequest) (_result *ListKeyVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeyVersionsResponse{}
	_body, _err := client.ListKeyVersionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeyVersionWithOptions(request *CreateKeyVersionRequest, runtime *util.RuntimeOptions) (_result *CreateKeyVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.DoRequest(tea.String("CreateKeyVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKeyVersion(request *CreateKeyVersionRequest) (_result *CreateKeyVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyVersionResponse{}
	_body, _err := client.CreateKeyVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceWithOptions(request *DescribeServiceRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeServiceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeService"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeService(request *DescribeServiceRequest) (_result *DescribeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceResponse{}
	_body, _err := client.DescribeServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAliasWithOptions(request *UpdateAliasRequest, runtime *util.RuntimeOptions) (_result *UpdateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateAliasResponse{}
	_body, _err := client.DoRequest(tea.String("UpdateAlias"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAlias(request *UpdateAliasRequest) (_result *UpdateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAliasResponse{}
	_body, _err := client.UpdateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourceWithOptions(request *UntagResourceRequest, runtime *util.RuntimeOptions) (_result *UntagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourceResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResource(request *UntagResourceRequest) (_result *UntagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourceResponse{}
	_body, _err := client.UntagResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourceWithOptions(request *TagResourceRequest, runtime *util.RuntimeOptions) (_result *TagResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourceResponse{}
	_body, _err := client.DoRequest(tea.String("TagResource"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResource(request *TagResourceRequest) (_result *TagResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourceResponse{}
	_body, _err := client.TagResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ScheduleKeyDeletionWithOptions(request *ScheduleKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *ScheduleKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.DoRequest(tea.String("ScheduleKeyDeletion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ScheduleKeyDeletion(request *ScheduleKeyDeletionRequest) (_result *ScheduleKeyDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ScheduleKeyDeletionResponse{}
	_body, _err := client.ScheduleKeyDeletionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceTagsWithOptions(request *ListResourceTagsRequest, runtime *util.RuntimeOptions) (_result *ListResourceTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.DoRequest(tea.String("ListResourceTags"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceTags(request *ListResourceTagsRequest) (_result *ListResourceTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListResourceTagsResponse{}
	_body, _err := client.ListResourceTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListKeysWithOptions(request *ListKeysRequest, runtime *util.RuntimeOptions) (_result *ListKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListKeysResponse{}
	_body, _err := client.DoRequest(tea.String("ListKeys"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListKeys(request *ListKeysRequest) (_result *ListKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeysResponse{}
	_body, _err := client.ListKeysWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAliasesByKeyIdWithOptions(request *ListAliasesByKeyIdRequest, runtime *util.RuntimeOptions) (_result *ListAliasesByKeyIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.DoRequest(tea.String("ListAliasesByKeyId"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAliasesByKeyId(request *ListAliasesByKeyIdRequest) (_result *ListAliasesByKeyIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliasesByKeyIdResponse{}
	_body, _err := client.ListAliasesByKeyIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAliasesWithOptions(request *ListAliasesRequest, runtime *util.RuntimeOptions) (_result *ListAliasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListAliasesResponse{}
	_body, _err := client.DoRequest(tea.String("ListAliases"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAliases(request *ListAliasesRequest) (_result *ListAliasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAliasesResponse{}
	_body, _err := client.ListAliasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeyMaterialWithOptions(request *ImportKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *ImportKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.DoRequest(tea.String("ImportKeyMaterial"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportKeyMaterial(request *ImportKeyMaterialRequest) (_result *ImportKeyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyMaterialResponse{}
	_body, _err := client.ImportKeyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetParametersForImportWithOptions(request *GetParametersForImportRequest, runtime *util.RuntimeOptions) (_result *GetParametersForImportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.DoRequest(tea.String("GetParametersForImport"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetParametersForImport(request *GetParametersForImportRequest) (_result *GetParametersForImportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetParametersForImportResponse{}
	_body, _err := client.GetParametersForImportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateDataKeyWithOptions(request *GenerateDataKeyRequest, runtime *util.RuntimeOptions) (_result *GenerateDataKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.DoRequest(tea.String("GenerateDataKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateDataKey(request *GenerateDataKeyRequest) (_result *GenerateDataKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GenerateDataKeyResponse{}
	_body, _err := client.GenerateDataKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EncryptWithOptions(request *EncryptRequest, runtime *util.RuntimeOptions) (_result *EncryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EncryptResponse{}
	_body, _err := client.DoRequest(tea.String("Encrypt"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Encrypt(request *EncryptRequest) (_result *EncryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EncryptResponse{}
	_body, _err := client.EncryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableKeyWithOptions(request *EnableKeyRequest, runtime *util.RuntimeOptions) (_result *EnableKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &EnableKeyResponse{}
	_body, _err := client.DoRequest(tea.String("EnableKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableKey(request *EnableKeyRequest) (_result *EnableKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableKeyResponse{}
	_body, _err := client.EnableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableKeyWithOptions(request *DisableKeyRequest, runtime *util.RuntimeOptions) (_result *DisableKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DisableKeyResponse{}
	_body, _err := client.DoRequest(tea.String("DisableKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableKey(request *DisableKeyRequest) (_result *DisableKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableKeyResponse{}
	_body, _err := client.DisableKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKeyWithOptions(request *DescribeKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKeyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKey(request *DescribeKeyRequest) (_result *DescribeKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKeyResponse{}
	_body, _err := client.DescribeKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeyMaterialWithOptions(request *DeleteKeyMaterialRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyMaterialResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteKeyMaterial"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyMaterial(request *DeleteKeyMaterialRequest) (_result *DeleteKeyMaterialResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyMaterialResponse{}
	_body, _err := client.DeleteKeyMaterialWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAliasWithOptions(request *DeleteAliasRequest, runtime *util.RuntimeOptions) (_result *DeleteAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteAliasResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteAlias"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlias(request *DeleteAliasRequest) (_result *DeleteAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAliasResponse{}
	_body, _err := client.DeleteAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DecryptWithOptions(request *DecryptRequest, runtime *util.RuntimeOptions) (_result *DecryptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DecryptResponse{}
	_body, _err := client.DoRequest(tea.String("Decrypt"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Decrypt(request *DecryptRequest) (_result *DecryptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DecryptResponse{}
	_body, _err := client.DecryptWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateKeyWithOptions(request *CreateKeyRequest, runtime *util.RuntimeOptions) (_result *CreateKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateKeyResponse{}
	_body, _err := client.DoRequest(tea.String("CreateKey"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateKey(request *CreateKeyRequest) (_result *CreateKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateKeyResponse{}
	_body, _err := client.CreateKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAliasWithOptions(request *CreateAliasRequest, runtime *util.RuntimeOptions) (_result *CreateAliasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAliasResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAlias"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlias(request *CreateAliasRequest) (_result *CreateAliasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAliasResponse{}
	_body, _err := client.CreateAliasWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelKeyDeletionWithOptions(request *CancelKeyDeletionRequest, runtime *util.RuntimeOptions) (_result *CancelKeyDeletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.DoRequest(tea.String("CancelKeyDeletion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-01-20"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (_result *CancelKeyDeletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelKeyDeletionResponse{}
	_body, _err := client.CancelKeyDeletionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
