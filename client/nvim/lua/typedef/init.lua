local drivers = require("typedef.drivers")
local M = {}

function M.setup()
    vim.notify("Hello World", vim.log.levels.INFO)
    drivers.register()
end

return M
