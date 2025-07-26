package crypto

import (
	"encoding/base64"

	"github.com/Cyrof/govault/internal/model"
)

// convert crypto -> meta
func (c *Crypto) ToMeta() model.Meta {
	return model.Meta{
		Salt: base64.StdEncoding.EncodeToString(c.Salt),
		Hash: base64.StdEncoding.EncodeToString(c.Key),
	}
}
