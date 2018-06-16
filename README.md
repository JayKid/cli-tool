# cli-tool
A Go experiment to build a cli-tool so that my long-ass aliases are not a pain anymore

## Things I would like to implement

+ Read (and later write) from JSON
+ Add aliases from the tool itself and save them to the JSON file
+ Make it a bit shorter to run aliases (if only no dashes needed...)

## Build

`go build -o t`

## Run

`./t -r {your_alias}`

## List aliases

`./t -l`
