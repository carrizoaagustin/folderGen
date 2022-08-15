package folder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type folder struct {
	Name       string   `json:"Name"`
	SubFolders []folder `json:"SubFolders"`
}

func CreateTemplateOfDirectory(pathDirectory string, nameTemplate string, createLocation string) {
	folderStruct := folder{
		Name:       "",
		SubFolders: nil,
	}
	createFolderStructOfDirectory(pathDirectory, &folderStruct)
	log.Println(folderStruct)
	jsonByte, err := json.Marshal(&folderStruct)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(fmt.Sprint(createLocation+"/"+nameTemplate+".json"), jsonByte, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

}

func CreateDirectoryOfCreationFile(pathCreationFile string) {
	jsonData, err := ioutil.ReadFile(pathCreationFile)
	if err != nil {
		fmt.Println(err)
	}
	var creationFile CreationFile
	err = json.Unmarshal(jsonData, &creationFile)
	if err != nil {
		fmt.Println(err)
	}

	for _, RootPath := range creationFile.RootPaths {
		os.MkdirAll(RootPath, os.ModePerm)
		createDirectoryOfTemplate(creationFile.TemplatePath, RootPath)
	}
}

//Aux Functions
func getFoldersOfDirectory(pathDirectory string) []string {
	var folders []string
	fileInfo, err := ioutil.ReadDir(pathDirectory)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range fileInfo {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders
}

func createDirectoryOfTreeFolder(treeFolder folder, createLocation string) {
	if createLocation != "" {
		treeFolder.Name = fmt.Sprint(createLocation, "/", treeFolder.Name)
	}
	err := os.Mkdir(treeFolder.Name, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		fmt.Println(err)
	}
	if treeFolder.SubFolders != nil {
		for _, oneSubFolder := range treeFolder.SubFolders {
			oneSubFolder.Name = fmt.Sprint(treeFolder.Name, "/", oneSubFolder.Name)
			createDirectoryOfTreeFolder(oneSubFolder, "")
		}
	}
}

func createFolderStructOfDirectory(pathDirectory string, root *folder) {
	folderList := getFoldersOfDirectory(pathDirectory)
	for _, folderExtracted := range folderList {
		newPathDirection := pathDirectory + "/" + folderExtracted
		newFolder := folder{Name: folderExtracted}
		createFolderStructOfDirectory(newPathDirection, &newFolder)
		root.SubFolders = append(root.SubFolders, newFolder)
	}
}

func createDirectoryOfTemplate(pathTemplate string, createLocation string) {
	jsonData, err := ioutil.ReadFile(pathTemplate)
	if err != nil {
		fmt.Println(err)
	}
	var treeFolder folder
	err = json.Unmarshal(jsonData, &treeFolder)
	if err != nil {
		fmt.Println(err)
	}
	createDirectoryOfTreeFolder(treeFolder, createLocation)
}
