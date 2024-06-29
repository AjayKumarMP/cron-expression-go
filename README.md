
# Cron Expression Parser

This Go program parses a cron expression and expands its fields into individual values. It also extracts and displays the command associated with the cron expression.

## Features

- Parses standard cron expressions
- Expands fields for minutes, hours, day of the month, month, and day of the week
- Displays the parsed command

## Usage

The program expects a cron expression as a single command line argument.

## Example

Given the cron expression `*/15 0 1,15 * 1-5 /usr/bin/find`, the program will produce the following output:

```
minute 0 15 30 45 
hour 0 
day of month 1 15 
month 1 2 3 4 5 6 7 8 9 10 11 12 
day of week 1 2 3 4 5 
command /usr/bin/find
```

## Requirements

- Go 1.16 or later

## Setup

1. **Install Go:**

   Make sure you have Go installed on your machine. You can download it from [https://golang.org/dl/](https://golang.org/dl/).

3. **Build the Project:**

   Build the Go program using the following command:

   ```sh
   go build -o cron-parser cron_parser.go
   ```

4. **Run the Program:**

   Execute the compiled program with a cron expression as an argument:

   ```sh
   ./cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
   ```

## Project Structure

```
cron-parser/
├── cron_parser.go
└── README.md
```

## Code Overview

### `cron_parser.go`

This file contains the main logic for parsing and expanding the cron expression. It includes the following functions:

- `ParseCron`: Parses the cron expression string into a `Cron` struct.
- `expand`: Expands a cron field into a list of values.
- `printExpanded`: Prints the expanded field values.
- `main`: The main entry point of the program that reads the command line argument, parses it, and displays the expanded fields.

### Example Usage

To run the program, use the following command:

```sh
./cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```

This will output:

```
minute 0 15 30 45 
hour 0 
day of month 1 15 
month 1 2 3 4 5 6 7 8 9 10 11 12 
day of week 1 2 3 4 5 
command /usr/bin/find
```

## Contributing

If you want to contribute to this project, please fork the repository and submit a pull request with your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
