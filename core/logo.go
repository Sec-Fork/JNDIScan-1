package core

import "fmt"

func PrintLogo(authors []interface{}) {
	template := "       ___   ______  _________                \n" +
		"      / / | / / __ \\/  _/ ___/_________ _____ \n" +
		" __  / /  |/ / / / // / \\__ \\/ ___/ __ `/ __ \\\n" +
		"/ /_/ / /|  / /_/ // / ___/ / /__/ /_/ / / / /\n" +
		"\\____/_/ |_/_____/___//____/\\___/\\__,_/_/ /_/ \n" +
		"    coded by %s"
	logo := fmt.Sprintf(template, authors...)
	fmt.Println(logo)
}
