local copy = require "@std.copy"
local set = {}

function set.from(t1)
  local s1 = {}
  for _, value in ipairs(t1) do
    s1[value] = true
  end
  return s1
end

function set.union(s1, s2)
  local s3 = copy(s1)
  for value, _ in pairs(s2) do
    s3[value] = true
  end
  return s3
end

function set.diff(s1, s2)
  local s3 = copy(s1)
  for value, _ in pairs(s2) do
    s3[value] = nil
  end
end

function set.has(s1, value)
  return s1[value] or false
end

return set
