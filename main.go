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
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	// Demander chemin
	//	- Si pas de chemin rechercher dans dossier actuel
	fmt.Print("Path du dossier à scanner (si vide on prend le path de l'exe) : ")
	scanner.Scan()
	dirPath := scanner.Text()
	fmt.Print("Rechercher : ")
	scanner.Scan()
	inputSearch := scanner.Text()
	fmt.Print("Remplacer par : ")
	scanner.Scan()
	inputReplace := scanner.Text()
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

	fmt.Println("Appuyez sur Entrée pour quitter le programme...")
	fmt.Scanln()

}

func findAndReplaceString(dirPath, searchString, replaceString string) {
	files, errFiles := os.ReadDir(dirPath)
	countRename := 0
	dir, _ := os.Executable()
	exeName := filepath.Base(dir)
	//Vérification path correct
	if errFiles != nil {
		log.Fatal(errFiles)
	}
	for _, file := range files {
		if file.Name() != exeName {
			if strings.Contains(file.Name(), searchString) {
				newName := strings.ReplaceAll(file.Name(), searchString, replaceString)
				oldPath := filepath.Join(dirPath, file.Name())
				newPath := filepath.Join(dirPath, newName)
				errRename := os.Rename(oldPath, newPath)
				if errRename != nil {
					log.Fatal(errRename)
				}
				fmt.Println("**", file.Name(), "remoner en", newName)
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

// *********** fonction pour déterminer s'il s'agis d'une dossier ou non
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
