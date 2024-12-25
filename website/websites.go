package website

import (
	"github.com/83Beyond/recruitment/models"
	"github.com/83Beyond/recruitment/network"
	"github.com/83Beyond/recruitment/website/Finance/ChinaLife"
	"github.com/83Beyond/recruitment/website/Finance/PICC"
	"github.com/83Beyond/recruitment/website/Finance/TaiPing"
)

var (
	AllWebsites = []network.Request{
		network.Request{
			Name:   "国寿",
			URL:    "https://chinalife.zhiye.com/api/Jobad/GetJobAdPageList",
			Method: "POST",
			Headers: map[string]string{
				"Accept":             "application/json, text/javascript, */*; q=0.01",
				"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8,pt;q=0.7",
				"Cache-Control":      "no-cache",
				"Connection":         "keep-alive",
				"Content-Type":       "application/json; charset=UTF-8",
				"Origin":             "https://chinalife.zhiye.com",
				"Pragma":             "no-cache",
				"Referer":            "https://chinalife.zhiye.com/custom/intern?hideMenu=1",
				"Sec-Fetch-Dest":     "empty",
				"Sec-Fetch-Mode":     "cors",
				"Sec-Fetch-Site":     "same-origin",
				"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
				"X-Requested-With":   "XMLHttpRequest",
				"sec-ch-ua":          "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
				"sec-ch-ua-mobile":   "?0",
				"sec-ch-ua-platform": "\"Windows\"",
			},
			PostData: map[string]any{
				"Category": []string{
					"3",
				},
				"SpecialType": 0,
				//"PageIndex":   1,
				"PageSize": 10,
				"DisplayFields": []string{
					"Category",
					"Kind",
					"LocId",
					"Org",
					"HeadCount",
					"Station",
					"EndTime",
					"PostDate",
					"Salary",
					"ClassificationOne",
					"ClassificationTwo",
				},
			},

			CurrentPage: 0,
			StartPage:   0,
			ScrapePage:  models.ScrapPage,
			Timestamp:   "",

			PageKey:       "PageIndex",
			TypeSelectKey: "Category.0",

			TypeSelectMap: map[string]string{
				"实习招聘": "3",
				"校园招聘": "2",
			},
			KeywordKey: "KeyWords",

			Parse: ChinaLife.Parse,
		},
		network.Request{
			Name:   "人保",
			URL:    "https://picc.zhiye.com/api/Jobad/GetJobAdPageList",
			Method: "POST",
			Headers: map[string]string{
				"Accept":          "application/json, text/javascript, */*; q=0.01",
				"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8,pt;q=0.7",
				"Cache-Control":   "no-cache",
				"Connection":      "keep-alive",
				"Content-Type":    "application/json; charset=UTF-8",
				"Origin":          "https://picc.zhiye.com",
				"Pragma":          "no-cache",
				//"Referer": "https://picc.zhiye.com/custom/campus?hideAll=true&ky=&c1=&c2=&d=&c=",
				"Sec-Fetch-Dest":     "empty",
				"Sec-Fetch-Mode":     "cors",
				"Sec-Fetch-Site":     "same-origin",
				"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
				"X-Requested-With":   "XMLHttpRequest",
				"sec-ch-ua":          "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
				"sec-ch-ua-mobile":   "?0",
				"sec-ch-ua-platform": "\"Windows\"",
			},
			PostData: map[string]any{
				"Category": []string{
					"2",
				},
				"SpecialType": 0,
				//"PageIndex": 1,
				"PageSize": 10,
				"DisplayFields": []string{
					"Category",
					"Kind",
					"LocId",
					"Org",
					"HeadCount",
					"Station",
					"EndTime",
					"PostDate",
					"Salary",
					"ClassificationOne",
					"ClassificationTwo",
				},
			},

			CurrentPage: 0,
			StartPage:   0,
			ScrapePage:  models.ScrapPage,
			Timestamp:   "",

			PageKey:       "PageIndex",
			TypeSelectKey: "Category.0",

			TypeSelectMap: map[string]string{
				"实习招聘": "3",
				"校园招聘": "2",
			},
			KeywordKey: "KeyWords",

			Parse: PICC.Parse,
		},
		network.Request{
			Name:   "太平",
			URL:    "https://cntp.zhiye.com/api/Jobad/GetJobAdPageList",
			Method: "POST",
			Headers: map[string]string{
				"Accept":             "application/json",
				"Accept-Language":    "zh-CN,zh;q=0.9,en;q=0.8,pt;q=0.7",
				"Cache-Control":      "no-cache",
				"Connection":         "keep-alive",
				"Content-Type":       "application/json",
				"EagleEye-TraceID":   "cff945a9-0059-a740-f0e2-c06661fcf04b",
				"Origin":             "https://cntp.zhiye.com",
				"Pragma":             "no-cache",
				"Referer":            "https://cntp.zhiye.com/intern/jobs",
				"Sec-Fetch-Dest":     "empty",
				"Sec-Fetch-Mode":     "cors",
				"Sec-Fetch-Site":     "same-origin",
				"User-Agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
				"X-Requested-With":   "xmlhttprequest",
				"langType":           "zh_CN",
				"sec-ch-ua":          "\"Google Chrome\";v=\"131\", \"Chromium\";v=\"131\", \"Not_A Brand\";v=\"24\"",
				"sec-ch-ua-mobile":   "?0",
				"sec-ch-ua-platform": "\"Windows\"",
			},
			PostData: map[string]any{
				"PageIndex": 1,
				"PageSize":  15,
				"Category": []string{
					"3",
				},
				"KeyWords":    "",
				"SpecialType": 0,
				"PortalId":    "",
				"DisplayFields": []string{
					"Kind",
					"LocId",
					"Org",
					"PostDate",
					"WorkWeChatQrCode",
				},
			},

			CurrentPage: 0,
			StartPage:   0,
			ScrapePage:  models.ScrapPage,
			Timestamp:   "",

			PageKey:       "PageIndex",
			TypeSelectKey: "Category.0",

			TypeSelectMap: map[string]string{
				"实习招聘": "3",
				"校园招聘": "2",
			},
			KeywordKey: "KeyWords",

			Parse: TaiPing.Parse,
		},
	}
)
