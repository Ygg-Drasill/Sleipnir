(module
  (global $Foo_Has_Run (mut i32) (i32.const 0))
  (global $Foo.y (mut i32) (i32.const 0))
  (global $Foo.x (mut i32) (i32.const 0))

  (func $Bar
    (local $a i32)
    (local $b i32)
    global.get $Foo_Has_Run
    (if
      (then
        nop
      )
      (else
        return
      )
    )
    (global.get $Foo.x)
    (local.set $a)
    (global.get $Foo.y)
    (local.set $b)

    (local.get $a)
    (local.get $b)

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