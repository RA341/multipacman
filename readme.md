# Multiplayer Pacman

A real-time multiplayer Pacman game built using Flutter, Phaser game engine, and Go.

🎮 **[Play Now](https://multipacman.radn.dev/)** 🎮

<div align="center">
  <img src="img/lobby.png" alt="Game Lobby" width="400" />
  <img src="img/game.png" alt="Gameplay Screenshot" width="400" />
</div>

## Stack

- **Flutter**: Cross-platform UI framework for the game client
- **Phaser**: game engine for javascript
- **Go**: backend server handling game logic and state management

## Selfhost

Ensure Docker is installed.

```bash
docker run -p 11200:11200 ghcr.io/ra341/multipacman:main
```

Then open your browser and navigate to `http://localhost:11200`

## Build

### Prerequisites

- Flutter SDK and Go installed locally

```bash
# Clone the repository
git clone https://github.com/yourusername/multipacman.git
cd multipacman/server

# Run the Go server
go run core/main.go
```

#### Frontend

```bash
cd client

# Get Flutter dependencies
flutter pub get

# Run in development mode
flutter run -d chrome  # For web
# Or
flutter run  # For mobile devices
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Disclaimer

This project is a fan-made implementation of Pacman for educational and portfolio purposes only. Pacman and all related
characters, sounds, and assets are trademarks of Bandai Namco Entertainment (formerly Namco). This project is not
affiliated with, endorsed by, or connected to Bandai Namco Entertainment in any way.

All game mechanics, visual styles, and character designs inspired by the original Pacman are used under fair use for
educational purposes. No copyright infringement is intended.

