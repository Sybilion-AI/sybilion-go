# ApiV1ForecastsIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Artifacts** | Pointer to [**[]ForecastArtifactMeta**](ForecastArtifactMeta.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EurCentsFinal** | Pointer to **NullableInt64** | Final settled charge for the job, in EUR cents. Null until settled. | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**PipelineError** | Pointer to **map[string]interface{}** | Present for some settled failed or canceled jobs when the pipeline wrote &#x60;error.json&#x60; under &#x60;forecasts/{id}/&#x60; or &#x60;forecasts/{id}/output/&#x60; in the artifact bucket. Shape is defined by the pipeline; omitted when missing or unreadable (bodies larger than 64 KiB are omitted). May be a JSON object or array depending on the pipeline.  | [optional] 
**PipelineType** | Pointer to **string** | Today only \&quot;forecast\&quot;; future pipeline types will appear here. | [optional] 
**RunId** | Pointer to **NullableString** | Internal run id (omitted for jobs that haven&#39;t started yet). | [optional] 
**Settled** | Pointer to **bool** |  | [optional] 
**SettledAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** | Internal workflow id (omitted for jobs that haven&#39;t started yet). | [optional] 

## Methods

### NewApiV1ForecastsIdGet200Response

`func NewApiV1ForecastsIdGet200Response() *ApiV1ForecastsIdGet200Response`

NewApiV1ForecastsIdGet200Response instantiates a new ApiV1ForecastsIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ForecastsIdGet200ResponseWithDefaults

`func NewApiV1ForecastsIdGet200ResponseWithDefaults() *ApiV1ForecastsIdGet200Response`

NewApiV1ForecastsIdGet200ResponseWithDefaults instantiates a new ApiV1ForecastsIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifacts

`func (o *ApiV1ForecastsIdGet200Response) GetArtifacts() []ForecastArtifactMeta`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ApiV1ForecastsIdGet200Response) GetArtifactsOk() (*[]ForecastArtifactMeta, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ApiV1ForecastsIdGet200Response) SetArtifacts(v []ForecastArtifactMeta)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ApiV1ForecastsIdGet200Response) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiV1ForecastsIdGet200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiV1ForecastsIdGet200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiV1ForecastsIdGet200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiV1ForecastsIdGet200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEurCentsFinal

`func (o *ApiV1ForecastsIdGet200Response) GetEurCentsFinal() int64`

GetEurCentsFinal returns the EurCentsFinal field if non-nil, zero value otherwise.

### GetEurCentsFinalOk

`func (o *ApiV1ForecastsIdGet200Response) GetEurCentsFinalOk() (*int64, bool)`

GetEurCentsFinalOk returns a tuple with the EurCentsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEurCentsFinal

`func (o *ApiV1ForecastsIdGet200Response) SetEurCentsFinal(v int64)`

SetEurCentsFinal sets EurCentsFinal field to given value.

### HasEurCentsFinal

`func (o *ApiV1ForecastsIdGet200Response) HasEurCentsFinal() bool`

HasEurCentsFinal returns a boolean if a field has been set.

### SetEurCentsFinalNil

`func (o *ApiV1ForecastsIdGet200Response) SetEurCentsFinalNil(b bool)`

 SetEurCentsFinalNil sets the value for EurCentsFinal to be an explicit nil

### UnsetEurCentsFinal
`func (o *ApiV1ForecastsIdGet200Response) UnsetEurCentsFinal()`

UnsetEurCentsFinal ensures that no value is present for EurCentsFinal, not even an explicit nil
### GetJobId

`func (o *ApiV1ForecastsIdGet200Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApiV1ForecastsIdGet200Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApiV1ForecastsIdGet200Response) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApiV1ForecastsIdGet200Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPipelineError

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineError() map[string]interface{}`

GetPipelineError returns the PipelineError field if non-nil, zero value otherwise.

### GetPipelineErrorOk

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineErrorOk() (*map[string]interface{}, bool)`

GetPipelineErrorOk returns a tuple with the PipelineError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineError

`func (o *ApiV1ForecastsIdGet200Response) SetPipelineError(v map[string]interface{})`

SetPipelineError sets PipelineError field to given value.

### HasPipelineError

`func (o *ApiV1ForecastsIdGet200Response) HasPipelineError() bool`

HasPipelineError returns a boolean if a field has been set.

### SetPipelineErrorNil

`func (o *ApiV1ForecastsIdGet200Response) SetPipelineErrorNil(b bool)`

 SetPipelineErrorNil sets the value for PipelineError to be an explicit nil

### UnsetPipelineError
`func (o *ApiV1ForecastsIdGet200Response) UnsetPipelineError()`

UnsetPipelineError ensures that no value is present for PipelineError, not even an explicit nil
### GetPipelineType

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineType() string`

GetPipelineType returns the PipelineType field if non-nil, zero value otherwise.

### GetPipelineTypeOk

`func (o *ApiV1ForecastsIdGet200Response) GetPipelineTypeOk() (*string, bool)`

GetPipelineTypeOk returns a tuple with the PipelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineType

`func (o *ApiV1ForecastsIdGet200Response) SetPipelineType(v string)`

SetPipelineType sets PipelineType field to given value.

### HasPipelineType

`func (o *ApiV1ForecastsIdGet200Response) HasPipelineType() bool`

HasPipelineType returns a boolean if a field has been set.

### GetRunId

`func (o *ApiV1ForecastsIdGet200Response) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ApiV1ForecastsIdGet200Response) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ApiV1ForecastsIdGet200Response) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ApiV1ForecastsIdGet200Response) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### SetRunIdNil

`func (o *ApiV1ForecastsIdGet200Response) SetRunIdNil(b bool)`

 SetRunIdNil sets the value for RunId to be an explicit nil

### UnsetRunId
`func (o *ApiV1ForecastsIdGet200Response) UnsetRunId()`

UnsetRunId ensures that no value is present for RunId, not even an explicit nil
### GetSettled

`func (o *ApiV1ForecastsIdGet200Response) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *ApiV1ForecastsIdGet200Response) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *ApiV1ForecastsIdGet200Response) SetSettled(v bool)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *ApiV1ForecastsIdGet200Response) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetSettledAt

`func (o *ApiV1ForecastsIdGet200Response) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *ApiV1ForecastsIdGet200Response) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *ApiV1ForecastsIdGet200Response) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *ApiV1ForecastsIdGet200Response) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### SetSettledAtNil

`func (o *ApiV1ForecastsIdGet200Response) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *ApiV1ForecastsIdGet200Response) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetStatus

`func (o *ApiV1ForecastsIdGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiV1ForecastsIdGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiV1ForecastsIdGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiV1ForecastsIdGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWorkflowId

`func (o *ApiV1ForecastsIdGet200Response) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *ApiV1ForecastsIdGet200Response) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *ApiV1ForecastsIdGet200Response) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *ApiV1ForecastsIdGet200Response) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *ApiV1ForecastsIdGet200Response) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *ApiV1ForecastsIdGet200Response) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


