-- A simple table that contains all submodules of std
return {
  copy = require "@std.copy",
  map = require "@std.map",
  migrator = require "@std.migrator",
  sequence = require "@std.sequence",
  set = require "@std.set",
  system = require "@std.system",
  rhel = require "@std.rhel",
  deb = require "@std.deb",
}
