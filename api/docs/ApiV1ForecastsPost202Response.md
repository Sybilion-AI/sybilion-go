# ApiV1ForecastsPost202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | Unique job id — use this to poll status and download artifacts. | [optional] 
**PollUrl** | Pointer to **string** | Convenience URL for polling this job&#39;s status. | [optional] 
**RunId** | Pointer to **string** | Internal run identifier (opaque; useful for support). | [optional] 
**Workflow** | Pointer to **string** | Internal workflow identifier (opaque; useful for support). | [optional] 

## Methods

### NewApiV1ForecastsPost202Response

`func NewApiV1ForecastsPost202Response() *ApiV1ForecastsPost202Response`

NewApiV1ForecastsPost202Response instantiates a new ApiV1ForecastsPost202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ForecastsPost202ResponseWithDefaults

`func NewApiV1ForecastsPost202ResponseWithDefaults() *ApiV1ForecastsPost202Response`

NewApiV1ForecastsPost202ResponseWithDefaults instantiates a new ApiV1ForecastsPost202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *ApiV1ForecastsPost202Response) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApiV1ForecastsPost202Response) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApiV1ForecastsPost202Response) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApiV1ForecastsPost202Response) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetPollUrl

`func (o *ApiV1ForecastsPost202Response) GetPollUrl() string`

GetPollUrl returns the PollUrl field if non-nil, zero value otherwise.

### GetPollUrlOk

`func (o *ApiV1ForecastsPost202Response) GetPollUrlOk() (*string, bool)`

GetPollUrlOk returns a tuple with the PollUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollUrl

`func (o *ApiV1ForecastsPost202Response) SetPollUrl(v string)`

SetPollUrl sets PollUrl field to given value.

### HasPollUrl

`func (o *ApiV1ForecastsPost202Response) HasPollUrl() bool`

HasPollUrl returns a boolean if a field has been set.

### GetRunId

`func (o *ApiV1ForecastsPost202Response) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *ApiV1ForecastsPost202Response) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *ApiV1ForecastsPost202Response) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *ApiV1ForecastsPost202Response) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetWorkflow

`func (o *ApiV1ForecastsPost202Response) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ApiV1ForecastsPost202Response) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ApiV1ForecastsPost202Response) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ApiV1ForecastsPost202Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


