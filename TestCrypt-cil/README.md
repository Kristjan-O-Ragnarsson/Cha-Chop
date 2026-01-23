# TestCrypt-cil

## About 

TestCrypt-cil is a program to run multiple iterations of test for binaries designed to dest latency of some kind of input processing in my case encryption and decryption. 
It was made to improve my testing following some advice found in [Systems Benchmarking Crimes by Gernot Heiser](https://gernot-heiser.org/benchmarking-crimes.html).

## Build

```bash
git clone git@github.com:Kristjan-O-Ragnarsson/TestCrypt-cil.git

cd TestCrypt-cli

go build
```


## Usage

````
Usage: Testcrypt [OPTIONS]

Options:


    -r              Set the amount of iterations that will be timed 
    
    -binary         set the name of the binary to test
    
    -no-warm-up     use if you want to skip warm up - an extra 5% iterations run before the 
                    timed ones to fill cache
                    
    -path           Set the path for the binary and test data
````

## Possible Improvments

- Random generation of Data using seed
- Total-time measurement
- figure/plot generation

