# EC-blog

Simple blog posts visor where you can see and search all the available posts.

## Getting Started

Too run the server go to the file and execute the EC-blog file.

### Prerequisites

You should provide a Mysql database with:

```
CREATE TABLE IF NOT EXISTS posts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title  VARCHAR(50) NOT NULL,
    body TEXT,
    created IMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) 
```

Also provide through env variables:
* DB_NAME Name of the database
* DB_USER User for the dabase
* DB_PASSWORD Password for the user

In order to contribute you need to download the package [Gorilla Mux](https://github.com/gorilla/mux) and the driver for [Mysql](https://github.com/go-sql-driver/mysql)

## Running the tests

To run the test go to the project folder and run 
```
## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
