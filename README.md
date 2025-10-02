# Yank

**Yank** is a simple, fast, and cross-platform package manager for Windows and Linux, designed to make installing, updating, and managing software easy and efficient.

## Features

- **Cross-Platform:** Works seamlessly on both Windows and Linux.
- **Fast & Lightweight:** Written in Go for speed and reliability.
- **Easy to Use:** Intuitive CLI for installing, updating, and removing packages.
- **Extensible:** Easily add new repositories or package sources.

## Installation

### Windows

Download the latest release from the [Releases](https://github.com/Woeter69/yank/releases) page and run the installer.

### Linux

Download the binary from the [Releases](https://github.com/Woeter69/yank/releases) page or use the provided install script:

```sh
curl -sSL https://github.com/Woeter69/yank/releases/latest/download/yank-linux-amd64 -o yank
chmod +x yank
sudo mv yank /usr/local/bin/
```

## Usage

```sh
# List available packages
yank list

# Install a package
yank install <package-name>

# Update a package
yank update <package-name>

# Remove a package
yank remove <package-name>
```

## Contributing

Contributions are welcome! Please open issues or submit pull requests for bug fixes, new features, or improvements.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

**Yank**: The package manager for everyone.
