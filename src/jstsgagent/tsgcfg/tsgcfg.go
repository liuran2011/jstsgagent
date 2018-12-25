package tsgcfg

import (
	"encoding/json"
	"io/ioutil"
	log "jstsgagent/tsglog"
    "time"
)

const cfgFilePath="/etc/jstsgagent/jstsgagent.conf"
const TimeLayout="2006-01-02 15:04:05 MST"

type HeartBeatConf struct {
	Port int `json:"port"`
	Interval int `json:"interval"`
}

type CtrlPlatformConf struct {
	HeartBeat *HeartBeatConf `json:"heartbeat"`
    Port int `json:"port"`
	Address string `json:"address"`
}

type DataPlatformConf struct {
	Port int `json:"port"`
	Address string `json:"address"`
}

type BaseInfoRegisterConf struct {
    URL string `json:"url"`
    VendorName string `json:"vendor_name"`
    VendorSecret string `json:"vendor_secret"`
}

type BaseInfoConf struct {
	SN string `json:"sn"`
	AddrArea string `json:"addr_area"`
	Address string `json:"address"`
	VendorCode string `json:"vendor_code"`
	CompanyName string `json:"company_name"`
	Industry string `json:"industry"`
	ServiceDays int `json:"service_days"`
	ServiceStartTime string `json:"service_start_time"`
	OrderTime string `json:"order_time"`
	InstallTime string `json:"install_time"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type JSTsgAgentConf struct {
	CtrlPlatform *CtrlPlatformConf `json:"control_platform"`
	DataPlatform *DataPlatformConf `json:"data_platform"`
    BaseInfoRegister *BaseInfoRegisterConf `json:"base_info_register"`
	BaseInfo *BaseInfoConf `json:"base_info"`
}

func baseInfoRegisterSanityCheck(info *BaseInfoRegisterConf) {
    if info==nil {
        log.Fatal("no base_info_register section!")
    }

    if info.URL=="" {
        log.Fatal("invalid url in base_info_register section!")
    }

    if info.VendorName=="" {
        log.Fatal("invalid vendor_name on base_info_regsiter section!")
    }

    if info.VendorSecret=="" {
        log.Fatal("invalid vendor_secret on base_info_register section!")
    }
}

func baseInfoSanityCheck(info *BaseInfoConf) {
    if info==nil {
        log.Fatal("no base_info section!")
    }

	if info.SN=="" {
		log.Fatal("invalid sn!")
	}
	
	if info.AddrArea=="" {
		log.Fatal("invalid addr_area!")
	}

	if info.Address=="" {
		log.Fatal("invalid address!")
	}

	if info.VendorCode=="" {
		log.Fatal("invalid vendor_code!")
	}

	if info.CompanyName=="" {
		log.Fatal("invalid company_name!")	
	}

	if info.Industry=="" {
		log.Fatal("invalid industry!")
	}

	if info.ServiceDays==0 {
		log.Fatal("invalid service_days!")
	}

	if info.ServiceStartTime=="" {
		log.Fatal("invalid service_start_time!")
	}

    _,err:=time.Parse(TimeLayout,info.ServiceStartTime)
    if err!=nil {
        log.Fatal("invalid service_start_time format,error:%s YEAR-MONTH-DAY HOUR:MINUTE:SECOND ZONE!",
            err)
    }

	if info.OrderTime=="" {
		log.Fatal("invalid order_time!")
	}

    _,err=time.Parse(TimeLayout,info.OrderTime)
    if err!=nil {
        log.Fatal("invalid order_time format,error:%s YEAR-MONTH-DAY HOUR:MINUTE:SECOND ZONE!",err)
    }

	if info.InstallTime=="" {
		log.Fatal("invalid install_time!")
	}

    _,err=time.Parse(TimeLayout,info.InstallTime)
    if err!=nil {
        log.Fatal("invalid install_time format, error:%s YEAR-MONTH-DAY HOUR:MINUTE:SECOND ZONE!",err)
    }

	if info.UserName=="" {
		log.Fatal("invalid user_name!")
	}

	if info.Password=="" {
		log.Fatal("invalid password!")
	}
}

func ctrlPlatformSanityCheck(conf *CtrlPlatformConf) {
    if conf==nil {
        log.Fatal("no control_platform section!")
    }

	if conf.Address=="" {
		log.Fatal("invalid control platform address!")
	}

    if conf.Port==0 {
        log.Fatal("invalid control platform port!")
    }

	if conf.HeartBeat.Port==0 {
		log.Fatal("invalid control platform heartbeat port!")
	}

	if conf.HeartBeat.Interval<5 {
		log.Fatal("invalid control platform heartbeat interval!")
	}
}

func dataPlatformSanityCheck(conf *DataPlatformConf) {
    if conf==nil {
        log.Fatal("no data_platform section!")
    }

	if conf.Port==0 {
		log.Fatal("invalid dataplatform port!")
	}

	if conf.Address=="" {
		log.Fatal("invalid dataplatform address!")
	}
}

var CONF JSTsgAgentConf 

func ReadConfig() {
	data,err:=ioutil.ReadFile(cfgFilePath)
	if err!=nil {
		log.Fatal("open config file:%s failed.error:%s",cfgFilePath,err)
	}

	err=json.Unmarshal(data,&CONF)
	if err!=nil {
		log.Fatal("parse config file:%s failed.error:%s",cfgFilePath,err)
	}

	dataPlatformSanityCheck(CONF.DataPlatform)
	ctrlPlatformSanityCheck(CONF.CtrlPlatform)
    baseInfoRegisterSanityCheck(CONF.BaseInfoRegister)
	baseInfoSanityCheck(CONF.BaseInfo)
}

