/*
 * flows_filtered Mitmproxy2Swagger
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package woolworths

type Apiv3uitrolleyupdateItems struct {
	Stockcode float64 `json:"stockcode,omitempty"`
	Diagnostics string `json:"diagnostics,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Source string `json:"source,omitempty"`
	OfferId *interface{} `json:"offerId,omitempty"`
	SearchTerm string `json:"searchTerm,omitempty"`
	ProfileId *interface{} `json:"profileId,omitempty"`
	PriceLevel *interface{} `json:"priceLevel,omitempty"`
}