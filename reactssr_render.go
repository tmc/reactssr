package reactssr

import (
	"rogchap.com/v8go"
)

// render renders the CRA output bundle.
func (r *Renderer) render() (string, error) {
	outputHTML := ""
	reactssr := v8go.NewObjectTemplate(r.isolate)
	render := v8go.NewFunctionTemplate(r.isolate, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()
		if len(args) > 0 {
			outputHTML = args[0].String()
		}
		return nil
	})
	reactssr.Set("render", render)
	global := v8go.NewObjectTemplate(r.isolate)
	global.Set("reactssr", reactssr)
	ctx := v8go.NewContext(r.isolate, global)
	if _, err := ctx.RunScript(r.scriptSource, r.Path); err != nil {
		return "", err
	}
	return outputHTML, nil
}
