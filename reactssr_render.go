package reactssr

import (
	"rogchap.com/v8go"
)

// render renders the CRA output bundle.
func (r *Renderer) render() (string, error) {
	iso, err := v8go.NewIsolate()
	if err != nil {
		return "", err
	}
	console, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return "", err
	}
	result := ""
	render, err := v8go.NewFunctionTemplate(iso, func(info *v8go.FunctionCallbackInfo) *v8go.Value {
		args := info.Args()
		if len(args) > 0 {
			result = args[0].String()
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	console.Set("render", render)
	global, err := v8go.NewObjectTemplate(iso)
	if err != nil {
		return "", err
	}
	global.Set("reactssr", console)
	ctx, err := v8go.NewContext(iso, global)
	if err != nil {
		return "", err
	}
	if _, err := ctx.RunScript(r.scriptSource, r.Path); err != nil {
		return "", err
	}
	return result, nil
}
