package reactssr

import "rogchap.com/v8go"

// Renderer renders a React applicationt to HTML.
type Renderer struct {
	Path string

	scriptSource string
	ctx          *v8go.Context
}

// NewServerSideRenderer creates a new server side renderer from a JavaScript bundle file.
func NewServerSideRenderer(path string) (*Renderer, error) {
	// TODO: perform validation(s) on path.
	iso, err := v8go.NewIsolate()
	if err != nil {
		return nil, err
	}
	ctx, err := v8go.NewContext(iso)
	if err != nil {
		return nil, err
	}
	r := &Renderer{
		Path: path,
		ctx:  ctx,
	}
	return r, r.loadScriptSource()
}

// Render renders the provided path to HTML.
func (r *Renderer) Render() (string, error) {
	return r.render()
}

// RenderFast renders the provided path to HTML.
func (r *Renderer) RenderFast() ([]byte, error) {
	return r.renderFast()
}
