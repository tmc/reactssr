#!/bin/bash
# setup-ttest-cra-app.sh - creates a fresh create-react-app app.

if [ "${1:-}" == "clean" ]; then
   echo "cleaning cra output"
   rm -rf ./testdata/test-app-1
fi

test -f ./testdata/test-app-1/package.json || yarn create react-app ./testdata/test-app-1

# Copy in the Server Side Rendering entrypoint.
cp ./js/index.ssr.jsx ./testdata/test-app-1/src/
cp ./js/react-shim.js ./testdata/test-app-1/src/

(
    cd ./testdata/test-app-1 || exit
    test -d node_modules/react || yarn
    npx esbuild \
       src/index.ssr.jsx \
       --inject:src/react-shim.js \
       --bundle \
       --sourcemap \
       --outfile=build/out.js \
       --loader:.js=jsx \
       --loader:.svg=text \
       --define:process.env.NODE_ENV=\"production\"
)
