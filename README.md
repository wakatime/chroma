# Chroma — A general purpose syntax highlighter in pure Go
[![Golang Documentation](https://godoc.org/github.com/alecthomas/chroma?status.svg)](https://godoc.org/github.com/alecthomas/chroma) [![CI](https://github.com/alecthomas/chroma/actions/workflows/ci.yml/badge.svg)](https://github.com/alecthomas/chroma/actions/workflows/ci.yml) [![Slack chat](https://img.shields.io/static/v1?logo=slack&style=flat&label=slack&color=green&message=gophers)](https://invite.slack.golangbridge.org/)

This is a fork for detecting languages in wakatime-cli using Chroma Lexers.

Chroma is based heavily on [Pygments](http://pygments.org/), and includes
translators for Pygments lexers and styles.

### Identifying the language

To highlight code, you'll first have to identify what language the code is
written in. There are three primary ways to do that:

1. Detect the language from its filename.

    ```go
    lexer := lexers.Match("foo.go")
    ```

3. Explicitly specify the language by its Chroma syntax ID (a full list is available from `lexers.Names()`).

    ```go
    lexer := lexers.Get("go")
    ```

3. Detect the language from its content.

    ```go
    lexer := lexers.Analyse("package main\n\nfunc main()\n{\n}\n")
    ```

In all cases, `nil` will be returned if the language can not be identified.

```go
if lexer == nil {
  lexer = lexers.Fallback
}
```

At this point, it should be noted that some lexers can be extremely chatty. To
mitigate this, you can use the coalescing lexer to coalesce runs of identical
token types into a single token:

```go
lexer = chroma.Coalesce(lexer)
```
