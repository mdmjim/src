package Rsystem

import(
	"net/http"
    "io/ioutil"
)

type HtmlGrabber struct{
	 client *http.Client
}

func (grabber *HtmlGrabber) GetResponse(req *http.Request) string{
	response,_ :=grabber.client.Do(req)
    if response.StatusCode == 200 {
        body, _ := ioutil.ReadAll(response.Body)
        bodystr := string(body); 
		return bodystr;
    }
	return ""
}

func (grabber *HtmlGrabber) GetHtml(url string,charset string) string{
	
	reqest, _ := http.NewRequest("GET", url, nil) 
    reqest.Header.Set("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
    //reqest.Header.Set("Accept-Charset","GBK,utf-8;q=0.7,*;q=0.3")
    reqest.Header.Set("Accept-Encoding","gzip,deflate,sdch")
    //reqest.Header.Set("Accept-Language","zh-CN,zh;q=0.8")
    reqest.Header.Set("Cache-Control","max-age=0")
    reqest.Header.Set("Connection","keep-alive")
	reqest.Header.Set("User-Agent","Mozilla/5.0 (X11; Linux i686) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.83 Safari/535.11")
     // "Cookie": "_gauges_unique_day=1; _gauges_unique_month=1; _gauges_unique_year=1; _gauges_unique=1; _gauges_unique_hour=1"
	return grabber.GetResponse(reqest)
}

