local Map = require("typedef.helpers.map")
local Select = require("typedef.helpers.select")
local splice = require("typedef.helpers.splice")

local EXIT_BUTTON = "[Exit]"
local INSERT_BUTTON = "[Insert]"
local CLIPBOARD_BUTTON = "[Copy to clipboard]"

local function create_lines(options)
    local lines = {
        "Select input format",
        "(*) json",
        "",
        "Select output format",
        --- NOTE: Injected using splice()
        "",
        INSERT_BUTTON,
        CLIPBOARD_BUTTON,
        EXIT_BUTTON,
    }
    return splice(lines, options, 5)
end

local M = {
    --- @type Panel
    panel = nil,
}
M.__index = M

---@param panel Panel
function M.new(panel)
    return setmetatable({
        panel = panel,
    }, M)
end

function M:execute()
    local keymap = Map.new()
    keymap:set("1", 1)
    local output_options_offset = 4
    local output_options = Select.new({ "Go", "Typescript", "Zod", "JSDoc" })
    local lines = create_lines(output_options:print())

    self.panel:add_keymap("q", function()
        self.panel:close()
    end)

    self.panel:add_keymap("<CR>", function(event)
        if lines[event.current_line] == INSERT_BUTTON then
            vim.notify("Generate code and insert into buffer")
            self.panel:close()
            return
        end

        if lines[event.current_line] == EXIT_BUTTON then
            self.panel:close()
            return
        end

        if lines[event.current_line] == CLIPBOARD_BUTTON then
            vim.notify("Generate code and save to clipboard")
            self.panel:close()
            return
        end

        output_options:select(event.current_line - output_options_offset)
        lines = create_lines(output_options:print())
        self.panel:render(lines)
    end)

    self.panel:render(lines)
end

return M
