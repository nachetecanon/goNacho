package main
import "fmt"

func location(city string) (string, string) {
	var region,continent string
	switch city {
	case "Los Angeles", "LA", "Santa Monica":
		region,continent = "California", "America"
	case "Madrid", "Aranjuez":
		region,continent = "Madrid", "Europa"
	default:
		region,continent = "Unknown", "Unknown"
	}
	return region,continent
}

func main() {
	fmt.Println("hola mundo!")
	region,continent:=location("Madrid")
	fmt.Printf("Madrid data: %s, %s\n",region,continent);
}

