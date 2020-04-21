# SM-808 README
### Author: Leslie Pedro
### Date: April 2020
### Version: 0.0.1

## About SM-808
The SM-808 is a simulated drum machine. 

## Running SM-808
Download the solution to your gopath.

```shell script
 go build
``` 

```shell script
./sm808 -bpm <bpm> -duration <song-duration> ## arguments optional
```

### optional args:
- bpm - the beats per minute at which to play the song. Must be between 1 and 999. Default is 128 bpm.
- duration - the amount of song to play (in seconds). Must be between 1 and 999. Default is 30 seconds.

## Tests
Tests were written for the following packages:
- config
- sound

To run the tests, navigate to the package you would like to test and run:
```shell script
go test
```