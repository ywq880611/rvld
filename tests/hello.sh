#!/bin/bash

test_name=$(basename $0 .sh)
tmp_path=out/tests/$test_name
mkdir -p "$tmp_path"

cat << EOF | riscv64-linux-gnu-gcc-10 -o "$tmp_path"/a.o -c -xc -
#include <stdio.h>

int main(){
    printf("Hello, world\n");
    return 0;
}
EOF

./rvld "$tmp_path"/a.o