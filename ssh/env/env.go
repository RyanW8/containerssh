package env

import (
	"fmt"
	"github.com/janoszen/containerssh/backend"
	"github.com/janoszen/containerssh/ssh/request"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type requestMsg struct {
	Name  string
	Value string
}

func onSetEnvRequest(request *requestMsg, session backend.Session) error {
	logrus.Trace(fmt.Sprintf("Set env request: %s=%s", request.Name, request.Value))
	return session.SetEnv(request.Name, request.Value)
}

var RequestTypeHandler = request.TypeHandler{
	GetRequestObject: func() interface{} { return &requestMsg{} },
	HandleRequest: func(request interface{}, reply request.Reply, channel ssh.Channel, session backend.Session) {
		err := onSetEnvRequest(request.(*requestMsg), session)
		if err != nil {
			logrus.Tracef("Failed env request (%s)", err)
			reply(false, nil)
		} else {
			reply(true, nil)
		}
	},
}
