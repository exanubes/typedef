'use strict'

/**
 * @param {import("../indexdb/repositories/codegen-input").CodegenInputRepository} repository
 * @param {import("../hasher/hasher").HasherService} hash_service
 * @returns {RequestCache}
 * */
export function create_request_cache(repository, hash_service) {
    return {
        async put(input, target, output) {
            try {
                // TODO: Add TTL parameter
                const input_hash_result = await hash_service.hash(JSON.parse(input))
                const output_hash_result = await hash_service.hash(output)
                if (output_hash_result.status == "ok" && input_hash_result.status == "ok") {
                    await repository.write({
                        input,
                        target,
                        output,
                        input_hash: input_hash_result.value,
                        output_hash: output_hash_result.value,
                    })
                } else {
                    if (input_hash_result.err) {
                        console.error("[RequestCache] Failed to hash input: " + input_hash_result.err)
                    }

                    if (output_hash_result.err) {
                        console.error("[RequestCache] Failed to hash output: " + output_hash_result.err)
                    }
                }
            } catch (error) {
                console.error("[RequestCache] Failed to write to cache: " + error)
            }
        },

        async get(input) {
            try {
                const hash_result = await hash_service.hash(JSON.parse(input))
                if (hash_result.status == "ok") {
                    return repository.find(hash_result.value)
                }

                return null
            } catch (error) {
                console.error(`[RequestCache] Failed to read from cache: ` + error)
                return null
            }
        }
    }
}

/**
 * @typedef {Object} RequestCache
 * @property {(input: string, target: string, output: string)=>Promise<void>} put
 * @property {(input: string)=>Promise<import("../indexdb/repositories/codegen-input").CachedCodegenCommand|null>} get
 * */
