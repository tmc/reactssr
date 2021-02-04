package reactssr

import "rogchap.com/v8go"

// render renders the CRA output bundle.
func (r *Renderer) renderFast() (string, error) {
	reactssr, err := v8go.NewObjectTemplate(r.isolate)
	if err != nil {
		return "", err
	}
	outputHTML := ""
	render, err := v8go.NewFunctionTemplate(r.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()
		if len(args) > 0 {
			outputHTML = args[0].String()
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	reactssr.Set("render", render)
	r.global.Set("reactssr", reactssr)
	ctx, err := v8go.NewContext(r.isolate, r.global)
	if err != nil {
		return "", err
	}
	if _, err := ctx.RunScript(r.scriptSource, r.Path); err != nil {
		return "", err
	}
	return outputHTML, nil
}
