import './wasm_exec.js'

// NOTE: Go ctor is exposed globally by wasm_exec.js
const go = new Go()

let ready = false

let wasm_ready_resolve;
const wasmReady = new Promise(resolve => wasm_ready_resolve = resolve)
WebAssembly.instantiateStreaming(
    fetch("main.wasm"),
    go.importObject
).then(result => {
    go.run(result.instance) // NOTE: never returns
    ready = true
    wasm_ready_resolve()
})

onmesssage = async (event) => {
    await wasmReady;

    const response = rpc(event.data)
    postMessage(response)
}
