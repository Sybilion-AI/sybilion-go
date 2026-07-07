# HealthGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Overall health — \&quot;ok\&quot; when all components are healthy, \&quot;degraded\&quot; otherwise. | [optional] 
**Components** | Pointer to [**map[string]HealthGet200ResponseComponentsValue**](HealthGet200ResponseComponentsValue.md) | Per-component status map. Keys are functional names (e.g. data_storage, workflow_engine). | [optional] 

## Methods

### NewHealthGet200Response

`func NewHealthGet200Response() *HealthGet200Response`

NewHealthGet200Response instantiates a new HealthGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthGet200ResponseWithDefaults

`func NewHealthGet200ResponseWithDefaults() *HealthGet200Response`

NewHealthGet200ResponseWithDefaults instantiates a new HealthGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthGet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthGet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthGet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthGet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponents

`func (o *HealthGet200Response) GetComponents() map[string]HealthGet200ResponseComponentsValue`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HealthGet200Response) GetComponentsOk() (*map[string]HealthGet200ResponseComponentsValue, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HealthGet200Response) SetComponents(v map[string]HealthGet200ResponseComponentsValue)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *HealthGet200Response) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


