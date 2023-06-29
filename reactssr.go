package reactssr

import "rogchap.com/v8go"

// Renderer renders a React application to HTML.
type Renderer struct {
	Path string

	scriptSource string

	isolate        *v8go.Isolate
	reactssrObject *v8go.ObjectTemplate
}

// NewServerSideRenderer creates a new server side renderer from a JavaScript bundle file.
func NewServerSideRenderer(path string) (*Renderer, error) {
	// TODO: perform validation(s) on path.

	iso := v8go.NewIsolate()
	// The "reactssr" global injected into the v8 isolate's global namespace.
	reactssrObj := v8go.NewObjectTemplate(iso)
	r := &Renderer{
		Path: path,

		isolate:        iso,
		reactssrObject: reactssrObj,
	}
	return r, r.loadScriptSource()
}

// Render renders the provided path to HTML.
func (r *Renderer) Render() (string, error) {
	return r.render()
}
