package shortener

type RedirectSerializer interface {
	Decode(data []byte) (*Redirect, error)
	Encode(r *Redirect) ([]byte, error)
}
