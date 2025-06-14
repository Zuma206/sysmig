local std = {}

-- Create a basic migrator
function std.migrator(name, func)
  return { name = name, func = func }
end

-- A blank migrator that does absolutely nothing
std.nothing = std.migrator("nothing", function()
  local script = "# std.nothing"
  return {
    migration = script,
    next_state = nil,
    sync = script,
  }
end)

return std
