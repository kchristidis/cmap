/*
Package cmap introduces a thead-safe capacity-constrained hash table that
evicts key/value pairs  according to the LRU (least recently used) policy.

It is technically a thin wrapper around Go's built-in map type.
*/
package cmap
