/*
Copyright 2014 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package palm implements parallel architecture-friendly latch-free
modifications (PALM).  Details can be found here:

http://cs.unc.edu/~sewall/palm.pdf

The primary purpose of the tree is to efficiently batch operations
in such a way that locks are not required.  This is most beneficial
for in-memory indices.  Otherwise, the operations have typical B-tree
time complexities.

You primarily see the benefits of multithreading in availability and
bulk operations.

Benchmarks:

BenchmarkReadAndWrites-8	    			2000	    943051 ns/op
BenchmarkSimultaneousReadsAndWrites-8	     300	   4354136 ns/op
BenchmarkBulkAdd-8	     					 200	   5617104 ns/op
BenchmarkAdd-8	  						  200000	      8976 ns/op
BenchmarkBulkAddToExisting-8	     	     100	  20682933 ns/op
BenchmarkGet-8	 						 2000000	       629 ns/op
BenchmarkBulkGet-8	    					5000	    217915 ns/op

*/
package palm

import "github.com/Workiva/go-datastructures/common"

// BTree is the interface returned from this package's constructor.
type BTree interface {
	// Insert will insert the provided keys into the tree.
	Insert(...common.Comparator)
	// Get will return a key matching the associated provided
	// key if it exists.
	Get(...common.Comparator) common.Comparators
	// Len returns the number of items in the tree.
	Len() uint64
	// Dispose will clean up any resources used by this tree.  This
	// must be called to prevent a memory leak.
	Dispose()
}
