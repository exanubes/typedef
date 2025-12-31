local ChainReader = {}
ChainReader.__index = ChainReader

---@param readers InputReader[]
function ChainReader.new(readers)
    return setmetatable({ readers = readers }, ChainReader)
end

function ChainReader:read()
    for index, reader in ipairs(self.readers) do
        local result = reader:read()
        vim.notify("index: " .. index .. "result: " .. result)
        if result ~= "" then
            return result
        end
    end

    return "No valid json objects"
end

return ChainReader
