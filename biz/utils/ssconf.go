package utils

import (
	"bufio"
	"os"
	"path"
	"regexp"
	"strings"
	"unicode"
)

var resolvePattern = regexp.MustCompile("\\{\\{\\s*([^\\}]+)\\s*\\}\\}")

func getWholeLine(scanner *bufio.Scanner) (string, bool) {
	var (
		wholeLine string
		flag      bool
	)
	for scanner.Scan() {
		flag = true
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.HasSuffix(line, "\\") {
			line = line[:len(line)-1]
			wholeLine += line
			continue
		}
		wholeLine += line
		break
	}
	return wholeLine, flag
}

func isConfigSep(r rune) bool {
	if r == ':' || r == '=' || unicode.IsSpace(r) {
		return true
	}
	return false
}

func parseConfigKeyValue(line string) (key, value string) {
	line = strings.Split(strings.Split(line, "#")[0], ";")[0]
	line = strings.TrimSpace(line)
	kEnd := -1
	vBegin := -1
	for i, c := range line {
		isSep := isConfigSep(c)
		if kEnd == -1 {
			if isSep {
				kEnd = i
			}
			continue
		}
		if !isSep {
			vBegin = i
			break
		}
	}
	if vBegin == -1 {
		return "", ""
	}
	return line[:kEnd], line[vBegin:]
}

func resolveConf(conf map[string]string) {
	found := true
	for found {
		found = false
		for k, v := range conf {
			idx := resolvePattern.FindAllStringIndex(v, 1000)
			lastIdx := 0
			finalV := ""
			for i := range idx {
				st := idx[i][0]
				ed := idx[i][1]
				if st >= lastIdx {
					finalV += v[lastIdx:st]
					lastIdx = st
				}
				key := strings.Trim(v[st+2:ed-2], " ")
				if realV, ok := conf[key]; ok {
					finalV += realV
				} else {
					// todo
				}
				lastIdx = ed
			}
			if lastIdx > 0 {
				if lastIdx < len(v) {
					finalV += v[lastIdx:]
				}
				conf[k] = finalV
				found = true
			}
		}
	}
}

func LoadConfigFile(filename string) (ret map[string]string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ret = make(map[string]string)

	baseName := path.Dir(filename)
	scanner := bufio.NewScanner(file)
	for {
		if line, ok := getWholeLine(scanner); ok {
			if key, value := parseConfigKeyValue(line); key == "include" {
				includeFile := value
				if !strings.HasPrefix(includeFile, "/") {
					includeFile = baseName + "/" + includeFile
				}
				includeConf, err := LoadConfigFile(includeFile)
				if err != nil {
					continue
				}
				for k, v := range includeConf {
					ret[k] = v
				}
			} else if len(key) > 0 {
				ret[key] = value
			}
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	resolveConf(ret)
	return
}
