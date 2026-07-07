# JobSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**EurCentsFinal** | Pointer to **NullableInt64** | Final settled charge for the job in EUR cents. Null until the job reaches a terminal state. | [optional] 
**JobId** | **string** |  | 
**PipelineType** | **string** | Pipeline that produced this job — currently always &#x60;forecast&#x60;. | 
**RunId** | Pointer to **string** | Opaque internal run identifier. Omitted for jobs that have not started yet; include in support requests. | [optional] 
**Settled** | **bool** | True once the job has reached a terminal state and the charge has been posted. | 
**SettledAt** | Pointer to **NullableTime** |  | [optional] 
**Status** | **string** |  | 
**TerminalReason** | Pointer to **NullableString** | Human-readable failure message for &#x60;failed&#x60; or &#x60;canceled&#x60; jobs; null for non-terminal statuses or cleanly-canceled jobs. | [optional] 
**WorkflowId** | Pointer to **string** | Opaque internal workflow identifier. Omitted for jobs that have not started yet; include in support requests. | [optional] 

## Methods

### NewJobSummary

`func NewJobSummary(createdAt time.Time, jobId string, pipelineType string, settled bool, status string, ) *JobSummary`

NewJobSummary instantiates a new JobSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSummaryWithDefaults

`func NewJobSummaryWithDefaults() *JobSummary`

NewJobSummaryWithDefaults instantiates a new JobSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *JobSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JobSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JobSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEurCentsFinal

`func (o *JobSummary) GetEurCentsFinal() int64`

GetEurCentsFinal returns the EurCentsFinal field if non-nil, zero value otherwise.

### GetEurCentsFinalOk

`func (o *JobSummary) GetEurCentsFinalOk() (*int64, bool)`

GetEurCentsFinalOk returns a tuple with the EurCentsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEurCentsFinal

`func (o *JobSummary) SetEurCentsFinal(v int64)`

SetEurCentsFinal sets EurCentsFinal field to given value.

### HasEurCentsFinal

`func (o *JobSummary) HasEurCentsFinal() bool`

HasEurCentsFinal returns a boolean if a field has been set.

### SetEurCentsFinalNil

`func (o *JobSummary) SetEurCentsFinalNil(b bool)`

 SetEurCentsFinalNil sets the value for EurCentsFinal to be an explicit nil

### UnsetEurCentsFinal
`func (o *JobSummary) UnsetEurCentsFinal()`

UnsetEurCentsFinal ensures that no value is present for EurCentsFinal, not even an explicit nil
### GetJobId

`func (o *JobSummary) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobSummary) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobSummary) SetJobId(v string)`

SetJobId sets JobId field to given value.


### GetPipelineType

`func (o *JobSummary) GetPipelineType() string`

GetPipelineType returns the PipelineType field if non-nil, zero value otherwise.

### GetPipelineTypeOk

`func (o *JobSummary) GetPipelineTypeOk() (*string, bool)`

GetPipelineTypeOk returns a tuple with the PipelineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineType

`func (o *JobSummary) SetPipelineType(v string)`

SetPipelineType sets PipelineType field to given value.


### GetRunId

`func (o *JobSummary) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *JobSummary) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *JobSummary) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *JobSummary) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetSettled

`func (o *JobSummary) GetSettled() bool`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *JobSummary) GetSettledOk() (*bool, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *JobSummary) SetSettled(v bool)`

SetSettled sets Settled field to given value.


### GetSettledAt

`func (o *JobSummary) GetSettledAt() time.Time`

GetSettledAt returns the SettledAt field if non-nil, zero value otherwise.

### GetSettledAtOk

`func (o *JobSummary) GetSettledAtOk() (*time.Time, bool)`

GetSettledAtOk returns a tuple with the SettledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledAt

`func (o *JobSummary) SetSettledAt(v time.Time)`

SetSettledAt sets SettledAt field to given value.

### HasSettledAt

`func (o *JobSummary) HasSettledAt() bool`

HasSettledAt returns a boolean if a field has been set.

### SetSettledAtNil

`func (o *JobSummary) SetSettledAtNil(b bool)`

 SetSettledAtNil sets the value for SettledAt to be an explicit nil

### UnsetSettledAt
`func (o *JobSummary) UnsetSettledAt()`

UnsetSettledAt ensures that no value is present for SettledAt, not even an explicit nil
### GetStatus

`func (o *JobSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTerminalReason

`func (o *JobSummary) GetTerminalReason() string`

GetTerminalReason returns the TerminalReason field if non-nil, zero value otherwise.

### GetTerminalReasonOk

`func (o *JobSummary) GetTerminalReasonOk() (*string, bool)`

GetTerminalReasonOk returns a tuple with the TerminalReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminalReason

`func (o *JobSummary) SetTerminalReason(v string)`

SetTerminalReason sets TerminalReason field to given value.

### HasTerminalReason

`func (o *JobSummary) HasTerminalReason() bool`

HasTerminalReason returns a boolean if a field has been set.

### SetTerminalReasonNil

`func (o *JobSummary) SetTerminalReasonNil(b bool)`

 SetTerminalReasonNil sets the value for TerminalReason to be an explicit nil

### UnsetTerminalReason
`func (o *JobSummary) UnsetTerminalReason()`

UnsetTerminalReason ensures that no value is present for TerminalReason, not even an explicit nil
### GetWorkflowId

`func (o *JobSummary) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *JobSummary) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *JobSummary) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *JobSummary) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


