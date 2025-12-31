local json_driver = require("typedef.drivers.json")
local M = {}

function M.register()
    json_driver.register()
end

return M
