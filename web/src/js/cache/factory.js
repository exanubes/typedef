export function create_database() {
    const connection_manager = create_connection_manager({
        namespace: "typedef",
        version: 1,
        schema(database, version) {
            if (version < 1) {
                const store = database.createObjectStore(config.table_name, {
                    keyPath: config.key_path,
                });

                for (const [key, indexConfig] of config.indexes) {
                    store.createIndex(key, key, indexConfig);
                }
            }
        }
    })
    return create_indexdb_client(connection_manager)
}
