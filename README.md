# Pokémon CLI

A command-line Pokédex written in Go.
This project is based on the [Boot.dev](https://boot.dev) course *“Build a Pokedex in Go”*.
It lets you explore Pokémon locations, catch Pokémon, and inspect your collection — all from the terminal.

---

## ✨ Features

* 🌍 **Explore** locations from the Pokémon world using the [PokéAPI](https://pokeapi.co/)
* 🎒 **Catch** Pokémon with a dynamic catch chance formula:

  ```
  catchChance = 75 * exp(-0.004 * BaseExperience) %
  ```
* 📜 **List** Pokémon you’ve caught in your personal Pokédex
* 🔎 **Inspect** detailed stats of your Pokémon
* 🗺️ **Navigate** through paginated location data with `map` and `mapb`
* ⚡ **Caching**:

  * Results from `map` and `explore` are cached for a few minutes
  * This reduces redundant API calls and speeds up revisits
* ⏳ Session-based persistence — your caught Pokémon stay with you until you exit

---

## 🛠️ Installation

Clone the repository:

```bash
git clone https://github.com/sebasukodo/pokemon-cli.git
cd pokemon-cli
```

Build the binary:

```bash
go build -o pokedex
```

Run the CLI:

```bash
./pokedex
```

---

## 📖 Usage

Start the CLI and type commands at the prompt:

```
help          Show all available commands
exit          Exit the Pokédex
map           Display the next 20 locations
mapb          Display the previous 20 locations
explore NAME  Explore a location (by name or ID)
catch NAME    Try to catch a Pokémon
inspect NAME  Display stats for one of your caught Pokémon
pokedex       List all caught Pokémon
```

---

## 🚀 Planned Improvements

* Improve the `map` and `explore` commands (better formatting, smarter pagination)
* Expand available commands (release Pokémon, save/load Pokédex to disk, etc.)
* Add richer Pokémon stats and battle simulation

---

## 📝 Credits

* Built while following the [Boot.dev](https://boot.dev) course *“Build a Pokedex in Go”*
* Pokémon data provided by the [PokéAPI](https://pokeapi.co/)