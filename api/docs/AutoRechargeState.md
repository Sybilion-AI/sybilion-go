# AutoRechargeState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BelowEurCents** | **int64** | When the available balance drops below this many EUR cents, a recharge is triggered. | 
**Enabled** | **bool** | Whether auto-recharge is active for this account. | 
**HasStripeCustomer** | **bool** | Whether a Stripe customer record exists (required for auto-recharge to run). | 
**MeterCents** | **int64** | EUR cents charged via auto-recharge in the current UTC calendar month. | 
**MeterMonth** | Pointer to **NullableString** | UTC month start for &#x60;meter_cents&#x60;; null if no auto-recharge has run this month. | [optional] 
**MonthlyCapCents** | **int64** | Maximum EUR cents that may be charged via auto-recharge per UTC calendar month. 0 &#x3D; no cap. | 
**TargetEurCents** | **int64** | Balance target after a successful recharge, in EUR cents. | 

## Methods

### NewAutoRechargeState

`func NewAutoRechargeState(belowEurCents int64, enabled bool, hasStripeCustomer bool, meterCents int64, monthlyCapCents int64, targetEurCents int64, ) *AutoRechargeState`

NewAutoRechargeState instantiates a new AutoRechargeState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRechargeStateWithDefaults

`func NewAutoRechargeStateWithDefaults() *AutoRechargeState`

NewAutoRechargeStateWithDefaults instantiates a new AutoRechargeState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBelowEurCents

`func (o *AutoRechargeState) GetBelowEurCents() int64`

GetBelowEurCents returns the BelowEurCents field if non-nil, zero value otherwise.

### GetBelowEurCentsOk

`func (o *AutoRechargeState) GetBelowEurCentsOk() (*int64, bool)`

GetBelowEurCentsOk returns a tuple with the BelowEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelowEurCents

`func (o *AutoRechargeState) SetBelowEurCents(v int64)`

SetBelowEurCents sets BelowEurCents field to given value.


### GetEnabled

`func (o *AutoRechargeState) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRechargeState) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRechargeState) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasStripeCustomer

`func (o *AutoRechargeState) GetHasStripeCustomer() bool`

GetHasStripeCustomer returns the HasStripeCustomer field if non-nil, zero value otherwise.

### GetHasStripeCustomerOk

`func (o *AutoRechargeState) GetHasStripeCustomerOk() (*bool, bool)`

GetHasStripeCustomerOk returns a tuple with the HasStripeCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasStripeCustomer

`func (o *AutoRechargeState) SetHasStripeCustomer(v bool)`

SetHasStripeCustomer sets HasStripeCustomer field to given value.


### GetMeterCents

`func (o *AutoRechargeState) GetMeterCents() int64`

GetMeterCents returns the MeterCents field if non-nil, zero value otherwise.

### GetMeterCentsOk

`func (o *AutoRechargeState) GetMeterCentsOk() (*int64, bool)`

GetMeterCentsOk returns a tuple with the MeterCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterCents

`func (o *AutoRechargeState) SetMeterCents(v int64)`

SetMeterCents sets MeterCents field to given value.


### GetMeterMonth

`func (o *AutoRechargeState) GetMeterMonth() string`

GetMeterMonth returns the MeterMonth field if non-nil, zero value otherwise.

### GetMeterMonthOk

`func (o *AutoRechargeState) GetMeterMonthOk() (*string, bool)`

GetMeterMonthOk returns a tuple with the MeterMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterMonth

`func (o *AutoRechargeState) SetMeterMonth(v string)`

SetMeterMonth sets MeterMonth field to given value.

### HasMeterMonth

`func (o *AutoRechargeState) HasMeterMonth() bool`

HasMeterMonth returns a boolean if a field has been set.

### SetMeterMonthNil

`func (o *AutoRechargeState) SetMeterMonthNil(b bool)`

 SetMeterMonthNil sets the value for MeterMonth to be an explicit nil

### UnsetMeterMonth
`func (o *AutoRechargeState) UnsetMeterMonth()`

UnsetMeterMonth ensures that no value is present for MeterMonth, not even an explicit nil
### GetMonthlyCapCents

`func (o *AutoRechargeState) GetMonthlyCapCents() int64`

GetMonthlyCapCents returns the MonthlyCapCents field if non-nil, zero value otherwise.

### GetMonthlyCapCentsOk

`func (o *AutoRechargeState) GetMonthlyCapCentsOk() (*int64, bool)`

GetMonthlyCapCentsOk returns a tuple with the MonthlyCapCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCapCents

`func (o *AutoRechargeState) SetMonthlyCapCents(v int64)`

SetMonthlyCapCents sets MonthlyCapCents field to given value.


### GetTargetEurCents

`func (o *AutoRechargeState) GetTargetEurCents() int64`

GetTargetEurCents returns the TargetEurCents field if non-nil, zero value otherwise.

### GetTargetEurCentsOk

`func (o *AutoRechargeState) GetTargetEurCentsOk() (*int64, bool)`

GetTargetEurCentsOk returns a tuple with the TargetEurCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEurCents

`func (o *AutoRechargeState) SetTargetEurCents(v int64)`

SetTargetEurCents sets TargetEurCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


