# AlertItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Human-readable name of the dataset or index that triggered the alert. | [optional] 
**News** | Pointer to [**[]NewsItemV1**](NewsItemV1.md) | Related news articles driving this alert. | [optional] 
**PctChange** | Pointer to **float64** | Percentage change that triggered the alert (negative &#x3D; decline, positive &#x3D; surge). | [optional] 
**Trending** | Pointer to **bool** | Whether this alert is currently trending across the platform. | [optional] 

## Methods

### NewAlertItemV1

`func NewAlertItemV1() *AlertItemV1`

NewAlertItemV1 instantiates a new AlertItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertItemV1WithDefaults

`func NewAlertItemV1WithDefaults() *AlertItemV1`

NewAlertItemV1WithDefaults instantiates a new AlertItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertItemV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertItemV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertItemV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AlertItemV1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNews

`func (o *AlertItemV1) GetNews() []NewsItemV1`

GetNews returns the News field if non-nil, zero value otherwise.

### GetNewsOk

`func (o *AlertItemV1) GetNewsOk() (*[]NewsItemV1, bool)`

GetNewsOk returns a tuple with the News field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNews

`func (o *AlertItemV1) SetNews(v []NewsItemV1)`

SetNews sets News field to given value.

### HasNews

`func (o *AlertItemV1) HasNews() bool`

HasNews returns a boolean if a field has been set.

### GetPctChange

`func (o *AlertItemV1) GetPctChange() float64`

GetPctChange returns the PctChange field if non-nil, zero value otherwise.

### GetPctChangeOk

`func (o *AlertItemV1) GetPctChangeOk() (*float64, bool)`

GetPctChangeOk returns a tuple with the PctChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPctChange

`func (o *AlertItemV1) SetPctChange(v float64)`

SetPctChange sets PctChange field to given value.

### HasPctChange

`func (o *AlertItemV1) HasPctChange() bool`

HasPctChange returns a boolean if a field has been set.

### GetTrending

`func (o *AlertItemV1) GetTrending() bool`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *AlertItemV1) GetTrendingOk() (*bool, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *AlertItemV1) SetTrending(v bool)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *AlertItemV1) HasTrending() bool`

HasTrending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


