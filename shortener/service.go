package shortener

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/teris-io/shortid"
)

type RedirectService interface {
	Find(code string) (*Redirect, error)
	Create(r *Redirect) (*Redirect, error)
}

type redirectService struct {
	repo     RedirectRepository
	validate *validator.Validate
}

func NewRedirectService(repo RedirectRepository) RedirectService {
	validate := validator.New()
	return &redirectService{
		repo:     repo,
		validate: validate,
	}
}

// Find implements RedirectService
func (svc *redirectService) Find(code string) (*Redirect, error) {
	return svc.repo.Find(code)
}

// Create implements RedirectService
func (svc *redirectService) Create(r *Redirect) (*Redirect, error) {
	if err := svc.validate.Struct(r); err != nil {
		return nil, err
	}
	r.Code = shortid.MustGenerate()
	r.CreatedAt = time.Now().UTC().Unix()
	return svc.repo.Create(r)
}
