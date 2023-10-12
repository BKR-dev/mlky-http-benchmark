#!/usr/bin/env bash
#
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

for ROUTER in "${binary_array[@]}"; do
	printf "Starting $ROUTER in its own process...\n"
	parallel ::: ./$ROUTER 
done
}

buildBinaries
startRouters
