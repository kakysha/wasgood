package templates

// Page is a base page viewmodel
type Page struct {
	Title       string
	Description string
	Keywords    string
	OgImage     string
	Flashes     []interface{}
}
