# adamctl

My professional swiss army knife / toolkit.

## Usage

### `date`

```
❯ adamctl date
Sun Jun 18 12:45:32 PDT 2023

UTC: Sun Jun 18 19:45:32 UTC 2023
Los Angeles: Sun Jun 18 12:45:32 PDT 2023
Denver: Sun Jun 18 13:45:32 MDT 2023
New York: Sun Jun 18 15:45:32 EDT 2023
Turkey: Sun Jun 18 22:45:32 +03 2023
```

```
❯ adamctl date "Sun Jun 18 13:18:41 PDT 2023"
Sun Jun 18 13:18:41 PDT 2023

UTC: Sun Jun 18 20:18:41 UTC 2023
...
```

```
❯ echo ""Sun Jun 18 13:18:41 PDT 2023"" | adamctl date
Sun Jun 18 13:18:41 PDT 2023

UTC: Sun Jun 18 20:18:41 UTC 2023
...
```

Date strings are parsed using [araddon/dateparse](https://github.com/araddon/dateparse):
```
❯ go run main.go date "2012-08-03 18:31:59.257000000 +0000 UTC"
Fri Aug  3 18:31:59 +0000 2012

UTC: Fri Aug  3 18:31:59 UTC 2012
...
```