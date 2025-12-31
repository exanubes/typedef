local codegen = require("typedef.domain.codegen")
local Codegen = require("typedef.infrastructure.codegen")
local SelectionReader = require("typedef.infrastructure.reader.selection")
local M = {}

---@param server JsonRpcServer
function M.register(server)
    vim.api.nvim_create_user_command("TypedefJson", function(opts)
        local arg = opts.args
        local format = codegen.parse_format(arg)
        if not format then
            vim.notify("[TypedefJson] invalid format: " .. arg, vim.log.levels.ERROR)
            return
        end
        -- server:start()

        local input_reader = SelectionReader.new()
        local codegen_repository = Codegen.new(server)
        local input = input_reader:read()
        vim.notify("[TypedefJson] input: " .. input, vim.log.levels.INFO)
        -- codegen_repository:generate(input, "json", format)

        -- vim.notify("[TypedefJson] transform json to " .. format, vim.log.levels.INFO)
    end, { nargs = 1 })
end

return M
