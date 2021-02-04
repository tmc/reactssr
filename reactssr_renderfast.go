package reactssr

// render renders the CRA output bundle.
func (r *Renderer) renderFast() ([]byte, error) {
	v, err := r.ctx.RunScript(r.scriptSource, r.Path)
	if err != nil {
		return nil, err
	}
	return []byte(v.String()), nil
}
