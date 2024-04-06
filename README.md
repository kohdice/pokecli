# pokecli

Command-line tool for Pokédex.

## 1. Table of Contents

- [1. Table of Contents](#1-table-of-contents)
- [2. About This Repository](#2-about-this-repository)
- [3. Usage](#3-usage)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)

## 2. About This Repository

This is a repository for a command-line tool that
retrieves and displays Pokémon data from an API.  
The API used is the [pokeapi-graphql](https://github.com/kohdice/pokeapi-graphql).

## 3. Usage

### Prerequisites

Set up [pokeapi-graphql](https://github.com/kohdice/pokeapi-graphql) following
[this section](https://github.com/kohdice/pokeapi-graphql?tab=readme-ov-file#3-usage)
of `README.md`.

### Installation

Download binary from Release page
or install with go install command.

```bash
go install github.com/kohdice/pokecli@latest
```

## 4. Demo

### Get Pokémon data by National Pokédex number

```bash
pokecli number 25
```

### Get Pokémon data by Name

```bash
pokecli name ピカチュウ
```
