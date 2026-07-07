# EuroTranche

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**ExpiresAt** | **time.Time** |  | 
**Id** | **string** |  | 
**InitialEurCents** | **int64** | Original size of this tranche in EUR cents. | 
**RemainingEurCents** | **int64** | Unconsumed balance remaining in this tranche, in EUR cents. | 
**Source** | **string** | Origin of the tranche — one of &#x60;signup_trial&#x60;, &#x60;stripe&#x60;, &#x60;partner&#x60;, &#x60;legacy&#x60;. Other labels may appear for custom grants. | 

## Methods

### NewEuroTranche

`func NewEuroTranche(createdAt time.Time, expiresAt time.Time, id string, initialEurCents int64, remainingEurCents int64, source string, ) *EuroTranche`

NewEuroTranche instantiates a new EuroTranche object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEuroTrancheWithDefaults

`func NewEuroTrancheWithDefaults() *EuroTranche`

NewEuroTrancheWithDefaults instantiates a new EuroTranche object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *EuroTranche) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EuroTranche) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EuroTranche) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *EuroTranche) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *EuroTranche) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *EuroTranche) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *EuroTranche) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EuroTranche) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EuroTranche) SetId(v string)`

SetId sets Id field to given value.


### GetInitialEurCents

`func (o *EuroTranche) GetInitialEurCents() int64`

GetInitialEurCents returns the InitialEurCents field if non-nil, zero value otherwise.

### GetInitialEurCentsOk

`func (o *EuroTranche) GetInitialEurCentsOk() (*int64, bool)`

GetInitialEurCentsOk returns a tuple with the InitialEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialEurCents

`func (o *EuroTranche) SetInitialEurCents(v int64)`

SetInitialEurCents sets InitialEurCents field to given value.


### GetRemainingEurCents

`func (o *EuroTranche) GetRemainingEurCents() int64`

GetRemainingEurCents returns the RemainingEurCents field if non-nil, zero value otherwise.

### GetRemainingEurCentsOk

`func (o *EuroTranche) GetRemainingEurCentsOk() (*int64, bool)`

GetRemainingEurCentsOk returns a tuple with the RemainingEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingEurCents

`func (o *EuroTranche) SetRemainingEurCents(v int64)`

SetRemainingEurCents sets RemainingEurCents field to given value.


### GetSource

`func (o *EuroTranche) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *EuroTranche) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *EuroTranche) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


