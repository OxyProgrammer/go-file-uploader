# go-file-uploader

---

```


                                         ██████╗   ██████╗      ███████╗ ██╗ ██╗      ███████╗     ██╗   ██╗ ██████╗  ██╗       ██████╗   █████╗  ██████╗  ███████╗ ██████╗
                                        ██╔════╝  ██╔═══██╗     ██╔════╝ ██║ ██║      ██╔════╝     ██║   ██║ ██╔══██╗ ██║      ██╔═══██╗ ██╔══██╗ ██╔══██╗ ██╔════╝ ██╔══██╗
                                        ██║  ███╗ ██║   ██║     █████╗   ██║ ██║      █████╗       ██║   ██║ ██████╔╝ ██║      ██║   ██║ ███████║ ██║  ██║ █████╗   ██████╔╝
                                        ██║   ██║ ██║   ██║     ██╔══╝   ██║ ██║      ██╔══╝       ██║   ██║ ██╔═══╝  ██║      ██║   ██║ ██╔══██║ ██║  ██║ ██╔══╝   ██╔══██╗
                                        ╚██████╔╝ ╚██████╔╝     ██║      ██║ ███████╗ ███████╗     ╚██████╔╝ ██║      ███████╗ ╚██████╔╝ ██║  ██║ ██████╔╝ ███████╗ ██║  ██║
                                         ╚═════╝   ╚═════╝      ╚═╝      ╚═╝ ╚══════╝ ╚══════╝      ╚═════╝  ╚═╝      ╚══════╝  ╚═════╝  ╚═╝  ╚═╝ ╚═════╝  ╚══════╝ ╚═╝  ╚═╝

```

Welcome to **go-file-uploader**, a sophisticated file processing API service crafted using Golang and powered by the GORM library. 🚀 This application reads massive datasets from CSV files, processes them through various concurrent methods, and uploads them into an SQLite database.

## Key Features

- **Efficient CSV Processing**: Handle CSV files with up to 10 million rows effortlessly.
- **Multiple Processing Strategies**: Implement solutions using a variety of techniques, including batching and multiprocessing.
- **SQLite Database Integration**: Seamlessly store data with GORM.
- **Performance Monitoring**: Measure memory usage and elapsed time for process optimization.

## Installation

1. **Prerequisites**: Ensure you have [Go 1.23.4](https://golang.org/doc/install) and [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/) installed.
2. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/go-file-uploader.git
   cd go-file-uploader
   ```

## Build the project

```
go build
```
## Run the application
```
./go-file-uploader
```

## API Endpoints

The application provides the following endpoints, each representing a different method of processing:

| Endpoint                       | Description                                                                                | Memory Usage (bytes) | Elapsed Time (ns)   |
|--------------------------------|--------------------------------------------------------------------------------------------|----------------------|---------------------|
| `GET /solution-one`            | Load all data and insert in batches. 📦                                                    | 3,296,896,352        | 18,968,041,000      |
| `GET /solution-two`            | Read line by line and insert in batches of 10,000. 📃                                      | 3,851,704            | 26,142,007,900      |
| `GET /solution-three`          | Use multiprocessing with a worker pool for reading and writing. 🛠️                        | 5,289,472            | 24,760,674,200      |
| `GET /solution-four`           | Pipeline approach on worker pool for reading, transforming, and writing. 🚀                | 4,515,240            | 23,637,820,200      |


## System Requirements

- **Processor**: AMD Ryzen 9 7900 12-Core Processor @ 3.70 GHz 🖥️
- **Memory**: 32 GB RAM 💾

## Contributing

Feel free to submit issues, fork the repository, and create pull requests. Contributions are always welcome and appreciated! 🎉

## License

This project is licensed under the MIT License. 📜
