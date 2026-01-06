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
    --- @type InputReader
    input = nil,
    --- @type CodegenRepository
    codegen = nil,
    --- @type OutputWriter
    output = nil,
}

M.__index = M

---@param panel Panel
---@param input_reader InputReader
---@param codegen CodegenRepository
---@param output_writer OutputWriter
function M.new(panel, input_reader, codegen, output_writer)
    return setmetatable({
        panel = panel,
        input = input_reader,
        codegen = codegen,
        output = output_writer,
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
            local input = self.input:read()
            local output = output_options:selected()
            local response = self.codegen:generate(input, "json", output)

            response:on_success(function(result)
                self.output:write(result.code)
                -- self.panel:close()
            end)

            response:on_error(function(error)
                vim.notify("ERROR(" .. error.code .. "): " .. error.message)
            end)

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
