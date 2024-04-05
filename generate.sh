#!/bin/bash

rm -rf ./compiler/gocc

gocc -no_lexer -a -v -o "./compiler/gocc" compiler/yggdrasill.bnf
