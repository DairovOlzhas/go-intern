package main



import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dairovolzhas/go-intern/lesson6/book_store"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath = ""
	config = book_store.Config{}
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "path",
			Aliases: []string{"p"},
			Usage:   "-path [--p] Path to json book_store",
			//Required:    true,
			Destination: &configPath,
		},
	}
)

func main(){
	app := &cli.App{
		Flags:flags,
		Name: "greet",
		Usage: "dairov olzhas",
		Action: run,
	}


	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if err := startServer(); err != nil {
		return err
	}
	return nil
}

func startServer() error{


	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	bookStore, err := book_store.CreateBookStore(config.PathToBookStore)
	if err != nil {
		return err
	}
	endpoints, err := book_store.CreateEndPointFactory(bookStore)
	if err != nil {
		return err
	}

	router.Methods("GET").Path("/").HandlerFunc(endpoints.BooksListHandler())
	router.Methods("POST").Path("/").HandlerFunc(endpoints.BooksCreateHandler())
	router.Methods("GET").Path("/{id}").HandlerFunc(endpoints.BookGetHandler("id"))
	router.Methods("PUT").Path("/{id}").HandlerFunc(endpoints.BookUpdateHandler("id"))
	router.Methods("DELETE").Path("/{id}").HandlerFunc(endpoints.BookDeleteHandler("id"))
	router.Methods("POST").Path("/save").HandlerFunc(endpoints.SaveBookStoreHandler(config.PathToBookStore))

	fmt.Println("Server Started")

	go func() {
		http.ListenAndServe(":"+config.Port, router)
	}()

	c := make(chan os.Signal)
	d := make(chan bool)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		d <- true
	}()
	<-d
	err = bookStore.SaveBookStore(config.PathToBookStore)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("\r- Good bay!")
	}
	os.Exit(1)


	return nil
}

func SetupCloseHandler(bookStore *book_store.BookStoreClass) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		err := bookStore.SaveBookStore(config.PathToBookStore)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("\r- Good bay!")
		}
		os.Exit(1)
	}()
}

