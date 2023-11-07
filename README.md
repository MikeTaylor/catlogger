# catlogger - categorical logging for Go

Copyright (C) 2023
[Index Data ApS.](https://www.indexdata.com/).

This software is distributed under the terms of the Apache License, Version 2.0. See the file [`LICENSE`](LICENSE) for more information.

<!-- md2toc -l 2 README.md -->
* [Introduction](#introduction)
* [Example usage](#example-usage)
* [API](#api)
    * [Constructor](#constructor)
    * [Logging](#logging)
    * [Category inquiry](#category-inquiry)
* [Provenance](#provenance)
* [Author](#author)


## Introduction

`catlogger` is a tiny library to do categorical logging for Go. 

By "categorical" or "category-based", we mean that categorical-logger can easily be configured to include or exclude different classes of log message (e.g. "app", "core", "calc", "okapiUrl", "supercalifragilisticexpialidocious",
whatever we decide to include). This is much more flexible than the traditional approach of escalating levels DEBUG -> INFO -> WARN -> ERROR -> FATAL. (It can also of course be turned off completely in production.)

Category names may not contain commas, and typically consist entirely of letters, with multiple-word categories spelled in camelCase.


## Example usage

```
import "github.com/MikeTaylor/catlogger"
var logger *catlogger.Logger
logger = catlogger.MakeLogger("listen,action")
logger.Log("config", fmt.Sprintf("%+v", cfg)) // Does not emit a message
logger.Log("listen", fmt.Sprintf("Listen port %d", port)) // Emits a message
```

This pulls in the library and creates a logger which is configured to emit messages in the categories "listen" and "action". Two messages are logged: one in the "config" category (for which no output will be generated since that category is not configured in the present logger) and one in the "listen" category (which _will_ be emitted).


## API

The API is gratifyingly small: a single class with a constructor and two methods (one of which is rarely if ever needed).

### Constructor

The constructor returns a logging object which carries a small amount
of configuration. That configuration is passed in when the object
is created. The constructor arguments, all optional, are:

1. `categories`. A comma-separated list of zero or more short strings, each of which is the name of a logging category. There is no predefined list of such categories: each application is welcome to make up its own.

2. `prefix`. If provided, a short string which is emitted at the beginning of each log message.

3. `timestamp`. A boolean. If true, then an ISO-formatted timestamp is included in each log message.

### Logging

All logging is done with a single method, `logger.log(STRING, ...VALUES)`. Its first argument is a string naming one of the application's logging categories, and the subsequent arguments are values to be included in the log message. The message is emitted only if the specified category is one of those configured in the logger.

Output is always to standard error.

### Category inquiry

You can ask a logger whether it has a particular category enabled using `logger.hasCategory(cat)`. This is a rather ugly back-door, but it's necessary for cases where another library does its own logging and you need to create a predicate for it based on the logger's categories.


## Provenance

This module is our old friend [`categorical-logger`](https://github.com/openlibraryenvironment/categorical-logger) (formerly [`@folio/stripes-logger`](https://github.com/folio-org/stripes-logger)), written in JavaScript, which traces its ancestry back to [the YAZ Toolkit's `yaz_log` function](https://github.com/indexdata/yaz/blob/0def0ed772fe6eda3b75404fca2c0f02d818892c/src/log.c#L487-L527).

For [`dumbass-news`](https://github.com/MikeTaylor/dumbass-news) I did [a simple port from the JavaScript version to Go](https://github.com/MikeTaylor/dumbass-news/tree/main/src/catlogger). Rather than replicate that in subsequent Go programs, I decided to release it as its own tiny package and depend on that.


## Author

Mike Taylor,
[Index Data ApS.](https://www.indexdata.com/)
Email
<mike@indexdata.com>

