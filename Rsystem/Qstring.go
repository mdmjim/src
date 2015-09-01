package Rsystem

import(
	//"fmt"
	"strings"
	"strconv"
	"regexp"
	"errors"
)

type Qstring struct{
	Val string
}

func (str * Qstring) GetLen() int { 
	return len(str.Val);
}

func (str *Qstring) HasPrefix(pat string)bool{
	if(str.GetLen()==0){
		return false;
	}
	return strings.HasPrefix(str.Val,pat);
}

func (str *Qstring) HasSuffix(pat string)bool{
	if(str.GetLen()==0){
		return false;
	}
	return strings.HasSuffix(str.Val,pat);
}

func (str *Qstring) Index(pat string) int{
	if(str.GetLen()==0){
		return -1;
	}
	return strings.Index(str.Val,pat)
}

func (str *Qstring) LastIndex(pat string) int{
	if(str.GetLen()==0){
		return -1;
	}
	return strings.LastIndex(str.Val,pat);
}


func (str *Qstring) Contains(pat string)bool{ 
	return str.Index(pat)>=0;
}


func (str *Qstring) Split(pat string) []string{
	if(str.GetLen()==0){
		return nil;
	}
	return strings.Split(str.Val,pat);
}

//*************** Chain funcs ************************



func(str *Qstring) Repeat(count int) *Qstring{
	if(str.GetLen()>0){
		str.Val=strings.Repeat(str.Val,count);
	}
	return str;
}

func (str *Qstring) Replace(old_val string,new_val string) *Qstring{
	if(str.GetLen()>0){
		str.Val=strings.Replace(str.Val,old_val,new_val,-1)
	}
	return str;
}


func (str *Qstring) ToLower() *Qstring{
	if(str.GetLen()>0){
	   str.Val= strings.ToLower(str.Val);
	}
	return str;
}

func (str *Qstring) ToUpper() *Qstring{
	if(str.GetLen()>0){
	   str.Val= strings.ToUpper(str.Val);
	}
	return str;
}

func (str *Qstring) TrimeCutters(cutters string) *Qstring{
	if(str.GetLen()>0){
		str.Val=strings.Trim(str.Val,cutters);
	}
	return str;
}

func (str *Qstring) TrimeSpace() *Qstring{
	if(str.GetLen()>0){
		str.Val=strings.TrimSpace(str.Val);
	}
	return str;
}

func (str *Qstring) JoinLeft(val string) *Qstring{
	str.Val=val + str.Val;
	return str;
}

func (str *Qstring) JoinRight(val string) *Qstring{
	str.Val=str.Val+val;
	return str;
}

func (str *Qstring) AppendBool(val bool) *Qstring{
	str.Val=string(strconv.AppendBool([]byte(str.Val),val))
	return str;
}

func (str *Qstring) AppendFloat(val float64,perc int) *Qstring{
	str.Val=string(strconv.AppendFloat([]byte(str.Val),val,'E',perc,64))
	return str;
}

func (str *Qstring) AppendInt(val int64) *Qstring{
	str.Val=string(strconv.AppendInt([]byte(str.Val),val,10))
	return str;
}

func (str *Qstring) AppendIntBase(val int64,base int) *Qstring{
	str.Val=string(strconv.AppendInt([]byte(str.Val),val,base))
	return str;
}

func (str *Qstring) SubStr(start int ,end int) *Qstring{
	var len=str.GetLen();
	if(start>=len||end<start){
		str.Val="";
		return str;
	} 
	if(end >=len){
		end=len-1
	}
	
	str.Val=string(str.Val[start:end+1])
	
	return str;
}


func (str *Qstring) SubAfter(val string) *Qstring{
	var ind=str.Index(val)
	if(ind<0){
		str.Val=""
	}else{
		str.SubStr(ind+len(val),str.GetLen())
	}
	return str;
}

func (str *Qstring) SubLastAfter(val string) *Qstring{
	var ind=str.LastIndex(val)
	if(ind<0){
		str.Val=""
	}else{
		str.SubStr(ind+len(val),str.GetLen())
	}
	return str;
}


func (str *Qstring) SubBefore(val string) *Qstring{
	var ind=str.Index(val)
	if(ind<0){
		str.Val=""
	}else{
		str.SubStr(0,ind-1)
	}
	return str;
}

func (str *Qstring) SubLastBefore(val string) *Qstring{
	var ind=str.LastIndex(val)
	if(ind<0){
		str.Val=""
	}else{
		str.SubStr(0,ind-1)
	}
	return str;
}

//*************     Convertor    ****************


func (str *Qstring) ToInt(base int) (int64,error){ 
	return strconv.ParseInt(str.Val,base,64)
}

func (str *Qstring) ToBool() (bool,error){ 
	return strconv.ParseBool(str.Val)
}

func (str *Qstring) ToFloat() (float64,error){ 
	return strconv.ParseFloat(str.Val,64)
}

var ExIntReg=regexp.MustCompile("\\d+")

func (str *Qstring) ExInt()(int64,error){
	if(str.GetLen()==0){
		return 0,errors.New("Len==0")
	}
	t :=ExIntReg.FindString(str.Val) 
	return strconv.ParseInt(t,10,64) 
}

var ExFloatReg=regexp.MustCompile("\\d+\\.?\\d*")

func (str *Qstring) ExFloat()(float64,error){
	if(str.GetLen()==0){
		return 0,errors.New("Len==0")
	}
	t :=ExFloatReg.FindString(str.Val) 
	return strconv.ParseFloat(t,64)
}




