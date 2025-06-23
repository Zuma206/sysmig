local serialize = require "@std.serialize"
local sequence = require "@std.sequence"
local entries = require "@std.entries"

return function(files)
  return sequence("std.files", entries(files), {
    migration = {
      add = function(file)
        return "# Adding file " .. file[2] .. " to " .. file[1]
      end
    }
  }, serialize)
end
