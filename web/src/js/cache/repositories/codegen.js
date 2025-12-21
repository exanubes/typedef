import { CODEGEN_SCHEMA } from "../schema/codegen"

/**
 * @param {import("../types").Client} client 
 * @returns {CodegenInputRepository}
 * */
export function create_codegen_input_repository(client) {
    return {
        async find(id) {
            /** @type {TableRow} */
            const response = await client.find(id, CODEGEN_SCHEMA.key_path, CODEGEN_SCHEMA.table_name)
            console.log("find(id): ", response)
            /** @type {CachedCodegenCommand} */
            const result = {}

            result.hash = response.id
            result.input = response.canonical_input
            result.target = response.target
            result.output_hash = response.output_hash
            return result
        },
        async write(payload) {
            /** @type {TableRow} */
            const req = {}
            req.id = payload.input_hash
            req.canonical_input = payload.input
            req.target = payload.target
            req.output_hash = payload.output_hash

            await client.write(req, CODEGEN_SCHEMA.table_name)

            return void 0
        }
    }
}

/**
 * @typedef {Object} CodegenInputRepository
 * @property {(id: string)=>Promise<CachedCodegenCommand>} find
 * @property {(payload: SaveCodegenCommandPayload)=>Promise<void>} write
 * */

/**
 * @typedef {Object} CachedCodegenCommand
 * @property {string} hash 
 * @property {string} input
 * @property {string} target
 * @property {string} output_hash
 * */

/**
 * @typedef {Object} SaveCodegenCommandPayload
 * @property {string} input_hash 
 * @property {string} input
 * @property {string} target
 * @property {string} output_hash
 * */

/**
 * @typedef {Object} TableRow private type not meant to be used outside of the repository
 * @property {string} id hash of the canonical input
 * @property {string} canonical_input
 * @property {string} target
 * @property {string} output_hash
 * */
