(module
  (global $Foo_Has_Run (mut i32) (i32.const 0))
  (global $Foo.y (mut i32) (i32.const 0))
  (global $Foo.x (mut i32) (i32.const 0))

  (func $Bar
    global.get $Foo_Has_Run
    (if
      (then
      	return
      )
    )
    global.get $Foo.x
    global.get $Foo.y
    i32.add
    return
  )


  (func $Foo
    i32.const 1
    global.set $Foo.y
    i32.const 2
    global.set $Foo.x


    i32.const 1
    global.set $Foo_Has_Run
    call $Bar
  )
)