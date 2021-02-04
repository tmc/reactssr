# reactssr

A library to perform Server Side Rendering of React apps.

[![Project status](https://img.shields.io/github/release/tmc/go-reactssr.svg?style=flat-square)](https://github.com/tmc/go-reactssr/releases/latest)
[![Build Status](https://github.com/tmc/go-reactssr/workflows/test/badge.svg)](https://github.com/tmc/go-reactssr/actions?query=workflow%3Atest)
[![Go Report Card](https://goreportcard.com/badge/tmc/go-reactssr?cache=0)](https://goreportcard.com/report/tmc/go-reactssr)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/tmc/go-reactssr)

## Example usage

Given a bundle produced from an additional entrypoint to your application such as [./js/index.ssr.jsx](./js/index.ssr.jsx):

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

This file should be bundled, for exaple, wwith [esbuild](https://esbuild.github.io/) as so:

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



