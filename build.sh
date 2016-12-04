#!/bin/bash

if [ -n "$1" ]; then
	maingo=${1}.go
else
	maingo=$(ls -l | grep -o -P '(?<=\s)\d+.go(?=$)')
fi
base=${maingo%.go}
mainexe=${base}
mainarc=${base}.a
mainsrc=${base}.go

list=$(ls *.go | grep -v "$maingo")
echo compile -pack -o m.a $list
compile -pack -o m.a $list

echo compile -I . -pack -o $mainarc $mainsrc
compile -I . -pack -o $mainarc $mainsrc

echo link -L . -o $mainexe $mainarc
link -L . -o $mainexe $mainarc

rm -f *.a


