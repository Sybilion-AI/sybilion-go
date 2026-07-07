# Filters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **[]int32** | Thematic category ids to filter by; each must be an integer **1–9999** inclusive. | [optional] 
**Limit** | Pointer to **int32** | Maximum number of items to return. When omitted, a per-environment default is applied (100 by default). The maximum accepted value is operator-configurable (default 1000). | [optional] 
**Regions** | Pointer to **[]int32** | Geographic region ids to filter by; each must be an integer **1–9999** inclusive. | [optional] 

## Methods

### NewFilters

`func NewFilters() *Filters`

NewFilters instantiates a new Filters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiltersWithDefaults

`func NewFiltersWithDefaults() *Filters`

NewFiltersWithDefaults instantiates a new Filters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *Filters) GetCategories() []int32`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *Filters) GetCategoriesOk() (*[]int32, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *Filters) SetCategories(v []int32)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *Filters) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetLimit

`func (o *Filters) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Filters) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Filters) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *Filters) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRegions

`func (o *Filters) GetRegions() []int32`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Filters) GetRegionsOk() (*[]int32, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Filters) SetRegions(v []int32)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Filters) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


