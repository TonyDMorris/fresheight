## What is this ?

within this repo there are 2 git submodules which comprise the producer and consumer of the application.

## How to demo?

to demo both applications running together follow these instructions 


```bash
cd producer
```

```bash
go run ./cmd/main.go --number-of-groups=1000000 --batch-size=5000 --interval=1 --output-directory=../events
```


open a new terminal and navigate to the consumer 


```bash
cd ../consumer
```

```bash
go run cmd/main.go --input-directory=../events
````


expected output 
```bash
"Viewed": 8140848 
 "Interacted": 813878 
 "Click-Through": 895521 
```

## dependencies 

github.com/google/uuid

is used to generate uuids for the events everything else is provided by the standard lib 

## tests
Unfortunatley I wasn't able to implement any tests due to personal time constraints for the test but I will be happy to talk about what tests I would write and implement hopefully in the next step of the interview .

