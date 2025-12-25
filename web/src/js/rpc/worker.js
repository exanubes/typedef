import './wasm_exec.js'
import wasm_url from 'url:./main.wasm'

// NOTE: Go ctor is attached to global context by wasm_exec.js
const go = new Go()

let ready = false

let wasm_ready_resolve;



const wasmReady = new Promise(resolve => wasm_ready_resolve = resolve);
(async () => {
    const response = await fetch(wasm_url)
    const bytes = await response.arrayBuffer()

    const { instance } = await WebAssembly.instantiate(bytes, go.importObject)
    go.run(instance) // NOTE: never returns
    wasm_ready_resolve()
})()

self.onmessage = async (event) => {
    await wasmReady;
    const response = self.rpc(event.data)
    self.postMessage(response)
}

[{
    "id": 1,
    "title": "Harry Potter",
    "user": {
        "id": 1,
        "name": "John"
    },
    "author": {
        "id": 2,
        "name": "Tom"
    },
    "numbers": [1, 2, 3],
    "mixed": [1, "2", false, { "id": 3, "name": "Simon" }],
    "float": 69.420,
    "cool": true
}]
