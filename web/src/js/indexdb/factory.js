'use strict'

import { create_connection_manager, create_indexdb_client } from "./client";
import { CODEGEN_INPUTS_SCHEMA } from "./schema/codegen";

export function create_database() {
    const connection_manager = create_connection_manager({
        namespace: "typedef",
        version: 1,
        schema(database, version) {
            if (version < 1) {
                const store = database.createObjectStore(CODEGEN_INPUTS_SCHEMA.table_name, {
                    keyPath: CODEGEN_INPUTS_SCHEMA.key_path,
                });

                for (const [key, indexConfig] of CODEGEN_INPUTS_SCHEMA.indexes) {
                    store.createIndex(key, key, indexConfig);
                }
            }
        }
    })
    return create_indexdb_client(connection_manager)
}
