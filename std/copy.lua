-- Creates a shallow copy of the table t1
return function(t1)
  local t2 = {}
  for key, value in pairs(t1) do
    t2[key] = value
  end
  return t2
end
