local path = require "@std.path"

local file = {}

function file.pass_paths(func)
  return function(file)
    local destination = path(file[1])
    local source = path(file[2])
    return func(source, destination)
  end
end

function file.remove(_, destination)
  return "rm -rf " .. destination
end

return file
