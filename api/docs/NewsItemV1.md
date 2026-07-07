# NewsItemV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Topical category of the article (e.g. &#x60;world&#x60;, &#x60;business&#x60;, &#x60;energy&#x60;). | [optional] 
**Description** | Pointer to **string** | Short summary of the article. | [optional] 
**PublishedAt** | Pointer to **time.Time** | Publication timestamp (RFC 3339 / ISO 8601). | [optional] 
**SourceName** | Pointer to **string** | Name of the publication or media outlet. | [optional] 
**Title** | Pointer to **string** | Headline of the news article. | [optional] 
**Trending** | Pointer to **bool** | Whether this article is currently trending across the platform. | [optional] 
**Url** | Pointer to **string** | Canonical URL of the article. | [optional] 

## Methods

### NewNewsItemV1

`func NewNewsItemV1() *NewsItemV1`

NewNewsItemV1 instantiates a new NewsItemV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewsItemV1WithDefaults

`func NewNewsItemV1WithDefaults() *NewsItemV1`

NewNewsItemV1WithDefaults instantiates a new NewsItemV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *NewsItemV1) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NewsItemV1) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NewsItemV1) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NewsItemV1) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *NewsItemV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewsItemV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewsItemV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewsItemV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPublishedAt

`func (o *NewsItemV1) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *NewsItemV1) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *NewsItemV1) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *NewsItemV1) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetSourceName

`func (o *NewsItemV1) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *NewsItemV1) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *NewsItemV1) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *NewsItemV1) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### GetTitle

`func (o *NewsItemV1) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewsItemV1) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NewsItemV1) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NewsItemV1) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTrending

`func (o *NewsItemV1) GetTrending() bool`

GetTrending returns the Trending field if non-nil, zero value otherwise.

### GetTrendingOk

`func (o *NewsItemV1) GetTrendingOk() (*bool, bool)`

GetTrendingOk returns a tuple with the Trending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrending

`func (o *NewsItemV1) SetTrending(v bool)`

SetTrending sets Trending field to given value.

### HasTrending

`func (o *NewsItemV1) HasTrending() bool`

HasTrending returns a boolean if a field has been set.

### GetUrl

`func (o *NewsItemV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NewsItemV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NewsItemV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NewsItemV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


