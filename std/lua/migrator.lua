-- Create a basic migrator
return function(name, func)
  return { name = name, func = func }
end
