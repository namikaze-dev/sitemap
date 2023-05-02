# Sitemap Generator

This is a sitemap generator built with Go and designed to create sitemaps for websites. It has the ability to parse links from HTML and supports specifying the max depth the sitemap builder will search to. It was developed using test-driven development and has good test coverage.

## Installation

To install and run the sitemap generator:
Requires Go to be installed.

1. Clone this repository to your local machine
3. Run `go build .` to build the project
4. Run the executable file generated

## Usage

The sitemap generator can be used with the following command-line arguments

- `-url`: The URL of the website to generate the sitemap for
- `-depth`: The maximum depth the sitemap builder will search to

Example usage:

```
./sitemap-generator -url https://example.com -depth 3
```

## Contributing

If you would like to contribute to the sitemap generator, please:

1. Fork this repository
2. Create a new branch
3. Make your changes and commit them
4. Submit a pull request

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
