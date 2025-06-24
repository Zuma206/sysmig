local sequence = require "@std.sequence"
local apt = {}

function apt.packages(desired_packages)
  return sequence("std.deb.packages", desired_packages, {
    migration = {
      added = function(added)
        return "sudo apt install -y " .. table.concat(added, " ")
      end,
      removed = function(removed)
        return "sudo apt remove -y " .. table.concat(removed, " ")
      end
    },
    sync = [[sudo apt update
sudo apt upgrade -y]]
  })
end

return apt
