'use strict'

/**
 * @param {(input: object)=>string} serializer
 * @returns {HasherService}
 * */
export function create_hasher(serializer) {
    return {
        async hash(input) {
            try {
                const canonical = serializer(input)
                const data = new TextEncoder().encode(canonical)
                const buffer = await crypto.subtle.digest("SHA-256", data)
                const hash_array = [...new Uint8Array(buffer)]
                const hex = hash_array.map(b => b.toString(16).padStart(2, "0")).join("")

                return {
                    status: "ok",
                    value: hex
                }
            } catch (error) {
                return {
                    status: "error",
                    err: new HasherException(error)
                }
            }
        }
    }
}

class HasherException extends Error {
    constructor(cause) {
        super("Failed to hash input", { cause })
    }
}

/**
 * @typedef {Object} HasherService
 * @property {(input: object)=>Promise<HashResult>} hash
 * */

/**
 * @typedef {HashSuccess | HashFail} HashResult
 * */

/**
 * @typedef {Object} HashSuccess
 * @property {"ok"} status
 * @property {string} value
 * */

/**
 * @typedef {Object} HashFail
 * @property {"error"} status
 * @property {Error} err
 * */
