package main  

import (
	"bufio"
	"log"
	"fmt"
	kval "github.com/kval-access-language/kval-boltdb"
	"os"
	"strings"
	"strconv"
)
	
type book struct {
    title string
    author string
    quantity int
}

func checkerror(err error) bool {
	if err != nil {
		log.Println("FYI there is an error!", err)
		return true
	}
	return false
}

func inputFunction(inputString string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(inputString)
	text, err := reader.ReadString('\n')
	checkerror(err)
	text = strings.TrimSpace(text)
	return text
}

func addNewBook() book {

	var newBook book 

	newBook.title = inputFunction("Enter title of book:")

	newBook.author = inputFunction("Enter author of book:")

	// TODO: Can this go in a validation function of its own?
	q := inputFunction("Enter quantity of book:")

	q2, err := strconv.Atoi(q)
	if err != nil {
		// TODO: May want to handle explicitly...
		log.Println(err)
	}

	newBook.quantity = q2

	return newBook
} 

func insertValue(kb kval.Kvalboltdb, author string, key string, value string) {

	//Lets do a test insert...
	_, err := kval.Query(kb, "INS library >> books >> " + author + " >>>> " + key+ " :: " + value)
	if err != nil {
		//work with your error
	}
	// else: start working with you res struct
}

func storeBook(nb book) {
	kb, err := kval.Connect("library.bolt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening bolt database: %#v", err)
		os.Exit(1)
	}
	defer kval.Disconnect(kb)

	insertValue(kb, nb.author, "title", nb.title)
	insertValue(kb, nb.author, "quantity", string(nb.quantity)) 	// TODOL quanityy is going missing...
}

func main() {

	text := ""

	for (text != "quit") {

		switch text {
		case "new":
			nb := addNewBook()
			storeBook(nb)
		case "hello":
			fmt.Println("hello again friend what is your input?")
		}



		text = inputFunction("Input please?")



	}

}