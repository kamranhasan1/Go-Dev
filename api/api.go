package api
import
(
"encoding/json"
"net/http"
)

type CoinBalanceParms struct{
	Username string
}

type CoinBalanceResponse struct{
	//Success Code, Usually 150
	Code int

	//Account Balance
	Balance int64
}

type Error struct{
	//Error code
	Code int 

	//Error message
	Message string 
}