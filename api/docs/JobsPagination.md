# JobsPagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** | Page size echoed back from the request. | 
**Order** | **string** | Sort direction echoed back from the request. | 
**Page** | **int32** | 1-indexed current page number echoed back from the request. | 
**Sort** | **string** | Column the rows are sorted by, echoed back from the request. | 
**Total** | **int64** | Total matching rows for the authenticated user (full set, not just this page). | 
**TotalPages** | **int64** | ceil(total / limit). Zero when total is zero. | 

## Methods

### NewJobsPagination

`func NewJobsPagination(limit int32, order string, page int32, sort string, total int64, totalPages int64, ) *JobsPagination`

NewJobsPagination instantiates a new JobsPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobsPaginationWithDefaults

`func NewJobsPaginationWithDefaults() *JobsPagination`

NewJobsPaginationWithDefaults instantiates a new JobsPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *JobsPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *JobsPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *JobsPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOrder

`func (o *JobsPagination) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *JobsPagination) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *JobsPagination) SetOrder(v string)`

SetOrder sets Order field to given value.


### GetPage

`func (o *JobsPagination) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *JobsPagination) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *JobsPagination) SetPage(v int32)`

SetPage sets Page field to given value.


### GetSort

`func (o *JobsPagination) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *JobsPagination) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *JobsPagination) SetSort(v string)`

SetSort sets Sort field to given value.


### GetTotal

`func (o *JobsPagination) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *JobsPagination) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *JobsPagination) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetTotalPages

`func (o *JobsPagination) GetTotalPages() int64`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *JobsPagination) GetTotalPagesOk() (*int64, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *JobsPagination) SetTotalPages(v int64)`

SetTotalPages sets TotalPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


