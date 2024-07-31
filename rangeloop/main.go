package main

type Product struct {
	Name     string
	Url      string
	Pictures []Picture
	Sections []Section
}
type Picture struct {
	Name     string
	SectCode string
	Image    string
}
type Section struct {
	Id      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func main() {
	product := Product{Sections: []Section{
		{Id: 1, Code: "sec1", Name: "sector 1", Content: "content decription", Image: ""},
		{Id: 2, Code: "sec2", Name: "sector 2", Content: "content decription 2", Image: ""},
	}}
	newSects := []Section{}
	newSects = append(newSects, product.Sections...)

	for i := range newSects {
		sect := &newSects[i] // Get a reference to the current section
		for _, b := range product.Pictures {
			if b.SectCode == sect.Code {
				sect.Image = b.Image
			}
		}
	}
	product.Sections = newSects
}
