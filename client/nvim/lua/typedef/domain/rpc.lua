---@class Connection
---@field close fun(self: Connection)
---@field is_active fun(self:Connection): boolean
---@field send fun(self:Connection, message: string)

---@class Rpc
---@field connect fun(self: Rpc, message_handler: fun(data: string)): Connection

---@class Server
---@field start fun(self: Server)

---@class RequestValidator
---@field validate fun(message: table): boolean
---

---@class RpcClient
---@field send fun(self: Server, method: string, payload: table): Promise
