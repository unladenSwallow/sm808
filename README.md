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
- util
- sound

To run the tests, navigate to the package you would like to test and run:
```shell script
go test
```

### Notes on submission:
As a backend programmer I chose to focus on the timing aspect of the assignment rather than the display. 
I could have set up a handler to run a browser program with images rather than text appearing in each step.
However, I thought that implementing the timing feature via a goroutine and ticker would do more to show
my advanced experience with Go. I hope you enjoy my submission. I look forward to chatting more!