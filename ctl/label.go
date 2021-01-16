package main

import (
	"bufio"
	// "encoding/json"
	"fmt"
	// m "github.com/halfrost/LeetCode-Go/ctl/models"
	// "github.com/halfrost/LeetCode-Go/ctl/util"
	"github.com/halfrost/LeetCode-Go/ctl/util"
	"github.com/spf13/cobra"
	"io"
	"os"
	"regexp"
	// "sort"
	// "strconv"
	"errors"
	"io/ioutil"
	"strings"
)

var (
	chapterOneFileOrder = []string{"_index", "Data_Structure", "Algorithm"}
	chapterTwoFileOrder = []string{"_index", "Array", "String", "Two_Pointers", "Linked_List", "Stack", "Tree", "Dynamic_Programming", "Backtracking", "Depth_First_Search", "Breadth_First_Search",
		"Binary_Search", "Math", "Hash_Table", "Sort", "Bit_Manipulation", "Union_Find", "Sliding_Window", "Segment_Tree", "Binary_Indexed_Tree"}
	chapterThreeFileOrder = []string{"_index", "Segment_Tree", "UnionFind", "LRUCache", "LFUCache"}
	preNextHeader         = "----------------------------------------------\n<div style=\"display: flex;justify-content: space-between;align-items: center;\">\n"
	preNextFotter         = "</div>"
	delLine               = "----------------------------------------------\n"
	delHeader             = "<div style=\"display: flex;justify-content: space-between;align-items: center;\">"
	delLabel              = "<[a-zA-Z]+.*?>([\\s\\S]*?)</[a-zA-Z]*?>"
	delFooter             = "</div>"

	//ErrNoFilename is thrown when the path the the file to tail was not given
	ErrNoFilename = errors.New("You must provide the path to a file in the \"-file\" flag.")

	//ErrInvalidLineCount is thrown when the user provided 0 (zero) as the value for number of lines to tail
	ErrInvalidLineCount = errors.New("You cannot tail zero lines.")
)

func newLabelCommand() *cobra.Command {
	mc := &cobra.Command{
		Use:   "label <subcommand>",
		Short: "Label related commands",
	}
	//mc.PersistentFlags().StringVar(&logicEndpoint, "endpoint", "localhost:5880", "endpoint of logic service")
	mc.AddCommand(
		newAddPreNext(),
		newDeletePreNext(),
	)
	return mc
}

func newAddPreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-pre-next",
		Short: "Add pre-next label",
		Run: func(cmd *cobra.Command, args []string) {
			addPreNext()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func newDeletePreNext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del-pre-next",
		Short: "Delete pre-next label",
		Run: func(cmd *cobra.Command, args []string) {
			delPreNext()
		},
	}
	// cmd.Flags().StringVar(&alias, "alias", "", "alias")
	// cmd.Flags().StringVar(&appId, "appid", "", "appid")
	return cmd
}

func addPreNext() {
	var (
		exist bool
		err   error
		res   []byte
	)
	// add Chpater one
	for index, v := range chapterOneFileOrder {
		tmp := ""
		if index == 0 {
			tmp = "\n\n" + delLine + fmt.Sprintf("<p align = \"right\"><a href=\"https://books.halfrost.com/leetcode/ChapterOne/%v/\">下一页➡️</a></p>\n", chapterOneFileOrder[index+1])
		} else if index == len(chapterOneFileOrder)-1 {
			tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterOne/%v/\">⬅️上一页</a></p>\n", chapterOneFileOrder[index-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterTwo/\">下一章➡️</a></p>\n") + preNextFotter
		} else if index == 1 {
			tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterOne/\">⬅️上一页</a></p>\n") + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterOne/%v/\">下一页➡️</a></p>\n", chapterOneFileOrder[index+1]) + preNextFotter
		} else {
			tmp = "\n\n" + preNextHeader + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterOne/%v/\">⬅️上一页</a></p>\n", chapterOneFileOrder[index-1]) + fmt.Sprintf("<p><a href=\"https://books.halfrost.com/leetcode/ChapterOne/%v/\">下一页➡️</a></p>\n", chapterOneFileOrder[index+1]) + preNextFotter
		}
		exist, err = needAdd(fmt.Sprintf("../website/content/ChapterOne/%v.md", v))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 当前没有上一页和下一页，才添加
		if !exist && err == nil {
			res, err = eofAdd(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), tmp)
			if err != nil {
				fmt.Println(err)
				return
			}
			util.WriteFile(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), res)
		}
	}
}

func eofAdd(filePath string, labelString string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				output = append(output, []byte(labelString)...)
				output = append(output, []byte("\n")...)
				return output, nil
			}
			return nil, err
		}
		output = append(output, line...)
		output = append(output, []byte("\n")...)
	}
}

func delPreNext() {
	// del Chpater one
	for index, v := range chapterOneFileOrder {
		lineNum := 5
		if index == 0 {
			lineNum = 2
		}
		exist, err := needAdd(fmt.Sprintf("../website/content/ChapterOne/%v.md", v))
		if err != nil {
			fmt.Println(err)
			return
		}
		// 存在才删除
		if exist && err == nil {
			removeLine(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), lineNum+1)
		}
		// res, err := eofDel(fmt.Sprintf("../website/content/ChapterOne/%v.md", v))
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		// util.WriteFile(fmt.Sprintf("../website/content/ChapterOne/%v.md", v), res)
	}
}

func needAdd(filePath string) (bool, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return false, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return false, nil
			}
			return false, err
		}
		if ok, _ := regexp.Match(delHeader, line); ok {
			return true, nil
		} else if ok, _ := regexp.Match(delLabel, line); ok {
			return true, nil
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func eofDel(filePath string) ([]byte, error) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader, output := bufio.NewReader(f), []byte{}
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				return output, nil
			}
			return nil, err
		}
		if ok, _ := regexp.Match(delLine, line); ok {
			reg := regexp.MustCompile(delLine)
			newByte := reg.ReplaceAll(line, []byte(" "))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delHeader, line); ok {
			reg := regexp.MustCompile(delHeader)
			newByte := reg.ReplaceAll(line, []byte(" "))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delLabel, line); ok {
			reg := regexp.MustCompile(delLabel)
			newByte := reg.ReplaceAll(line, []byte(" "))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else if ok, _ := regexp.Match(delFooter, line); ok {
			reg := regexp.MustCompile(delFooter)
			newByte := reg.ReplaceAll(line, []byte(" "))
			output = append(output, newByte...)
			output = append(output, []byte("\n")...)
		} else {
			output = append(output, line...)
			output = append(output, []byte("\n")...)
		}
	}
}

func removeLine(path string, lineNumber int) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	info, _ := os.Stat(path)
	mode := info.Mode()
	array := strings.Split(string(file), "\n")
	array = array[:len(array)-lineNumber-1]
	ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
	fmt.Println("remove line successful")
}
