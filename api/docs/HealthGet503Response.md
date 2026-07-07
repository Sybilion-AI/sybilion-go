# HealthGet503Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**map[string]HealthGet503ResponseComponentsValue**](HealthGet503ResponseComponentsValue.md) |  | [optional] 

## Methods

### NewHealthGet503Response

`func NewHealthGet503Response() *HealthGet503Response`

NewHealthGet503Response instantiates a new HealthGet503Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthGet503ResponseWithDefaults

`func NewHealthGet503ResponseWithDefaults() *HealthGet503Response`

NewHealthGet503ResponseWithDefaults instantiates a new HealthGet503Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *HealthGet503Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthGet503Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthGet503Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthGet503Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComponents

`func (o *HealthGet503Response) GetComponents() map[string]HealthGet503ResponseComponentsValue`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *HealthGet503Response) GetComponentsOk() (*map[string]HealthGet503ResponseComponentsValue, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *HealthGet503Response) SetComponents(v map[string]HealthGet503ResponseComponentsValue)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *HealthGet503Response) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


