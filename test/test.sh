#!/usr/bin/env bash

for i in {1..100}; do
  grpcurl -import-path ../proto/account -proto user.proto -d '{"email":"'${i}'@gmail.com", "password":"hogehogehoge", "name":"今", "company_id":"2"}' -plaintext localhost:50051 account.UserService.RegisterUser
done
