/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A BlockRequest is utilized to make a block request on the /block endpoint.
type BlockRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	BlockIdentifier *PartialBlockIdentifier `json:"block_identifier"`
}
