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
    go.run(instance)
    wasm_ready_resolve()
})()

self.onmessage = async (event) => {
    await wasmReady;
    self.rpc(JSON.stringify(event.data))
}

