import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './app';
import "./wasm/wasm_exec"

function wasmURL() {
    return "./static/proxx.wasm";
}

// Go from ./wasm/wasm_exec
const go = new globalThis.Go();
WebAssembly
    .instantiateStreaming(fetch(wasmURL()), go.importObject)
    .then(function (result) {
        go.run(result.instance);
    });

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
