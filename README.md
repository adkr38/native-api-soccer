### Building a native API from web soccer data.

Goal is to serve data from different, player, team and statgroups.

Currently working on: **Api Setup**

mysql> SELECT <br>
    p.name, <br>
    po.age, <br>
    po.position, <br>
    CAST(pp.dribbles_succeeded / pp.dribbles AS decimal(16,2)) AS dribbling_success_rate <br>
    FROM player_possession pp <br>
    JOIN players p ON pp.id = p.id <br>
    JOIN player_overall po ON pp.id = po.id <br>
    WHERE pp.dribbles >= 30 <br>
    ORDER BY 4 DESC <br>
    LIMIT 10; <br>



| Name             | Age | Position | Dribbling Success Rate |
|------------------|-----|----------|-----------------------|
| Cheick Konaté    | 18  | DF       | 0.81                  |
| Thomas Partey    | 29  | MF       | 0.78                  |
| Ellyes Skhiri    | 27  | MF       | 0.75                  |
| Youri Tielemans  | 25  | MF       | 0.74                  |
| Pau Torres       | 25  | DF       | 0.74                  |
| Nemanja Matić    | 33  | MF       | 0.73                  |
| Nicolás González | 20  | MF       | 0.72                  |
| Niklas Schmidt   | 24  | MF       | 0.72                  |
| Aurélien Tchouaméni  | 22  | MF   | 0.72                  |
| Daniel Parejo    | 33  | MF       | 0.71                  |

<br>

10 rows in set (0.01 sec)

