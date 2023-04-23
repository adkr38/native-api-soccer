package scrape

import (
	"database/sql"
	"fmt"
	"net/http"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-sql-driver/mysql"
)

func Scrape(stat statenums.StatEnum, f func(string) string){
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

  
  var headerNames []string
  var filteredCols []Column
  var cleanColumns []Column
  keys := make([]int,0,len(headers))
  for k:= range headers{
      keys = append(keys, k)
    }

  sort.Ints(keys)

  for _, k:= range keys{
      if !(len(headers[k].ColName)>0){
        continue
      }
      if utils.Contains(headerNames,headers[k].ColName){
        intChar, err := strconv.Atoi(string(headers[k].ColName[len(headers[k].ColName)-1]))
        if err != nil{
          var repeatedString string;
          for i:=0 ;i<10;i++{
            repeatedString = headers[k].ColName[:len(headers[k].ColName)] + strconv.Itoa(intChar + i+1)
            if utils.Contains(headerNames,repeatedString){
              continue
            }
            break
          }
          headers[k].SetColName(repeatedString)
          headerNames = append(headerNames, headers[k].ColName)
          filteredCols = append(filteredCols,*headers[k])
            
        } else{
          numberAtEnd,err := strconv.Atoi(string(headers[k].ColName[len(headers[k].ColName)]))
            if err != nil{
                fmt.Printf("Error parsing int -> %v\n",err)
                return //do something
          }
          convertedString := headers[k].ColName[:len(headers[k].ColName)] + strconv.Itoa(1 + numberAtEnd)
          if !(len(convertedString) > 0 ){
            continue
          }
          headers[k].SetColName(convertedString)
          headerNames = append(headerNames, headers[k].ColName)
          filteredCols = append(filteredCols,*headers[k])
        }

        }else{
        headerNames = append(headerNames, headers[k].ColName)
        filteredCols = append(filteredCols,*headers[k])
      }



        }

      for _, h := range filteredCols{
        convertedString := f(h.ColName)
        if len(convertedString) > 1{
          h.SetColName(convertedString)
          cleanColumns = append(cleanColumns,h)
      }

    }


    mySqlConfig := mysql.Config{
        User:   "root",
        Passwd: "",
        Net:    "tcp",
        Addr:   "localhost",
        DBName: "SOCCER",
    }

    dsn := mySqlConfig.FormatDSN()
    db,err := sql.Open("mysql",dsn)
    if err != nil{
        fmt.Printf("Error formatting DSN-> %v\n",err)
  }

    var mySqlError error = (&Column{}).ToMySql(cleanColumns,stat,db)
    if mySqlError != nil{
        fmt.Printf("Error adding data -> %v\n",err)
        return //do something
  }
    

  case err := <-errChan:
    fmt.Println("Error:",err)

    }

}

