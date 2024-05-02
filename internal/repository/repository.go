package repository

type UrlGetter interface {
	GetUrl(hash string) (string, error)
}

type UrlSetter interface {
	SetUrl(hash string, url string) error
}

type UrlRepository interface {
	UrlGetter
	UrlSetter
}

type Repositories struct {
	UrlRepository UrlRepository
}
