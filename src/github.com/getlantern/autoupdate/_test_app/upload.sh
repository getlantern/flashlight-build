#/bin/bash

for i in $(seq 1 4); do
	sed s/"internalVersion =.*"/"internalVersion = $i"/g main.go > main.go.tmp
	mv main.go.tmp main.go
	go build -o main.v$i
	releasetool -config equinox.yaml -arch amd64 -os darwin -version $i -channel stable -source main.v$i
done
