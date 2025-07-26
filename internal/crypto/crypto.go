package crypto

type Crypto struct {
	Salt []byte
	Key  []byte
}

func NewCrypto() *Crypto {
	return &Crypto{}
}
