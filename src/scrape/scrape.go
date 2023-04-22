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
    doubleHeader := statsTable.Find("thead").Find("tr").Last()
    ths := doubleHeader.Find("th")
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

  fmt.Println(headers[0].ColName)
  
  
  var headerNames []string
  var filteredCols []Column
  var cleanColumns []Column

  for _, h:= range headers{
      if !(len(h.ColName)>0){
        continue
      }
      if utils.Contains(headerNames,h.ColName){
        intChar, err := strconv.Atoi(string(h.ColName[len(h.ColName)-1]))
        if err != nil{
          var repeatedString string;
          for i:=0 ;i<10;i++{
            repeatedString = h.ColName[:len(h.ColName)] + strconv.Itoa(intChar + i+1)
            if utils.Contains(headerNames,repeatedString){
              continue
            }
            break
          }
          h.SetColName(repeatedString)
          headerNames = append(headerNames, h.ColName)
          filteredCols = append(filteredCols,*h)
            
        } else{
          numberAtEnd,err := strconv.Atoi(string(h.ColName[len(h.ColName)]))
            if err != nil{
                fmt.Printf("Error parsing int -> %v\n",err)
                return //do something
          }
          convertedString := h.ColName[:len(h.ColName)] + strconv.Itoa(1 + numberAtEnd)
          if !(len(convertedString) > 0 ){
            continue
          }
          h.SetColName(convertedString)
          headerNames = append(headerNames, h.ColName)
          filteredCols = append(filteredCols,*h)
        }

        }else{
        headerNames = append(headerNames, h.ColName)
        filteredCols = append(filteredCols,*h)
      }



        }

      for _, h := range filteredCols{
        convertedString := ConvertPassingStatHeader(h.ColName)
        if len(convertedString) > 1{
          h.SetColName(convertedString)
          cleanColumns = append(cleanColumns,h)
      }

    }

    var csvErr error

    if csvErr = (&Column{}).ExportColumnsToCsv(cleanColumns,"data.csv"); csvErr != nil{
      fmt.Println("Error exporting csv: ",csvErr)
    }


  case err := <-errChan:
    fmt.Println("Error:",err)

    }

}

