# AlertsRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextEnriched** | **bool** | When true, treat the supplied metadata as already context-enriched. | 
**DateFrom** | Pointer to **string** | Optional start date bound for alert detection (YYYY-MM-DD). | [optional] 
**DateTo** | Pointer to **string** | Optional end date bound for alert detection (YYYY-MM-DD). | [optional] 
**Filters** | Pointer to [**Filters**](Filters.md) | Optional. &#x60;limit&#x60; controls the number of alerts returned (**0–1000**, default **100**). &#x60;categories[]&#x60; and &#x60;regions[]&#x60; narrow the alert universe; each must be an integer **1–9999**. Values are not verified against catalog APIs.  | [optional] 
**Metadata** | [**TimeseriesMetadata**](TimeseriesMetadata.md) |  | 

## Methods

### NewAlertsRequestV1

`func NewAlertsRequestV1(contextEnriched bool, metadata TimeseriesMetadata, ) *AlertsRequestV1`

NewAlertsRequestV1 instantiates a new AlertsRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsRequestV1WithDefaults

`func NewAlertsRequestV1WithDefaults() *AlertsRequestV1`

NewAlertsRequestV1WithDefaults instantiates a new AlertsRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextEnriched

`func (o *AlertsRequestV1) GetContextEnriched() bool`

GetContextEnriched returns the ContextEnriched field if non-nil, zero value otherwise.

### GetContextEnrichedOk

`func (o *AlertsRequestV1) GetContextEnrichedOk() (*bool, bool)`

GetContextEnrichedOk returns a tuple with the ContextEnriched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextEnriched

`func (o *AlertsRequestV1) SetContextEnriched(v bool)`

SetContextEnriched sets ContextEnriched field to given value.


### GetDateFrom

`func (o *AlertsRequestV1) GetDateFrom() string`

GetDateFrom returns the DateFrom field if non-nil, zero value otherwise.

### GetDateFromOk

`func (o *AlertsRequestV1) GetDateFromOk() (*string, bool)`

GetDateFromOk returns a tuple with the DateFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFrom

`func (o *AlertsRequestV1) SetDateFrom(v string)`

SetDateFrom sets DateFrom field to given value.

### HasDateFrom

`func (o *AlertsRequestV1) HasDateFrom() bool`

HasDateFrom returns a boolean if a field has been set.

### GetDateTo

`func (o *AlertsRequestV1) GetDateTo() string`

GetDateTo returns the DateTo field if non-nil, zero value otherwise.

### GetDateToOk

`func (o *AlertsRequestV1) GetDateToOk() (*string, bool)`

GetDateToOk returns a tuple with the DateTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTo

`func (o *AlertsRequestV1) SetDateTo(v string)`

SetDateTo sets DateTo field to given value.

### HasDateTo

`func (o *AlertsRequestV1) HasDateTo() bool`

HasDateTo returns a boolean if a field has been set.

### GetFilters

`func (o *AlertsRequestV1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AlertsRequestV1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AlertsRequestV1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AlertsRequestV1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetMetadata

`func (o *AlertsRequestV1) GetMetadata() TimeseriesMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertsRequestV1) GetMetadataOk() (*TimeseriesMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertsRequestV1) SetMetadata(v TimeseriesMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


