// How it's works? We have fileCounter variable, which contains allFiles map that contains found files from folder.
// Doubles map, which contains only duplicates files. Also contains RWMutex for sync access to allFiles map from several
// gourutines.
// If at start it settled -delete flag, all found duplicate-files will be deleted after precision.
// If folder isn't exist will be thrown an exception.
package main

import (
	"flag"
	"fmt"
	"golang_2/ex8/double_finder/counter"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	fileCounter     = counter.NewCounters()
	deleteDuplicate = flag.Bool("delete", false, "Mark duplicate files in folder to delete")
)

func main() {
	flag.Parse()
	fmt.Println("Input folder where needed to find duplicates:")

	var folder string // = "C:\\Users\\nktnl\\test"
	fmt.Scanln(&folder)

	if !exists(folder) {
		log.Fatalf("Folder %s doesn't exist", folder)
	}
	// sync search all files in folder and subfolders
	listFilesWithWG(folder)
	fileCounter.WaitGroupWait()
	// show results
	for k, v := range fileCounter.FindAll() {
		fmt.Println(k, v)
	}

	foundDuplicates := fileCounter.FindDoubles()
	if len(foundDuplicates) > 0 {
		fmt.Println("-------Found duplicate-files:-------")
		for k := range foundDuplicates {
			fmt.Println(k)
		}
	}

	// delete duplicates if needed
	if *deleteDuplicate {
		var verifyDeleteInput string = "n"
		fmt.Println("Do you really want to delete these duplicate files(Y-yes, sure; N-no, leave them)?: ")
		fmt.Scanln(&verifyDeleteInput)

		if strings.ToUpper(verifyDeleteInput) == "Y" {
			for d := range fileCounter.GetDoubles() {
				e := os.Remove(d.FullPath)
				if e != nil {
					log.Fatal(e)
				}
			}
			fmt.Println("Duplicates are deleted!")
		}
		fmt.Println("Have a good day")
	}
}

// Method listFiles provides searching files inside folder, and saving them to allFiles
func listFiles(path string) {

	if string(path[len(path)-1:]) != "\\" {
		path = path + "\\"
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			listFilesWithWG(path + f.Name())
		} else {
			fileCounter.Add(
				counter.FullFileInfo{
					Name:     f.Name(),
					Size:     f.Size(),
					FullPath: path + f.Name(),
				},
			)
		}
	}
}

// Method listFilesWithWG allows to find files inside folder with Wait Group
func listFilesWithWG(folder string) {
	fileCounter.WaitGroupAdd(1)
	go func() {
		listFiles(folder)
		fileCounter.WaitGroupDone()
	}()
}

// Method exists checks is folder existing or not
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
