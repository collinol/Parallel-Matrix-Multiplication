# Parallel Implementation of Dense Matrix Chain Multiplication
`go run src/matrix_multiplication/matrixmultiply.go `
```
Usage:
    -f(n) =value : Specifies directory size. Takes arguments
        s: Small (dimensions between 2 and 10)  
        m: Medium (dimensions between 10 and 500)  
        l: Large (dimensions between 500 and 1000)  
    default = s
    (note) size refers to dimensions of matrices contained within, not the number of files within the directory.  
    if files of that size have already been created, include the n like -fn  
    this will tell the program to bypass creating new files for that directory. -f by itself will create a new set of matrices
   
    -p=value : Specifies the upper bound limit on the number of go routines allowed.  
        Leaving out this flag will run the sequential version.   
        Because the parallel version requires at least 2 threads, running with -p=<less than 3> will run the sequential version

    -a : Specifies whether or not to optimize the multiplication operations with the matrix chain algorithm. No value required, -a will set it to true

    --help : provides usage statement
     
```

See Report for further details about implementation
