Remove-Item ./pkg/gocc -Recurse -Force

gocc -no_lexer -a -v -o "./pkg/gocc" pkg/yggdrasill.bnf
