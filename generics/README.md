# generics experiments in Go

See reference: https://go.dev/blog/intro-generics

## When / not to use generics
Also watch good advice from Ian Lance Taylor on when
and when *not* to use generics.

See: https://youtu.be/Pa_e9EeCdy8?t=1255

* Write code, don't design types. Start by writing functions.

* Avoid writing boiler plate code.

## Use generics, when:

* Functions that work on slices, maps and channels of any element type.
 Eg. keys(m map[K]V) K

* General purpose data structure.
 Eg. SiuYin's idea for a priority set (like a heap but with only unique elements)
* When operating on type parameters, prefer functions to methods.
 Eg. For priority sets, pass a *less function*, rather than require a less *method* in the data structure.
* When a method all looks the same.
 Counter-example: area of rectangle is different from area of a circle.
 While area of a square or circle both take one parameter, their computation is different. side squared vs pi time radius squared.

## Do not use generics, when:

* You're just calling a method. Use interfaces instead.
* When the common method have different implementations, eg. area of a cicle vs area of a rectangle.
* Use reflection when the type does not have methods and the implementation is different for each type.



## main.go
Hello world with computing the area of a shape.

## double
Generic double method demonstrating doubling inventory.

