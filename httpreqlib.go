package httpreqlib

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	//"io/ioutil"
	"io"
	"log"
)

func word_count(words []string) map[string]int { // (word is the key and frequency is the value)
	word_freq := make(map[string]int)
	for _, word := range words {
		_, ok := word_freq[word]
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}

type word_struct struct {
	freq int
	word string
}

type by_freq []word_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq > a[j].freq }

//func convert(data [][]byte) []string {
//dataint := int(data[0][0])
//	s := make([]string, len(data))
//	for row := range data {
//		s[row] = string(data[row])
//	}
//	return s
//}
// import "http"

type HTTPReq struct {
	URL           string
	sortDirection bool
	//out           string
}

func NewHTTP_Req(urlToUse string, sortDirectionToUSe bool) *HTTPReq {
	instance := &HTTPReq{
		URL:           urlToUse,
		sortDirection: sortDirectionToUSe,
	}
	return instance
}

//func (instance HTTPReq) GetUrl (urlToUse string) {
//	instance.URL = urlToUse
//}

//func (instance HTTPReq) GetSortDirection (sortDirectionToUSe bool) {
//	instance.sortDirection = sortDirectionToUSe
//}

func (instance *HTTPReq) MakeReq() string { //(urlToUse string, sortDirectionToUSe bool) {

	//instance.URL := urlToUse
	//instance.sortDirection := sortDirectionToUSe
	var out string
	var client http.Client
	resp, err := client.Get(instance.URL) //("http://192.168.8.8")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)

		bodyString = strings.ToLower(bodyString)

		//bodyRange := len(bodyString)
		//log.Print(bodyString)
		//log.Print(bodyRange)
		//log.Print(bodyString[0])
		//log.Print(bodyString[bodyRange-3])
		//input, _ := io.ReadAll(resp.Body)

		words, _ := regexp.Compile("[A-Za-z]+")

		//fmt.Println("total words: ", len(words.FindAll(bodyBytes, -1)))

		tempArr := words.FindAllString(bodyString, -1)
		word_map := word_count(tempArr)

		pws := new(word_struct)
		struct_slice := make([]word_struct, len(word_map))
		ix := 0
		for key, val := range word_map {
			pws.freq = val
			pws.word = key
			struct_slice[ix] = *pws
			ix++
		}

		fmt.Println("Words in response body sorted by frequency:")
		// sorting slice of structers by field freq
		sort.Sort(by_freq(struct_slice))
		if instance.sortDirection {
			for ix := 0; ix < len(struct_slice); ix++ {
				//fmt.Printf("%v\n", struct_slice[ix])				// is it ok or we have to return smth
				out += strconv.Itoa(struct_slice[ix].freq)
				out += "\t"
				out += struct_slice[ix].word
				out += "\n"
			}
		} else {
			for ix := len(struct_slice) - 1; ix >= 0; ix-- {
				out += strconv.Itoa(struct_slice[ix].freq)
				out += "\t"
				out += struct_slice[ix].word
				out += "\n"

			}
			//delete*(pws)
		}

	}
	return out // ok, now we return smth
}
