/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ConstructionHashRequest is the input to the `/construction/hash` endpoint.
type ConstructionHashRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	SignedTransaction string `json:"signed_transaction"`
}