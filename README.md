# Go API client for swagger-go

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./swagger-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://foursee-polyrhythm.herokuapp.com:443/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DevicesApi* | [**CreateDevice**](docs/DevicesApi.md#createdevice) | **Post** /devices | Create a Device
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /devices/{deviceId} | Delete a Device
*DevicesApi* | [**GetDevice**](docs/DevicesApi.md#getdevice) | **Get** /devices/{deviceId} | Get a Device
*DevicesApi* | [**GetDevices**](docs/DevicesApi.md#getdevices) | **Get** /devices | List All Devices
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Put** /devices/{deviceId} | Update a Device
*MessagesApi* | [**CreateMessage**](docs/MessagesApi.md#createmessage) | **Post** /messages | Create a Message
*MessagesApi* | [**DeleteMessage**](docs/MessagesApi.md#deletemessage) | **Delete** /messages/{messageId} | Delete a Message
*MessagesApi* | [**GetMessage**](docs/MessagesApi.md#getmessage) | **Get** /messages/{messageId} | Get a Message
*MessagesApi* | [**GetMessages**](docs/MessagesApi.md#getmessages) | **Get** /messages | List All Messages
*PairingRequestsApi* | [**CreatePairingRequest**](docs/PairingRequestsApi.md#createpairingrequest) | **Post** /pairing_requests | Create a PairingRequest
*PairingRequestsApi* | [**DeletePairingRequest**](docs/PairingRequestsApi.md#deletepairingrequest) | **Delete** /pairing_requests/{pairingRequestId} | Delete a PairingRequest
*PairingRequestsApi* | [**GetPairingRequest**](docs/PairingRequestsApi.md#getpairingrequest) | **Get** /pairing_requests/{pairingRequestId} | Get a PairingRequest
*PairingRequestsApi* | [**GetPairingRequests**](docs/PairingRequestsApi.md#getpairingrequests) | **Get** /pairing_requests | List All PairingRequests
*PairingRequestsApi* | [**UpdatePairingRequest**](docs/PairingRequestsApi.md#updatepairingrequest) | **Put** /pairing_requests/{pairingRequestId} | Update a PairingRequest
*PairingsApi* | [**CreatePairing**](docs/PairingsApi.md#createpairing) | **Post** /pairings | Create a Pairing
*PairingsApi* | [**DeletePairing**](docs/PairingsApi.md#deletepairing) | **Delete** /pairings/{pairingId} | Delete a Pairing
*PairingsApi* | [**GetPairing**](docs/PairingsApi.md#getpairing) | **Get** /pairings/{pairingId} | Get a Pairing
*PairingsApi* | [**GetPairings**](docs/PairingsApi.md#getpairings) | **Get** /pairings | List All Pairings
*PairingsApi* | [**UpdatePairing**](docs/PairingsApi.md#updatepairing) | **Put** /pairings/{pairingId} | Update a Pairing
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /users | Create a User
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{userId} | Delete a User
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{userId} | Get a User
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /users | List All Users
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /users/{userId} | Update a User


## Documentation For Models

 - [Device](docs/Device.md)
 - [DevicesList](docs/DevicesList.md)
 - [Message](docs/Message.md)
 - [MessagesList](docs/MessagesList.md)
 - [PaginationData](docs/PaginationData.md)
 - [Pairing](docs/Pairing.md)
 - [PairingRequest](docs/PairingRequest.md)
 - [PairingRequestCreate](docs/PairingRequestCreate.md)
 - [PairingRequestsList](docs/PairingRequestsList.md)
 - [PairingsList](docs/PairingsList.md)
 - [User](docs/User.md)


## Documentation For Authorization

## Auth0
- **Type**: OAuth
- **Flow**: implicit
- **Authorization URL**: https://polyrhythm.auth0.com/authorize
- **Scopes**: 
 - **openid**: 
 - **profile**: 
 - **email**: 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```
## agent_auth_key
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author



