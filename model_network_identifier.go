/*
 * Rosetta
 *
 * Build Once. Integrate Your Blockchain Everywhere.
 *
 * API version: 1.4.10
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The network_identifier specifies which network a particular object is associated with.
type NetworkIdentifier struct {
	Blockchain string `json:"blockchain"`
	// If a blockchain has a specific chain-id or network identifier, it should go in this field. It is up to the client to determine which network-specific identifier is mainnet or testnet.
	Network string `json:"network"`
	SubNetworkIdentifier *SubNetworkIdentifier `json:"sub_network_identifier,omitempty"`
}
