package handler

import (
	"github.com/mono83/xray"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
	"github.com/mono83/xray/args"
)

type MakeDecision struct {
	Ray xray.Ray
}

func (h MakeDecision) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Ray.Info("Make decision")

	bts, err := ioutil.ReadAll(r.Body)

	if err != nil {
		sendResponse(w, err)
		return
	}

	var req makeDecisionRequest
	if err := json.Unmarshal(bts, &req); err != nil {
		sendResponse(w, err)
		return
	}

	var accept, reject, manual bool
	switch req.EnforcedDecision {
	case "reject":
		reject = true
		break
	case "manual":
		manual = true
		break
	default:
		accept = true
	}

	score := 90
	if len(req.EnforcedScore) > 0 {
		h.Ray.Info("Enforced score :score", args.String{N: "score", V: req.EnforcedScore})
		if i, err := strconv.Atoi(req.EnforcedScore); err == nil {
			score = i
		} else {
			sendResponse(w, err)
			return
		}
	}
	reason := ""
	if len(req.EnforcedReason) > 0 {
		h.Ray.Info("Enforced reason :val", args.String{N: "val", V: req.EnforcedReason})
		reason = req.EnforcedReason
	}

	sendResponse(w, makeDecisionResponse{
		RequestID: time.Now().Second(),
		Score:     score,
		Reject:    reject,
		Manual:    manual,
		Accept:    accept,
		Action:    "makeDecision",
		Reason:    reason,
	})
}

type makeDecisionRequest struct {
	Type             string `json:"type"`
	PaymentMethod    string `json:"payment_method"`
	TransactionID    string `json:"transaction_id"`
	PaymentAccountID string `json:"payment_account_id"`
	EnforcedDecision string `json:"custom_fraud_decision"`
	EnforcedReason   string `json:"custom_fraud_reject_reason"`
	EnforcedScore    string `json:"custom_fraud_score"`
}

type makeDecisionResponse struct {
	RequestID int    `json:"requestId"`
	Score     int    `json:"score"`
	Reject    bool   `json:"reject"`
	Manual    bool   `json:"manual"`
	Accept    bool   `json:"accept"`
	Action    string `json:"action"`
	Reason    string `json:"reason"`
}
