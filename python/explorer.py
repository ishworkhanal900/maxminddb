#!/usr/bin/env python3
import maxminddb
import json


input_db = input("Enter path of mmdb: ")

reader = maxminddb.open_database(input_db)

#result = reader.get('57.135.152.12')
#result = reader.get(sys.argv[1])

input_ip=input("Enter the IP: ")
result = reader.get(input_ip)


print(f"Result:{result}")

reader.close()