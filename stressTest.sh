#!/bin/bash

# chi server stress test
echo "stress testing chi server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:3000/healthCheck

#echo server stress test
echo "stress testing echo server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:3443/healthCheck

# fiber stress test
echo "stress testing fiber server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:7443/healthCheck

# gin stress test
echo "stress testing gin server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:8443/healthCheck

# gorilla mux stress test
echo "stress testing gorilla mux server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://127.0.0.1:5443/healthCheck

# httprouter stress test
echo "stress testing httprouter server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:9000/healthCheck

# standard lib http stress test
echo "stress testing standard lib http server with 1,000,000 requests"
hey -m "GET" -n 1000000 http://localhost:5228/healthCheck
