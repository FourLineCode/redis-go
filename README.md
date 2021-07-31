# Redis-Go: A command line Redis implementation

This is a simple redis clone written in golang. its a cli tool which has some basic \
features of redis in memory database system. This project was built for fun and \
learning new language features.

# Run CLI

Clone the project and run the following commands

```
make run
```

# Documentation

### Available actions

-   **GET**

    -   Usage: `GET {key}`
    -   Description: gets the corresponding value for the key and prints to the console.

-   **SET**

    -   Usage: `SET {key} {value}`
    -   Description: sets the value of the given key to the given value.

-   **DEL**

    -   Usage: `DEL {key}`
    -   Description: deletes the entry for the given key.

-   **HGET**

    -   Usage: `HGET {key} {field}`
    -   Description: gets the value for the field of corresponding hash key.

-   **HSET**

    -   Usage: `HSET {key} {field} {value}`
    -   Description: sets the value of the field of corresponding hash key.

-   **HGETALL**

    -   Usage: `HGETALL {key}`
    -   Description: gets all the values of the correspoding hash key.

-   **HDEL**

    -   Usage: `HDEL {key}`
    -   Description: deletes the entry for the given hash key.

-   **LPUSH**

    -   Usage: `LPUSH {key} {value}`
    -   Description: pushes value on the left of the list of given key.

-   **RPUSH**

    -   Usage: `RPUSH {key} {value}`
    -   Description: pushes value on the right of the list of given key.

-   **LPOP**

    -   Usage: `LPOP {key}`
    -   Description: gets one value from the left of the given list and deletes the value.

-   **RPOP**

    -   Usage: `RPOP {key}`
    -   Description: gets one value from the right of the given list and deletes the value.

-   **LINDEX**

    -   Usage: `LINDEX {key} {index}`
    -   Description: gets the value of the given index from the corresponding list key.

-   **LRANGE**

    -   Usage: `LRANGE {key} {start} {end}`
    -   Description: gets the list of values from start to end index of the corresponding list key.

-   **LDEL**

    -   Usage: `LDEL {key}`
    -   Description: deletes the entry for the given list key.

-   **EXIT**

    -   Usage: `EXIT`
    -   Description: exits the program.
