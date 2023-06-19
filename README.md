# adamctl

My professional swiss army knife / toolkit.

## Usage

### `date`

```
❯ adamctl date
 The time in various places
 PLACE        OFFSET  DATE
 Raw          -7:00   Sun Jun 18 12:45:32 PDT 2023
 Local        -7:00   Sun Jun 18 12:45:32 PDT 2023
 UTC          +0:00   Sun Jun 18 19:45:32 UTC 2023
 Los Angeles  -7:00   Sun Jun 18 12:45:32 PDT 2023
 Denver       -6:00   Sun Jun 18 13:45:32 MDT 2023
 New York     -4:00   Sun Jun 18 15:45:32 EDT 2023
 Turkey       +3:00   Sun Jun 18 22:45:32 +03 2023
```

```
❯ adamctl date "Sun Jun 18 13:18:41 PDT 2023"
 The time in various places
 PLACE        OFFSET  DATE
 Raw          -7:00   Sun Jun 18 13:18:41 PDT 2023
 Local        -7:00   Sun Jun 18 13:18:41 PDT 2023
 UTC          +0:00   Sun Jun 18 20:18:41 UTC 2023
...
```

```
❯ echo "Sun Jun 18 13:18:41 PDT 2023" | adamctl date
 The time in various places
 PLACE        OFFSET  DATE
 Raw          -7:00   Sun Jun 18 13:18:41 PDT 2023
 Local        -7:00   Sun Jun 18 13:18:41 PDT 2023
 UTC          +0:00   Sun Jun 18 20:18:41 UTC 2023
...
```

Date strings are parsed using [araddon/dateparse](https://github.com/araddon/dateparse):
```
❯ adamctl date "2012-08-03 18:31:59.257000000 +0000 UTC"
 The time in various places
 PLACE        OFFSET  DATE
 Raw          +0:00   Fri Aug  3 18:31:59 +0000 2012
 Local        -7:00   Fri Aug  3 11:31:59 PDT 2012
 UTC          +0:00   Fri Aug  3 18:31:59 UTC 2012
...
```