-- A simple table that contains all submodules of std
return {
  copy = require "@std.copy",
  map = require "@std.map",
  migrator = require "@std.migrator",
  sequence = require "@std.sequence",
  Set = require "@std.Set",
  system = require "@std.system",
  rhel = require "@std.rhel",
  deb = require "@std.deb",
  serialize = require "@std.serialize",
  files = require "@std.files",
  path = require "@std.path",
  dir = require "@std.dir",
  file = require "@std.file",
  symlinks = require "@std.symlinks"
}
