modulo_hash
===========

a hash based on simple modulo 

this is a very simple module:

given a list of target strings, a source element can be fairly mapped on to one of those
target strings by taking the crc32 sum of the source element and getting the modulo of this
value with the length of the list of target strings. 

this is useful for load balancing, for example. run the tests to see how this function
provides a good distribution of maps of sources to targets.