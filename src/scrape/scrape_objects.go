package scrape;

import(
  "fmt"
  "encoding/csv"
  "os"
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
  fmt.Println(header)
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
         row[j] = "NA"
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
