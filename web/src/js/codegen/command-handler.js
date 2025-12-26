'use strict';

import { create_json_rpc_request } from '../libs/jsonrpc';
import { parse_codegen_response } from './commands';

/**
 * @param {import("../rpc/client").RpcClient} client
 * @returns {import("./commands").CodegenCommandHandler}
 * */
export function create_rpc_codegen_command_handler(client) {
    return {
        /**
         *
         * @param command
         */
        async send(command) {
            try {
                const request = create_json_rpc_request('codegen', command);
                const response = await client.send(request);
                return { status: 'ok', data: parse_codegen_response(response) };
            } catch (error) {
                return {
                    status: 'error',
                    err: new Error('Error while making request', { cause: error }),
                };
            }
        },
    };
}
