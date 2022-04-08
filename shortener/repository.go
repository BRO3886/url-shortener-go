package shortener

type RedirectRepository interface {
	Find(code string) (*Redirect, error)
	Create(r *Redirect) (*Redirect, error)
}
