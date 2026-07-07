# UsageEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsyncJobId** | Pointer to **NullableString** | The &#x60;async_jobs.id&#x60; (UUID) for async pipeline charges; null for synchronous endpoint charges. | [optional] 
**CreatedAt** | **string** | Timestamp of the charge (ISO 8601). | 
**CreditsCharged** | **int64** | Whole credits debited for this row before EUR conversion. | 
**Endpoint** | Pointer to **string** | Billing route key (e.g. &#x60;drivers&#x60;, &#x60;alerts&#x60;, or a forecast pipeline type). | [optional] 
**EurCentsCharged** | **int64** | EUR cents debited for this row (1 EUR &#x3D; 100 cents). | 
**Id** | **int64** | Auto-increment row identifier. | 
**Units** | **int64** | Metered quantity — item count for per-unit pricing, 1 for flat-fee strategies. | 

## Methods

### NewUsageEvent

`func NewUsageEvent(createdAt string, creditsCharged int64, eurCentsCharged int64, id int64, units int64, ) *UsageEvent`

NewUsageEvent instantiates a new UsageEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEventWithDefaults

`func NewUsageEventWithDefaults() *UsageEvent`

NewUsageEventWithDefaults instantiates a new UsageEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsyncJobId

`func (o *UsageEvent) GetAsyncJobId() string`

GetAsyncJobId returns the AsyncJobId field if non-nil, zero value otherwise.

### GetAsyncJobIdOk

`func (o *UsageEvent) GetAsyncJobIdOk() (*string, bool)`

GetAsyncJobIdOk returns a tuple with the AsyncJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncJobId

`func (o *UsageEvent) SetAsyncJobId(v string)`

SetAsyncJobId sets AsyncJobId field to given value.

### HasAsyncJobId

`func (o *UsageEvent) HasAsyncJobId() bool`

HasAsyncJobId returns a boolean if a field has been set.

### SetAsyncJobIdNil

`func (o *UsageEvent) SetAsyncJobIdNil(b bool)`

 SetAsyncJobIdNil sets the value for AsyncJobId to be an explicit nil

### UnsetAsyncJobId
`func (o *UsageEvent) UnsetAsyncJobId()`

UnsetAsyncJobId ensures that no value is present for AsyncJobId, not even an explicit nil
### GetCreatedAt

`func (o *UsageEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UsageEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UsageEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreditsCharged

`func (o *UsageEvent) GetCreditsCharged() int64`

GetCreditsCharged returns the CreditsCharged field if non-nil, zero value otherwise.

### GetCreditsChargedOk

`func (o *UsageEvent) GetCreditsChargedOk() (*int64, bool)`

GetCreditsChargedOk returns a tuple with the CreditsCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCharged

`func (o *UsageEvent) SetCreditsCharged(v int64)`

SetCreditsCharged sets CreditsCharged field to given value.


### GetEndpoint

`func (o *UsageEvent) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageEvent) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageEvent) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UsageEvent) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetEurCentsCharged

`func (o *UsageEvent) GetEurCentsCharged() int64`

GetEurCentsCharged returns the EurCentsCharged field if non-nil, zero value otherwise.

### GetEurCentsChargedOk

`func (o *UsageEvent) GetEurCentsChargedOk() (*int64, bool)`

GetEurCentsChargedOk returns a tuple with the EurCentsCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEurCentsCharged

`func (o *UsageEvent) SetEurCentsCharged(v int64)`

SetEurCentsCharged sets EurCentsCharged field to given value.


### GetId

`func (o *UsageEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageEvent) SetId(v int64)`

SetId sets Id field to given value.


### GetUnits

`func (o *UsageEvent) GetUnits() int64`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *UsageEvent) GetUnitsOk() (*int64, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *UsageEvent) SetUnits(v int64)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


