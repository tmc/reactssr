package reactssr

import "io/ioutil"

func (r *Renderer) loadScriptSource() error {
	buf, err := ioutil.ReadFile(r.Path)
	if err != nil {
		return err
	}
	r.scriptSource = string(buf)
	return nil
}
