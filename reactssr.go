package reactssr

import "rogchap.com/v8go"

// Renderer renders a React application to HTML.
type Renderer struct {
	Path string

	scriptSource string

	isolate        *v8go.Isolate
	global         *v8go.ObjectTemplate
	reactssrObject *v8go.ObjectTemplate
}

// NewServerSideRenderer creates a new server side renderer from a JavaScript bundle file.
func NewServerSideRenderer(path string) (*Renderer, error) {
	// TODO: perform validation(s) on path.

	iso, err := v8go.NewIsolate()
	if err != nil {
		return nil, err
	}
	// The "global" global injected into the v8 isolate's global namespace.
	global, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return nil, err
	}
	// The "reactssr" global injected into the v8 isolate's global namespace.
	reactssrObj, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return nil, err
	}
	r := &Renderer{
		Path: path,

		isolate:        iso,
		global:         global,
		reactssrObject: reactssrObj,
	}
	return r, r.loadScriptSource()
}

// Render renders the provided path to HTML.
func (r *Renderer) Render() (string, error) {
	return r.render()
}

// RenderFast renders the provided path to HTML.
func (r *Renderer) RenderFast() (string, error) {
	return r.renderFast()
}
