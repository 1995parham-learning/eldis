# The Elahe's Redis implementation

## Introduction

This is the question, there are two commands "GET" and "SET", "SET" command gets a key and it's value and puts them in
the memory as you can guess the "GET" method get a key and returns its related value.There is a threshold for memory and
when the number of pairs in the memory reaches the threshold it flushes the data into a file and so on, when it wants to
search for a value it first searches the memory then the latest file and so on.

## Search

How to understand which file is newer and which one older?<br/>
I used time at the end of the file name.<br/>
What if the file is too big?<br/>
For better performance I don't read whole the file I use binary search to read lines.<br/>
