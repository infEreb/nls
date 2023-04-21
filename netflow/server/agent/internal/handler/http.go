package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/infEreb/nls/netflow/server/agent/internal"
	"github.com/infEreb/nls/netflow/server/agent/pkg/agent"
)

type HandlerHTTP struct {
	ctrl *internal.Controller
}

func NewHandlerHTTP(ctrl *internal.Controller) *HandlerHTTP {
	return &HandlerHTTP{
		ctrl: ctrl,
	}
}

// func (h *Handler) GetAgents(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
	
// 	res := agent.NewResponse()
// 	ags, _ := h.ctrl.FindAll(ctx)
// 	res.Resp = ags

// 	w.WriteHeader(http.StatusOK)
// 	if err := json.NewEncoder(w).Encode(ags); err != nil {
// 		log.Printf("response encode error: %v\n", err)
// 	}
// }

func (h *HandlerHTTP) GetAgents(ctx *gin.Context) {
	res := agent.NewResponse()
	ags, _ := h.ctrl.FindAll(ctx)
	res.Resp = ags

	ctx.JSON(http.StatusOK, res)
}

func (h *HandlerHTTP) GetAgent(ctx *gin.Context) {
	id := ctx.Param("id")
	res := agent.NewResponse()

	a, err := h.ctrl.FindByID(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		res.Errors = append(res.Errors, err.Error())
		if errors.As(err, &agent.ErrorNotFound{}) {
			code = http.StatusNotFound
		}
		ctx.JSON(code, res)
		return
	}

	res.Resp = a
	ctx.JSON(http.StatusOK, res)
}

func (h *HandlerHTTP) PostAgent(ctx *gin.Context) {
	res := agent.NewResponse()

	a := &agent.Agent{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(a); err != nil {
		errs := fmt.Errorf("request decode error: %v\n", err)
		res.Errors = append(res.Errors, errs.Error())
		log.Print(errs)
		ctx.JSON(http.StatusInternalServerError, res)
		
		return
	}
	
	ag, err := h.ctrl.CreateOne(ctx, a)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
	}
	res.Resp = ag
	ctx.JSON(http.StatusOK, res)
}

func (h *HandlerHTTP) PutAgent(ctx *gin.Context) {
	id := ctx.Param("id")
	res := agent.NewResponse()

	ag, err := h.ctrl.FindByID(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		res.Errors = append(res.Errors, err.Error())
		if errors.As(err, &agent.ErrorNotFound{}) {
			code = http.StatusNotFound
		}
		ctx.JSON(code, res)
		return
	}

	a := &agent.Agent{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(a); err != nil {
		errs := fmt.Errorf("request decode error: %v\n", err)
		res.Errors = append(res.Errors, errs.Error())
		log.Print(errs)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	a.ID = ag.ID
	au, err := h.ctrl.UpdateOne(ctx, a)
	if err != nil {
		res.Errors = append(res.Errors, err.Error())
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Resp = au
	ctx.JSON(http.StatusOK, res)
}

func (h *HandlerHTTP) DeleteAgent(ctx *gin.Context) {
	id := ctx.Param("id")
	res := agent.NewResponse()

	c, err := h.ctrl.DeleteOne(ctx, id)
	if err != nil {
		code := http.StatusInternalServerError
		res.Errors = append(res.Errors, err.Error())
		if errors.As(err, &agent.ErrorNotFound{}) {
			code = http.StatusNotFound
		}
		ctx.JSON(code, res)
		return
	}

	res.Resp = c
	ctx.JSON(http.StatusOK, res)
}