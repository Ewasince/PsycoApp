package environment

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"
)

type environmentVariables struct {
	PORT            string
	DEBUG           bool
	MIGRATIONS_PATH string
	DB_PATH         string
}

var Env environmentVariables = environmentVariables{}

func init() {
	var err error
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// PORT
	Env.PORT = getEnv("PORT", "8181")

	// DEBUG
	if Env.DEBUG, err = strconv.ParseBool(getEnv("DEBUG", "false")); err != nil {
		panic(err)
	}

	// MIGRATIONS_PATH
	Env.MIGRATIONS_PATH = "file://" + getEnv("MIGRATIONS_PATH", "")

	// DB_PATH
	Env.DB_PATH = getEnv("DATABASE_PATH", "")

	printEnvVariables()
}

func getEnv(key, defaultValue string) string {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}
	if defaultValue != "" {
		return defaultValue
	}
	panic("No value for " + key)
}

func printEnvVariables() {
	fmt.Println("Environment Variables:")

	// Получаем значение и тип структуры
	val := reflect.ValueOf(Env)
	fields := reflect.VisibleFields(reflect.TypeOf(Env))

	// Итерируемся по полям структуры
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, field := range fields {
		fieldValue := val.FieldByName(field.Name)
		fmt.Fprintf(w, "%s=\t%v\n", field.Name, fieldValue.Interface())
	}
	w.Flush()
}