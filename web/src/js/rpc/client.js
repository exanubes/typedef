'use strict'

/**
 * @returns {RpcClient}
 * */
export function create_rpc_client() {
    const worker = new Worker(
        new URL("./worker.js", import.meta.url),
        {
            type: "module"
        })

    const pending_requests = new Map()

    worker.addEventListener("message", (event) => {
        const response = JSON.parse(event.data);
        const entry = pending_requests.get(response.id)

        if (!entry) return;


        if (response.error) {
            entry.reject(response.error)
        } else {
            entry.resolve(response)
        }
    })

    worker.addEventListener("error", (error) => {
        for (const { reject } of pending_requests.values()) {
            reject(error)
        }

        pending_requests.clear()
    })

    return {
        async send(request, timeout_ms = 5_000) {
            return new Promise((resolve, reject) => {
                if (pending_requests.has(request.id)) {
                    throw new Error(`Duplicate request id ${request.id}`)
                }

                const timer_id = setTimeout(() => {
                    pending_requests.delete(request.id)
                    reject(new Error("RPC timeout"))
                }, timeout_ms)


                pending_requests.set(request.id, {
                    resolve: (value) => {
                        clearTimeout(timer_id)
                        pending_requests.delete(request.id)
                        resolve(value)
                    },
                    reject: (error) => {
                        clearTimeout(timer_id)
                        pending_requests.delete(request.id)
                        reject(error)
                    }
                })
                worker.postMessage(request)
            })
        },
        close() {
            for (const { reject } of pending_requests.values()) {
                reject(new Error("Worker terminated."))
            }

            pending_requests.clear()
            worker.terminate()
        }
    }
}

/**
 * @typedef {Object} RpcClient 
 * @property {SendRPCCommand} send
 * */

/**
 * @callback SendRPCCommand
 * @param {import("../libs/jsonrpc").JSONRPCRequest<unknown>} request
 * @returns {Promise<import("../libs/jsonrpc").JSONRPCResponse<unknown>>}
 * */
