# HackerNews Go

[![Build Status](https://github.com/JoseThen/hn/workflows/Build%20Master/badge.svg)](https://github.com/JoseThen/hn/blob/master/.github/workflows/build_master.yaml)

[![Test Status](https://github.com/JoseThen/hn/workflows/Test%20Master/badge.svg)](https://github.com/JoseThen/hn/blob/master/.github/workflows/test_master.yaml)

A HackerNews client writtent in Go

---

## Usage

To get the top stories

```
hn -top
```

To get most recent (new) stories

```
hn -new
```

-   The default number of stories for both subcommands is 10. You can change that by using the `-count` flag
-   If you do not build the executable then you can simply do `go run main.go -count 30`
-   Note that `-top` will be set by default if no flag is set

## Output

When running a command you will get 2 columns for Title and URL in your terminal:
![Image of HN GO output](images/output.png)
