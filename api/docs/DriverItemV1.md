# DriverItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriverName** | Pointer to **string** | Human-readable name of the dataset. | [optional] 
**HashId** | Pointer to **string** | Stable identifier for the dataset; use to reference this driver across requests. | [optional] 
**Score** | Pointer to **float64** | Relevance score indicating how well this dataset explains your timeseries (higher is more relevant). | [optional] 

## Methods

### NewDriverItemV1

`func NewDriverItemV1() *DriverItemV1`

NewDriverItemV1 instantiates a new DriverItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverItemV1WithDefaults

`func NewDriverItemV1WithDefaults() *DriverItemV1`

NewDriverItemV1WithDefaults instantiates a new DriverItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriverName

`func (o *DriverItemV1) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *DriverItemV1) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *DriverItemV1) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *DriverItemV1) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### GetHashId

`func (o *DriverItemV1) GetHashId() string`

GetHashId returns the HashId field if non-nil, zero value otherwise.

### GetHashIdOk

`func (o *DriverItemV1) GetHashIdOk() (*string, bool)`

GetHashIdOk returns a tuple with the HashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashId

`func (o *DriverItemV1) SetHashId(v string)`

SetHashId sets HashId field to given value.

### HasHashId

`func (o *DriverItemV1) HasHashId() bool`

HasHashId returns a boolean if a field has been set.

### GetScore

`func (o *DriverItemV1) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DriverItemV1) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DriverItemV1) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *DriverItemV1) HasScore() bool`

HasScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


