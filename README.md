# Fuzzy matching: Go vs Python

## Background 

I have been recently tasked with perform fuzzy matching of two large lists of companies. Noticing that Python was struggling with a large amount of calculations, I decided to run a quick test comparing its performance using the [`fuzzywuzzy`](https://github.com/seatgeek/fuzzywuzzy) library, which is more or less the de facto library for fuzzy matching in Python, against the same code in Golang, [using a port](https://github.com/paul-mannino/go-fuzzywuzzy) of the very same library, which is linked from within the original Seatgeek repository. The code should read the lists from CSV files, grab the necessary columns and output a result file with two new columns: The closest name match found and the score to its similarity. 

## Source files

For this exercise I am using the S&P500 companies list in two different files that are slightly different, so despite there will be some 100% matches, there will still be a lot of fuzziness. The original file is `constituents.csv` and the file that contains the matches is `sp500.csv`

## Additional notes

In the Python version I have used the standard `csv` library rather than `Pandas` to read the files. Likewise, the Go code also uses `encoding/csv` without any wrapper. 

## Results

### Python:
    ruben@MacBook > ~/Code/fuzzy > time python3 fuzzy.py
    python3 fuzzy.py  19,51s user 0,11s system 98% cpu 19,873 total

### Golang:

With `go run`

    ruben@MacBook > ~/Code/fuzzy > time go run fuzzy.go
    go run fuzzy.go  74,11s user 1,42s system 102% cpu 1:13,43 total

Compiling first with `go build` doesn't make much of a difference

    ruben@MacBook > ~/Code/fuzzy > time ./fuzzy  
    ./fuzzy  75,04s user 1,28s system 98% cpu 1:17,51 total


### Conclusion

For this exercise, Python was faster than Go. 
HOWEVER, for an insight about why is it faster (mostly blaming it to the `go-fuzzywuzzy` implementation), see this answer:

https://codereview.stackexchange.com/questions/236189/are-there-ways-to-speed-up-this-string-fuzzy-matching-in-golang/

## UPDATE!

The author of `go-fuzzywuzzy` kindly applied the fixes described in the codereview answer, and now it outperforms Python:

    ruben@MacBook > ~/Code/fuzzy > time go run fuzzy.go
    go run fuzzy.go  11,31s user 0,60s system 102% cpu 11,614 total
