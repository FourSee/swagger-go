/*
 * PolyRhythm
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// The root of the Message type's schema.
type Message struct {
	Id string `json:"id"`
	// An encrypted payload representing the content.
	Content string `json:"content"`
	// Exit code, if provided
	ExitCode int32 `json:"exit_code,omitempty"`
	// Device ID of the sender
	SenderId string `json:"sender_id"`
	// Recipient device IDs
	RecieverId string `json:"reciever_id"`
	// Plaintext title. Required to be plaintext due to mobile limitations on display text
	Title string `json:"title"`
}