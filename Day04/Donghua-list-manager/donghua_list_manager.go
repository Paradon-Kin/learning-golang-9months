package main

import (
	"fmt"
)

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
		fmt.Print("Your choice (1-7):")
		fmt.Scan(&choice)
		fmt.Println()

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
			viewByStatus()
		case 6:
			deleteDonghua()
		case 7:
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
		{title: "Avatar", episodes: 12, rating: 8.5, status: "Completed"},
		{title: "Click", episodes: 11, rating: 9.0, status: "Watching"},
		{title: "Scissor", episodes: 30, rating: 8.7, status: "Completed"},
		{title: "Aventure", episodes: 30, rating: 8.7, status: "Plan to watch"},
	}
}

func showMenu() {
	fmt.Println("Main Menu")
	fmt.Println("1. Add Donghua")
	fmt.Println("2. View All")
	fmt.Println("3. Search")
	fmt.Println("4. Update Rating")
	fmt.Println("5. View by status")
	fmt.Println("6. Delete")
	fmt.Println("7. Exit")
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
	fmt.Println("2. Completed")
	fmt.Println("3. Plan to watch")

	var statusChoice int
	fmt.Print("Your choice (1-3)")
	fmt.Scan(&statusChoice)

	switch statusChoice {
	case 1:
		d.status = "Watching"
	case 2:
		d.status = "Completed"
	case 3:
		d.status = "Plan to watch"
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

func viewByStatus() {
	fmt.Println("View By status")
	fmt.Println("1. Watching")
	fmt.Println("2. Completed")
	fmt.Println("3. Plan to watch")
	fmt.Println("-----------------------------")
	fmt.Println()

	var choice int
	fmt.Print("Your choice (1-3):")
	fmt.Scan(&choice)
	fmt.Println()

	var searchStatus string

	switch choice {
	case 1:
		searchStatus = "Watching"
	case 2:
		searchStatus = "Completed"
	case 3:
		searchStatus = "Plan to watch"
	default:
		fmt.Println("Invalid choice")
		return
	}

	found := false
	for _, d := range DonghuaList {
		if d.status == searchStatus {
			found = true
			fmt.Printf("Donghua name:%s\n", d.title)
			fmt.Printf("Rating:%.1f\n", d.rating)
			fmt.Printf("Status:%s\n", d.status)
			fmt.Println("--------------------------------")
		}
	}

	if !found {
		fmt.Println("Not found")
	}
}

func deleteDonghua() {
	var searchTitle string
	fmt.Println("--- ลบข้อมูล Donghua ---")
	fmt.Print("กรอกชื่อเรื่องที่ต้องการลบ: ")
	fmt.Scan(&searchTitle)

	found := false

	for index, d := range DonghuaList {

		if d.title == searchTitle {
			found = true
			DonghuaList = append(DonghuaList[:index], DonghuaList[index+1:]...)
			fmt.Printf("Deleted '%s' successfully\n", d.title)
			break
		}
	}

	if !found {
		fmt.Printf("Not found '%s'\n", searchTitle)
	}
}
