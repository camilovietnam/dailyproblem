package handlers

import (
	"fmt"
	"github.com/dailyproblem/942-count-parentheses/responder"
	"github.com/dailyproblem/942-count-parentheses/rinkeby"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func MinRemove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		text := chi.URLParam(r, "string")
		contract, auth, err := rinkeby.LoadContract(r.Context(), r)
		if err != nil {
			fmt.Printf("[❌] load rinkeby contract: %v\n", err)
			responder.WriteJson(w, http.StatusInternalServerError, nil)
			return
		}

		res, err := contract.CountWrong(&bind.CallOpts{
			From: auth.From,
		}, text)
		if err != nil {
			fmt.Println("[❌] call contract CountWrong")
			responder.WriteJson(w, http.StatusInternalServerError, nil)
			return
		}

		responder.WriteJson(w, http.StatusOK, res)
	}
}
