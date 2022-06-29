# {{classname}}

All URIs are relative to *https://www.woolworths.com.au*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiUiV2BootstrapGet**](DefaultApi.md#ApiUiV2BootstrapGet) | **Get** /api/ui/v2/bootstrap | GET bootstrap
[**ApiV3UiAuthenticationOtpPost**](DefaultApi.md#ApiV3UiAuthenticationOtpPost) | **Post** /api/v3/ui/authentication/otp | POST otp
[**ApiV3UiCheckoutDigitalpayInstrumentsGet**](DefaultApi.md#ApiV3UiCheckoutDigitalpayInstrumentsGet) | **Get** /api/v3/ui/checkout/digitalpay/instruments | GET instruments
[**ApiV3UiFulfilmentFeestructureIdGet**](DefaultApi.md#ApiV3UiFulfilmentFeestructureIdGet) | **Get** /api/v3/ui/fulfilment/feestructure/{id} | GET feestructure by id
[**ApiV3UiFulfilmentWindowsGet**](DefaultApi.md#ApiV3UiFulfilmentWindowsGet) | **Get** /api/v3/ui/fulfilment/windows | GET windows
[**ApiV3UiMylistsGet**](DefaultApi.md#ApiV3UiMylistsGet) | **Get** /api/v3/ui/mylists | GET mylists
[**ApiV3UiOrderAmendingGet**](DefaultApi.md#ApiV3UiOrderAmendingGet) | **Get** /api/v3/ui/order/amending | GET amending
[**ApiV3UiPackagingPreferencesGet**](DefaultApi.md#ApiV3UiPackagingPreferencesGet) | **Get** /api/v3/ui/PackagingPreferences | GET packagingpreferences
[**ApiV3UiPackagingPreferencesIdPost**](DefaultApi.md#ApiV3UiPackagingPreferencesIdPost) | **Post** /api/v3/ui/PackagingPreferences/{id} | POST packagingpreferences by id
[**ApiV3UiTrolleyUpdatePost**](DefaultApi.md#ApiV3UiTrolleyUpdatePost) | **Post** /api/v3/ui/trolley/update | POST trolley update
[**ApisUiCheckoutConfirmOrderGet**](DefaultApi.md#ApisUiCheckoutConfirmOrderGet) | **Get** /apis/ui/Checkout/ConfirmOrder | GET confirmorder
[**ApisUiCheckoutDigitalPayInitializeCreditCardIframePost**](DefaultApi.md#ApisUiCheckoutDigitalPayInitializeCreditCardIframePost) | **Post** /apis/ui/Checkout/DigitalPay/InitializeCreditCardIframe | POST initializecreditcardiframe
[**ApisUiCheckoutDigitalPayPaymentPost**](DefaultApi.md#ApisUiCheckoutDigitalPayPaymentPost) | **Post** /apis/ui/Checkout/DigitalPay/Payment | POST payment
[**ApisUiCheckoutGet**](DefaultApi.md#ApisUiCheckoutGet) | **Get** /apis/ui/Checkout | GET checkout
[**ApisUiDeliveryDeliveryInfoGet**](DefaultApi.md#ApisUiDeliveryDeliveryInfoGet) | **Get** /apis/ui/Delivery/DeliveryInfo | GET deliveryinfo
[**ApisUiDeliveryOptionsGet**](DefaultApi.md#ApisUiDeliveryOptionsGet) | **Get** /apis/ui/Delivery/Options | GET options
[**ApisUiDeliveryOptionsPost**](DefaultApi.md#ApisUiDeliveryOptionsPost) | **Post** /apis/ui/Delivery/Options | POST options
[**ApisUiFulfilmentPost**](DefaultApi.md#ApisUiFulfilmentPost) | **Post** /apis/ui/Fulfilment | POST fulfilment
[**ApisUiInspirationCardsGet**](DefaultApi.md#ApisUiInspirationCardsGet) | **Get** /apis/ui/inspiration/cards | GET cards
[**ApisUiLoginLoginwithcredentialPost**](DefaultApi.md#ApisUiLoginLoginwithcredentialPost) | **Post** /apis/ui/login/loginwithcredential | POST loginwithcredential
[**ApisUiPastshopsListGet**](DefaultApi.md#ApisUiPastshopsListGet) | **Get** /apis/ui/pastshops/list | GET list
[**ApisUiPiesCategoriesWithSpecialsGet**](DefaultApi.md#ApisUiPiesCategoriesWithSpecialsGet) | **Get** /apis/ui/PiesCategoriesWithSpecials | GET piescategorieswithspecials
[**ApisUiProductHaveyouforgottenGet**](DefaultApi.md#ApisUiProductHaveyouforgottenGet) | **Get** /apis/ui/product/haveyouforgotten | GET haveyouforgotten
[**ApisUiSearchProductsPost**](DefaultApi.md#ApisUiSearchProductsPost) | **Post** /apis/ui/Search/products | POST products
[**ApisUiSearchSuggestionGet**](DefaultApi.md#ApisUiSearchSuggestionGet) | **Get** /apis/ui/Search/suggestion | GET suggestion
[**ApisUiSeoMetatagsPost**](DefaultApi.md#ApisUiSeoMetatagsPost) | **Post** /apis/ui/SeoMetatags | POST seometatags
[**ApisUiSettingsGet**](DefaultApi.md#ApisUiSettingsGet) | **Get** /apis/ui/settings | GET settings
[**ApisUiShopperAuthenticationMethodGet**](DefaultApi.md#ApisUiShopperAuthenticationMethodGet) | **Get** /apis/ui/shopper/AuthenticationMethod | GET authenticationmethod
[**ApisUiTrolleyGet**](DefaultApi.md#ApisUiTrolleyGet) | **Get** /apis/ui/Trolley | GET trolley
[**ApisUiTrolleyRemovePost**](DefaultApi.md#ApisUiTrolleyRemovePost) | **Post** /apis/ui/Trolley/Remove | POST remove
[**ApisUiTrolleyUpdateOrRemoveItemsPost**](DefaultApi.md#ApisUiTrolleyUpdateOrRemoveItemsPost) | **Post** /apis/ui/Trolley/UpdateOrRemoveItems | POST updateorremoveitems
[**ApisUiUserPreferencesHasDoneCheckoutOnboardingPost**](DefaultApi.md#ApisUiUserPreferencesHasDoneCheckoutOnboardingPost) | **Post** /apis/ui/UserPreferences/HasDoneCheckoutOnboarding | POST hasdonecheckoutonboarding
[**ApisUiV2SearchCountGet**](DefaultApi.md#ApisUiV2SearchCountGet) | **Get** /apis/ui/v2/Search/count | GET count

# **ApiUiV2BootstrapGet**
> InlineResponse200 ApiUiV2BootstrapGet(ctx, )
GET bootstrap

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiAuthenticationOtpPost**
> InlineResponse2006 ApiV3UiAuthenticationOtpPost(ctx, optional)
POST otp

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApiV3UiAuthenticationOtpPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApiV3UiAuthenticationOtpPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AuthenticationOtpBody**](AuthenticationOtpBody.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiCheckoutDigitalpayInstrumentsGet**
> InlineResponse20020 ApiV3UiCheckoutDigitalpayInstrumentsGet(ctx, optional)
GET instruments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApiV3UiCheckoutDigitalpayInstrumentsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApiV3UiCheckoutDigitalpayInstrumentsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isForPaymentAgreement** | **optional.String**|  | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiFulfilmentFeestructureIdGet**
> InlineResponse20024 ApiV3UiFulfilmentFeestructureIdGet(ctx, id)
GET feestructure by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiFulfilmentWindowsGet**
> InlineResponse20017 ApiV3UiFulfilmentWindowsGet(ctx, optional)
GET windows

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApiV3UiFulfilmentWindowsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApiV3UiFulfilmentWindowsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **areaId** | **optional.Float64**|  | 
 **fulfilmentMethod** | **optional.String**|  | 
 **addressId** | **optional.Float64**|  | 
 **suburbId** | **optional.Float64**|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiMylistsGet**
> InlineResponse20013 ApiV3UiMylistsGet(ctx, )
GET mylists

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiOrderAmendingGet**
> InlineResponse2007 ApiV3UiOrderAmendingGet(ctx, )
GET amending

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiPackagingPreferencesGet**
> InlineResponse20025 ApiV3UiPackagingPreferencesGet(ctx, )
GET packagingpreferences

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiPackagingPreferencesIdPost**
> InlineResponse20025 ApiV3UiPackagingPreferencesIdPost(ctx, id, optional)
POST packagingpreferences by id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**|  | 
 **optional** | ***DefaultApiApiV3UiPackagingPreferencesIdPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApiV3UiPackagingPreferencesIdPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of interface{}**](interface{}.md)|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiV3UiTrolleyUpdatePost**
> InlineResponse20014 ApiV3UiTrolleyUpdatePost(ctx, optional)
POST trolley update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApiV3UiTrolleyUpdatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApiV3UiTrolleyUpdatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TrolleyUpdateBody**](TrolleyUpdateBody.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiCheckoutConfirmOrderGet**
> InlineResponse20029 ApisUiCheckoutConfirmOrderGet(ctx, optional)
GET confirmorder

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiCheckoutConfirmOrderGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiCheckoutConfirmOrderGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderId** | **optional.Float64**|  | 

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiCheckoutDigitalPayInitializeCreditCardIframePost**
> InlineResponse20021 ApisUiCheckoutDigitalPayInitializeCreditCardIframePost(ctx, )
POST initializecreditcardiframe

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiCheckoutDigitalPayPaymentPost**
> InlineResponse20027 ApisUiCheckoutDigitalPayPaymentPost(ctx, optional)
POST payment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiCheckoutDigitalPayPaymentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiCheckoutDigitalPayPaymentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DigitalPayPaymentBody**](DigitalPayPaymentBody.md)|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiCheckoutGet**
> InlineResponse20016 ApisUiCheckoutGet(ctx, )
GET checkout

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiDeliveryDeliveryInfoGet**
> InlineResponse20023 ApisUiDeliveryDeliveryInfoGet(ctx, )
GET deliveryinfo

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiDeliveryOptionsGet**
> InlineResponse20018 ApisUiDeliveryOptionsGet(ctx, )
GET options

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiDeliveryOptionsPost**
> InlineResponse20019 ApisUiDeliveryOptionsPost(ctx, optional)
POST options

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiDeliveryOptionsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiDeliveryOptionsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of DeliveryOptionsBody**](DeliveryOptionsBody.md)|  | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiFulfilmentPost**
> InlineResponse20022 ApisUiFulfilmentPost(ctx, optional)
POST fulfilment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiFulfilmentPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiFulfilmentPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UiFulfilmentBody**](UiFulfilmentBody.md)|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiInspirationCardsGet**
> InlineResponse20012 ApisUiInspirationCardsGet(ctx, optional)
GET cards

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiInspirationCardsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiInspirationCardsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **url** | **optional.String**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiLoginLoginwithcredentialPost**
> InlineResponse2005 ApisUiLoginLoginwithcredentialPost(ctx, optional)
POST loginwithcredential

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiLoginLoginwithcredentialPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiLoginLoginwithcredentialPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of LoginLoginwithcredentialBody**](LoginLoginwithcredentialBody.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiPastshopsListGet**
> InlineResponse20028 ApisUiPastshopsListGet(ctx, )
GET list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiPiesCategoriesWithSpecialsGet**
> InlineResponse2002 ApisUiPiesCategoriesWithSpecialsGet(ctx, )
GET piescategorieswithspecials

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiProductHaveyouforgottenGet**
> []Object ApisUiProductHaveyouforgottenGet(ctx, )
GET haveyouforgotten

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiSearchProductsPost**
> InlineResponse20011 ApisUiSearchProductsPost(ctx, optional)
POST products

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiSearchProductsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiSearchProductsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SearchProductsBody**](SearchProductsBody.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiSearchSuggestionGet**
> InlineResponse2009 ApisUiSearchSuggestionGet(ctx, optional)
GET suggestion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiSearchSuggestionGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiSearchSuggestionGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **optional.String**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiSeoMetatagsPost**
> InlineResponse2003 ApisUiSeoMetatagsPost(ctx, optional)
POST seometatags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiSeoMetatagsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiSeoMetatagsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UiSeoMetatagsBody**](UiSeoMetatagsBody.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiSettingsGet**
> []InlineResponse2001 ApisUiSettingsGet(ctx, )
GET settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiShopperAuthenticationMethodGet**
> InlineResponse2004 ApisUiShopperAuthenticationMethodGet(ctx, )
GET authenticationmethod

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiTrolleyGet**
> InlineResponse20015 ApisUiTrolleyGet(ctx, )
GET trolley

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiTrolleyRemovePost**
> InlineResponse2008 ApisUiTrolleyRemovePost(ctx, optional)
POST remove

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiTrolleyRemovePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiTrolleyRemovePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TrolleyRemoveBody**](TrolleyRemoveBody.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiTrolleyUpdateOrRemoveItemsPost**
> InlineResponse20026 ApisUiTrolleyUpdateOrRemoveItemsPost(ctx, optional)
POST updateorremoveitems

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiTrolleyUpdateOrRemoveItemsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiTrolleyUpdateOrRemoveItemsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TrolleyUpdateOrRemoveItemsBody**](TrolleyUpdateOrRemoveItemsBody.md)|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiUserPreferencesHasDoneCheckoutOnboardingPost**
> ApisUiUserPreferencesHasDoneCheckoutOnboardingPost(ctx, optional)
POST hasdonecheckoutonboarding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiUserPreferencesHasDoneCheckoutOnboardingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiUserPreferencesHasDoneCheckoutOnboardingPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserPreferencesHasDoneCheckoutOnboardingBody**](UserPreferencesHasDoneCheckoutOnboardingBody.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApisUiV2SearchCountGet**
> InlineResponse20010 ApisUiV2SearchCountGet(ctx, optional)
GET count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiApisUiV2SearchCountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiApisUiV2SearchCountGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchTerm** | **optional.String**|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

