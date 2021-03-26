/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Transactions contain an array of Operations that are attributable to the same TransactionIdentifier.
type Transaction struct {
	TransactionIdentifier *TransactionIdentifier `json:"transaction_identifier"`
	Operations []Operation `json:"operations"`
	RelatedTransactions []RelatedTransaction `json:"related_transactions,omitempty"`
	// Transactions that are related to other transactions (like a cross-shard transaction) should include the tranaction_identifier of these transactions in the metadata.
	Metadata *interface{} `json:"metadata,omitempty"`
}