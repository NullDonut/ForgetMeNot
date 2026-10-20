# ForgetMeNot
A bot pet with a squishy brain that learns facts you teach it, then forgets them along an Ebbinghaus-style forgetting curve unless you quiz it.

## Features
 - Teach Momo facts, memory strengths decays as R=e^(-t/S)
 - Quiz Mode picks the weakest memory, correct answers grow stability of memory S, wrong answers reduce memory stability.
 - Random Dreams and fact mashups (10%-15% times)
 - Pet Level grows with more facts and more stability.Decreases if pet is not mantained for a while.
 - Persistence to 'momo_memories.json'
 - Colours... (Might not work on windows and i am lazy to fix it)

 ## Download and Run
Grab a binary from [Releases](../../releases) page,then :
```
chmod +x forgetmenot-linux-amd64
./forgetmenot-linux-amd64
```

## Build from source

### PreRequistes
 - Go compiler, v1.27.1 and higher 

### Process 

1. Git clone
```
git clone https://github.com/NullDonut/ForgetMeNot.git/
```

2. Go to the directory and build it 
```
cd ForgetMeNot && go build .
```

3.Run the executable ;)
```
./forgetmenot
```


## Idea and Inspiration 
- My pyschology subject
- Go-ship website idea generator and exact idea is given below

ForgetMeNot: a bot pet with a squishy brain that learns facts you teach it, then forgets them using the real Ebbinghaus forgetting curve; memories decay over time unless you quiz it to reinforce them. Fun: it randomly dreams, mumbling half-remembered facts, blurts hilarious mashups when confused, and levels up as memories strengthen. Save/load your pet between sessions. Packages: fmt, bufio, time, math, encoding/json.


## Testing (Using go tests)
go test ./...

## License
UnLicense

