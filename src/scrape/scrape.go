package scrape

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"soccer-go/src/statenums"
	"soccer-go/src/utils"
	"sort"
	"strconv"
	"strings"
    "time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// TEAMS StatEnum =  "TEAMS"
// COMPETITION_PARTICIPATION StatEnum =  "COMPETITION_PARTICIPATIONç"
// PLAYERS StatEnum = "PLAYERS"

func Scrape(stat statenums.StatEnum, f func(string) string){

    url := string(stat)
    // var url string

        defer func(){
        url = string(stat)
    }()

    switch(stat){
    case statenums.TEAMS:
        url = string(statenums.TEAM_OVERALL)
        break
    case statenums.COMPETITION_PARTICIPATION:
        url = "TODO"
            break
        case statenums.PLAYERS:
            url = string(statenums.PLAYER_OVERALL)
            break

        }

  fmt.Println("Checking on ",url)
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
        headers[i].AddValue(strings.ReplaceAll(td.(string),",",""))
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

	err = godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}


    mySqlConfig := mysql.Config{
        User:  os.Getenv("MYSQL_USER"),
        Passwd: os.Getenv("MYSQL_PASSWORD"),
        Net:   "tcp",
        Addr:  os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") ,
        DBName:os.Getenv("MYSQL_DBNAME"),
    }

    dsn := mySqlConfig.FormatDSN()
    db,err := sql.Open("mysql",dsn)
    if err != nil{
        fmt.Printf("Error formatting DSN-> %v\n",err)
  }
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s LIMIT 1",statenums.GetEnumName(stat)))
	if err != nil {
		fmt.Printf("%v",err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		fmt.Printf("%v",err)
	}

    var acceptedCols []Column
    for _,col := range cleanColumns{
        for _, okString := range columns{
            if col.ColName == okString{
                acceptedCols = append(acceptedCols, col)
                continue
            }
        }
    }

    var mySqlError error = (&Column{}).ToMySql(acceptedCols,stat,db)
    if mySqlError != nil{
        fmt.Printf("Error adding data -> %v\n",mySqlError)
        return
  }



  case err := <-errChan:
    fmt.Println("Error received on channel:",err)

    }

}

func PostScrapingMySqlSetup(){
    err :=  godotenv.Load()
    if err != nil{
        fmt.Println("Unable to read .env file!")
    }

    mySqlConfig := mysql.Config{
        User:  os.Getenv("MYSQL_USER"),
        Passwd: os.Getenv("MYSQL_PASSWORD"),
        Net:   "tcp",
        Addr:  os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") ,
        DBName:os.Getenv("MYSQL_DBNAME"),
        MultiStatements: true,
    }


    dsn := mySqlConfig.FormatDSN()
    db,err := sql.Open("mysql",dsn)
    if err != nil{
        fmt.Printf("Error formatting DSN-> %v\n",err)
  }

  defer db.Close()
  _, err = db.Exec(`
  UPDATE players p
  JOIN teams t ON p.team_name = t.team_name
  SET p.team_id = t.id;

  UPDATE player_overall ps
  JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
  SET ps.id = p.id,
  ps.team_id = p.team_id;

  UPDATE player_defensive ps
  JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
  SET ps.id = p.id,
  ps.team_id = p.team_id;

  UPDATE player_shooting ps
  JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
  SET ps.id = p.id,
  ps.team_id = p.team_id;

  UPDATE player_passing ps
  JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
  SET ps.id = p.id,
  ps.team_id = p.team_id;

  UPDATE player_possession ps
  JOIN players p ON ps.player_name = p.name AND ps.team_name = p.team_name
  SET ps.id = p.id,
  ps.team_id = p.team_id;

  ALTER TABLE player_overall
  DROP COLUMN team_name,
  DROP COLUMN player_name;

  ALTER TABLE player_defensive
  DROP COLUMN team_name,
  DROP COLUMN player_name;

  ALTER TABLE player_shooting
  DROP COLUMN team_name,
  DROP COLUMN player_name;

  ALTER TABLE player_passing
  DROP COLUMN team_name,
  DROP COLUMN player_name;

  ALTER TABLE player_possession
  DROP COLUMN team_name,
  DROP COLUMN player_name;

  ALTER TABLE players
  DROP COLUMN team_name;

  UPDATE team_overall tt
  JOIN teams t ON tt.team_name = t.team_name
  SET tt.id = t.id,
  tt.id = t.id;

  UPDATE team_defensive tt
  JOIN teams t ON tt.team_name = t.team_name
  SET tt.id = t.id,
  tt.id = t.id;

  UPDATE team_shooting tt
  JOIN teams t ON tt.team_name = t.team_name
  SET tt.id = t.id,
  tt.id = t.id;

  UPDATE team_possession tt
  JOIN teams t ON tt.team_name = t.team_name
  SET tt.id = t.id,
  tt.id = t.id;


  UPDATE team_passing tt
  JOIN teams t ON tt.team_name = t.team_name
  SET tt.id = t.id,
  tt.id = t.id;


  ALTER TABLE team_overall
  DROP COLUMN team_name;

  ALTER TABLE team_defensive
  DROP COLUMN team_name;

  ALTER TABLE team_shooting
  DROP COLUMN team_name;

  ALTER TABLE team_passing
  DROP COLUMN team_name;

  ALTER TABLE team_possession
  DROP COLUMN team_name;
  `)
  if err != nil{
      fmt.Println(err)
  }

}

func ScrapeMain(){

    utils.TimeFunc(Scrape,statenums.TEAMS,ConvertTeamHeader)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYERS,ConvertPlayerHeader)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYER_OVERALL,ConvertPlayerOverallStatHeader)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYER_PASSING,ConvertPassingStatHeaderPlayer)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYER_SHOOTING,ConvertShootingStatHeaderPlayer)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYER_POSSESSION,ConvertPossesionHeaderPlayer)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.PLAYER_DEFENSIVE,ConvertDefensiveStatHeaderPlayer)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.TEAM_OVERALL,ConvertTeamOverallStatHeader)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.TEAM_PASSING,ConvertPassingStatHeaderTeam)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.TEAM_SHOOTING,ConvertShootingStatHeaderTeam)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.TEAM_POSSESSION,ConvertPossesionHeaderTeam)
    time.Sleep(5*time.Second)

    utils.TimeFunc(Scrape,statenums.TEAM_DEFENSIVE,ConvertDefensiveStatHeaderTeam)
    time.Sleep(5*time.Second)

}

