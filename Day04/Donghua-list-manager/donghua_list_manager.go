package main

import "fmt"

type Donghua struct {
	title    string
	episodes int
	rating   float64
	status   string
}

var DonghuaList []Donghua

func main() {
	fmt.Println("==============================================")
	fmt.Println("          Donghua List Manager V1.0           ")
	fmt.Println("==============================================")
	fmt.Println()

	addSampleData()

	for {
		showMenu()

		var choice int
		fmt.Print("Your choice (1-6):")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addDonghua()
		case 2:
			viewAllDonghua()
		case 3:
			searchDonghua()
		case 4:
			updateRating()
		case 5:
			fmt.Println("Thank you for using Donghua Manager")
			return
		default:
			fmt.Println("Invalid choice. Please enter 1-6")
		}
		fmt.Println()
	}
}

func addSampleData() {
	DonghuaList = []Donghua{
		{title: "The King's Avatar", episodes: 12, rating: 8.5, status: "Completed"},
		{title: "Link Click", episodes: 11, rating: 9.0, status: "Watching"},
		{title: "Scissor Seven", episodes: 30, rating: 8.7, status: "Completed"},
	}
}

func showMenu() {
	fmt.Println("Main Menu")
	fmt.Println("1. Add Donghua")
	fmt.Println("2. View All")
	fmt.Println("3. Search")
	fmt.Println("4. Update Rating")
	fmt.Println("5. Exit")
	fmt.Println()
}

func addDonghua() {
	var d Donghua

	fmt.Println("  Add New Domghua  ")
	fmt.Println("-------------------")

	fmt.Print("Tile:")
	fmt.Scan(&d.title)

	fmt.Print("Episodes:")
	fmt.Scan(&d.episodes)

	if d.episodes < 1 {
		fmt.Println("Episodes must be at least 1")
		return
	}

	fmt.Print("Rating (0-10):")
	fmt.Scan(&d.rating)

	if d.rating < 0 || d.rating > 10 {
		fmt.Println("Rating must be between 0 and 10")
		return
	}

	fmt.Println("Status:")
	fmt.Println("1. Watching")
	fmt.Println("2. Complete")
	fmt.Println("3. Plan to watch")

	var statusChoice int
	fmt.Print("Your choice (1-3)")
	fmt.Scan(&statusChoice)

	switch statusChoice {
	case 1:
		d.status = "Watching"
	case 2:
		d.status = "Complete"
	case 3:
		d.status = "Plan to wathch"
	default:
		fmt.Println("Invalid status, setting to 'Plan to watch'")
		d.status = "Plan to watch"
	}

	DonghuaList = append(DonghuaList, d)

	fmt.Println("Donghua added successfully")
	fmt.Printf("Add %s successfully", d.title)

}

func viewAllDonghua() {
	fmt.Println("  All Donghua  ")
	fmt.Println("---------------")

	if len(DonghuaList) == 0 {
		fmt.Println("No Donghua in the list yet")
		return
	}

	for index, d := range DonghuaList {
		fmt.Printf("\n#%d\n", index+1)
		fmt.Printf("Title:%s\n", d.title)
		fmt.Printf("Eppisodes:%d\n", d.episodes)
		fmt.Printf("Rating:%.1f\n", d.rating)
		fmt.Printf("Status:%s\n", d.status)
		fmt.Println("-------------------------")
	}

	fmt.Printf("\nTotal:%d donghua", len(DonghuaList))
}

func searchDonghua() {
	var searchTitle string

	fmt.Print("Enter title to search:")
	fmt.Scan(&searchTitle)
	fmt.Println()

	found := false

	for index, d := range DonghuaList {
		if d.title == searchTitle {
			found = true
			fmt.Printf("Found at positon #%d\n", index+1)
			fmt.Println("------------------------------")
			fmt.Printf("Title:%s\n", d.title)
			fmt.Printf("Episodes:%d\n", d.episodes)
			fmt.Printf("Rating:%.1f\n", d.rating)
			fmt.Printf("Status:%s\n", d.status)
			break
		}
	}

	if !found {
		fmt.Printf("'%s' not found in the list\n", searchTitle)
	}
}

func updateRating() {
	var searchTitle string
	var newRating float64

	fmt.Print("Enter title to update:")
	fmt.Scan(&searchTitle)

	found := false

	for i := range DonghuaList {
		if DonghuaList[i].title == searchTitle {
			found = true
			fmt.Printf("\nCurrent rating:%.1f\n", DonghuaList[i].rating)
			fmt.Print("Enter new rating:")
			fmt.Scan(&newRating)

			if newRating < 0 || newRating > 10 {
				fmt.Println("Rating must be between 0 and 10")
				return
			}

			DonghuaList[i].rating = newRating

			fmt.Printf("Update rating for '%s' to %.1f\n", searchTitle, newRating)
			break
		}
	}
	if !found {
		fmt.Printf("'%s' not found\n", searchTitle)
	}
}
