#!/usr/bin/env bash


# [GET]
wrk -c 100 -d 10s -t100 http://localhost:8080/v1/student

# [POST]
wrk -c 100 -d 10s -t100 --script=post.lua --latency http://localhost:8080/v1/student

# [GET]
wrk -c 100 -d 10s -t100 http://localhost:8080/v1/student/user_1

# [PUT]
wrk -c 100 -d 10s -t100 --script=put.lua --latency http://localhost:8080/v1/student/user_1

# [DELETE]
wrk -c 100 -d 10s -t100 --script=delete.lua --latency http://localhost:8080/v1/student/user_1

