package scrape

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"soccer-go/src/statenums"
	"strings"
)

type Column struct{
  ColName string
  Values []interface{}
}

func (c *Column) SetColName(colName string){
  c.ColName = colName
}

func (c *Column) AddValue(v interface{}){
  switch val := v.(type){
  case string:
    c.Values = append(c.Values,val)
  case int:
    c.Values = append(c.Values,val)
  case float64:
    c.Values = append(c.Values,val)
  default:
    fmt.Printf("Unsupported value type %T",val)

  }

}

func (c *Column) GetValues()[]interface{}{
  return c.Values
}

func (c *Column) ExportColumnsToCsv(columns []Column, filename string) error{
  if len(columns) == 0 {
    return fmt.Errorf("Empty column slice")
  }

  file,err := os.Create(filename)
  if err != nil{
      fmt.Printf("Error creating csv -> %v\n",err)
      return err
  }

  defer file.Close()

  writer := csv.NewWriter(file)

  header := make([]string, len(columns))
  for i, col:= range columns{
    header[i] = col.ColName
  }
  err = writer.Write(header)
  if err != nil{
      fmt.Printf("Error writing header -> %v\n",err)
      return err
  }

  var maxRows int = 0;
  for _,col := range columns{
    if len(col.Values) > maxRows{
      maxRows = len(col.Values)
    }

  }
  rows := make([][]string,maxRows)
  for i := 0; i< maxRows; i++{
    row := make([]string, len(columns))
    for j,col := range columns{
      if i >= len(col.Values){
         row[j] = "NULL"
      }else{
        row[j] = col.Values[i].(string)
    }

    }

    rows = append(rows, row)
    err = writer.Write(row)

  }

  writer.Flush()
  return nil

}


func(c *Column) ToMySql(columns []Column, statType statenums.StatEnum, db *sql.DB) error{
  var queryArgs []interface{}
  var valueHolders []string
  colNames := make([]string,0,len(columns))
  for _,c := range columns{
    colNames = append(colNames, c.ColName)
  }

  for i:= 0; i<len(columns); i++{
    valueHolders = append(valueHolders, "?")
  }
  var placeHolderStr string = strings.Join(valueHolders,",")

  var query string = fmt.Sprintf(`
    INSERT INTO %s (%s)
    VALUES (%s);
    `, statenums.GetEnumName(statType), strings.Join(colNames,","), placeHolderStr)

  statement, err := db.Prepare(query)
  if err != nil{
      fmt.Printf("Error -> %v\n",err)
      return err
  }
  defer statement.Close()

   for i := 0; i < len(columns[0].Values); i++ {
        queryArgs = nil // Reset queryArgs for each row

        for _, col := range columns {
            if i >= len(col.Values) {
                // queryArgs = append(queryArgs, nil) // Add null value for missing values
                queryArgs = append(queryArgs, nil) // Add null value for missing or empty values
            } else {
                if len(col.Values[i].(string)) == 0{
                    queryArgs = append(queryArgs, nil) // Add null value for missing or empty values
                    continue
                }
                queryArgs = append(queryArgs, strings.Trim(col.Values[i].(string)," \t\n"))
            }
        }

        fmt.Println(query)
        fmt.Println(queryArgs...)

        _, err = statement.Exec(queryArgs...)
        if err != nil {
            match,_:=  regexp.MatchString("1062",err.Error())
            if match{
                continue
            }
            fmt.Printf("Error adding value -> %v\n", err.Error() )
            return err
        }
    }
  return nil

}
