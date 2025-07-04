local Set = {}

local function default_key(value)
  return value
end

function Set:create(values, key)
  local set = { values = {}, key = key or default_key }
  setmetatable(set, { __index = self })
  set:add_values(values)
  return set
end

function Set:add_values(values)
  for _, value in ipairs(values) do
    self:add(value)
  end
end

function Set:add(value)
  self.values[self.key(value)] = { value = value }
end

function Set:to_table()
  local result = {}
  for _, value in pairs(self.values) do
    table.insert(result, value.value)
  end
  return result
end

function Set:copy()
  return Set:create(self:to_table(), self.key)
end

function Set:delete(value)
  self.values[self.key(value)] = nil
end

function Set:diff(other)
  local diff = self:copy()
  for _, value_ref in pairs(other.values) do
    diff:delete(value_ref.value)
  end
  return diff
end

function Set:size()
  local size = 0
  for _, _ in pairs(self.values) do
    size = size + 1
  end
  return size
end

return Set
