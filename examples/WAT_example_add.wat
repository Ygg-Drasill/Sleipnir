(module
(func $Bar (param $a i64) (param $b i64) (result i64)
  local.get $a
  local.get $b
  i64.add)
(func $Foo
  (local $x i64)
  (local $y i64)
  i64.const 1
  local.set $x
  i64.const 2
  local.set $y
  local.get $x
  local.get $y
  call $Bar
  drop
  )
)