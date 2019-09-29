# Plain stupid diff algorythm

It doesn't really matter how dumb we are with it as soon as we can do the following. Given documents A and B:

```
B - A = diff
A + diff = B
```

To be pedantic we can expand all key deletions and potentially explode diff size. We'll get this then:

```
B + (-diff) = A
```

Two objects are the same, if their structure and values are the same.

If we're traversing an object, the logic looks like following.

```
For each key:
 - if value types are different, remove old key, add new key
 - if types are the same:
   * Primitive types: if values are equal, ignore, otherwise update value
   * Objects: recursively descend
   * Arrays: Array diff algrorhythm!
```

## Array diff algorythm

I see two ways there. First is to treat and array as an object and compare it by keys (see previous chapter).
Simplicity is there, but we can get a very inefficient diff as a results. Consider two arrays:

```
[1,2,3,4]
[2,3,4,5]
```

Object algorythm will lead to a full object replacement, however, we can treat key deletion no as a simple
value removal, but also as a shifting operation, so that we don't have any gaps. If we're smart enough,
we will get:

```
-.0 # delete first element
+0.3=5 # append to the end of the array. 3 is fine since removal of the first item will lead to the shit of the rest of elements
```
