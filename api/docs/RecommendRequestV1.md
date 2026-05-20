# RecommendRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**Filters**](Filters.md) | Optional. Same validation as **&#x60;POST /api/v1/forecasts&#x60;**: each **&#x60;categories[]&#x60;** and **&#x60;regions[]&#x60;** entry is an integer **1–9999** (inclusive); optional **&#x60;limit&#x60;** is **0–10000**. Values are not verified against catalog APIs.  | [optional] 
**RecencyFactor** | **float64** |  | 
**Timeseries** | Pointer to **map[string]float32** | Optional. Map of YYYY-MM-DD date keys to numeric observation values. When supplied, all keys must parse as YYYY-MM-DD and all values must be finite. Unlike &#x60;/forecasts&#x60;, this endpoint is frequency-agnostic — there is no monthly alignment, gap detection, or 60-point minimum. When omitted, the handler does not forward the field to the upstream Recommend service at all.  | [optional] 
**TimeseriesMetadata** | [**TimeseriesMetadata**](TimeseriesMetadata.md) |  | 
**Version** | **string** | Recommend pipeline version. Closed set; only v1 is supported today. Used **locally** to select the per-version validator and is **not forwarded** to the upstream Recommend service.  | 

## Methods

### NewRecommendRequestV1

`func NewRecommendRequestV1(recencyFactor float64, timeseriesMetadata TimeseriesMetadata, version string, ) *RecommendRequestV1`

NewRecommendRequestV1 instantiates a new RecommendRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecommendRequestV1WithDefaults

`func NewRecommendRequestV1WithDefaults() *RecommendRequestV1`

NewRecommendRequestV1WithDefaults instantiates a new RecommendRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *RecommendRequestV1) GetFilters() Filters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RecommendRequestV1) GetFiltersOk() (*Filters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RecommendRequestV1) SetFilters(v Filters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RecommendRequestV1) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRecencyFactor

`func (o *RecommendRequestV1) GetRecencyFactor() float64`

GetRecencyFactor returns the RecencyFactor field if non-nil, zero value otherwise.

### GetRecencyFactorOk

`func (o *RecommendRequestV1) GetRecencyFactorOk() (*float64, bool)`

GetRecencyFactorOk returns a tuple with the RecencyFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecencyFactor

`func (o *RecommendRequestV1) SetRecencyFactor(v float64)`

SetRecencyFactor sets RecencyFactor field to given value.


### GetTimeseries

`func (o *RecommendRequestV1) GetTimeseries() map[string]float32`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *RecommendRequestV1) GetTimeseriesOk() (*map[string]float32, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *RecommendRequestV1) SetTimeseries(v map[string]float32)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *RecommendRequestV1) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetTimeseriesMetadata

`func (o *RecommendRequestV1) GetTimeseriesMetadata() TimeseriesMetadata`

GetTimeseriesMetadata returns the TimeseriesMetadata field if non-nil, zero value otherwise.

### GetTimeseriesMetadataOk

`func (o *RecommendRequestV1) GetTimeseriesMetadataOk() (*TimeseriesMetadata, bool)`

GetTimeseriesMetadataOk returns a tuple with the TimeseriesMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseriesMetadata

`func (o *RecommendRequestV1) SetTimeseriesMetadata(v TimeseriesMetadata)`

SetTimeseriesMetadata sets TimeseriesMetadata field to given value.


### GetVersion

`func (o *RecommendRequestV1) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RecommendRequestV1) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RecommendRequestV1) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


