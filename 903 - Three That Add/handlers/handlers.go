package handlers

import (
	"fmt"
	"github.com/camilovietnam/threethatadd/respond"
	"github.com/camilovietnam/threethatadd/rinkeby"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/go-chi/chi/v5"
	"math/big"
	"net/http"
	"strconv"
)

type HTTPResponse struct {
	Message string `json:"message"`
}

func AddItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder := respond.NewHTTPResponder(w)
		param := chi.URLParam(r, "item")

		item, err := strconv.Atoi(param)
		if err != nil {
			fmt.Printf("[❌] convert item to int: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		contract, auth, err := rinkeby.LoadContract(r)
		if err != nil {
			fmt.Printf("[❌] load rinkeby contract: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		_, err = contract.AddItem(&bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
		}, big.NewInt(int64(item)))
		if err != nil {
			fmt.Printf("[❌] call method to add item: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		responder.WriteJSON(http.StatusOK, HTTPResponse{
			Message: "The item was added, the transaction will take some time to appear in the network",
		})
	}
}

func GetAllItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder := respond.NewHTTPResponder(w)

		contract, auth, err := rinkeby.LoadContract(r)
		if err != nil {
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}
		defer r.Body.Close()

		items, err := contract.GetAllItems(&bind.CallOpts{
			From: auth.From,
		})
		if err != nil {
			fmt.Printf("[❌] call contract method getAllItems: %v", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		responder.WriteJSON(http.StatusOK, items)
	}
}

func FindThree() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder := respond.NewHTTPResponder(w)
		sum, err := strconv.Atoi(chi.URLParam(r, "sum"))
		if err != nil {
			fmt.Printf("[❌] convert sum to int: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		contract, auth, err := rinkeby.LoadContract(r)
		if err != nil {
			fmt.Printf("[❌] load contract: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		response, err := contract.ThreeThatAdd(&bind.CallOpts{
			From:    auth.From,
			Context: r.Context(),
		}, big.NewInt(int64(sum)))
		if err != nil {
			fmt.Printf("[❌] call method threeThatAdd: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		responder.WriteJSON(http.StatusOK, HTTPResponse{
			Message: response,
		})
	}
}

func DeleteAllItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responder := respond.NewHTTPResponder(w)
		contract, auth, err := rinkeby.LoadContract(r)
		if err != nil {
			fmt.Printf("[❌] load contract: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		_, err = contract.ResetItems(&bind.TransactOpts{
			From:   auth.From,
			Signer: auth.Signer,
		})
		if err != nil {
			fmt.Printf("[❌] call method reset items: %v\n", err)
			responder.WriteJSON(http.StatusInternalServerError, HTTPResponse{
				Message: "internal server error, try again later",
			})
			return
		}

		responder.WriteJSON(http.StatusOK, HTTPResponse{
			Message: "The list was reset, the transaction will take some time to appear in the network",
		})
	}
}
