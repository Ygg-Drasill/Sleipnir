package generator

var (
	imports = `(import "console" "log" (func $_log (param i32)))
(import "screeps" "move" (func $_move (param i32)))
(import "screeps" "set" (func $_set (param i32)))
(import "screeps" "get" (func $_get (result i32)))`
)
