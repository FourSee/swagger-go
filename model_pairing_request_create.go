/*
 * PolyRhythm
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiGo

// Used to create a PairingRequest
type PairingRequestCreate struct {
	// A base64 encoded public key
	PublicKey string `json:"public_key"`
	// For now, must be `non_push`. Only non-push devices can initiate a pairing request.
	Platform string `json:"platform"`
	// A user-friendly device name
	DeviceName string `json:"device_name"`
}
