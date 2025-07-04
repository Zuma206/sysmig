local migrator = require "@std.migrator"

local function get_initial_state(migrators)
  if migrators.initial_state then
    return migrators.initial_state.func(nil).next_state
  end
  return {}
end

-- A migrator multiplexer
-- Takes a list of migrators, splits the state, and combines the resolution
return function(config)
  local migrators = config.migrators or config
  return migrator("std.system", function(current_state)
    local sys_resolution = { migration = "", sync = "", next_state = {} }
    current_state = current_state or get_initial_state(migrators)
    for _, i_migrator in ipairs(migrators) do
      local resolution = i_migrator.func(current_state[i_migrator.name])
      sys_resolution.migration = sys_resolution.migration .. "\n" .. resolution.migration
      sys_resolution.sync = sys_resolution.sync .. "\n" .. resolution.sync
      sys_resolution.next_state[i_migrator.name] = resolution.next_state
    end
    return sys_resolution
  end)
end
