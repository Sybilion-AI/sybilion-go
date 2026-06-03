# CategoryListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]CategoryItemV1**](CategoryItemV1.md) | Complete category listing, sorted by id ascending. No pagination. | 

## Methods

### NewCategoryListResponse

`func NewCategoryListResponse(items []CategoryItemV1, ) *CategoryListResponse`

NewCategoryListResponse instantiates a new CategoryListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryListResponseWithDefaults

`func NewCategoryListResponseWithDefaults() *CategoryListResponse`

NewCategoryListResponseWithDefaults instantiates a new CategoryListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *CategoryListResponse) GetItems() []CategoryItemV1`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CategoryListResponse) GetItemsOk() (*[]CategoryItemV1, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CategoryListResponse) SetItems(v []CategoryItemV1)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


