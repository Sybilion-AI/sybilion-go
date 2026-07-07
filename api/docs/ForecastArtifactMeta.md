# ForecastArtifactMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | MIME type of the artifact (e.g. &#x60;application/json&#x60;). | [optional] 
**Href** | Pointer to **string** | Relative URL to stream via &#x60;GET /api/v1/forecasts/{id}/artifacts/{name}&#x60;. | [optional] 
**Name** | Pointer to **string** | Artifact filename (e.g. &#x60;forecast.json&#x60;, &#x60;backtest_metrics.json&#x60;). | [optional] 
**Size** | Pointer to **int64** | File size in bytes. | [optional] 

## Methods

### NewForecastArtifactMeta

`func NewForecastArtifactMeta() *ForecastArtifactMeta`

NewForecastArtifactMeta instantiates a new ForecastArtifactMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastArtifactMetaWithDefaults

`func NewForecastArtifactMetaWithDefaults() *ForecastArtifactMeta`

NewForecastArtifactMetaWithDefaults instantiates a new ForecastArtifactMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ForecastArtifactMeta) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ForecastArtifactMeta) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ForecastArtifactMeta) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ForecastArtifactMeta) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetHref

`func (o *ForecastArtifactMeta) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ForecastArtifactMeta) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ForecastArtifactMeta) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ForecastArtifactMeta) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *ForecastArtifactMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ForecastArtifactMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ForecastArtifactMeta) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ForecastArtifactMeta) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *ForecastArtifactMeta) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ForecastArtifactMeta) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ForecastArtifactMeta) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ForecastArtifactMeta) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


