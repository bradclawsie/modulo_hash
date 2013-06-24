## About

The so-called modulo hash is a simple implementation of a technique to even distributed sources to a 
list of targets. Typically, the source would be a job or request to be assigned to a single host in
a list of hosts. The technique is simple: we hash the source string to a int, and this int is mod'd with
the number of targets in the targets list. 

## Installing

   $ go get github.com/bradclawsie/modulo_hash

## Docs

   $ go doc github.com/bradclawsie/modulo_hash

## Examples

The included unit test file contains examples of the various ways the modulo hash can be used.

