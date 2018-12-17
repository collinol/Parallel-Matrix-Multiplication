# Parallel Implementation of Dense Matrix Chain Multiplication
`go run src/matrix_multiplication/matrixmultiply.go `
```
Usage:
    -f=value : Specifies directory size. Takes arguments
        s: Small
        m: Medium
        l: Large
    default = s
    (note) size refers to dimensions of matrices contained within, not the number of files within the directory.

    -p=value : Specifies the upper bound limit on the number of go routines allowed. Leaving out this flag will run the sequential version. Because the parallel version requires at least 2 threads, running with -p=<less than 3> will run the sequential version

    -a : Specifies whether or not to optimize the multiplication operations with the matrix chain algorithm. No value required, -a will set it to true

    --help : provides usage statement
     
```

See Report for further details about implementation