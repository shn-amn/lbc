# lbc
## A tiny fizz-buzz server in Go

This service exposes two endpoint, `/fizzbuzz` and `/statistics`.
The first endpoint plays a little fizz-buzz game with you,
and the second one tells you how most people like to play the game.

In this exercise, I have only used Go's standard library.

## The game of fizz-buzz
The classical game of fizz-buzz consists of counting up from one
while substituting every multiple of 3 by the word "fizz",
every multiple of 5 by "buzz",
and every multiple of both by "fizzbuzz".

The generalised version exposed by this service takes five parameters:
- a positive integer `int1` that replaces 3 in the classical version,
- a positive integer `int2` that replaces 5,
- a string `str1` that replaces "fizz",
- a string `str2` that replaces "buzz",
- and a positive integer `limit`.
It then plays the game for you from 1 up to (and including) `limit`.

### The route `/fizzbuzz`
To an `GET` request with the five query parameters listed above,
the service responds with a JSON array containing
`"<str1>"` for every multiple of `int1`,
`"<str2>"` for every multiple of `int2`,
`"<str1><str2>"` for every multiple of both,
and `"<n>"` for every other integer `<n>`,
going from 1 to `<limit>` including both ends.
As in this example:

```bash
$ curl "localhost:8080/fizzbuzz?int1=5&int2=7&str1=hello&str2=world&limit=50"

["1","2","3","4","hello","6","world","8","9","hello","11","12","13","world","hello","16","17","18","19","hello","world","22","23","24","hello","26","27","world","29","hello","31","32","33","34","helloworld","36","37","38","39","hello","41","world","43","44","hello","46","47","48","world","hello"]
```

### Errors
There are no default values for parameters.
So if any of the parameters is missing there will be an error.
Also, the numerical parameters in the query must be positive decimal integers greater than zero.
If any of these conditions is violated,
the response will contain the status code 400 for bad request,
and a body with description of the errors in form of a JSON array of strings.

> In case of **multiple errors** care is taken not to stop at the first error
  and make an exhaustive list of all the irregularities in the input.

```bash
$ curl "localhost:8080/fizzbuzz?int1=0&int2=abc&str1=hello"

["Parameter 'int1' must be a positive integer but provided value '0' is not.","Parameter 'int2' must be a positive integer but provided value 'abc' is not.","Parameter 'limit' is not specified.","Parameter 'str2' is not specified."]
```

To every HTTP method other than get, the response is the typical `404 page not found`
(without the JSON content type header).
