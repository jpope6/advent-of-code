package.path = package.path .. ";../?.lua"
local util = require("util")

local function main()
    local lines = util.read_lines("input.txt")

    -- Print each line (for verification)
    for i, line in ipairs(lines) do
        print(string.format("Line %d: %s", i, line))
    end
end

main()
