# LFU-Cache-Go

LFU cache with runtime complexity of O(1)

> Cache eviction algorithms are used widely in operating systems, databases
and other systems that use caches to speed up execution by caching data
that is used by the application. There are many policies such as MRU
(Most  Recently  Used),  MFU  (Most  Frequently  Used),  LRU  (Least  Re-
cently  Used)  and  LFU  (Least  Frequently  Used)  which  each  have  their
advantages and drawbacks and are hence used in specific scenarios.  By
far,  the most widely used algorithm is LRU, both for its O(1) speed of
operation as well as its close resemblance to the kind of behaviour that
is expected by most applications.

# Implementation

This LFU algorithm has a runtime complexity of O(1) for each of the dictionary operations (insertion, lookup and deletion) that can be performed on an LFU cache.

This insert operation has runtime of O(1) because every item being added to the cache does it so with the frequency of access of 1.
Therefore, we search for the frequency node of 1 and if it exists we add the new item to it otherwise we create the frequency node before adding to it.
