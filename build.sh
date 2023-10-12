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
	printf "Building $ROUTER_NAME as $ROUTER_NAME-router\n"
	go build -tags $ROUTER_NAME -o $ROUTER_NAME-router
done
printf "Done building all binaries.\n"

# execute the binaries (DIFICULTY IMPOSSIBLE)
printf "Starting all binaries via GNU parallel\n"
for f in main_*
do
	ROUTER_NAME=$(echo "$f" | cut -d"_" -f2 | cut -d"." -f1)
	printf "Starting $ROUTER_NAME-router in its own process...\n"
	parallel ::: ./$ROUTER_NAME-router
done
