/*
 * flows_filtered Mitmproxy2Swagger
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package woolworths

type InlineResponse20026 struct {
	Totals *interface{} `json:"Totals,omitempty"`
	AvailableItems []interface{} `json:"AvailableItems,omitempty"`
	UnavailableItems []Object `json:"UnavailableItems,omitempty"`
	ExceededAtpLimitItems []Object `json:"ExceededAtpLimitItems,omitempty"`
	UnavailableByMarketOutOfStockItems []Object `json:"UnavailableByMarketOutOfStockItems,omitempty"`
	RestrictedItems []Object `json:"RestrictedItems,omitempty"`
	BundleItems []Object `json:"BundleItems,omitempty"`
	RestrictedByDeliveryMethodItems []Object `json:"RestrictedByDeliveryMethodItems,omitempty"`
	RestrictedByDeliPlattersItems []Object `json:"RestrictedByDeliPlattersItems,omitempty"`
	ExceededSupplyLimitItems []Object `json:"ExceededSupplyLimitItems,omitempty"`
	DeliveryFee *interface{} `json:"DeliveryFee,omitempty"`
	MarketShippingFees *interface{} `json:"MarketShippingFees,omitempty"`
	MarketDeliveryDetails *interface{} `json:"MarketDeliveryDetails,omitempty"`
	PotentialOrderPromotions []interface{} `json:"PotentialOrderPromotions,omitempty"`
	Message *interface{} `json:"Message,omitempty"`
	RestrictedByDeliveryMethodMessage *interface{} `json:"RestrictedByDeliveryMethodMessage,omitempty"`
	RestrictedByDeliPlattersMessage *interface{} `json:"RestrictedByDeliPlattersMessage,omitempty"`
	RestrictedByDeliveryWindowMessage *interface{} `json:"RestrictedByDeliveryWindowMessage,omitempty"`
	ProductGroupSupplyLimitWindowWarningMessage *interface{} `json:"ProductGroupSupplyLimitWindowWarningMessage,omitempty"`
	TrolleyItemCount float64 `json:"TrolleyItemCount,omitempty"`
	TotalTrolleyItemQuantity float64 `json:"TotalTrolleyItemQuantity,omitempty"`
	TotalWowDeliveryTrolleyItemQuantity float64 `json:"TotalWowDeliveryTrolleyItemQuantity,omitempty"`
	MaximumOrderQuantity float64 `json:"MaximumOrderQuantity,omitempty"`
	WowRewardsSummary *interface{} `json:"WowRewardsSummary,omitempty"`
	TaxRate float64 `json:"TaxRate,omitempty"`
	Boosted bool `json:"Boosted,omitempty"`
}
