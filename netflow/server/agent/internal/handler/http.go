package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/infEreb/nls/netflow/server/agent/internal"
	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
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

func (h *HandlerHTTP) GetAgents(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ags, err := h.ctrl.FindAll(ctx)
	if err != nil && errors.As(err, &serror.ErrorNotFound{}) {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ags); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}
func (h *HandlerHTTP) GetAgent(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	as, err := h.ctrl.FindByID(ctx, id)
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
	if err := json.NewEncoder(w).Encode(as); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
}

func (h *HandlerHTTP) PostAgent(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	a := &agent.Agent{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if a.Hostname == "" || a.Status == ""  {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("any of hostname, status is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	ag, err := h.ctrl.CreateOne(ctx, a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ag); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) PutAgent(w http.ResponseWriter, r *http.Request, id string) {
	ctx := r.Context()

	ag, err := h.ctrl.FindByID(ctx, id)
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

	a := &agent.Agent{}
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		err = fmt.Errorf("request decode error: %v\n", err)
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	// required fileds check
	if a.Hostname == "" || a.Status == ""  {
		w.WriteHeader(http.StatusBadRequest)
		err := serror.ErrorNotFound{Err: fmt.Errorf("any of hostname, status is empty")}
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	a.ID = ag.ID
	au, err := h.ctrl.UpdateOne(ctx, a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(err.Error()); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(au); err != nil {
		log.Printf("response encode error: %v\n", err)
	}
	return
}

func (h *HandlerHTTP) DeleteAgent(w http.ResponseWriter, r *http.Request, id string) {
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

// func (h *HandlerHTTP) GetAgents(ctx *gin.Context) {
// 	res := agent.NewResponse()
// 	ags, _ := h.ctrl.FindAll(ctx)
// 	res.Resp = ags

// 	ctx.JSON(http.StatusOK, res)
// }

// func (h *HandlerHTTP) GetAgent(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	res := agent.NewResponse()

// 	a, err := h.ctrl.FindByID(ctx, id)
// 	if err != nil {
// 		code := http.StatusInternalServerError
// 		res.Errors = append(res.Errors, err.Error())
// 		if errors.As(err, &agent.ErrorNotFound{}) {
// 			code = http.StatusNotFound
// 		}
// 		ctx.JSON(code, res)
// 		return
// 	}

// 	res.Resp = a
// 	ctx.JSON(http.StatusOK, res)
// }

// func (h *HandlerHTTP) PostAgent(ctx *gin.Context) {
// 	res := agent.NewResponse()

// 	a := &agent.Agent{}
// 	if err := json.NewDecoder(ctx.Request.Body).Decode(a); err != nil {
// 		errs := fmt.Errorf("request decode error: %v\n", err)
// 		res.Errors = append(res.Errors, errs.Error())
// 		log.Print(errs)
// 		ctx.JSON(http.StatusInternalServerError, res)

// 		return
// 	}

// 	ag, err := h.ctrl.CreateOne(ctx, a)
// 	if err != nil {
// 		res.Errors = append(res.Errors, err.Error())
// 	}
// 	res.Resp = ag
// 	ctx.JSON(http.StatusOK, res)
// }

// func (h *HandlerHTTP) PutAgent(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	res := agent.NewResponse()

// 	ag, err := h.ctrl.FindByID(ctx, id)
// 	if err != nil {
// 		code := http.StatusInternalServerError
// 		res.Errors = append(res.Errors, err.Error())
// 		if errors.As(err, &agent.ErrorNotFound{}) {
// 			code = http.StatusNotFound
// 		}
// 		ctx.JSON(code, res)
// 		return
// 	}

// 	a := &agent.Agent{}
// 	if err := json.NewDecoder(ctx.Request.Body).Decode(a); err != nil {
// 		errs := fmt.Errorf("request decode error: %v\n", err)
// 		res.Errors = append(res.Errors, errs.Error())
// 		log.Print(errs)
// 		ctx.JSON(http.StatusInternalServerError, res)
// 		return
// 	}

// 	a.ID = ag.ID
// 	au, err := h.ctrl.UpdateOne(ctx, a)
// 	if err != nil {
// 		res.Errors = append(res.Errors, err.Error())
// 		ctx.JSON(http.StatusInternalServerError, res)
// 		return
// 	}

// 	res.Resp = au
// 	ctx.JSON(http.StatusOK, res)
// }

// func (h *HandlerHTTP) DeleteAgent(ctx *gin.Context) {
// 	id := ctx.Param("id")
// 	res := agent.NewResponse()

// 	c, err := h.ctrl.DeleteOne(ctx, id)
// 	if err != nil {
// 		code := http.StatusInternalServerError
// 		res.Errors = append(res.Errors, err.Error())
// 		if errors.As(err, &agent.ErrorNotFound{}) {
// 			code = http.StatusNotFound
// 		}
// 		ctx.JSON(code, res)
// 		return
// 	}

// 	res.Resp = c
// 	ctx.JSON(http.StatusOK, res)
// }
