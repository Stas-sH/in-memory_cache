# in-memory_cache
it is my package for data storage in the map[key] value 
key type is string
value type is interface{}

metod NewMemory() :
create and return new Memory type var

metod Set(key string, value interface{}):
adds new data

metod Get(key string) :
returns data by key;
if there is no data - returns an error

metod Delete(key string):
delete value by key (if key exists)
