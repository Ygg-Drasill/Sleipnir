(module
  (global $Bar.c (mut i32) (i32.const 0))
  (global $Foo.y (mut i32) (i32.const 0))
  (global $Foo.x (mut i32) (i32.const 0))

  (func $Bar
    call $Foo
    global.get $Foo.x
    global.get $Foo.y
    i32.add
    ;; drop / return
  )

  (func $Foo
    i32.const 1
    global.set $Foo.y
    i32.const 2
    global.set $Foo.x
  )
)
