package scrape

import (
	"fmt"
	"net/http"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(stat statenums.StatEnum){
  url := string(stat)
  respChan, errChan := make(chan *http.Response), make(chan error)
  go func(){
    resp,err := http.Get(url)
    if err != nil{
        errChan <- err
        fmt.Printf("Error -> %v\n",err)
        return //do something
    }
    respChan <- resp
  }()

  select{
  case resp := <- respChan:
    defer resp.Body.Close()

    doc,err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil{
        fmt.Printf("Error -> %v\n",err)
        return //do something
    }
    tables := doc.Find("table")
    statsTable := tables.First()
    var headers = make(map[int]*Column)
    ths := statsTable.Find("thead").Find("tr").Find("th")
    bodyRows := statsTable.Find("tbody").Find("tr")

    ths.Each(func(i int, s *goquery.Selection) {
      if i==0{
        return 
      } //remove annoying Rk column set as th
      if headers[i-1] == nil{
        headers[i-1] = &Column{}
      }
      headers[i-1].SetColName(strings.ToLower(s.Text()))
    })

    var allRows []interface{}

    bodyRows.Each(func(i int, s *goquery.Selection){
      var values []interface{}
      s.Find("td").Each(func(j int, s *goquery.Selection) {
        values = append(values,s.Text())
      })
      allRows = append(allRows, values)
    })

  for _, row := range allRows{
      vals := row.([]interface{})
      for i,td := range vals{
        headers[i].AddValue(td)
      }

    }

  var headerNames []string
  var filteredCols []Column
  for _, h:= range headers{
      if !(len(h.ColName)>0){
        fmt.Println("Empty colname")
        continue
      }
      if utils.Contains(headerNames,h.ColName){
        intChar, err := strconv.Atoi(string(h.ColName[len(h.ColName)-1]))
        if err != nil{
          repeatedString := h.ColName[:len(h.ColName)-1] + strconv.Itoa(intChar + 1)
          fmt.Println(repeatedString)
          convertedString := ConvertPassingStatHeader(repeatedString)
          if !(len(convertedString) > 0 ){
            continue
          }
          h.SetColName(convertedString)
          headerNames = append(headerNames, h.ColName)
          filteredCols = append(filteredCols,*h)
            
        } else{
          convertedString := ConvertPassingStatHeader(h.ColName + strconv.Itoa(1))
          if !(len(convertedString) > 0 ){
            continue
          }
          h.SetColName(convertedString)
          headerNames = append(headerNames, h.ColName)
          filteredCols = append(filteredCols,*h)
        }

        }else{
        convertedString := ConvertPassingStatHeader(h.ColName)
          if !(len(convertedString) > 0){
            continue
          }

        h.SetColName(convertedString)
        headerNames = append(headerNames, h.ColName)
        filteredCols = append(filteredCols,*h)


      }



    }

      fmt.Println(filteredCols)

    



  case err := <-errChan:
    fmt.Println("Error:",err)

    }

}

