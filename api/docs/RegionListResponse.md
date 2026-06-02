# RegionListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]RegionItemV1**](RegionItemV1.md) | Complete region listing, sorted by id ascending. No pagination. | 

## Methods

### NewRegionListResponse

`func NewRegionListResponse(items []RegionItemV1, ) *RegionListResponse`

NewRegionListResponse instantiates a new RegionListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionListResponseWithDefaults

`func NewRegionListResponseWithDefaults() *RegionListResponse`

NewRegionListResponseWithDefaults instantiates a new RegionListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *RegionListResponse) GetItems() []RegionItemV1`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RegionListResponse) GetItemsOk() (*[]RegionItemV1, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RegionListResponse) SetItems(v []RegionItemV1)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


