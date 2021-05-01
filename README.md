# Redis-Go: A command line Redis implementation

This is a simple redis clone written in golang. its a cli tool which has some basic \
features of redis in memory database system. This project was built for fun and \
learning new language features.

# Run CLI

Clone the project and run the following command

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
    -   Description: deletes the entry for the given key
