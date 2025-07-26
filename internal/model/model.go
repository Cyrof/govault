package model

type Meta struct {
	Salt string `json:"salt"`
	Hash string `json:"hash"`
}

// use secret struct for scalability in the future
type Secret struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
