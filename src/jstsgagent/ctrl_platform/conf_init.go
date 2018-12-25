package ctrl_platform

import (
    log "jstsgagent/tsglog"
    cfg "jstsgagent/tsgcfg"
    codec "jstsgagent/codec"
    "time"
    "net"
)

func confInitOnce() bool {
    addr:=fmt.Sprintf("%s:%d",cfg.CONF.CtrlPlatform.Address,
            cfg.CONF.CtrlPlatform.Port)

    conn,err:=net.Dial("tcp",addr)
    if err!=nil {
        log.Error("tcp connect: %s failed. error:%s",addr,err)
        return false
    }

    defer conn.Close()

    buf,err:=codec.NewConfInitMessage()
    
    len,err:=conn.Write(buf.Bytes())
    if err!=nil { 
        log.Error("write conf_init message failed. error:%s",err)
        return false
    }

    if len!=buf.Len() {
        log.Error("write len:%d != buf_len:%d",len,buf.Len())
        return false
    }

    resp,err:=codec.GetConfInitResponse(conn)
    if err!=nil {
        log.Error("recv conf_init response failed.error:%s",err)
        return false
    }

    if ! resp.IsResultOk()  {
        log.Error("conf_init response action code:%d != Result_OK",resp.ActionCode)
        return false
    }

    log.Info("CONF_INIT SUCCESS...")

    return true
}

func ConfInit() {
    for {
        flag:=confInitOnce()
        if flag==true {
            break
        }

        log.Info("retry conf_init 1 minute later!!!")

        time.Sleep(time.Duration(1)*time.Minute)
    }
}
