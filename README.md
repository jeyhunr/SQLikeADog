# SQLikeADog

A simple MySQL database management tool built with Go and Fyne. View and manage your MySQL databases with a clean graphical interface.

## Features

- View all databases and tables
- Browse table contents
- Execute SQL queries
- Create and edit tables
- Secure credential management
- Cross-platform GUI

## Installation

### Prerequisites
- Go 1.16 or later
- MySQL Server

### Quick Start

1. Clone and build:
```
git clone https://github.com/jeyhunr/SQLikeADog.git
cd SQLikeADog
go build -o SQLikeADog-APP cmd/sqlikeadog/main.go
```

2. Run the application:
```bash
./SQLikeADog
```

3. On first run, enter your MySQL credentials:
   - Host (default: localhost)
   - Port (default: 3306)
   - Username
   - Password

## Usage

1. **Connect to MySQL**
   - Enter your credentials on first launch
   - Credentials are securely saved for future use

2. **Browse Databases**
   - Select a database from the left sidebar
   - View list of tables in the selected database

3. **View Table Data**
   - Click on a table to view its contents
   - Data is displayed in a scrollable grid

4. **Manage Tables**
   - Create new tables using the '+' button
   - Edit tables using the document icon
   - Execute SQL queries using the list icon

## Development

Build for different platforms:

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o SQLikeADog.exe cmd/sqlikeadog/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o SQLikeADog-mac cmd/sqlikeadog/main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o SQLikeADog-linux cmd/sqlikeadog/main.go
```

## Contact

For questions and support, contact: [rahimli.net](https://rahimli.net)

## License

MIT License