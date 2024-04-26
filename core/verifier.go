package core

/**
 * Verifier Object
 * This object is responsible for verifying the downloaded file
 * using the etag or checksum provided by the server.
 */
type Verifier struct {
	file FileDetails
}
