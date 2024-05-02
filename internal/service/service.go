package service

type UrlGetter interface {
	GetUrl(hash string) (string, error)
}

type UrlSetter interface {
	SetUrl(url string) (string, error)
}

type UrlService interface {
	UrlGetter
	UrlSetter
}

type Services struct {
	UrlService UrlService
}
