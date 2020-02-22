package server

import (
	"io"
	"net"
	"net/http"

	"gogw/common"
	"gogw/logger"
	"gogw/schema"
)

const (
	BUFSIZE = 1024 * 1024
)

func (c *Client) HttpHandler(w http.ResponseWriter, req *http.Request) {
	msgPack, err := schema.ReadMsg(req.Body)
	if err != nil {
		logger.Error(err)
		return
	}

	if msgPack.MsgType == schema.MSG_TYPE_OPEN_CONN_REQUEST {
		msg, _ := msgPack.Msg.(*schema.OpenConnRequest)
		c.openConnHandler(msg, w, req)

	}else if msgPack.MsgType == schema.MSG_TYPE_MSG_COMMON_REQUEST {
		msg := <- c.MsgChann
		schema.WriteMsg(w, msg)
	}
}

func (c *Client) openConnHandler(msg *schema.OpenConnRequest, w http.ResponseWriter, req *http.Request) {
	if msg.Role == schema.ROLE_QUERY_CONNID {
		//Forward client: open a new conn
		msgPack := & schema.MsgPack {
			MsgType: schema.MSG_TYPE_OPEN_CONN_RESPONSE,
			Msg: & schema.OpenConnResponse {
				ConnId: "",
				Status: schema.STATUS_FAILED,
			},
		}

		var conn net.Conn
		var err error
		if conn, err = net.Dial(c.Protocol, c.SourceAddr); err == nil {
			connId := common.UUID("connid")
			msgPack.Msg = & schema.OpenConnResponse {
				ConnId: connId,
				Status: schema.STATUS_SUCCESS,
			}

			c.addConn(connId, conn)
		}

		schema.WriteMsg(w, msgPack)

	}else if msg.Role == schema.ROLE_READER {
		logger.Debug("reader: ", msg.ConnId)

		if conni, ok := c.Conns.Load(msg.ConnId); ok {
			conn, _ := conni.(*common.Conn)
			_, err := io.Copy(conn.Conn, req.Body)
			logger.Error(err)
		}	

	}else if msg.Role == schema.ROLE_WRITER {
		logger.Debug("writer: ", msg.ConnId)

		if conni, ok := c.Conns.Load(msg.ConnId); ok {
			conn, _ := conni.(*common.Conn)

			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			data := make([]byte, BUFSIZE)
			var err error 
			var n int 
			for {
				n, err = conn.Conn.Read(data)
				if err != nil {
					break
				}
				w.Write(data[:n])
				ww, _ := w.(http.Flusher)
				ww.Flush()
			}
			
			logger.Error(err)
		}

	}else {
		logger.Error("Unknown role", msg.Role)
	}
}