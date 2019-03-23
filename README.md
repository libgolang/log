# Logger with no dependencies

- It wraps Golang's `log` package
- Adds Levels


## Usage

Download it

```
go get -u github.com/libgolang/log

```


In Code

```
import "github.com/libgolang/log"


func main() {

	log.SetLevel(log.LevelDebug) // defaults to log.LevelInfo

	someVar := "example"

	log.Debug("message", someVar)
	log.Debugf("message: %s", someVar)

	log.Info("message", someVar)
	log.Infof("message: %s", someVar)

	log.Warn("message", someVar)
	log.Warnf("message: %s", someVar)

	log.Error("message", someVar)
	log.Errorf("message: %s", someVar)

	log.Fatal("message", someVar)
	log.Fatalf("message: %s", someVar)
}

```

