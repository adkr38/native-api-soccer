package scrape;

import(
  "fmt"
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
