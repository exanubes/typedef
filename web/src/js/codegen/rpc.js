'use strict'

import { validate_json_rpc_response } from "../libs/jsonrpc"

export function parse_codegen_response(response) {
    validate_json_rpc_response(response)
    return {
        code: response.result.code,
        format: response.result.format,
    }
}
