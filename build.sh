#!/usr/bin/env bash
#
routerArray=()
for e in main_*; do
	ROUTER=$(echo "$e" | cut -d"_" -f2 | cut -d"." -f1)
	routerArray+=("$ROUTER")
done

buildBinaries() {
for ROUTER in "${routerArray[@]}"; do
	ROUTER_NAME=$ROUTER-router
# if is not binary or does not exist, build
if [[ ! -x $ROUTER_NAME || ! -e $ROUTER_NAME ]]; then
	printf "Building $ROUTER as $ROUTER_NAME\n"
	go build -tags $ROUTER -o $ROUTER_NAME
else
	printf "$ROUTER_NAME already build, skipping build step\n"
fi
done
printf "Done building all binaries.\n"
}


startRouters() {
printf "Starting all binaries via GNU parallel\n"
for ROUTER in "${routerArray[@]}"; do
	ROUTER_NAME=$ROUTER-router
	printf "Starting $ROUTER in its own process...\n"
	parallel -j 50 ::: ./$ROUTER_NAME
done
}

buildBinaries
startRouters
