local migrator = require "@std.migrator"
local Set = require "@std.Set"
local map = require "@std.map"

local function get_group_script(list, handle, handle_no)
  if #list == 0 or handle == nil then
    return handle_no .. "\n"
  elseif handle then
    return handle(list) .. "\n"
  end
  return ""
end

local function get_iterative_script(list, handle)
  if handle then
    return table.concat(map(list, handle), "\n") .. "\n"
  end
  return ""
end

local function get_sequence_script(added, removed, handlers)
  if type(handlers) == "string" then
    return handlers
  else
    handlers = handlers or {}
    return get_group_script(added, handlers.added, handlers.no_added or "# Nothing to add") ..
        get_group_script(removed, handlers.removed, handlers.no_removed or "# Nothing to remove") ..
        get_iterative_script(added, handlers.add) ..
        get_iterative_script(removed, handlers.remove)
  end
end

-- Creates a migrator that acts upon a set, adding various commands for added/removed items
return function(name, desired_sequence, handlers, key)
  return migrator(name, function(current_sequence)
    current_sequence = current_sequence or {}
    local current = Set:create(current_sequence, key)
    local desired = Set:create(desired_sequence, key)
    local added = desired:diff(current):to_table()
    local removed = current:diff(desired):to_table()
    return {
      next_state = desired_sequence,
      migration = get_sequence_script(added, removed, handlers.migration),
      sync = get_sequence_script(added, removed, handlers.sync)
    }
  end)
end
