import { CODEGEN_SCHEMA } from "../schema/codegen"

/**
 * @param {import("../types").Client} client 
 * */
export function create_codegen_repository(client) {

    return {
        async find(id) {
            // TODO: build result object from response
            const result = {}
            const response = await client.find(id, CODEGEN_SCHEMA.key_path, CODEGEN_SCHEMA.table_name)

            return result
        },
        // TODO: define payload shape
        write(payload) {
            // TODO: validate shape based on CODEGEN_SCHEMA

            return client.write(payload, CODEGEN_SCHEMA.table_name)
        }
    }
}

