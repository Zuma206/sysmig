local sequence = require "@std.sequence"
local map = require "@std.map"

local function get_key(component)
  return component.name
end

return function(components)
  return sequence("std.components", components, {
    migration = {
      add = function(component)
        return component.mount or ""
      end,
      remove = function(component)
        return component.unmount or ""
      end
    },
    sync = table.concat(
      map(components, function(component)
        return component.sync or ""
      end),
      "\n"
    )
  }, get_key)
end
