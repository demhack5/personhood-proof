package cdn

import (
	"crypto/sha1"
	"encoding/hex"
)

//content-delivery-network

type Client interface {
	GetUniqueURL(hash string) (string, error)
}

type client struct {
}

var _ Client = &client{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetUniqueURL(hash string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(hash))
	sha1_hash := hex.EncodeToString(hasher.Sum(nil))
	return sha1_hash, nil
}
