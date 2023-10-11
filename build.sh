#!/usr/bin/env bash
#
#
# read all files beginning with main_${}
# execute go build -tags ${} -o ${}-router accordingly
# check if all hs been built: exit code!?
#bash
# and then just ./${}
#
# check if stuff is already built, no need to do twice
for f in main_*
do
	ROUTER_NAME=$(echo "$f" | cut -d"_" -f2 | cut -d"." -f1)
	go build -tags $ROUTER_NAME -o $ROUTER_NAME-router
done

# execute the binaries (DIFICULTY IMPOSSIBLE)
for f in main_*
do
	ROUTER_NAME=$(echo "$f" | cut -d"_" -f2 | cut -d"." -f1)
	parallel ::: ./$ROUTER_NAME-router
done
