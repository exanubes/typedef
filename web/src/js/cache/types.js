/**
 * @typedef {Object} Cache
 * @property {(key: string)=>CacheItem} get
 * @property {(key: string, payload: Record<string, unknown>)=>boolean} write
 * */

/**
 * @typedef {Object} Client
 * @property {(config: Config)=>Promise<any>} connect
 * @property {(id: string)=>Promise<object>} find
 * @property {(id: string, payload: object)=>Promise<void>} find
 * */


/**
 * @typedef {Object} Config
 * @property {boolean} auto_increment
 * @property {string} key_path
 * @property {Map<string, {}>} indexes
 * @property {string} table_name
 * @property {string} namespace
 * */

export { }
