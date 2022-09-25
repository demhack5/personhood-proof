package cdn

import (
	"crypto/sha1"
	"encoding/hex"

	jsoniter "github.com/json-iterator/go"
)

//content-delivery-network

type Client interface {
	GetUniqueURL(hash string) (string, error)
}

type client struct{}

var _ Client = &client{}

func NewClient() Client {
	return &client{}
}

func (c *client) GetUniqueURL(userhash string) (string, error) {
	hasher := sha1.New()
	hasher.Write([]byte(userhash))
	sha1_hash := hex.EncodeToString(hasher.Sum(nil))
	return sha1_hash, nil
}

func (c *client) GetTrafficInformation(url string) ([]byte, error) {
	return jsoniter.Marshal("")
}
