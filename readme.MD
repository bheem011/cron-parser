---

# Cron Parser

Cron Parser is a Go application that parses a cron expression and prints the schedule for cron job.

## Installation

 you have Go installed on your system. You can install it from the official [Go website](https://golang.org/doc/install).

## Usage

1. Clone the repository to your local machine or download zip file:
a. For clonning : 
   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory:

   ```bash
   cd cron-parser
   ```

3. Build the executable file:

   ```bash
   go build
   ```

4. Run the program with a cron expression as a command-line argument:

   ```bash
   ./cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
   ```

   Replace the cron expression within quotes with your desired expression.

## Cron Expression Format

The cron expression should follow the standard format:

```
<minute> <hour> <day of month> <month> <day of week> <command>
```

- `<minute>`: 0-59 or */n (increment)
- `<hour>`: 0-23 or */n (increment)
- `<day of month>`: 1-31, */n (increment), or comma-separated list
- `<month>`: 1-12 or */n (increment)
- `<day of week>`: 0-7 (both 0 and 7 represent Sunday) or */n (increment)
- `<command>`: any command

## Example

```
./cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```

This will parse the given cron expression and print the parsed information.

---
