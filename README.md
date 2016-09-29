go-levenshtein
==============

Levenshtein algorithm in Golang (with C ext)


Usage
-----

```go
levenshtein.Distance("kitten", "sitting")
=> 3

// Multibyte string support
levenshtein.Distance("あいうえお", "aいうえo")
=> 2
```


License
-------

This project is copyright by [Creasty](http://creasty.com), released under the MIT license.  
See `LICENSE.txt` file for details.
