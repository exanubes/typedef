'use strict'

export function create_rpc_client() {
    const worker = new Worker(
        new URL("./worker.js", import.meta.url),
        {
            type: "module"
        })

    const pending_requests = new Map()

    worker.addEventListener("message", (event) => {
        const response = JSON.parse(event.data); const entry = pending.get(response.id)

        if (!entry) return;

        pending.delete(response.id)

        if (response.error) {
            entry.reject(response.error)
        } else {
            entry.resolve(response.result)
        }
    })

    worker.addEventListener("error", (error) => {
        for (const { reject } of pending.values()) {
            reject(error)
        }

        pending.clear()
    })

    return {
        async send(request) {
            return new Promise((resolve, reject) => {
                pending.set(request.id, { resolve, reject })
                worker.postMessage(JSON.stringify(jsonRpcRequest))
            })
        },
        close() {
            for (const { reject } of pending.values()) {
                reject(new Error("Worker terminated."))
            }

            pending.clear()
            worker.terminate()
        }
    }
}

/**
 * @typedef {Object} RpcClient 
 * @property {Send} send
 * */

/**
 * @callback SendRPCCommand
 * @param {import("../libs/jsonrpc").JSONRPCRequest<unknown>} request
 * @returns {import("../libs/jsonrpc").JSONRPCResponse<unknown>}
 * */
