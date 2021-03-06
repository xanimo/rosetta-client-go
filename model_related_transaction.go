/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The related_transaction allows implementations to link together multiple transactions. An unpopulated network identifier indicates that the related transaction is on the same network.
type RelatedTransaction struct {
	NetworkIdentifier *NetworkIdentifier `json:"network_identifier,omitempty"`
	TransactionIdentifier *TransactionIdentifier `json:"transaction_identifier"`
	Direction *Direction `json:"direction"`
}
