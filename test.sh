#!/bin/bash
cd asm;
make all;
cd ..;
go test ./lib/cpu
