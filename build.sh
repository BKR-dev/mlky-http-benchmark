#!/usr/bin/env bash
#
#set -x
# check if stuff is already built, no need to do twice
buildBinaries() {
for f in main_*
do
	ROUTER_NAME=$(echo "$f" | cut -d"_" -f2 | cut -d"." -f1)
	printf "Building $ROUTER_NAME as $ROUTER_NAME-router\n"
	go build -tags $ROUTER_NAME -o $ROUTER_NAME-router
done
printf "Done building all binaries.\n"
}


startRouters() {
printf "Starting all binaries via GNU parallel\n"
binary_array=()
for f in main_*
do
	ROUTER_NAME=$(echo "$f" | cut -d"_" -f2 | cut -d"." -f1)
	binary_array+=("$ROUTER_NAME-router")
done

for thing in "${binary_array[@]}"
do 
	./${thing} &
done
}

buildBinaries
startRouters
