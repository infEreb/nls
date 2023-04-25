package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/infEreb/nls/netflow/server/app/packet/internal"
	"github.com/infEreb/nls/netflow/server/app/packet/pkg/packet"
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

func (h *HandlerHTTP) GetPackets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	packs, err := h.ctrl.FindAll(ctx)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(packs); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}
func (h *HandlerHTTP) GetPacket(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	ps, err := h.ctrl.FindByID(ctx, id)
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
	if err := json.NewEncoder(w).Encode(ps); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}

func (h *HandlerHTTP) PostPacket(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p := &packet.Packet{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if p.Content == "" || p.AgentID == "" || p.EthernetID == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("some of content, agent_id or ethernet_id is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	pck, err := h.ctrl.CreateOne(ctx, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pck); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) PutPacket(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	pck, err := h.ctrl.FindByID(ctx, id)
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

	p := &packet.Packet{}
	if err := json.NewDecoder(r.Body).Decode(p); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if p.Content == "" || p.AgentID == "" || p.EthernetID == "" {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("some of content, agent_id or ethernet_id is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}


	p.ID = pck.ID
	pu, err := h.ctrl.UpdateOne(ctx, p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(pu); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) DeletePacket(w http.ResponseWriter, r *http.Request, id string) {
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