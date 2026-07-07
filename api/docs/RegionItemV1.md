# RegionItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Integer identifier. Use this value in filters.regions[]. | 
**Latitude** | **float64** | Geographic latitude (0.0 when not applicable). | 
**Longitude** | **float64** | Geographic longitude (0.0 when not applicable). | 
**Name** | **string** | Human-readable region label. | 

## Methods

### NewRegionItemV1

`func NewRegionItemV1(id int32, latitude float64, longitude float64, name string, ) *RegionItemV1`

NewRegionItemV1 instantiates a new RegionItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionItemV1WithDefaults

`func NewRegionItemV1WithDefaults() *RegionItemV1`

NewRegionItemV1WithDefaults instantiates a new RegionItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionItemV1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionItemV1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionItemV1) SetId(v int32)`

SetId sets Id field to given value.


### GetLatitude

`func (o *RegionItemV1) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RegionItemV1) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RegionItemV1) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *RegionItemV1) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RegionItemV1) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RegionItemV1) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetName

`func (o *RegionItemV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegionItemV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegionItemV1) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


