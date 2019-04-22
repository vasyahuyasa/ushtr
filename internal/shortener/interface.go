package shortener

type Encoder interface {
	Encode(id uint64) string
}

type Decoder interface {
	Decode(code string) (uint64, error)
}

type EncoderDecoder interface {
	Encoder
	Decoder
}
