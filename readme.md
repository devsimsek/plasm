# Plasm - A cli entrypoint for your applications.

Plasm is a entrypoint for you to build your own cli applications. It provides a simple interface to build your own commands and subcommands.
I would love to call this project as a "cli framework" but it's not a framework, it's just a simple entrypoint for your cli applications. Maybe in the future, I will add more features to call it a framework.

## Developing Using Plasm

To create a new command, you need to create a new file in the `commands` directory. After the creation, you should register your command in the init function.
For example, let's create a new command called `hello` located at `commands/hello.go`:

```go
package commands

import (
	"fmt"

	"github.com/devsimsek/plasm/core"
	"github.com/devsimsek/plasm/types"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "hello",
		Alias: []string{"h"},
		Flags: []types.Flag{
			{Name: "name", Alias: []string{"n"}, Value: ""},
		},
		Description: "Say hello to someone",
		Handler:     func(flags []types.Flag) {
			name := core.FindFlag(flags, "name")
			if name == "" {
				fmt.Println("Hello, World!")
			} else {
				fmt.Printf("Hello, %s!\n", name)
			}
		},
	})
}
```

Its that simple! Now you can run your command by running `go run . hello` or `go run . h`.

To create a type/model, you can create a new file in the `types` directory. For example, let's create a new type called `User` located at `types/user.go` that leverages Gorm:

```go
package types

import (
	"gorm.io/gorm"
)

// User struct to represent a user
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
}

func init() {
	// Register the User model as a migration, so that it can be migrated to the database
	Migrations = append(Migrations, Migration{
		Name:  "User",
		Model: &User{},
		Seed:  []User{
			// Seeding is this easy!
			{Name: "John", Surname: "Doe", Email: "john@does.com"},
		},
	})
}

// TableName sets the table name for the User model in the database
func (t *User) TableName() string {
	return "users"
}

// GetAll fetches all users from the database
func (t *User) GetAll() ([]User, error) {
	var users []User
	result := DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetById fetches a single user by ID
func (t *User) GetById(id int) (*User, error) {
	var user User
	result := DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// Create inserts a new user into the database
func (t *User) Create() error {
	result := DB.Create(t)
	return result.Error
}

// Update updates an existing user in the database
func (t *User) Update() error {
	result := DB.Save(t)
	return result.Error
}

// Delete removes a user from the database
func (t *User) Delete() error {
	result := DB.Delete(t)
	return result.Error
}

// Find fetches users based on a query
func (t *User) Find(query string) ([]User, error) {
	var users []User
	result := DB.Where("name LIKE ?", "%"+query+"%").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// Count fetches the total number of users
func (t *User) Count() (int64, error) {
	var count int64
	result := DB.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
```

Now you can use the User model in your commands. For example, let's create a new command called `users` located at `commands/users.go`:

```go
package commands

import (
	"fmt"

	"github.com/devsimsek/plasm/core"
	"github.com/devsimsek/plasm/types"
)

func init() {
	core.RegisterCommand(types.Command{
		Name:  "find",
		Alias: []string{"f"},
		Flags: []types.Flag{
			{Name: "query", Alias: []string{"q"}, Value: ""},
		},
		Description: "Find users by name",
		Handler:     func(flags []types.Flag) {
			query := core.FindFlag(flags, "query")
			if query == "" {
				fmt.Println("Please provide a query")
				return
			}

			user := types.User{}
			users, err := user.Find(query)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, u := range users {
				fmt.Printf("ID: %d, Name: %s, Surname: %s, Email: %s\n", u.ID, u.Name, u.Surname, u.Email)
			}
		},
	})
}
```

Now you can run your command by running `go run . users find --query=John`.

## Installation

To install Plasm, you can run the following command:

```bash
    git clone https://github.com/devsimsek/plasm.git myapp
    cd myapp
    go mod tidy
```

## Run Locally

To run the project locally, you can run the following command:

```bash
    go run .
```

## Features

- Simple interface to build your own commands and subcommands
- Easy to use
- Easy to extend
- Gorm support for database operations
- Easy to use configuration

## Feedback

If you have any feedback, please reach out to me using the issues tab.
Also, I don't have a lot of time to maintain this project, so if you want to contribute, feel free to open a pull request.

## License

[MIT](https://devsimsek.mit-license.org/)
or
license.md file in the root directory.

## Authors

- [@devsimsek](https://github.com/devsimsek)
