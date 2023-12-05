# EventAPI Project

## 🚀 About the Project
EventAPI is a Go-based web server that generates random characters for different themed events like Christmas, Halloween, and Easter. It's designed to be lightweight, efficient, and easy to use.

## 📁 Structure
- `main.go`      : Sets up the server and routes.
- `handlers.go`  : Contains handlers for different events to generate characters.
- `character.go` : Defines the `Character` struct and character generation logic.
- `utils.go`     : Includes utility functions like CORS middleware and file reading.
- `go.mod`       : Module file for Go dependencies.
- `assets`       : Directory containing themed assets for character generation.

## 🏃‍♂️ How to Run
1. Clone the repository.
2. Navigate to the project directory.
3. Run `go build` and `./eventapi` to start the server.
4. Access the server at `http://localhost:8081`.

## 🎨 Endpoints
- `/christmas/`: Generates Christmas themed characters. 🎄
- `/halloween/`: Generates Halloween themed characters. 🎃
- `/easter/`   : Generates Easter themed characters.    🐰

## 🛠️ Built With
- Go (version 1.21.4)

## 📚 Further Documentation
- For more details on Go: [Official Go Documentation](https://golang.org/doc/)

## 🤝 Contributing
Contributions, issues, and feature requests are welcome!

## 🌟 Show your support
Give a ⭐️ if this project helped you!

---

💡 Remember to check out the `assets` directory for each event to customize character data! 🎅👻🐇

---

## Few DALL·E Visualizations
- 🎁 <br> ![image](https://github.com/KaNiuSii/Go-EventAPI/assets/123270897/c90b4f80-cb81-4fd8-884e-61732ef2a49d) <br>
- 🦇 <br> ![image](https://github.com/KaNiuSii/Go-EventAPI/assets/123270897/cc3abca7-befc-48c6-9e22-eb6b7b09dd51) <br>
- 🐤 <br> ![image](https://github.com/KaNiuSii/Go-EventAPI/assets/123270897/2e2794d3-7e6d-40e5-823c-a4ed9d60e541) <br>
