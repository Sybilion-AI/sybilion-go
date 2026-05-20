# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV1CategoriesGet**](DefaultAPI.md#ApiV1CategoriesGet) | **Get** /api/v1/categories | List categories
[**ApiV1DriversPost**](DefaultAPI.md#ApiV1DriversPost) | **Post** /api/v1/drivers | Recommend ranked driver-dataset candidates (synchronous proxy)
[**ApiV1ForecastsIdArtifactsNameGet**](DefaultAPI.md#ApiV1ForecastsIdArtifactsNameGet) | **Get** /api/v1/forecasts/{id}/artifacts/{name} | Stream a single output file (proxied from internal artifact store; no gs:// URLs)
[**ApiV1ForecastsIdGet**](DefaultAPI.md#ApiV1ForecastsIdGet) | **Get** /api/v1/forecasts/{id} | Poll forecast job status and artifact metadata (downloads via artifacts sub-path)
[**ApiV1ForecastsPost**](DefaultAPI.md#ApiV1ForecastsPost) | **Post** /api/v1/forecasts | Start async forecast job
[**ApiV1JobsGet**](DefaultAPI.md#ApiV1JobsGet) | **Get** /api/v1/jobs | Paginated list of the caller&#39;s async jobs
[**ApiV1MeGet**](DefaultAPI.md#ApiV1MeGet) | **Get** /api/v1/me | Current user, balances, and usage tier
[**ApiV1RegionsGet**](DefaultAPI.md#ApiV1RegionsGet) | **Get** /api/v1/regions | List regions
[**ApiV1UsageGet**](DefaultAPI.md#ApiV1UsageGet) | **Get** /api/v1/usage | Paginated usage (charged tasks)
[**HealthGet**](DefaultAPI.md#HealthGet) | **Get** /health | Health check



## ApiV1CategoriesGet

> CatalogListResponse ApiV1CategoriesGet(ctx).Execute()

List categories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1CategoriesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1CategoriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1CategoriesGet`: CatalogListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1CategoriesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1CategoriesGetRequest struct via the builder pattern


### Return type

[**CatalogListResponse**](CatalogListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1DriversPost

> ApiV1DriversPost(ctx).RecommendRequestV1(recommendRequestV1).Execute()

Recommend ranked driver-dataset candidates (synchronous proxy)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	recommendRequestV1 := *openapiclient.NewRecommendRequestV1(float64(123), *openapiclient.NewTimeseriesMetadata("Title_example"), "Version_example") // RecommendRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ApiV1DriversPost(context.Background()).RecommendRequestV1(recommendRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1DriversPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1DriversPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recommendRequestV1** | [**RecommendRequestV1**](RecommendRequestV1.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsIdArtifactsNameGet

> *os.File ApiV1ForecastsIdArtifactsNameGet(ctx, id, name).Execute()

Stream a single output file (proxied from internal artifact store; no gs:// URLs)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsIdArtifactsNameGet(context.Background(), id, name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsIdArtifactsNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsIdArtifactsNameGet`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsIdArtifactsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsIdArtifactsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsIdGet

> ApiV1ForecastsIdGet200Response ApiV1ForecastsIdGet(ctx, id).Execute()

Poll forecast job status and artifact metadata (downloads via artifacts sub-path)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsIdGet`: ApiV1ForecastsIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiV1ForecastsIdGet200Response**](ApiV1ForecastsIdGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1ForecastsPost

> ApiV1ForecastsPost202Response ApiV1ForecastsPost(ctx).ForecastRequestV1(forecastRequestV1).Execute()

Start async forecast job



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	forecastRequestV1 := *openapiclient.NewForecastRequestV1("Frequency_example", "PipelineVersion_example", float64(123), map[string]float32{"key": float32(123)}, *openapiclient.NewTimeseriesMetadata("Title_example")) // ForecastRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1ForecastsPost(context.Background()).ForecastRequestV1(forecastRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1ForecastsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1ForecastsPost`: ApiV1ForecastsPost202Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1ForecastsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1ForecastsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forecastRequestV1** | [**ForecastRequestV1**](ForecastRequestV1.md) |  | 

### Return type

[**ApiV1ForecastsPost202Response**](ApiV1ForecastsPost202Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1JobsGet

> ApiV1JobsGet200Response ApiV1JobsGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Status(status).PipelineType(pipelineType).Execute()

Paginated list of the caller's async jobs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	limit := int32(56) // int32 | Page size. Capped at 200. (optional) (default to 50)
	sort := "sort_example" // string |  (optional) (default to "created_at")
	order := "order_example" // string |  (optional) (default to "desc")
	status := "status_example" // string | Optional status filter. (optional)
	pipelineType := "pipelineType_example" // string | Optional pipeline type filter (today only \"forecast\" is emitted). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1JobsGet(context.Background()).Page(page).Limit(limit).Sort(sort).Order(order).Status(status).PipelineType(pipelineType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1JobsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1JobsGet`: ApiV1JobsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1JobsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1JobsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** | Page size. Capped at 200. | [default to 50]
 **sort** | **string** |  | [default to &quot;created_at&quot;]
 **order** | **string** |  | [default to &quot;desc&quot;]
 **status** | **string** | Optional status filter. | 
 **pipelineType** | **string** | Optional pipeline type filter (today only \&quot;forecast\&quot; is emitted). | 

### Return type

[**ApiV1JobsGet200Response**](ApiV1JobsGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1MeGet

> MeResponse ApiV1MeGet(ctx).Execute()

Current user, balances, and usage tier

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1MeGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1MeGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1MeGet`: MeResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1MeGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1MeGetRequest struct via the builder pattern


### Return type

[**MeResponse**](MeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1RegionsGet

> CatalogListResponse ApiV1RegionsGet(ctx).Execute()

List regions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1RegionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1RegionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1RegionsGet`: CatalogListResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1RegionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1RegionsGetRequest struct via the builder pattern


### Return type

[**CatalogListResponse**](CatalogListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1UsageGet

> ApiV1UsageGet200Response ApiV1UsageGet(ctx).Page(page).Limit(limit).Sort(sort).Order(order).Execute()

Paginated usage (charged tasks)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {
	page := int32(56) // int32 | 1-indexed page number. (optional) (default to 1)
	limit := int32(56) // int32 | Page size. Capped at 200. (optional) (default to 50)
	sort := "sort_example" // string | Column to sort by. (optional) (default to "id")
	order := "order_example" // string | Sort direction. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ApiV1UsageGet(context.Background()).Page(page).Limit(limit).Sort(sort).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV1UsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1UsageGet`: ApiV1UsageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ApiV1UsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiV1UsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | 1-indexed page number. | [default to 1]
 **limit** | **int32** | Page size. Capped at 200. | [default to 50]
 **sort** | **string** | Column to sort by. | [default to &quot;id&quot;]
 **order** | **string** | Sort direction. | [default to &quot;desc&quot;]

### Return type

[**ApiV1UsageGet200Response**](ApiV1UsageGet200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HealthGet

> map[string]interface{} HealthGet(ctx).Execute()

Health check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "go.sybilion.dev/sybilion/api"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.HealthGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HealthGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.HealthGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthGetRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

