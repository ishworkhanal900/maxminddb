# maxminddb

# Editor.md
Open source online Markdown editor.

## maxminddb
​
### Usage
#### Python: 
```bash
#chomd +x ./python/explorer.py
#./python/explorer.py
Enter path of mmdb: GeoLite2-City.mmdb  
Enter the IP: 8.8.8.8
Result:
{
    'continent': {'code': 'NA', 'geoname_id': 6255149, 'names': {'de': 'Nordamerika', 'en': 'North America', 'es': 'Norteamérica', 'fr': 'Amérique du Nord', 'ja': '北アメリカ', 'pt-BR': 'América do Norte', 'ru': 'Северная Америка', 'zh-CN': '北美洲'}}, 
    'country': {'geoname_id': 6252001, 'iso_code': 'US', 'names': {'de': 'Vereinigte Staaten', 'en': 'United States', 'es': 'Estados Unidos', 'fr': 'États Unis', 'ja': 'アメリカ', 'pt-BR': 'EUA', 'ru': 'США', 'zh-CN': '美国'}}, 
    'location': {'accuracy_radius': 1000, 'latitude': 37.751, 'longitude': -97.822, 'time_zone': 'America/Chicago'}, 
    'registered_country': {'geoname_id': 6252001, 'iso_code': 'US', 'names': {'de': 'Vereinigte Staaten', 'en': 'United States', 'es': 'Estados Unidos', 'fr': 'États Unis', 'ja': 'アメリカ', 'pt-BR': 'EUA', 'ru': 'США', 'zh-CN': '美国'}}
}

    
Java:
    #javac MmdbExplorer.class  (Optional if any modification required, create MmdbExplorer.class file )
    #java ./java/MmdbExplorer  ( You can directly run this skipping above step)
    Enter the path to the MMDB file:
    GeoLite2-City.mmdb
    Enter the IP address to look up:
    8.8.8.8
    City: null
    Country: United States
    Latitude: 37.751
    Longitude: -97.822

