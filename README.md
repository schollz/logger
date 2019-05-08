# logger

<img src="https://img.shields.io/badge/coverage-78%25-brightgreen.svg?style=flat-square" alt="Code coverage">&nbsp;<a href="https://travis-ci.org/schollz/logger"><img src="https://img.shields.io/travis/schollz/logger.svg?style=flat-square" alt="Build Status"></a>&nbsp;<a href="https://godoc.org/github.com/schollz/logger"><img src="http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square" alt="Go Doc"></a> 

Simplistic, opinionated logging for Golang

- Zero dependencies
- Global logger (with optional local logger)
- Leveled
- Useful defaults / i.e. zero-config
- Simple API



## Install

```
go get github.com/schollz/logger
```

## Usage 


```golang
package main

import (
	log "github.com/schollz/logger"
)

func main() {
	log.SetLevel("debug")
	log.Debug("hello, world")
}
```

## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
