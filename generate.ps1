Remove-Item ./compiler/gocc -Recurse -Force

gocc -no_lexer -a -v -o "./compiler/gocc" compiler/yggdrasill.bnf
