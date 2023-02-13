# Overview

Gives all combinations of an array/slice that contains unique items.

## The algorithm

This crate uses binary iteration to generate the possible combinations by
utilizing the 1s bits and 1s counts to get the combinations.

If given an array \[1, 2, 3\] it will loop from 1 to 1 << 3 (that is binary 1000
and decimal 8) which gives the 7 combinations of \[1, 2, 3\].

This gives the combinations: \[001, 010, 011, 100, 101, 110, 111\]
Which is: \[\[1\], \[2\], \[1, 2\], \[3\], \[1, 3\], \[2, 3\], \[1, 2, 3\]\]

## functionality

- all combinations
- all qualifying combinations
- positions of all qualifying combinations
- combinations of a particular length
- positions of combinations of a particular length
- positions of qualifying combinations of a particular length

_NOTE: positions is the position in the array of arrays of all combinations
possible, where the list is 1 indexed rather than 0. Therefore if you have
array \[1, 2\] and your total set is \[\[1\], \[2\], \[1, 2\]\] then position 1
is the array \[1\] and position 3 is the array \[1, 2\]_

## Why unique only

Providing an array with non-unique items, it will give you more than you want.

This package could have been made more robust to prevent duplication, but
that would reduce it less performant. If you require a package that
handles cases of non-unique arrays, please look elsewhere

### Example:

this fails:

```
 actual := Combinations(&[]int{1, 2, 2, 3}, 3);
 expected := [][]int{
    {1, 2, 2},
    {1, 2, 3},
    {2, 2, 3},
 }
 // would throw an error here:
 if actual != expected {
     throw error
 }
```

this passes:

```
 actual := Combinations(&[]int{1, 2, 2, 3}, 3);
 expected := [][]int{
    {1, 2, 2},
    {1, 2, 3},
    {1, 2, 3}, // note duplicate
    {2, 2, 3},
 }
if actual != expected {
    throw error
}
```
