local env = require("typedef.env")
local bin_path = ""

local function has_rpc_binary()
    if vim.fn.executable(bin_path) ~= 1 then
        return false
    end

    return true
end

local function get_binary_version()
    local result = vim.system({ bin_path, "version" }, { text = true }):wait()

    if result.code ~= 0 then
        return {}, false
    end

    local ok, info = pcall(vim.json.decode, result.stdout)

    if not ok then
        return {}, false
    end

    return { version = info.version, commit = info.commit_sha, created_at = info.build_time }, true
end

local function compare(expected, received)
    if expected.version ~= received.version then
        return false, "Expected: " .. expected.version .. ", received: " .. received.version
    end

    if expected.commit ~= received.commit then
        return false, "Expected: " .. expected.commit .. ", received: " .. received.commit
    end

    return true, ""
end

local function verify_binary()
    local has_binary = has_rpc_binary()
    if not has_binary then
        return false
    end

    local binary_version, ok = get_binary_version()

    if not ok then
        return false
    end

    local ok, err = compare(env, binary_version)

    if not ok then
        error(err)
    end

    return true
end

return function()
    local ok, result = pcall(verify_binary)

    if ok and result then
        --- NOTE: binary is installed and matches the version of the plugin. Do nothing
        return
    end
    --- TODO: download the binary from github releases
end
