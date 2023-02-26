#!/usr/bin/env python3
import maxminddb
import sys

reader = maxminddb.open_database('./GeoLite2-City.mmdb')

#result = reader.get('57.135.152.12')
#result = reader.get(sys.argv[1])

input_ip=input("Enter the IP: ")
result = reader.get(input_ip)


print(f"Result:{result}")

reader.close()