package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var restartChoice string
	optionText := `Options :
	- Recherche les dossiers et fichiers => 1 (default) 
	- Recherche seulement les dossiers => 2 
	- Recherche seulement les fichiers => 3 
	- Recherche les dossiers, les fichiers, les sous-dossiers et les sous-fichiers => 4 :`
	for {
		restartChoice = ""
		dirPath := askValue("Path du dossier à scanner (si vide on prend le path de l'exe) : ")
		inputSearch := askValue("Rechercher : ")
		inputReplace := askValue("Remplacer par : ")
		optionChoice := askValue(optionText)
		for checkOptionChoice(optionChoice) == false {
			optionChoice = askValue(optionText)
		}
		fmt.Println(optionChoice)
		if dirPath == "" {
			// _, dir, _, _ := runtime.Caller(0) //Mode débug
			dir, err := os.Executable() //Mode Prod
			if err != nil {
				fmt.Println("Erreur lors de la récupération du chemin de l'exécutable :", err)
				return
			}
			dirPath = filepath.Dir(dir)
		}
		fmt.Println("Path :", dirPath)
		findAndReplaceString(dirPath, inputSearch, inputReplace)

		fmt.Println(`Appuyez sur "y" pour relancer le programme ou Entrer pour le quitter`)
		fmt.Scanln(&restartChoice)
		if strings.ToLower(restartChoice) != "y" {
			break
		}
	}

}
func checkOptionChoice(optionChoice string) bool {
	switch optionChoice {
	case "1", "":
		fmt.Println("Option choisie : recherche les dossiers et fichiers.")
		return true
	case "2":
		fmt.Println("Option choisie : recherche seulement les dossiers.")
		return true
	case "3":
		fmt.Println("Option choisie : recherche seulement les fichiers.")
		return true
	case "4":
		fmt.Println("Option choisie : recherche les dossiers, les fichiers, les sous-dossiers et les sous-fichiers.")
		return true
	default:
		fmt.Println("Entrée incorrecte !!!")
		return false
	}
}
func askValue(question string) string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	fmt.Print(question)
	scanner.Scan()
	return scanner.Text()
}
func findAndReplaceString(dirPath, searchString, replaceString string) {
	files, errFiles := os.ReadDir(dirPath)
	countRename := 0
	//On récupère le nom de l'exe
	dir, _ := os.Executable()
	exeName := filepath.Base(dir)
	//Vérification path correct
	if errFiles != nil {
		log.Fatal(errFiles)
	}
	for _, file := range files {
		if file.Name() != exeName { //Sécuritée pour le pas renommer l'exe
			if strings.Contains(file.Name(), searchString) {
				newName := strings.ReplaceAll(file.Name(), searchString, replaceString)
				oldPath := filepath.Join(dirPath, file.Name())
				newPath := filepath.Join(dirPath, newName)
				errRename := os.Rename(oldPath, newPath)
				if errRename != nil {
					log.Fatal(errRename)
				}
				fmt.Println("**", file.Name(), " ==> ", newName)
				countRename++
			}
		}
	}
	if countRename > 1 {
		fmt.Println(countRename, "éléments ont été renommés.")
	} else if countRename == 1 {
		fmt.Println(countRename, "élément a été renommé.")
	} else {
		fmt.Println(countRename, "élément renommé.")
	}
}

// *********** fonction pour déterminer s'il s'agit d'un dossier ou non
//	************  Voir plus tard si utile
// func isDirectory(path string) bool {
// 	fileInfo, err := os.Stat(path)
// 	if err != nil {
// 		return false
// 	}

// 	return fileInfo.IsDir()
// }

//Test appel de la fonction isDirectory
// if isDirectory(filepath.Join(dirPath, file.Name())) {
// 	fmt.Printf("dossier \n")
// } else {
// 	fmt.Printf("pas dossier \n")
// }

// ***********
