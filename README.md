# Godex

A command-line Pokédex built in Go that lets you explore locations, encounter wild Pokémon, and catch 'em all — right from your terminal.

```
Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 -tentacool
 -tentacruel
 -magikarp
 -gyarados
 -shellos

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.
```

## Motivation

I built Godex to learn how to work with HTTP clients in Go and interact with public REST APIs. The [PokéAPI](https://pokeapi.co/) provided the perfect playground — it's free, well-documented, and let's be honest, Pokémon makes everything more fun.

Along the way, I implemented features like response caching with TTL-based expiration, pagination through API results, and a clean REPL interface. It turned out to be a great way to practice building a real CLI tool that actually *does* something.

## 🚀 Quick Start

### Clone and build

```bash
git clone https://github.com/pedroaguia8/Pokedex-cli.git
cd Pokedex-cli
go build -o godex
./godex
```

You'll be dropped into an interactive REPL where you can start exploring!

## 📖 Usage

### Available Commands

| Command | Description |
|---------|-------------|
| `help` | Display available commands |
| `map` | Show the next 20 location areas |
| `mapb` | Show the previous 20 location areas |
| `explore <area>` | List all Pokémon in a specific area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | View details of a caught Pokémon |
| `pokedex` | List all your caught Pokémon |
| `exit` | Exit the application |

### Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 -tentacool
 -magikarp
 -gyarados

Pokedex > catch gyarados
Throwing a Pokeball at gyarados...
gyarados escaped!

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
 -hp: 20
 -attack: 10
 -defense: 55
 -special-attack: 15
 -special-defense: 20
 -speed: 80
Types:
 - water

Pokedex > pokedex
 - magikarp
```

### How Catching Works

The catch probability is based on the Pokémon's base experience — weaker Pokémon like Magikarp are nearly guaranteed catches, while powerful Pokémon like Blissey have roughly a 1-in-6 chance of being caught.

### Caching

API responses are cached for 5 minutes to reduce redundant network calls and speed up repeated queries.

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/pedroaguia8/Pokedex-cli.git
cd Pokedex-cli
```

### Build the binary

```bash
go build -o godex
```

### Run the tests

```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.