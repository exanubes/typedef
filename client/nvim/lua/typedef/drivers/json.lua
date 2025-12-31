local codegen = require("typedef.domain.codegen")
local Codegen = require("typedef.infrastructure.codegen")
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

        server:start()
        local codegen_repository = Codegen.new(server)

        codegen_repository:generate(input, "json", format)

        vim.notify("[TypedefJson] transform json to " .. format, vim.log.levels.INFO)
    end, { nargs = 1 })
end

return M
