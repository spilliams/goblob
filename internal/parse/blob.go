package parse

// BlobParser represents an object that can parse blobs
type BlobParser struct {
	parsedObj interface{}
}

// NewBlobParser returns a new parser
func NewBlobParser() *BlobParser {
	return &BlobParser{}
}

// ParseString parses a string-encoded blob
func (b *BlobParser) ParseString(s string) error {
	return nil
}
