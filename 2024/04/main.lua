package.path = package.path .. ";../?.lua"
local util = require("util")

local word = "XMAS"

local function part1(lines)
    local count = 0

    -- horizontal
    for i = 1, #lines do
        local start_index = string.find(lines[i], word)
        local start_index_rev = string.find(lines[i], string.reverse(word))
        if start_index or start_index_rev then
            count = count + 1
        end
    end

    for i = 1, #lines[1] do
        local vertical_str = ""
        for j = 1, #lines do
            local char = lines[j]:sub(i, i)
            vertical_str = vertical_str .. char
        end

        local vert_start_index = string.find(vertical_str, word)
        local vert_start_index_rev = string.find(vertical_str, string.reverse(word))
        if vert_start_index or vert_start_index_rev then
            count = count + 1
        end
    end

    return count
end

local function part2(lines)
end

local function main()
    local lines = util.read_lines("input.txt")

    local result, elapsedTime = util.run_function_with_timer(part1, lines)
    print(string.format("\nResult: %d", result))
    print(string.format("\nElapsed Time: %.6f seconds\n", elapsedTime))
end

main()
