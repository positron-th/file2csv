package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type GConf struct {
	artist      string
	title       string
	description string
	keywords    string
	path        []string
	name        []string
}

var Conf *GConf

func LoadIni(cur_path string) {
	file, err := os.Open(cur_path + "/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	Conf = &GConf{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := strings.SplitN(strings.TrimSpace(scanner.Text()), "=", 2)
		//fmt.Println(txt)
		if len(txt) == 2 {
			cmd := strings.TrimSpace(txt[0])
			val := strings.TrimSpace(txt[1])

			switch cmd {
			case "artist":
				if val != "" {
					Conf.artist = val
				}
			case "title":
				if val != "" {
					Conf.title = val
				}
			case "description":
				if val != "" {
					Conf.description = val
				}
			case "keywords":
				if val != "" {
					Conf.keywords = val
				}
			case "file_path":
				//fmt.Println("file_path", val)
				if val != "" {
					v := strings.Split(val, ",")
					for i := range v {
						vr := strings.ToLower(strings.TrimSpace(v[i]))
						if vr != "" {
							p := vr
							if p[0:2] == "./" {
								p = cur_path + vr[1:]
							}
							Conf.path = append(Conf.path, p)
						}
					}
				}
			case "file_name":
				//fmt.Println("file_name", val)
				if val != "" {
					v := strings.Split(val, ",")
					for i := range v {
						vr := strings.ToLower(strings.TrimSpace(v[i]))
						if vr != "" {
							Conf.name = append(Conf.name, vr)
						}
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
