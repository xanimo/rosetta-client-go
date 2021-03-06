# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchTransactions**](SearchApi.md#SearchTransactions) | **Post** /search/transactions | [INDEXER] Search for Transactions

# **SearchTransactions**
> SearchTransactionsResponse SearchTransactions(ctx, body)
[INDEXER] Search for Transactions

`/search/transactions` allows the caller to search for transactions that meet certain conditions. Some conditions include matching a transaction hash, containing an operation with a certain status, or containing an operation that affects a certain account. `/search/transactions` is considered an \"indexer\" endpoint and Rosetta implementations are not required to complete it to adhere to the Rosetta spec. However, any Rosetta \"indexer\" MUST support this endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchTransactionsRequest**](SearchTransactionsRequest.md)|  | 

### Return type

[**SearchTransactionsResponse**](SearchTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

