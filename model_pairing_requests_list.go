/*
 * PolyRhythm
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PairingRequestsList struct {
	PairingRequests []PairingRequest `json:"pairing_requests"`
	Pagination PaginationData `json:"pagination"`
}