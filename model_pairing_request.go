/*
 * PolyRhythm
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapiGo
import (
	"time"
)

// The root of the PairingRequest type's schema.
type PairingRequest struct {
	// Unique ID for this pairing request.
	Id string `json:"id"`
	// Ideally, a human-readable name for the device that created the pairing request.  Will default to using the `hostname`.
	DeviceName string `json:"device_name"`
	// A URL-safe base62-encoded representation of the ID
	ShortId string `json:"short_id"`
	// Status of the pairing request
	Status string `json:"status"`
	// If this pairing request has been accepted, this will be the base64 encoded public cyrpto key of the device that accepted the pairing reques.
	AcceptedCryptoKey string `json:"accepted_crypto_key,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	// The URL to accept this pairing request.  Will use the `polyrhythm` scheme, and follow this pattern: `polyrhythm://pr/{short_id}`
	AcceptUrl string `json:"accept_url"`
	// IP address where the pairing request came from.
	IpAddress string `json:"ip_address"`
}
