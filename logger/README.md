
# Extended `*log.Logger` lib

## Import the library

```
import (
    ml "github.com/tomekwlod/utils/logger"
)

```

## Simple usage

```
func main() {
    multi := io.MultiWriter(file, os.Stdout)

    l := ml.New(
        os.Getenv("LOGGING_MODE"),
        log.New(multi, "", log.Ldate|log.Ltime|log.Lshortfile),
    )

    l.Debugln("Debug message")
    l.Println("Info")
}
```

## Todo
So far  there are only three additional methods to print out the debug. Later I may implement more levels, but for the most cases the debug is really enough.