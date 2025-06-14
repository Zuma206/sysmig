local copy = require "@std.copy"
local set = {}

-- Converts a sequence into a set
function set.from(t1)
  local s1 = {}
  for _, value in ipairs(t1) do
    s1[value] = true
  end
  return s1
end

-- Returns the union of two sets
function set.union(s1, s2)
  local s3 = copy(s1)
  for value, _ in pairs(s2) do
    s3[value] = true
  end
  return s3
end

-- Returns the difference of two sets
function set.diff(s1, s2)
  local s3 = copy(s1)
  for value, _ in pairs(s2) do
    s3[value] = nil
  end
  return s3
end

-- Checks if a set contains a value
function set.has(s1, value)
  return s1[value] or false
end

-- Converts a set back into a sequence
function set.to_sequence(s1)
  local t1 = {}
  for value, has in pairs(s1) do
    if has then
      table.insert(t1, value)
    end
  end
  return t1
end

return set
