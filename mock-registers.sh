#!/bin/bash

url="http://localhost:8000/godating-dealls/api/authenticate/register"
contentType="Content-Type: application/json"

for i in {1..50}
do
    email="u${i}@gmail.com"
    data='{
        "email": "'${email}'",
        "username": "u'${i}'",
        "password": "test1234",
        "full_name": "U '${i}'"
    }'

    curl --location "$url" --header "$contentType" --data-raw "$data"
    echo "Request $i completed with email $email"
done
