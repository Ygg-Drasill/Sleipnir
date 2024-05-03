#!/bin/bash

rm -rf ./pkg/gocc

gocc -no_lexer -a -v -o "./pkg/gocc" pkg/yggdrasill.bnf
