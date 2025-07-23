local sequence = require "@std.sequence"

return function(directories)
  return sequence("std.directories", directories, {
    migration = {
      add = function(directory)
        return "mkdir -p " .. directory
      end,
      remove = function(directory)
        return "rm -rf " .. directory
      end
    }
  })
end
