package base_info

import (
    log "jstsgagent/tsglog"
    cfg "jstsgagent/tsgcfg"
    "net/http"
    "net/url"
    "encoding/json"
    "strconv"
    "time"
    "fmt"
    "strings"
)

type baseInfo struct {
    SN string `json:"sn"`
    AddrArea string `json:"addr_area"`
    Address string `json:"addr"`
    VendorCode string `json:"vendor_code"`
    CompanyName string `json:"company_name"`
    Industry string `json:"industry"`
    ServiceDays string `json:"service_days"`
    ServiceStartTime int64 `json:"service_start_time"`
    OrderTime int64 `json:"order_time"`
    SetupTime int64 `json:"setup_time"`
    UserName string `json:"ss_user_name"`
    Password string `json:"ss_pwd"`
}

func fillBaseInfo() string {
    info:=baseInfo{}
    cfg_info:=cfg.CONF.BaseInfo

    info.SN=cfg_info.SN
    info.AddrArea=cfg_info.AddrArea
    info.Address=cfg_info.Address
    info.VendorCode=cfg_info.VendorCode
    info.CompanyName=cfg_info.CompanyName
    info.Industry=cfg_info.Industry
    info.ServiceDays=strconv.Itoa(cfg_info.ServiceDays)
    
    t,_:=time.Parse(cfg.TimeLayout,cfg_info.ServiceStartTime)
    info.ServiceStartTime=t.Unix()*1000
    
    t,_=time.Parse(cfg.TimeLayout,cfg_info.OrderTime)
    info.OrderTime=t.Unix()*1000

    t,_=time.Parse(cfg.TimeLayout,cfg_info.InstallTime)
    info.SetupTime=t.Unix()*1000

    info.UserName=cfg_info.UserName
    info.Password=cfg_info.Password

    s,err:=json.Marshal(info)
    if err!=nil {
        log.Fatal("json encode base info failed.error:%s",err)
    }

    return fmt.Sprintf("[%s]",s)
}

func registerOnce() bool {
    url_param:=url.Values{}
    url_param.Add("vendor_name",cfg.CONF.BaseInfoRegister.VendorName)
    url_param.Add("vendor_secret",cfg.CONF.BaseInfoRegister.VendorSecret)
    
    url:=fmt.Sprintf("%s?%s",cfg.CONF.BaseInfoRegister.URL,url_param.Encode())
    
    log.Info("Begine register to url:%s",url)

    resp,err:=http.Post(url,"application/json",strings.NewReader(fillBaseInfo()))
    if err!=nil {
        log.Error("register base info to:%s failed. error:%s",url,err)
        return false
    }

    defer resp.Body.Close()

    if resp.StatusCode!=200 {
        log.Error("register base info to:%s failed. resp code:%d != 200",
            url,resp.StatusCode)
        return false
    }

    log.Info("REGISTER TO URL:%s SUCCESS...",url)

    return true
}

func BaseInfoRegister() {
    for {
        flag:=registerOnce()  
        if flag==true {
            break
        }

        log.Error("retry register 5 seconds later!!!")

        time.Sleep(time.Duration(5)*time.Second)
    }
}
