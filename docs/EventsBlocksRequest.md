# EventsBlocksRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkIdentifier** | [***NetworkIdentifier**](NetworkIdentifier.md) |  | [default to null]
**Offset** | **int64** | offset is the offset into the event stream to sync events from. If this field is not populated, we return the limit events backwards from tip. If this is set to 0, we start from the beginning. | [optional] [default to null]
**Limit** | **int64** | limit is the maximum number of events to fetch in one call. The implementation may return &lt;&#x3D; limit events. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

