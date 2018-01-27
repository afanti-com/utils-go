package mobile

import "fmt"

func ExampleSearch() {
	if err := Init("./data/mobile.dat"); err != nil {
		fmt.Println(err.Error())
		return
	}

	mi, err := Search("15010117638")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n%s\n",
		mi.No7, mi.Province, mi.City, mi.CityAreaCode, mi.CityZipCode, mi.Company)

	// Output:
	// 1501011
	// 北京
	// 北京
	// 010
	// 100000
	// 中国移动
}
