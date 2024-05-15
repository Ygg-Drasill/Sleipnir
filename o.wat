(module
(import "console" "log" (func $log (param i32)))
(global $Foo_y (mut i32) (i32.const 0))
(global $Foo_x (mut i32) (i32.const 0))
(global $Foo_processed (mut i32) (i32.const 0))
(global $Bar_c (mut i32) (i32.const 0))
(global $Bar_processed (mut i32) (i32.const 0))
(global $Add0_result (mut i32) (i32.const 0))
(global $Add0_processed (mut i32) (i32.const 0))
(global $Print0_processed (mut i32) (i32.const 0))
(func $Foo
(global.set $Foo_y (i32.const 2))
(global.set $Foo_x (i32.const 1))
call $Add0
call $Bar
)
(func $Bar
(global.set $Bar_c (i32.const 0))
global.get $Foo_processed
(if (then nop) (else return))
global.get $Foo_processed
(if (then nop) (else return))
call $Add0
)
(func $Add0
global.get $Bar_processed
(if (then nop) (else return))
global.get $Foo_processed
(if (then nop) (else return))
global.get $Foo_x
global.get $Bar_c
i32.add
global.set $Add0_result
call $Print0
)
(func $Print0
global.get $Add0_processed
(if (then nop) (else return))
global.get $Add0_result
call $log
)
(func (export "root")
call $Foo
)
)