local codegen = require("typedef.domain.codegen")
local M = {}

function M.register()
    vim.api.nvim_create_user_command("TypedefJson", function(opts)
        local arg = opts.args
        local format = codegen.parse_format(arg)
        if not format then
            vim.notify("[TypedefJson] invalid format: " .. arg, vim.log.levels.ERROR)
            return
        end

        vim.notify("[TypedefJson] transform json to " .. format, vim.log.levels.INFO)
    end, { nargs = 1 })
end

return M
