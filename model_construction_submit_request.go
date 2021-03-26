/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The transaction submission request includes a signed transaction.
type ConstructionSubmitRequest struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier"`
	SignedTransaction string `json:"signed_transaction"`
}