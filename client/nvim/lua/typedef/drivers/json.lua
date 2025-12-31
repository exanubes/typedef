local codegen = require("typedef.domain.codegen")
local Codegen = require("typedef.infrastructure.codegen")
local SelectionReader = require("typedef.infrastructure.reader.selection")
local YankReader = require("typedef.infrastructure.reader.yank")
local ChainReader = require("typedef.infrastructure.reader.chain")
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
        local input_reader = ChainReader.new({ SelectionReader.new(opts.range > 0), YankReader.new() })
        local codegen_repository = Codegen.new(server)
        local input = input_reader:read()
        -- codegen_repository:generate(input, "json", format)
    end, { nargs = 1, range = true })
end

return M
