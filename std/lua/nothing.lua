local migrator = require "@std.migrator"

-- A migrator that does literally nothing
-- Here for testing purposes mainly
return migrator("std.nothing", function()
  local script = "# std.nothing"
  return {
    migration = script,
    next_state = nil,
    sync = script,
  }
end)
