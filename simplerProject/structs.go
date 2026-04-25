package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
)

type Product struct {
	Name  string
	Price float64
}

func ProductTotalValue() float64 {
	fmt.Println("enter a name and price of the product :")
	var pr Product
	fmt.Scan(&pr.Name, &pr.Price)
	fmt.Println("how much Quantity you want ?")
	var qty int
	fmt.Scan(&qty)
	total := pr.Price * float64(qty)
	return total
}

func GetProductTotalValues() {
	total := ProductTotalValue()
	if total > 2000 {
		fmt.Println("Congratulations! You are eligible for a discount.")
		fmt.Println("Discount Applied : 10%", total*0.1)
		total = total * 0.9
	}
	fmt.Println("total :", total)

}

type Profile struct {
	Name    string
	Age     int
	Address string
}

func isAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
func birthDay(user *Profile) {
	user.Age++
}

func charelocateUser(user *Profile, newCity string) {
	user.Address = newCity
}

func GetProfile() {
	fmt.Println("enter your name, age, and city :")
	var pr Profile
	fmt.Scan(&pr.Name, &pr.Age, &pr.Address)
	adult := isAdult(pr.Age)
	if adult {
		fmt.Printf("Welcome, %s! You are an adult.\n", pr.Name)
	}

	fmt.Println("change city ?")
	var verify string
	fmt.Scan(&verify)
	if verify == "yes" || verify == "y" {
		fmt.Println("enter new city name :")
		var newCity string
		fmt.Scan(&newCity)
		charelocateUser(&pr, newCity)
	}
	birthDay(&pr)
	fmt.Printf("Profile updated: Name: %s, Age: %d, Address: %s\n", pr.Name, pr.Age, pr.Address)

}

//	{
//	    "id": 4,
//	    "name": "Patricia Lebsack",
//	    "username": "Karianne",
//	    "email": "Julianne.OConner@kory.org",
//	    "address": {
//	      "street": "Hoeger Mall",
//	      "suite": "Apt. 692",
//	      "city": "South Elvis",
//	      "zipcode": "53919-4257",
//	      "geo": {
//	        "lat": "29.4572",
//	        "lng": "-164.2990"
//	      }
//	    },
//	    "phone": "493-170-9623 x156",
//	    "website": "kale.biz",
//	    "company": {
//	      "name": "Robel-Corkery",
//	      "catchPhrase": "Multi-tiered zero tolerance productivity",
//	      "bs": "transition cutting-edge web services"
//	    }
//	  },
type ApiUsers struct {
	ID       int    `"json:"id"`
	Name     string `"json:"name"`
	Username string `"json:"username"`
	Email    string `"json:"email"`
	Address  struct {
		Street  string `"json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func ApiUserProcess() {
	rep, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Println("something went wrongs ")
	}
	body, err := io.ReadAll(rep.Body)
	defer rep.Body.Close()
	if err != nil {
		fmt.Println("something went wrongs ")
	}
	users := []ApiUsers{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("something went wrongs ")
	}

	fmt.Println("Select Feature")
	fmt.Println("1. print all users name and email")
	fmt.Println("2. Search user by name")
	fmt.Println("3. Filter by Email Domain")
	fmt.Println("4. Domain count for each domain")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		for _, user := range users {
			fmt.Println("name : ", user.Name, " email : ", user.Email)
		}
	case 2:
		fmt.Println("Search user by name :")
		var searchName string
		fmt.Scan(&searchName)
		var found bool
		for _, user := range users {
			if user.Name == searchName {
				fmt.Println("name : ", user.Name, " email : ", user.Email)
				found = true
			}
		}
		if !found {
			fmt.Println("user not found")
		}
	case 3:
		fmt.Println("Filter by Email Domain :")
		var domain string
		fmt.Scan(&domain)
		var found bool
		for _, user := range users {
			if strings.Contains(user.Email, domain) {
				fmt.Println("name : ", user.Name, " email : ", user.Email)
				found = true
			}
		}
		if !found {
			fmt.Println("no user found with this domain")
		}
	case 4:
		domainCount := make(map[string]int)

		for _, user := range users {

			_, fullDomain, found := strings.Cut(user.Email, "@")
			if found {
				domain := strings.Split(fullDomain, ".")
				if len(domain) > 0 {
					domainCount[domain[0]]++
				}
			}
		}
		for domain, count := range domainCount {
			fmt.Printf("Domain: %s , Count : %d\n", domain, count)
		}
		sort.Slice(users, func(i, j int) bool {
			return users[i].ID > users[j].ID
		})
		for _, user := range users {
			fmt.Printf("ID : %d , Name : %s\n", user.ID, user.Name)
		}

	default:
		fmt.Println("invalid choice")
	}

}
