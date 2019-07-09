
# What's this ?

This program aims to be a HTTP proxy imitating the Docker Daemon REST API
However, the real backend is a Kubernetes Cluster

# How to use
```
go run main.go structs.go
curl http://localhost:3000/v1.24/containers/json | json_pp
```

You can even point your Docker CLI toward it !
```
$ docker -H localhost:3000 ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
d6f0fbc2-839        fieldbox-dev        ""                  3 hours ago         Running                                 evelop-app-0
f02ddccc-83a        2017-latest         ""                  About an hour ago   Running                                 evelop-mssql-0
d6d47bcd-839        3-management        ""                  3 hours ago         Running                                 evelop-rabbitmq-0
e354bade-839        harbor.ops.f        ""                  3 hours ago         Running                                 evelop-redis-6379-0
dd5efa20-839        harbor.ops.f        ""                  3 hours ago         Running                                 evelop-redis-6380-0
d6cd7f71-839        harbor.ops.f        ""                  3 hours ago         Running                                 evelop-riak-0
f0429ded-83a        fieldbox-mas        ""                  About an hour ago   Running                                 aster-geop-hq-1-app-0
65ef913e-836        2017-latest         ""                  10 hours ago        Running                                 aster-geop-hq-1-mssql-0
f00348f7-83a        3-management        ""                  About an hour ago   Running                                 aster-geop-hq-1-rabbitmq-0
79e14ec3-836        harbor.ops.f        ""                  10 hours ago        Running                                 aster-geop-hq-1-redis-6379-0
efa44a7a-83a        harbor.ops.f        ""                  About an hour ago   Running                                 aster-geop-hq-1-redis-6380-0
6c86965e-836        harbor.ops.f        ""                  10 hours ago        Running                                 aster-geop-hq-1-riak-0
70593239-836        fieldbox-mas        ""                  10 hours ago        Running                                 aster-sbm-stg-1-app-0
65e1073d-836        2017-latest         ""                  10 hours ago        Running                                 aster-sbm-stg-1-mssql-0
ef961fd1-83a        3-management        ""                  About an hour ago   Running                                 aster-sbm-stg-1-rabbitmq-0
83a90db9-836        harbor.ops.f        ""                  10 hours ago        Running                                 aster-sbm-stg-1-redis-6379-0
860c9147-836        harbor.ops.f        ""                  10 hours ago        Running                                 aster-sbm-stg-1-redis-6380-0
6a676d22-836        harbor.ops.f        ""                  10 hours ago        Running                                 aster-sbm-stg-1-riak-0
6e71080a-836        fieldbox-rel        ""                  10 hours ago        Running                                 review-app-0
78f22d89-836        2017-latest         ""                  10 hours ago        Running                                 review-mssql-0
f025ba18-83a        3-management        ""                  About an hour ago   Running                                 review-rabbitmq-0
efdada26-83a        harbor.ops.f        ""                  About an hour ago   Running                                 review-redis-6379-0
f00ab91c-83a        harbor.ops.f        ""                  About an hour ago   Running                                 review-redis-6380-0
6db6bcc2-836        harbor.ops.f        ""                  10 hours ago        Running                                 review-riak-0
ae80e5ef-803        bitnami/kube        ""                  4 days ago          Succeeded                               ebalance-pods-after-scale-up-at-day-1558933800-qmtw7
d941abcb-810        bitnami/kube        ""                  3 days ago          Succeeded                               ebalance-pods-after-scale-up-at-day-1559020200-5kwzp
041e9153-81d        bitnami/kube        ""                  2 days ago          Succeeded                               ebalance-pods-after-scale-up-at-day-1559106600-7466h
2e3b1ab0-829        bitnami/kube        ""                  34 hours ago        Succeeded                               ebalance-pods-after-scale-up-at-day-1559193000-h4j86
5aa1088a-836        bitnami/kube        ""                  10 hours ago        Succeeded                               ebalance-pods-after-scale-up-at-day-1559279400-s957z
efc6711a-83a        fieldbox-rel        ""                  About an hour ago   Running                                 elease-geop-hq-1-app-0
772388a9-836        2017-latest         ""                  10 hours ago        Running                                 elease-geop-hq-1-mssql-0
759723d1-836        3-management        ""                  10 hours ago        Running                                 elease-geop-hq-1-rabbitmq-0
72fd69dd-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-geop-hq-1-redis-6379-0
6b569883-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-geop-hq-1-redis-6380-0
685cb51e-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-geop-hq-1-riak-0
f016c4c0-83a        fieldbox-rel        ""                  About an hour ago   Running                                 elease-ke-hq-1-app-0
6d323e6c-836        2017-latest         ""                  10 hours ago        Running                                 elease-ke-hq-1-mssql-0
efade8f3-83a        3-management        ""                  About an hour ago   Running                                 elease-ke-hq-1-rabbitmq-0
6bcfb18a-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-ke-hq-1-redis-6379-0
6fd13ef9-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-ke-hq-1-redis-6380-0
efba2592-83a        harbor.ops.f        ""                  About an hour ago   Running                                 elease-ke-hq-1-riak-0
73ecb51a-836        fieldbox-rel        ""                  10 hours ago        Running                                 elease-psc-pa-1-app-0
eff61979-83a        2017-latest         ""                  About an hour ago   Running                                 elease-psc-pa-1-mssql-0
78ec3c9b-836        3-management        ""                  10 hours ago        Running                                 elease-psc-pa-1-rabbitmq-0
7baab728-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-psc-pa-1-redis-6379-0
f034ae80-83a        harbor.ops.f        ""                  About an hour ago   Running                                 elease-psc-pa-1-redis-6380-0
72e0ea0f-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-psc-pa-1-riak-0
7705f7f3-836        fieldbox-rel        ""                  10 hours ago        Running                                 elease-sbm-cdi-1-app-0
710df861-836        2017-latest         ""                  10 hours ago        Running                                 elease-sbm-cdi-1-mssql-0
6f62f7c1-836        3-management        ""                  10 hours ago        Running                                 elease-sbm-cdi-1-rabbitmq-0
7d57db78-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-sbm-cdi-1-redis-6379-0
7c435cfb-836        harbor.ops.f        ""                  10 hours ago        Running                                 elease-sbm-cdi-1-redis-6380-0
f03ab8a5-83a        harbor.ops.f        ""                  About an hour ago   Running                                 elease-sbm-cdi-1-riak-0
9a87e0c4-7e6        google/cloud        ""                  6 days ago          Succeeded                               cale-down-at-night-1558728000-v4ncc
03c9c56c-80b        google/cloud        ""                  3 days ago          Succeeded                               cale-down-at-night-1558987200-pkq6t
2ce150ec-818        google/cloud        ""                  2 days ago          Succeeded                               cale-down-at-night-1559073600-jtpts
58a6d486-824        google/cloud        ""                  43 hours ago        Succeeded                               cale-down-at-night-1559160000-mjqwq
87212982-831        google/cloud        ""                  19 hours ago        Succeeded                               cale-down-at-night-1559246400-7g262
47b21201-803        google/cloud        ""                  4 days ago          Succeeded                               cale-up-at-day-1558933200-rchrs
723e373c-810        google/cloud        ""                  3 days ago          Succeeded                               cale-up-at-day-1559019600-d99nv
9cee2bfc-81c        google/cloud        ""                  2 days ago          Succeeded                               cale-up-at-day-1559106000-2mq9v
ccbed24f-829        google/cloud        ""                  34 hours ago        Succeeded                               cale-up-at-day-1559192400-ktgzd
f2f68d29-836        google/cloud        ""                  10 hours ago        Succeeded                               cale-up-at-day-1559278800-w95jq
f048c4ae-83a        fieldbox-mas        ""                  About an hour ago   Running                                 raining-1-app-0
ef82f16e-83a        2017-latest         ""                  About an hour ago   Running                                 raining-1-mssql-0
75d46640-836        3-management        ""                  10 hours ago        Running                                 raining-1-rabbitmq-0
6c68f27d-836        harbor.ops.f        ""                  10 hours ago        Running                                 raining-1-redis-6379-0
71690cb6-836        harbor.ops.f        ""                  10 hours ago        Running                                 raining-1-redis-6380-0
69c9c6ae-836        harbor.ops.f        ""                  10 hours ago        Running                                 raining-1-riak-0
```
