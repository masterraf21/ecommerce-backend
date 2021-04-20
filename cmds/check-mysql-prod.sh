#!/bin/bash

docker exec -it mysql_prod_ordering bash -c "printenv | grep MYSQL_VERSION"