# The Elahe's Redis implementation (Eldis)

## Introduction

Eldis has two commands named "GET" and "SET". "SET" command gets a key, and it's value and puts them in
the memory as you can guess the "GET" method get a key and returns its related value.

There is a threshold for memory and
when the number of pairs in the memory reaches the threshold it flushes the data into a file and so on, when it wants to
search for a value it first searches the memory then the latest file and so on.

## Q&A

1. How to understand which file is newer and which one older?

> I used time at the end of the file name.

2. What if the file is too big?

> For better performance I don't read whole the file I use binary search to read lines.
