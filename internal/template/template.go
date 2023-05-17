package template

var (
	T *Template
)

type Template struct {
}

func NewTemplate() *Template {
	T = &Template{}
	return T
}

func GetTemplate() *Template {
	return T
}

func (t *Template) IsProduction() bool {
	return true
}

func (t *Template) GetFrontendPort() int {
	return 3000
}

func (t *Template) GetHost() string {
	return "127.0.0.1"
}

func (t *Template) GetPort() int {
	return 8080
}

func (t *Template) GetSSLPort() int {
	return 8443
}

func (t *Template) GetCertificate() string {
	return ""
}

func (t *Template) GetCertificateKey() string {
	return ""
}

func (t *Template) IsSSL() bool {
	return false
}
