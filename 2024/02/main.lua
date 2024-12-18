package.path = package.path .. ";../?.lua"
local util = require("util")

local function part1(lines)
    local res = 0

    for i = 1, #lines do
        local nums = {}
        for str in string.gmatch(lines[i], "%S+") do
            nums[#nums + 1] = tonumber(str)
        end

        local is_increasing = nums[2] - nums[1] > 0
        local prev = nums[1]
        for j = 2, #nums do
            local diff = nums[j] - prev
            if (is_increasing and diff < 0) or (not is_increasing and diff > 0) then
                break
            end

            diff = math.abs(diff)
            if diff < 1 or diff > 3 then
                break
            end

            if j == #nums then
                res = res + 1
            end

            prev = nums[j]
        end
    end

    return res
end

local function try_again(nums)
    for i = 1, #nums do
        local t = {}
        for j = 1, #nums do
            if j ~= i then
                t[#t + 1] = nums[j]
            end
        end

        if part1({ table.concat(t, " ") }) > 0 then
            return 1
        end
    end

    return 0
end


local function part2(lines)
    local res = 0

    for i = 1, #lines do
        local nums = {}
        for str in string.gmatch(lines[i], "%S+") do
            nums[#nums + 1] = tonumber(str)
        end

        local is_increasing = nums[2] - nums[1] > 0
        local prev = nums[1]
        for j = 2, #nums do
            local diff = nums[j] - prev
            if (is_increasing and diff < 0) or (not is_increasing and diff > 0) then
                res = res + try_again(nums)
                break
            end

            diff = math.abs(diff)
            if diff < 1 or diff > 3 then
                res = res + try_again(nums)
                break
            end

            if j == #nums then
                res = res + 1
            end

            prev = nums[j]
        end
    end

    return res
end

local function main()
    local lines = util.read_lines("input.txt")

    local result, elapsedTime = util.run_function_with_timer(part2, lines)
    print(string.format("\nResult: %d", result))
    print(string.format("\nElapsed Time: %.6f seconds\n", elapsedTime))
end

main()
