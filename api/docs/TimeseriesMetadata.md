# TimeseriesMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Extended context for the model, up to 2048 characters. More detail improves driver relevance. | [optional] 
**Keywords** | Pointer to **[]string** | Up to 20 semantic tags that help anchor the search to relevant datasets. | [optional] 
**Title** | **string** | Short identifier for the series, 20–511 characters. | 

## Methods

### NewTimeseriesMetadata

`func NewTimeseriesMetadata(title string, ) *TimeseriesMetadata`

NewTimeseriesMetadata instantiates a new TimeseriesMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeseriesMetadataWithDefaults

`func NewTimeseriesMetadataWithDefaults() *TimeseriesMetadata`

NewTimeseriesMetadataWithDefaults instantiates a new TimeseriesMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TimeseriesMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TimeseriesMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TimeseriesMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TimeseriesMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKeywords

`func (o *TimeseriesMetadata) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *TimeseriesMetadata) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *TimeseriesMetadata) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *TimeseriesMetadata) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetTitle

`func (o *TimeseriesMetadata) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimeseriesMetadata) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimeseriesMetadata) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


