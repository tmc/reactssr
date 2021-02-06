# reactssr

A Go package to perform Server Side Rendering of React apps.

[![Project status](https://img.shields.io/github/release/tmc/reactssr.svg?style=flat-square)](https://github.com/tmc/reactssr/releases/latest)
[![Build Status](https://github.com/tmc/reactssr/workflows/test/badge.svg)](https://github.com/tmc/reactssr/actions?query=workflow%3Atest)
[![Go Report Card](https://goreportcard.com/badge/tmc/reactssr?cache=0)](https://goreportcard.com/report/tmc/reactssr)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/tmc/reactssr)

## Example usage

Given a bundle produced from an additional entrypoint to your application such as [js/index.ssr.jsx](./js/index.ssr.jsx):

```jsx

import * as React from 'react';
import * as Server from 'react-dom/server'
import './index.css';
import App from './App';

const AppOutput = Server.renderToString(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// reactssr.render is the callback injected by the go runtime to pass the rendered application back.
reactssr.render(AppOutput);
```

This file should be bundled, for example, with [esbuild](https://esbuild.github.io/) as so:

```bash
npx esbuild \
   src/index.ssr.jsx \
   --inject:src/react-shim.js \
   --bundle \
   --sourcemap \
   --outfile=build/out.js \
   --loader:.js=jsx \
   --loader:.svg=text \
   --define:process.env.NODE_ENV=\"production\"
```

Then the following code will execute the bundle and load the results into a Go variable (for serving
to a client, for emple).

```go
r, _ := reactssr.NewServerSideRenderer("./testdata/test-app-1/build/out.js")
output, _ := r.Render()

// output contains the rendered html from the React application.
```

## How this works

`reactssr` works by executing a React application bundle with `reactssr.render` injected into the
global Javascript namespace.

In this example:

```js
reactssr.render(Server.renderToString(
  <React.StrictMode>
    <App />
  </React.StrictMode>
));
```

`reactssr.render` is a Go callback which allows the Javascript execution environment to pass
the rendered HTML and CSS between runtimes.


## Performance

This package includes benchmarks which are run in CI: [reactssr_test.go](./reactssr_test.go).

[Recent performance results](https://github.com/tmc/reactssr/runs/1828170002?check_suite_focus=true)

```sh
go test -v -run=XXX -benchmem -bench=.*
goos: linux
goarch: amd64
pkg: github.com/tmc/reactssr
BenchmarkRender
BenchmarkRender-2    	     464	   5855720 ns/op	    3459 B/op	      19 allocs/op
PASS
```

This indicates that it takes just under `6 milliconds` to render the current default output 
from [create-react-app](https://github.com/facebook/create-react-app). This is without any specific 
work towards optimizing the implementation and this output is easily cachable.
