This is what I set up connect-go and connect-web on a trial.

## Installation

```
$ cd connect-sandbox
$ buf mod update # this will update buf.lock. I don't know how to install without buf.lock updated.
$ go work sync # maybe same as `cd api && go mod tidy`?
```

## Regenerate code

```
$ cd connect-sandbox
$ buf generate --include-imports
```
