package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/infEreb/nls/netflow/server/ethernet/internal"
	"github.com/infEreb/nls/netflow/server/ethernet/pkg/ethernet"
	"github.com/infEreb/nls/netflow/server/pkg/serror"
)

type HandlerHTTP struct {
	ctrl *internal.Controller
}

func NewHandlerHTTP(ctrl *internal.Controller) *HandlerHTTP {
	return &HandlerHTTP{
		ctrl: ctrl,
	}
}

func (h *HandlerHTTP) GetEthernets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	eths, err := h.ctrl.FindAll(ctx)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eths); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}
func (h *HandlerHTTP) GetEthernet(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	eths, err := h.ctrl.FindByID(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &serror.ErrorNotFound{}) {
			code = http.StatusNotFound
		}
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eths); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}

func (h *HandlerHTTP) PostEthernet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	e := &ethernet.Ethernet{}
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if e.AgentID == "" || e.DevName == "" || e.Name == "" || e.Hardware == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("any of agent_id, dev_name, name, hardware is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	_, err := h.ctrl.GetAgent(ctx, e.AgentID)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &serror.ErrorNotFound{}) {
			code = http.StatusNotFound
		}

		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	eth, err := h.ctrl.CreateOne(ctx, e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eth); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) PutEthernet(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	eth, err := h.ctrl.FindByID(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &serror.ErrorNotFound{}) {
			code = http.StatusNotFound
		}

		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	e := &ethernet.Ethernet{}
	if err := json.NewDecoder(r.Body).Decode(e); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if e.AgentID == "" || e.DevName == "" || e.Name == "" || e.Hardware == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("any of agent_id, dev_name, name, hardware is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	e.ID = eth.ID
	eu, err := h.ctrl.UpdateOne(ctx, e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eu); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) DeleteEthernet(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	c, err := h.ctrl.DeleteOne(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		if errors.As(err, &serror.ErrorNotFound{}) {
			code = http.StatusNotFound
		}
		w.WriteHeader(code)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}