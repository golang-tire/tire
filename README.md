[![codecov](https://codecov.io/gh/golang-tire/tire/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-tire/tire)
[![Go Report Card](https://goreportcard.com/badge/github.com/golang-tire/tire)](https://goreportcard.com/report/github.com/golang-tire/tire)

### Golang tire cli

A `CLI` application to make create projects easier.

**_this repo is under heavy development. every thing can change before first stable release_**

#### Usage

```
tire --help
```

```

    tire generate golang projects

    Find more information at:
                https://github.com/golang-tire

Usage:
  tire [flags]
  tire [command]

Available Commands:
  embed       embed static codes into golang files
  gen         generate codes based on selected type
  help        Help about any command
  version     Print the tire cli version

Flags:
  -h, --help                    help for tire
      --profile string          Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex) (default "none")
      --profile-output string   Name of the file to write the profile to (default "profile.pprof")
```



