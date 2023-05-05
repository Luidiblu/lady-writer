# Lady Writer

[![Go Report Card](https://goreportcard.com/badge/github.com/luidiblu/lady-writer)](https://goreportcard.com/report/github.com/luidiblu/lady-writer)
[![License: MIT](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)

Lady Writer is a command-line interface (CLI) tool that leverages the power of OpenAI's GPT-3 to execute code snippets. It is inspired by the song "Lady Writer" by Dire Straits.

## Features

- Execute code snippets using GPT-3 by OpenAI
- Easy-to-use command-line interface
- Supports multiple programming languages (Golang for now)

## Installation

1. Install [Go](https://golang.org/dl/) if you haven't already.
2. Clone the repository:

```bash
git clone https://github.com/luidiblu/lady-writer.git
```

3. Build the project:

```bash
cd ladyriter
go build
```

4. Set the environment variable `OPENAI_API_KEY` with your OpenAI API key:

```bash
export OPENAI_API_KEY=your_openai_api_key
```

## Usage

```bash
./lady-writer "<code_snippet>"
```

Replace `<code_snippet>` with the code you want to execute.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://opensource.org/licenses/MIT)