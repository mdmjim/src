package UnitTest

import(
	"testing"
	"Rsystem"
)

func TestQstring(t *testing.T){
	var str=Rsystem.Qstring{" helloRsystemhelloworld12.3 "}
	t.Log(str ,"Len:",str.GetLen())
	t.Log("Has Prefix hell:",str.HasPrefix("hell")," Has Prefix well:",str.HasPrefix("well"))
    t.Log("Has suffix tem:",str.HasSuffix("tem")," Has suffix tom:",str.HasSuffix("tom"))
	t.Log("Index of lo:",str.Index("lo"),"Index of ao:",str.Index("ao"))
	t.Log("last index of lo:",str.LastIndex("lo")," last index of bo:",str.LastIndex("bo"))
	t.Log("split:",str.Split("h"))
	t.Log(str.Repeat(2))
	t.Log(str.Replace("he","we"))
	t.Log(str.ToLower())
	t.Log(str.ToUpper())
	t.Log(str.TrimeSpace())
	t.Log(str.TrimeCutters("WE"))
	t.Log(str.JoinLeft("This "))
	t.Log(str.JoinRight(" done").AppendBool( true) .AppendFloat(11.11111,2).JoinRight(" app 10 base int ").AppendInt(100).JoinRight("append 16 bse int:").AppendIntBase(100,16))
    str.Val="iphone 4.5$" 
	t.Log(str.SubAfter("ne").SubBefore("$").TrimeSpace().ToFloat())
	str.Val="iphone 4.5$"
	t.Log(str.ExFloat())
	str.Val="iphone 4.5$"
	t.Log(str.ExInt()) 
	str.Val="iphone ne 4.5$"
	t.Log(str.SubLastAfter("ne"))
	str.Val="iphone ne 4.5$"
	t.Log(str.SubLastBefore("ne"))
}
