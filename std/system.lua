local migrator = require "@std.migrator"

return function(migrators)
  return migrator("std.system", function(current_state)
    local sys_resolution = { migration = "", sync = "", next_state = {} }
    current_state = current_state or {}
    for _, i_migrator in ipairs(migrators) do
      local resolution = i_migrator.func(current_state[i_migrator.name])
      sys_resolution.migration = sys_resolution.migration .. "\n" .. resolution.migration
      sys_resolution.sync = sys_resolution.sync .. "\n" .. resolution.sync
      sys_resolution.next_state[i_migrator.name] = resolution.next_state
    end
    return sys_resolution
  end)
end
