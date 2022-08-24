/* // +build ignore */

package main

//#include <stdio.h>
//#include <stdlib.h>
//int strlenM(char *strIn)
//{
//	int i;
//	i = 0;
//	do {
//		if (strIn[i] == '\0') break;
//		i = i + 1;
//	} while (strIn[i] != '\0');
//	return i;
//}
//char *doDecrypt(char *strVar, int blnOdd)
//{
//	static char strRetDoDecrypt[9]; /* must be static for successful return */
//	int i, intLen, intError;
//	int intArr[8]; /* 0,...,7 indexes */
//	int intArrDecr[8]; /* 0,...,7 indexes */
//	int intTemp[8]; /* 0,...,7 indexes */
//	int intOdd[8][2]; /* 0,...,7  and 0,1 indexes */
//	int intEven[8][2]; /* 0,...,7 and 0,1 indexes */
//	int intMatrix[8][2]; /* 0,...,7 and 0,1 indexes */
//	for (i = 0; i < 9; i = i + 1) strRetDoDecrypt[i] = '\0';
//	intLen = strlenM(strVar); /* count of chars in strVar[], last '\0' NOT included */
//	if ((intLen == 0) || (intLen < 6) || (intLen > 8)) {
//		strRetDoDecrypt[0] = '\0'; /* empty string */
//		return strRetDoDecrypt;
//	}
//	for (i = 0; i < 8; i = i + 1) intArr[i] = 0;
//	intError = 0; /* 0 is False, 1(non zero) is True */
//	for (i = 0; i < intLen; i = i + 1) {
//		intArr[i] = strVar[i] - '0';
//		if (intArr[i] > 9) {
//			intError = 1;
//			break;
//		}
//	}
//	if (intError) {
//		strRetDoDecrypt[0] = '\0'; /* empty string */
//		return strRetDoDecrypt;
//	}
//	for (i = 0; i < 8; i = i + 1) intArrDecr[i] = 0;
//	for (i = 0; i < 8; i = i + 1) intTemp[i] = 0;
//	if (intLen == 6) {
//		intOdd[0][0] = 5; intOdd[0][1] = 4;
//		intOdd[1][0] = 2; intOdd[1][1] = 0;
//		intOdd[2][0] = 4; intOdd[2][1] = 3;
//		intOdd[3][0] = 6; intOdd[3][1] = 2;
//		intOdd[4][0] = 8; intOdd[4][1] = 5;
//		intOdd[5][0] = 1; intOdd[5][1] = 1;
//		intEven[0][0] = 8; intEven[0][1] = 5;
//		intEven[1][0] = 1; intEven[1][1] = 4;
//		intEven[2][0] = 5; intEven[2][1] = 0;
//		intEven[3][0] = 2; intEven[3][1] = 1;
//		intEven[4][0] = 9; intEven[4][1] = 3;
//		intEven[5][0] = 6; intEven[5][1] = 2;
//	}
//	else if (intLen == 7) {
//		intOdd[0][0] = 5; intOdd[0][1] = 5;
//		intOdd[1][0] = 2; intOdd[1][1] = 0;
//		intOdd[2][0] = 4; intOdd[2][1] = 3;
//		intOdd[3][0] = 6; intOdd[3][1] = 2;
//		intOdd[4][0] = 8; intOdd[4][1] = 6;
//		intOdd[5][0] = 1; intOdd[5][1] = 1;
//		intOdd[6][0] = 9; intOdd[6][1] = 4;
//		intEven[0][0] = 8; intEven[0][1] = 5;
//		intEven[1][0] = 1; intEven[1][1] = 6;
//		intEven[2][0] = 5; intEven[2][1] = 0;
//		intEven[3][0] = 2; intEven[3][1] = 4;
//		intEven[4][0] = 9; intEven[4][1] = 3;
//		intEven[5][0] = 6; intEven[5][1] = 1;
//		intEven[6][0] = 4; intEven[6][1] = 2;
//	}
//	else {
//		intOdd[0][0] = 5; intOdd[0][1] = 6;
//		intOdd[1][0] = 2; intOdd[1][1] = 0;
//		intOdd[2][0] = 4; intOdd[2][1] = 3;
//		intOdd[3][0] = 6; intOdd[3][1] = 2;
//		intOdd[4][0] = 8; intOdd[4][1] = 7;
//		intOdd[5][0] = 1; intOdd[5][1] = 1;
//		intOdd[6][0] = 9; intOdd[6][1] = 4;
//		intOdd[7][0] = 3; intOdd[7][1] = 5;
//		intEven[0][0] = 8; intEven[0][1] = 5;
//		intEven[1][0] = 1; intEven[1][1] = 6;
//		intEven[2][0] = 5; intEven[2][1] = 0;
//		intEven[3][0] = 2; intEven[3][1] = 4;
//		intEven[4][0] = 9; intEven[4][1] = 3;
//		intEven[5][0] = 6; intEven[5][1] = 7;
//		intEven[6][0] = 4; intEven[6][1] = 2;
//		intEven[7][0] = 7; intEven[7][1] = 1;
//	}
//
//	if (blnOdd) for (i = 0; i < 8; i = i + 1) {
//		intMatrix[i][0] = intOdd[i][0];
//		intMatrix[i][1] = intOdd[i][1];
//	}
//	else for (i = 0; i < 8; i = i + 1) {
//		intMatrix[i][0] = intEven[i][0];
//		intMatrix[i][1] = intEven[i][1];
//	}
//	for (i = 0; i < intLen; i = i + 1) {
//		intTemp[i] = intArr[i] - intMatrix[i][0];
//		if (intTemp[i] < 0)
//			intTemp[i] = intTemp[i] + 10;
//		intArrDecr[intMatrix[i][1]] = intTemp[i];
//	}
//	for (i = 0; i < intLen; i = i + 1) {
//		strRetDoDecrypt[i] = intArrDecr[i] + '0';
//	}
//	strRetDoDecrypt[i] = '\0';
//	return strRetDoDecrypt;
//}
import "C"
import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const pageNameIni = string("index.html") // string("index.html") or string("indexForm.html")
const dir = string("./build")            // string("./build") or string("./build_Test")

// project WinTicsCheckNoSslTEST new at 'http://10.8.194.3:9994/'
//const urlval = "https://unl.works:9999/" // for test with real data.
const urlval = "http://10.8.194.3:9994/"
const reqStringCheckVal = urlval + "?agent=58&type=2&command=checkval&ticket_number=" // + search;
const urlpay = "http://10.8.194.3:38000/"                                             // project PayTest ver. 4.0 NoSsl.
const strAgent string = "16"

var txnid int = 10000000

func handlerReq(w http.ResponseWriter, r *http.Request) {
	//dir := "./build"
	//fmt.Fprintf(w, "The URL.Path[1:] is %s", r.URL.Path[1:])
	// initial r.Method: "GET" r.URL.Path e.g. "/" or "/script.js" or "/external-link-52.png" or "/favicon.ico" r.RequestURI e.g. "/" or "/?first=1&second=2
	// after form method="get" r.Method: "GET" r.URL.Path e.g. /formAKpay" r.RequestURI e.g. "/formAKpay?game=KN&n1=2&n2=1&q=15"
	// after form method="post" r.Method: "POST" r.URL.Path "/formAKpay" r.RequestURI "/formAKpay" r.Body []uint8 with e.g. (length: 20, cap: 512)
	fmt.Println("r.Mehtod:", r.Method, "r.URL.Path:", r.URL.Path, "r.RequsetURI:", r.RequestURI)
	var filename string = ""
	// Processing GET or PUT method.
	if r.Method == "GET" {
		// initial r.Method: "GET" r.URL.Path e.g. "/" or "/script.js" or "/external-link-52.png" or "/favicon.ico" r.RequestURI e.g. "/" or "/?first=1&second=2
		// after form method="get" r.Method: "GET" r.URL.Path e.g. /formAKpay" r.RequestURI e.g. "/formAKpay?game=KN&n1=2&n2=1&q=15"
		if r.URL.Path == "/" { // if r.URL.Path is "/" then default initial page wll be used and r.RequestURI info will be ignored.
			filename = dir + "/" + pageNameIni
			filecontent, err := ioutil.ReadFile(filename)
			if err == nil {
				w.Header().Set("Content-Type", "text/html")
				fmt.Fprintf(w, "%s", filecontent)
			} else {
				fmt.Fprintf(w, "File not found %s", filename)
			}
		} else {
			if r.URL.Path == r.RequestURI { // files refered from initial page.
				filename = dir + r.URL.Path
				//ttt := r.URL.Path[1:] // get "" or "script.js" or "external-link-52.png" or "favicon.ico"
				//fmt.Println(ttt)
				// Static files processing thru http.ServeFile
				http.ServeFile(w, r, filename)
				/*
					// Static files processing manually.
					filecontent, err := ioutil.ReadFile(filename)
					if err == nil {
						if strings.LastIndex(filename, ".svg") != -1 {
							// w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
							w.Header().Set("Content-Type", "image/svg+xml")
						} else if strings.LastIndex(filename, ".css") != -1 {
							w.Header().Set("Content-Type", "text/css")
						} else if strings.LastIndex(filename, ".js") != -1 {
							w.Header().Set("Content-Type", "application/javascript")
						} else if strings.LastIndex(filename, ".json") != -1 {
							w.Header().Set("Content-Type", "application/json")
						} else if strings.LastIndex(filename, ".map") != -1 {
							w.Header().Set("Content-Type", "application/map")
						} else if strings.LastIndex(filename, ".ico") != -1 {
							w.Header().Set("Content-Type", "image/bmp")
						} else if strings.LastIndex(filename, ".png") != -1 {
							w.Header().Set("Content-Type", "image/png")
						} else if strings.LastIndex(filename, ".jpg") != -1 || strings.LastIndex(filename, ".jpeg") != -1 {
							w.Header().Set("Content-Type", "image/jpeg")
						} else if strings.LastIndex(filename, ".html") != -1 {
							w.Header().Set("Content-Type", "text/html")
						} else {
							w.Header().Set("Content-Type", "application/octet-stream")
						}
						fmt.Fprintf(w, "%s", filecontent)
					} else {
						fmt.Fprintf(w, "File not found %s", filename)
					}
				*/
				// End of Static files processing.
			} else { // r.URL.Path != r.RequestURI
				// form method="get" r.Method: "GET" r.URL.Path e.g. /formAKpay" r.RequestURI e.g. "/formAKpay?game=KN&n1=2&n2=1&q=15"
				//fmt.Fprintf(w, "Form called %s", r.RequestURI)
				fmt.Printf("Form called %s\n", r.RequestURI)
				/*
					t1 := "0123456789"
					t2 := strings.Index(t1, "5")
					t3 := strings.Index(t1, "8")
					t4 := t1[t2 : t3] // get substring starting from "5" up to not include "8"
					fmt.Println(t1, t2, t3, t4) // 0123456789 5 8 567
				*/
				var strSearch string = ""
				var strTicnum string = ""
				var strPage string = ""
				var pos1 int = -1
				var pos2 int = -1
				var val string = ""
				if strings.Index(r.RequestURI, "/formAKchk?") != -1 {
					//r.RequestURI e.g. "/formAKchk?q=123-12345678-1234567" - old format
					//r.RequestURI e.g.  /formAKchk?q=123-12345678-1234567&v=Y
					pos2 = strings.Index(r.RequestURI, "&v=")
					val = r.RequestURI[pos2+3:]
					pos1 = strings.Index(r.RequestURI, "/formAKchk?q=")
					strTicnum = r.RequestURI[pos1+13 : pos2]
					fmt.Printf("Form formAKchk called with ticket number %s %s\n", strTicnum, val)
					strPage = checkValTicket(strTicnum, val)
					w.Header().Set("Content-Type", "text/html")
					fmt.Fprintf(w, "%s", strPage)
				} else if strings.Index(r.RequestURI, "/formAKpay?") != -1 {
					pos1 = strings.Index(r.RequestURI, "/formAKpay?q=")
					strSearch = r.RequestURI[pos1+13:] // e.g. "6_1_1_a_04_05_09_12_34_51"
					//fmt.Fprintf(w, "Form formAKpay called with params %s", strSearch)
					strPage = buyTicket(strSearch)
					w.Header().Set("Content-Type", "text/html")
					fmt.Fprintf(w, "%s", strPage)
				} else if strings.Index(r.RequestURI, "/formAKresults?") != -1 {
					// e.g. formAKresults?g=2&d= or formAKresults?g=2&d=1001
					pos1 = strings.Index(r.RequestURI, "/formAKresults?g=")
					strSearch = r.RequestURI[pos1+17:]
					// e.g. 2&d= or 2&d=1001
					//fmt.Fprintf(w, "Form formAKresults called with params %s", strSearch)
					strPage = getResults(strSearch)
					w.Header().Set("Content-Type", "text/html")
					fmt.Fprintf(w, "%s", strPage)
				} else {
					fmt.Fprintf(w, "NB!!! Unknown form called with URI: %s", r.RequestURI)
				}
			}
		} // End of r.URL.Path != "/" processing
		// End of processing GET method.
		// Processing POST method for testing in ./build_Test dir only.
	} else if r.Method == "POST" {
		// after form method="post" r.Method: "POST" r.URL.Path "/formAKpay" r.RequestURI "/formAKpay" r.Body []uint8 with e.g. (length: 20, cap: 512)
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body) // body will be []uint8 with e.g. (length: 20, cap: 512) err will be <nil>
		// e.g. [103 97 109 101 61 75 78 38 110 49 61 49 38 110 50 61 50 38 113 61] <nil>
		//fmt.Println(body, err)
		var strPage string = ""
		if err == nil {
			strBody := string(body) // e.g. "game=KN&n1=1&n2=2&q=" or e.g. "game=SL&n1=&n2=2&q=qwerty"
			//fmt.Println(strBody)
			//fmt.Fprintf(w, "Unimplemeted POST method")
			strPage = "<!DOCTYPE html>"
			strPage = strPage + "<html lang=\"en\">"
			strPage = strPage + "<head>"
			strPage = strPage + "<meta charset=\"utf-8\" />"
			strPage = strPage + "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />"
			strPage = strPage + "<title>Ticket info</title>"
			strPage = strPage + "<style>"
			strPage = strPage + "#ticinfo {"
			//strPage = strPage + "width: 70%;"
			strPage = strPage + "margin: 3% 3% 3% 3%;"
			strPage = strPage + "background-color: #dfdbdb;"
			strPage = strPage + "border: thick solid black;"
			strPage = strPage + "outline: dashed red;"
			strPage = strPage + "}"
			strPage = strPage + "#ticback {"
			strPage = strPage + "display: block;"
			strPage = strPage + "width: 10%;"
			strPage = strPage + "margin: 1% 3% 1% 3%;"
			strPage = strPage + "padding: 1% 1% 1% 1%;"
			strPage = strPage + "color: white;"
			strPage = strPage + "background-color: blue;"
			strPage = strPage + "border: thin solid black;"
			strPage = strPage + "border-radius: 15%;"
			strPage = strPage + "text-decoration:none;"
			strPage = strPage + "}"
			strPage = strPage + "#tichdr {"
			strPage = strPage + "margin: 1% 3% 1% 3%;"
			//strPage = strPage + "padding: 1% 1% 1% 1%;"
			strPage = strPage + "}"
			strPage = strPage + "#ticket {"
			strPage = strPage + "display: block;"
			strPage = strPage + "margin: 1% 3% 1% 3%;"
			strPage = strPage + "padding: 1% 1% 1% 1%;"
			strPage = strPage + "background-color: white;"
			strPage = strPage + "border: thin solid black;"
			strPage = strPage + "}"
			//strPage = strPage + ""
			strPage = strPage + "</style>"
			strPage = strPage + "</head>"
			strPage = strPage + "<body>"
			strPage = strPage + "<div id=\"ticinfo\">"
			strPage = strPage + "<a id=\"ticback\" href=\"/\">Back</a>"
			strPage = strPage + "<h3 id=\"tichdr\">POST method called with body:</h3>"
			strPage = strPage + "<ul id=\"ticket\">"
			strPage = strPage + strBody
			strPage = strPage + "</ul>"
			strPage = strPage + "</div>"
			strPage = strPage + "<script>"
			strPage = strPage + "console.log('page body script started');"
			strPage = strPage + "console.log(document.getElementById('ticket').innerHTML);"
			strPage = strPage + "if (document.getElementById('ticket').innerHTML === 'Ожидайте ответа сервера...') document.location.reload();"
			strPage = strPage + "</script>"
			strPage = strPage + "</body>"
			strPage = strPage + "</html>"
		} else {
			strPage = err.Error()
		}
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "%s", strPage)
	} else {
		fmt.Fprintf(w, "Unimplemeted method: %s", r.Method)
	}
	// End of processing POST method.
}

func buyTicket(strSearch string) string { // strSearch e.g. "6_1_1_a_04_05_09_12_34_51"
	var reqStringPay string = urlpay + strCmd(strSearch)
	txnid = txnid + 1
	fmt.Println(reqStringPay)
	var bytRep []byte
	var err error
	var strXML string = ""
	var errMsg string = ""
	var strPage string = ""
	var ticinfo string = ""
	var result string = ""
	var txnID string = ""
	var game string = ""
	var gameName string = ""
	var date string = ""
	var time string = ""
	var agent string = ""
	var numOfDraws string = ""
	var board = [10]string{"", "", "", "", "", "", "", "", "", ""}
	var bType = [10]string{"", "", "", "", "", "", "", "", "", ""}
	var i int
	var sum string = ""
	var number string = ""
	var gguard string = ""
	var firstDraw string = ""
	var system string = ""
	var numUsed string = ""
	var stake string = ""
	res, err := http.Get(reqStringPay)
	if err != nil {
		//log.Fatal(err)
		//return sum
		//errMsg = "Connection failed"
	} else {
		bytRep, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			//log.Fatal(err)
		} else { // begin response body processing.
			strXML = string(bytRep)
			fmt.Printf("%s\n", strXML)
			/*
				var pos1 int = -1
				var pos2 int = -1
				pos1 = strings.Index(strXML, "<sum>")
				pos2 = strings.Index(strXML, "</sum>")
				if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9) {
					sum = strXML[pos1+5 : pos2]
				}
			*/
			var pos1 int = -1
			var pos2 int = -1
			//<result>0</result>
			//01234567890
			pos1 = strings.Index(strXML, "<result>")
			pos2 = strings.Index(strXML, "</result>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
				errMsg = "Wrong server response XML format"
				goto ExitBuyTicket
			}
			result = strXML[pos1+8 : pos2]
			if result != "0" {
				errMsg = "Server reply with result: " + result
				goto ExitBuyTicket
			}
			//<txn_id>1</txn_id>
			//0123456789
			pos1 = strings.Index(strXML, "<txn_id>")
			pos2 = strings.Index(strXML, "</txn_id>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
				errMsg = "Server response wrong txn_id format"
				goto ExitBuyTicket
			}
			txnID = strXML[pos1+8 : pos2]
			//<game>6</game>
			//0123456789
			pos1 = strings.Index(strXML, "<game>")
			pos2 = strings.Index(strXML, "</game>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+7)) {
				errMsg = "Server response wrong game format"
				goto ExitBuyTicket
			}
			game = strXML[pos1+6 : pos2]
			//<date>25.10.18</date>
			//012345678901234
			pos1 = strings.Index(strXML, "<date>")
			pos2 = strings.Index(strXML, "</date>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
				errMsg = "Server response wrong date format"
				goto ExitBuyTicket
			}
			date = strXML[pos1+6 : pos2]
			//<time>13:12:58</time>
			//012345678901234
			pos1 = strings.Index(strXML, "<time>")
			pos2 = strings.Index(strXML, "</time>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
				errMsg = "Server response wrong time format"
				goto ExitBuyTicket
			}
			time = strXML[pos1+6 : pos2]
			//<agent>02501267</agent>
			//0123456789012345
			pos1 = strings.Index(strXML, "<agent>")
			pos2 = strings.Index(strXML, "</agent>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+15)) {
				errMsg = "Server response wrong agent format"
				goto ExitBuyTicket
			}
			agent = strXML[pos1+7 : pos2]
			//<num_of_draws>1</num_of_draws>
			//0123456789012345
			pos1 = strings.Index(strXML, "<num_of_draws>")
			pos2 = strings.Index(strXML, "</num_of_draws>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+15)) {
				errMsg = "Server response wrong num_of_draws format"
				goto ExitBuyTicket
			}
			numOfDraws = strXML[pos1+14 : pos2]
			//<sum>3.00</sum>
			//0123456789
			pos1 = strings.Index(strXML, "<sum>")
			pos2 = strings.Index(strXML, "</sum>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
				errMsg = "Server response wrong sum format"
				goto ExitBuyTicket
			}
			sum = strXML[pos1+5 : pos2]
			//<number>298-75699630-0976316</number>
			//01234567890123456789012345678
			pos1 = strings.Index(strXML, "<number>")
			pos2 = strings.Index(strXML, "</number>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+28)) {
				errMsg = "Server response wrong number format"
				goto ExitBuyTicket
			}
			number = strXML[pos1+8 : pos2]
			//<gguard>425377</gguard>
			//012345678901234
			pos1 = strings.Index(strXML, "<gguard>")
			pos2 = strings.Index(strXML, "</gguard>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
				errMsg = "Server response wrong gguard format"
				goto ExitBuyTicket
			}
			gguard = strXML[pos1+8 : pos2]
			if game == "5" || game == "6" {
				//<first_draw>27.10.18</first_draw> <first_draw>05.06.11 - 14.06.11</first_draw>
				//012345678901234567890
				pos1 = strings.Index(strXML, "<first_draw>")
				pos2 = strings.Index(strXML, "</first_draw>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+20)) {
					errMsg = "Server response wrong first_draw format"
					goto ExitBuyTicket
				}
				firstDraw = strXML[pos1+12 : pos2]
				//<system>7</system>
				//0123456789
				pos1 = strings.Index(strXML, "<system>")
				pos2 = strings.Index(strXML, "</system>")
				if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9) {
					system = strXML[pos1+8 : pos2]
				}
			}
			if game == "2" || game == "4" {
				//<stake>1</stake>
				//012345678
				pos1 = strings.Index(strXML, "<stake>")
				pos2 = strings.Index(strXML, "</stake>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+8)) {
					errMsg = "Server response wrong stake format"
					goto ExitBuyTicket
				}
				stake = strXML[pos1+7 : pos2]
			}
			if game == "2" {
				//<num_used>2</num_used>
				//012345678901
				pos1 = strings.Index(strXML, "<num_used>")
				pos2 = strings.Index(strXML, "</num_used>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
					errMsg = "Server response wrong num_used format"
					goto ExitBuyTicket
				}
				numUsed = strXML[pos1+10 : pos2]
			}
			//board = "01_80" // for test only.
			//================== boards processing ===================================
			//<board1a>01 02 03 04 05 06 07 08 09 10 11 12</board1a>
			//<board1>01 02 03 04 05 06 07 08 09 10 11 12</board1>
			//<board1>123S</board1>
			//<board1a>15 18</board1a>
			//<board1>15 18</board1>
			//<board1>123S</board1>
			//0123456789012
			if strings.Index(strXML, "<board1") == -1 {
				errMsg = "Server response no boards info"
				goto ExitBuyTicket
			}
			var maxBords int = 6
			if game == "2" {
				maxBords = 2
			} else if game == "4" {
				maxBords = 10
			}
			var strBoardBeg string = ""
			var strBoardEnd string = ""
			for i = 1; i <= maxBords; i = i + 1 {
				strBoardBeg = "<board" + strconv.Itoa(i) + ">"
				strBoardEnd = "</board" + strconv.Itoa(i) + ">"
				pos1 = strings.Index(strXML, strBoardBeg)
				pos2 = strings.Index(strXML, strBoardEnd)
				if pos1 == -1 {
					strBoardBeg = "<board" + strconv.Itoa(i) + "a>"
					strBoardEnd = "</board" + strconv.Itoa(i) + "a>"
					pos1 = strings.Index(strXML, strBoardBeg)
					pos2 = strings.Index(strXML, strBoardEnd)
					if (pos1 != -1) && (pos2 != -1) {
						bType[i-1] = "АВТО"
					}
				}
				if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+len(strBoardBeg)+2) {
					board[i-1] = strXML[pos1+len(strBoardBeg) : pos2]
				} else {
					break
				}
			} // end of boards processing.
		} // end of response body processing.
	} // end of readAll(res.Body).
	//return "Form formAKpay called with params " + reqStringPay
ExitBuyTicket:
	if err != nil {
		errMsg = err.Error()
	}
	if errMsg != "" {
		if errMsg == "Server reply with result: 777" {
			ticinfo = "Ожидайте ответа сервера..."
		} else {
			ticinfo = errMsg
		}
		goto buyTicketPage
	}
	_ = fmt.Sprint(agent, system, numUsed) // to avoid unused warning.
	ticinfo = ticinfo + "<li>Україньска Національна Лотерея</li>"
	ticinfo = ticinfo + "<li>Дата: " + date + "</li>"
	ticinfo = ticinfo + "<li>Час: " + time + "</li>"
	if game == "2" {
		gameName = "Кено"
	} else if game == "4" {
		gameName = "Трійка"
	} else if game == "5" {
		gameName = "Максима"
	} else if game == "6" {
		gameName = "Супер Лото"
	} else {
		gameName = "Неведома"
	}
	ticinfo = ticinfo + "<li>Гра: " + gameName + "</li>"
	if numOfDraws != "" {
		ticinfo = ticinfo + "<li>Розіграшей: " + numOfDraws + "</li>"
	}
	if firstDraw != "" {
		ticinfo = ticinfo + "<li>Розіграш: " + firstDraw + "</li>"
	} else {
		ticinfo = ticinfo + "<li>Перший розіграш: " + date + "</li>"
	}
	ticinfo = ticinfo + "<li>Комбінації:</li>"
	ticinfo = ticinfo + "<li>------------------------------------------------------</li>"
	for i = 0; i < len(board); i = i + 1 {
		if board[i] == "" {
			break
		}
		if game == "4" {
			// e.g. 012S or 112B or 765A or 123Y
			if board[i][3:] == "S" {
				ticinfo = ticinfo + "<li>" + board[i][0:3] + " Точний" + "</li>"
			} else if board[i][3:] == "B" {
				ticinfo = ticinfo + "<li>" + board[i][0:3] + " Довільний" + "</li>"
			} else if board[i][3:] == "A" {
				ticinfo = ticinfo + "<li>" + board[i][0:3] + " Точний + Довільний" + "</li>"
			} else {
				ticinfo = ticinfo + "<li>" + board[i][0:3] + " Система" + "</li>"
			}
		} else {
			ticinfo = ticinfo + "<li>" + board[i] + " " + bType[i] + "</li>"
		}
	}
	ticinfo = ticinfo + "<li>------------------------------------------------------</li>"
	if stake != "" {
		ticinfo = ticinfo + "<li>Ставка: " + stake + "</li>"
	}
	ticinfo = ticinfo + "<li>Сума: " + sum + "</li>"
	ticinfo = ticinfo + "<li>Білет: " + decrGG(gguard) + " " + decrNum(number) + "</li>"
	ticinfo = ticinfo + "<li>txn_id: " + txnID + "</li>"
	// Create strPage:
buyTicketPage:
	strPage = "<!DOCTYPE html>"
	strPage = strPage + "<html lang=\"en\">"
	strPage = strPage + "<head>"
	strPage = strPage + "<meta charset=\"utf-8\" />"
	strPage = strPage + "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />"
	strPage = strPage + "<title>Ticket info</title>"
	strPage = strPage + "<style>"
	strPage = strPage + "#ticinfo {"
	//strPage = strPage + "width: 70%;"
	strPage = strPage + "margin: 3% 3% 3% 3%;"
	strPage = strPage + "background-color: #dfdbdb;"
	strPage = strPage + "border: thick solid black;"
	strPage = strPage + "outline: dashed red;"
	strPage = strPage + "}"
	strPage = strPage + "#ticback {"
	strPage = strPage + "display: block;"
	strPage = strPage + "width: 10%;"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "color: white;"
	strPage = strPage + "background-color: blue;"
	strPage = strPage + "border: thin solid black;"
	strPage = strPage + "border-radius: 15%;"
	strPage = strPage + "text-decoration:none;"
	strPage = strPage + "}"
	strPage = strPage + "#tichdr {"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	//strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "}"
	strPage = strPage + "#ticket {"
	strPage = strPage + "display: block;"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "background-color: white;"
	strPage = strPage + "border: thin solid black;"
	strPage = strPage + "}"
	//strPage = strPage + ""
	strPage = strPage + "</style>"
	strPage = strPage + "</head>"
	strPage = strPage + "<body>"
	strPage = strPage + "<div id=\"ticinfo\">"
	strPage = strPage + "<a id=\"ticback\" href=\"/\">Back</a>"
	strPage = strPage + "<h3 id=\"tichdr\">Запрос на покупку билета</h3>"
	strPage = strPage + "<ul id=\"ticket\">"
	strPage = strPage + ticinfo
	strPage = strPage + "</ul>"
	strPage = strPage + "</div>"
	strPage = strPage + "<script>"
	strPage = strPage + "console.log('page body script started');"
	strPage = strPage + "console.log(document.getElementById('ticket').innerHTML);"
	strPage = strPage + "if (document.getElementById('ticket').innerHTML === 'Ожидайте ответа сервера...') document.location.reload();"
	strPage = strPage + "</script>"
	strPage = strPage + "</body>"
	strPage = strPage + "</html>"
	//fmt.Println(strPage)
	return strPage
} // end of buyTicket

func strCmd(ticreq string) string {
	// e.g. '?agent=16&type=2&command=pay&date=20201020&txn_id=' + txn_id + '&game=6&num_of_draws=1&num_of_boards=1&sum=15.00&msisdn=0';
	var boardKeno int = 10
	var boardMx int = 10
	var boardSl int = 15
	var boardTr int = 3
	// ticreq e.g. "6_1_1_a_10_19_27_34_49_50"
	var reqArr []string = strings.Split(ticreq, "_") //ticreq.split("_")
	fmt.Printf("%q\n", reqArr)
	// e.g. (10) ["6" "1" "1" "a" "10" "19" "27" "34" "49" "50"]
	//           [game, draws, stake, auto/manual, ...]
	var strSearch string = ""
	if len(reqArr) < 3 {
		return ""
	}
	//let dtVar = new Date();
	var dtVar time.Time = time.Now()
	year, month, day := dtVar.Date()
	var strYear string = strconv.Itoa(year)        // e.g. "2021"
	var strMonth string = strconv.Itoa(int(month)) // e.g. "1" or "12"
	var strDay string = strconv.Itoa(day)          // e.g. "1" or "31"
	//fmt.Println(strYear, strMonth, strDay)
	if len(strMonth) == 1 {
		strMonth = "0" + strMonth
	}
	if len(strDay) == 1 {
		strDay = "0" + strDay
	}
	//fmt.Println(strYear, strMonth, strDay)
	strSearch = "?agent=" + strAgent + "&type=2&command=pay&date=" + strYear + strMonth + strDay + "&txn_id=" + strconv.Itoa(txnid)
	strSearch = strSearch + "&game=" + reqArr[0] + "&num_of_draws=" + reqArr[1]
	//fmt.Println(strSearch)
	var i, j, n, b int
	var draws, stake int
	//var err error
	var k int
	var sum int = 0
	if reqArr[0] == "2" { // Keno.
		if len(reqArr) < 6 {
			return ""
		}
		// ["2" "1" "1" "a" "03" "13" "15" "25" "35" "36" "38" "57" "59" "78" "a" "09" "10" "22" "29" "49" "50" "56" "62" "63" "71"]
		n = 0 // num used.
		for i = 4; i < len(reqArr); i = i + 1 {
			if (reqArr[i] != "a") && (reqArr[i] != "m") {
				n = n + 1 // num used.
			} else {
				break
			}
		}
		// '&num_used=10&stake=2&board1=02_11_15_24_33_44_55_66_77_80&board2a=01_12_16_23_34_45_56_65_78_79&sum=84.00&msisdn=380121234567'
		strSearch = strSearch + "&num_used=" + strconv.Itoa(n)
		strSearch = strSearch + "&stake=" + reqArr[2]
		//fmt.Println(strSearch)
		// ["2" "1" "1" "a" "03" "13" "15" "25" "35" "36" "38" "57" "59" "78" "a" "09" "10" "22" "29" "49" "50" "56" "62" "63" "71"]
		b = 0 // board number.
		for i = 3; i < len(reqArr); i = i + (n + 1) {
			b = b + 1 // board number.
			if reqArr[i] == "a" {
				strSearch = strSearch + "&board" + strconv.Itoa(b) + "a="
			} else {
				strSearch = strSearch + "&board" + strconv.Itoa(b) + "="
			}
			for j = 0; j < n; j = j + 1 {
				strSearch = strSearch + reqArr[i+1+j] // get number of board field.
				if j != (n - 1) {
					strSearch = strSearch + "_"
				}
			}
			sum = sum + boardKeno
		}
		//draws, stake
		draws, _ = strconv.Atoi(reqArr[1]) // number of draws, _ is used to ignore err return value.
		stake, _ = strconv.Atoi(reqArr[2]) // stake, _ is used to ignore err return value.
		sum = sum * draws * stake
		strSearch = strSearch + "&sum=" + strconv.Itoa(sum) + ".00"
		strSearch = strSearch + "&msisdn=0" // very last item.
		fmt.Println(strSearch)
		return strSearch // end of Keno.
	} else if reqArr[0] == "4" { // Triyka.
		if len(reqArr) < 7 {
			return ""
		}
		// e.g. ["4" "1" "1" "S" "1" "1" "1" "B" "1" "2" "2" "A" "1" "2" "3" "Y" "0" "1" "3"] or type B or A or Y.
		//           [game, draws, stake, type, ...]
		n = 3 // num used.
		// '&stake=2&board1=123S&board2=112B&board3=123A&board4=023Y&sum=60.00&msisdn=380503332211'
		strSearch = strSearch + "&stake=" + reqArr[2]
		b = 0
		var strCmb = [3]string{"0", "0", "0"} // array literal form used.
		for i = 3; i < len(reqArr); i = i + (n + 1) {
			b = b + 1
			strSearch = strSearch + "&board" + strconv.Itoa(b) + "="
			strSearch = strSearch + reqArr[i+1] + reqArr[i+2] + reqArr[i+3] + reqArr[i]
			strCmb[0] = reqArr[i+1]
			strCmb[1] = reqArr[i+2]
			strCmb[2] = reqArr[i+3]
			sum = sum + cmbPriceCalcType(strCmb, reqArr[i], boardTr)
		}
		//draws, stake
		draws, _ = strconv.Atoi(reqArr[1]) // number of draws, _ is used to ignore err return value.
		stake, _ = strconv.Atoi(reqArr[2]) // stake, _ is used to ignore err return value.
		sum = sum * draws * stake          // // number of draws and stake.
		strSearch = strSearch + "&sum=" + strconv.Itoa(sum) + ".00"
		strSearch = strSearch + "&msisdn=0" // very last item.
		fmt.Println(strSearch)
		return strSearch // end of Triyka.
	} else if reqArr[0] == "5" { // Maxima.
		if len(reqArr) < 9 {
			return ""
		}
		if (reqArr[3] == "sa") || (reqArr[3] == "sm") { // system.
			k = 0 // system number.
			for i = 4; i < len(reqArr); i = i + 1 {
				k = k + 1
			}
			strSearch = strSearch + "&system=" + strconv.Itoa(k)
			if reqArr[3] == "sa" {
				strSearch = strSearch + "&board1a="
			} else {
				strSearch = strSearch + "&board1="
			}
			for i = 4; i < len(reqArr); i = i + 1 {
				strSearch = strSearch + reqArr[i]
				if i != (len(reqArr) - 1) {
					strSearch = strSearch + "_"
				}
			}
			sum = sysCmbMX(k) * boardMx
			// end of system case.
		} else { // not system case
			k = 0
			for i = 3; i < len(reqArr); i = i + 6 {
				k = k + 1
				if reqArr[i] == "a" {
					strSearch = strSearch + "&board" + strconv.Itoa(k) + "a="
				} else {
					strSearch = strSearch + "&board" + strconv.Itoa(k) + "="
				}
				strSearch = strSearch + reqArr[i+1] + "_" + reqArr[i+2] + "_" + reqArr[i+3] + "_"
				strSearch = strSearch + reqArr[i+4] + "_" + reqArr[i+5]
				sum = sum + boardMx
			}
		} // end of not system.
		draws, _ = strconv.Atoi(reqArr[1]) // number of draws, _ is used to ignore err return value.
		sum = sum * draws
		strSearch = strSearch + "&sum=" + strconv.Itoa(sum) + ".00"
		strSearch = strSearch + "&msisdn=0" // very last item.
		return strSearch                    // // end of Maxima.
	} else if reqArr[0] == "6" { // Super Loto.
		if len(reqArr) < 10 {
			return ""
		}
		if (reqArr[3] == "sa") || (reqArr[3] == "sm") { // system.
			k = 0 // system number.
			for i = 4; i < len(reqArr); i = i + 1 {
				k = k + 1
			}
			strSearch = strSearch + "&system=" + strconv.Itoa(k)
			if reqArr[3] == "sa" {
				strSearch = strSearch + "&board1a="
			} else {
				strSearch = strSearch + "&board1="
			}
			for i = 4; i < len(reqArr); i = i + 1 {
				strSearch = strSearch + reqArr[i]
				if i != (len(reqArr) - 1) {
					strSearch = strSearch + "_"
				}
			}
			sum = sysCmbSL(k) * boardSl
			// end of system case.
		} else { // not system case
			k = 0
			for i = 3; i < len(reqArr); i = i + 7 {
				k = k + 1
				if reqArr[i] == "a" {
					strSearch = strSearch + "&board" + strconv.Itoa(k) + "a="
				} else {
					strSearch = strSearch + "&board" + strconv.Itoa(k) + "="
				}
				strSearch = strSearch + reqArr[i+1] + "_" + reqArr[i+2] + "_" + reqArr[i+3] + "_"
				strSearch = strSearch + reqArr[i+4] + "_" + reqArr[i+5] + "_" + reqArr[i+6]
				sum = sum + boardSl
			}
		} // end of not system.
		draws, _ = strconv.Atoi(reqArr[1]) // number of draws, _ is used to ignore err return value.
		sum = sum * draws
		strSearch = strSearch + "&sum=" + strconv.Itoa(sum) + ".00"
		strSearch = strSearch + "&msisdn=0" // very last item.
		return strSearch                    // end of Super Loto.
	} else {
		return ""
	}
}

// cmb, typ, boardPrice
func cmbPriceCalcType(cmb [3]string, typ string, boardPrice int) int {
	var nums int = 0
	var cmbprice int = 0
	nums = uniqnums(cmb)
	if typ == "S" {
		cmbprice = cmbprice + boardPrice
	} else if typ == "B" {
		if (nums == 2) || (nums == 3) {
			cmbprice = cmbprice + boardPrice
		} else {
			cmbprice = 0
		}
	} else if typ == "A" {
		if (nums == 2) || (nums == 3) {
			cmbprice = cmbprice + 2*boardPrice
		} else {
			cmbprice = 0
		}
	} else if typ == "Y" {
		if nums == 2 {
			cmbprice = cmbprice + 3*boardPrice
		} else if nums == 3 {
			cmbprice = cmbprice + 6*boardPrice
		} else {
			cmbprice = 0
		}
	}
	return cmbprice
}

func uniqnums(cmb [3]string) int {
	var nums int = 1
	if cmb[0] != cmb[1] {
		nums = nums + 1
	}
	if cmb[0] != cmb[2] {
		nums = nums + 1
	}
	if (nums > 1) && (cmb[1] == cmb[2]) {
		nums = nums - 1
	}
	return nums
}

func sysCmbMX(sys int) int {
	if sys == 6 {
		return 6
	} else if sys == 7 {
		return 21
	} else if sys == 8 {
		return 56
	} else if sys == 9 {
		return 126
	} else if sys == 10 {
		return 252
	} else if sys == 11 {
		return 462
	} else { // sys == 12
		return 792
	}
}

func sysCmbSL(sys int) int {
	if sys == 7 {
		return 7
	} else if sys == 8 {
		return 28
	} else if sys == 9 {
		return 84
	} else if sys == 10 {
		return 210
	} else if sys == 11 {
		return 462
	} else { // sys == 12
		return 924
	}
}

func getResults(strSearch string) string {
	// http://10.8.194.3:38000/?agent=16&type=2&command=scheck&game=5&draw_results
	// http://10.8.194.3:38000/?agent=16&type=2&command=scheck&game=5&draw_results&draw=1
	var strPage string = ""
	var reptext string = "Wrong results request format"
	var game string = ""
	var gamename string = ""
	var draw string = ""
	var drawnum = [2]string{"", ""}
	var date = [2]string{"", ""}
	var winnignumbers = [2]string{"", ""}
	var reqStringResults string = ""
	var bytRep []byte
	var err error
	var res *http.Response
	var errMsg string = ""
	var strXML string = ""
	var result string = ""
	var pos1 int = -1
	var pos2 int = -1
	// strSearch e.g. 2&d= or 2&d=1001
	pos1 = 0
	pos2 = strings.Index(strSearch, "&d=")
	if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+1) {
		game = strSearch[pos1:pos2]
	} else {
		goto getResultsCreatePage
	}
	pos1 = strings.Index(strSearch, "&d=")
	if pos1 != -1 {
		draw = strSearch[pos1+3:]
	} else {
		goto getResultsCreatePage
	}
	reptext = ""
	reqStringResults = urlpay
	if draw == "" {
		reqStringResults = reqStringResults + "?agent=16&type=2&command=scheck&game=" + game + "&draw_results"
	} else {
		reqStringResults = reqStringResults + "?agent=16&type=2&command=scheck&game=" + game + "&draw_results&draw=" + draw
	}
	res, err = http.Get(reqStringResults)
	if err != nil {
		//log.Fatal(err)
		//return sum
		//errMsg = "Connection failed"
	} else {
		bytRep, err = ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			//log.Fatal(err)
		} else { // begin response body processing.
			strXML = string(bytRep)
			var pos1 int = -1
			var pos2 int = -1
			//<result>0</result>
			//01234567890
			pos1 = strings.Index(strXML, "<result>")
			pos2 = strings.Index(strXML, "</result>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
				errMsg = "Wrong server response XML format"
				goto ExitGetResults
			}
			result = strXML[pos1+8 : pos2]
			if result != "0" {
				errMsg = "Server reply with result: " + result
				goto ExitGetResults
			}
			//<game>2</game>
			//01234567
			pos1 = strings.Index(strXML, "<game>")
			pos2 = strings.Index(strXML, "</game>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+7)) {
				errMsg = "Wrong server response game format"
				goto ExitGetResults
			}
			game = strXML[pos1+6 : pos2]
			if game == "2" {
				gamename = "Кено"
			} else if game == "4" {
				gamename = "Трійка"
			} else if game == "5" {
				gamename = "Лото Максима"
			} else if game == "6" {
				gamename = "Супер Лото"
			}
			//<draw_num>1</draw_num>
			//012345678901
			pos1 = strings.Index(strXML, "<draw_num>")
			pos2 = strings.Index(strXML, "</draw_num>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
				errMsg = "Wrong server response draw_num format"
				goto ExitGetResults
			}
			drawnum[0] = strXML[pos1+10 : pos2]
			//<date>24.10.18</date>
			//012345678901234
			pos1 = strings.Index(strXML, "<date>")
			pos2 = strings.Index(strXML, "</date>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
				errMsg = "Wrong server response date format"
				goto ExitGetResults
			}
			date[0] = strXML[pos1+6 : pos2]
			//<winnig_numbers>123</winnig_numbers>
			//01234567890123456789
			pos1 = strings.Index(strXML, "<winnig_numbers>")
			pos2 = strings.Index(strXML, "</winnig_numbers>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+19)) {
				errMsg = "Wrong server response winnig_numbers format"
				goto ExitGetResults
			}
			winnignumbers[0] = strXML[pos1+16 : pos2]
			// Previous draw info if any.
			//<previous_draw>
			//0123456789
			pos1 = strings.Index(strXML, "<previous_draw>")
			if pos1 == -1 {
				goto ExitGetResults
			}
			//<draw_num>1</draw_num>
			//012345678901
			pos1 = strings.LastIndex(strXML, "<draw_num>")
			pos2 = strings.LastIndex(strXML, "</draw_num>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
				errMsg = "Wrong server response previous draw_num format"
				goto ExitGetResults
			}
			drawnum[1] = strXML[pos1+10 : pos2]
			//<date>24.10.18</date>
			//012345678901234
			pos1 = strings.LastIndex(strXML, "<date>")
			pos2 = strings.LastIndex(strXML, "</date>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
				errMsg = "Wrong server response previous date format"
				goto ExitGetResults
			}
			date[1] = strXML[pos1+6 : pos2]
			//<winnig_numbers>123</winnig_numbers>
			//01234567890123456789
			pos1 = strings.LastIndex(strXML, "<winnig_numbers>")
			pos2 = strings.LastIndex(strXML, "</winnig_numbers>")
			if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+19)) {
				errMsg = "Wrong server response previous winnig_numbers format"
				goto ExitGetResults
			}
			winnignumbers[1] = strXML[pos1+16 : pos2]
		} // end of response body processing.
	} // end of connection processing.
ExitGetResults:
	if err != nil {
		errMsg = err.Error()
	}
	if errMsg != "" {
		reptext = errMsg
		goto getResultsCreatePage
	}
	reptext = reptext + "<li>Україньска Національна Лотерея</li>"
	reptext = reptext + "<li>Результаты</li>"
	reptext = reptext + "<li>Гра: " + gamename + "</li>"
	reptext = reptext + "<li>Розіграш: " + drawnum[0] + "</li>"
	reptext = reptext + "<li>Дата: " + date[0] + "</li>"
	reptext = reptext + "<li>Номера: " + winnignumbers[0] + "</li>"
	if drawnum[1] != "" {
		reptext = reptext + "<li>Розіграш: " + drawnum[1] + "</li>"
		reptext = reptext + "<li>Дата: " + date[1] + "</li>"
		reptext = reptext + "<li>Номера: " + winnignumbers[1] + "</li>"
	}
	// Create Page:
getResultsCreatePage:
	strPage = strPage + "<!DOCTYPE html>"
	strPage = strPage + "<html lang=\"en\">"
	strPage = strPage + "<head>"
	strPage = strPage + "<meta charset=\"utf-8\"/>"
	strPage = strPage + "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>"
	strPage = strPage + "<title>Draw info</title>"
	strPage = strPage + "<style>"
	strPage = strPage + "#ticinfo {"
	//strPage = strPage + "width: 70%;"
	strPage = strPage + "margin: 3% 3% 3% 3%;"
	strPage = strPage + "background-color: #dfdbdb;"
	strPage = strPage + "border: thick solid black;"
	strPage = strPage + "outline: dashed red;"
	strPage = strPage + "}"
	strPage = strPage + "#ticback {"
	strPage = strPage + "display: block;"
	strPage = strPage + "width: 10%;"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "color: white;"
	strPage = strPage + "background-color: blue;"
	strPage = strPage + "border: thin solid black;"
	strPage = strPage + "border-radius: 15%;"
	strPage = strPage + "text-decoration:none;"
	strPage = strPage + "}"
	strPage = strPage + "#tichdr {"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	//strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "}"
	strPage = strPage + "#ticket {"
	strPage = strPage + "display: block;"
	strPage = strPage + "margin: 1% 3% 1% 3%;"
	strPage = strPage + "padding: 1% 1% 1% 1%;"
	strPage = strPage + "background-color: white;"
	strPage = strPage + "border: thin solid black;"
	strPage = strPage + "}"
	//strPage = strPage + ""
	strPage = strPage + "</style>"
	strPage = strPage + "</head>"
	strPage = strPage + "<body>"
	strPage = strPage + "<div id=\"ticinfo\">"
	strPage = strPage + "<a id=\"ticback\" href=\"/\">Back</a>"
	strPage = strPage + "<ul id=\"ticket\">"
	strPage = strPage + reptext
	strPage = strPage + "</ul>"
	strPage = strPage + "</div>"
	//strPage = strPage + "<script>"
	//strPage = strPage + "console.log("page body script started"");
	//strPage = strPage + "console.log(document.getElementById("ticket").innerHTML);");
	//strPage = strPage + "if (document.getElementById("ticket").innerHTML === "Ожидайте ответа сервера...") document.location.reload();");
	//strPage = strPage + "</script>"
	strPage = strPage + "</body>"
	strPage = strPage + "</html>"
	return strPage
}

func checkValTicket(strTicnum string, val string) string {
	//var err error
	//var errMsg string = "Unknown error"
	if val == "N" {
		var errMsg string = ""
		var result string = ""
		var sum string = "-999" // initial as error.
		var bytRep []byte
		res, err := http.Get(reqStringCheckVal + strTicnum)
		if err != nil {
			//log.Fatal(err)
			//return sum
			//errMsg = "Connection failed"
		} else {
			bytRep, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				//log.Fatal(err)
				//return sum
				//errMsg = "XML reply reading failed"
			} else {
				strXML := string(bytRep)
				fmt.Printf("%s\n", strXML)
				var pos1 int = -1
				var pos2 int = -1
				//<result>0</result>
				//01234567890
				pos1 = strings.Index(strXML, "<result>")
				pos2 = strings.Index(strXML, "</result>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
					errMsg = "Wrong server response XML format"
					goto ExitValChkTicket
				}
				result = strXML[pos1+8 : pos2] // get result string numerical code.
				if result != "0" {
					errMsg = "Server reply with result: " + result
					goto ExitValChkTicket
				}
				pos1 = strings.Index(strXML, "<sum>")
				pos2 = strings.Index(strXML, "</sum>")
				if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9) {
					sum = strXML[pos1+5 : pos2]
				}
			}
		}
	ExitValChkTicket:
		//ticinfo = "Выплата пока не реализована."
		var ticinfo string = ""
		if err != nil {
			errMsg = err.Error()
		}
		if errMsg != "" {
			if errMsg == "Server reply with result: 777" {
				ticinfo = "Ожидайте ответа сервера..."
			} else {
				ticinfo = errMsg
			}
			goto valChkTicketPage
		}
		if sum == "-1.00" {
			ticinfo = "Большой выигрыш!!!."
		} else if sum == "-2.00" {
			ticinfo = "Билет уже выплачен."
		} else if sum == "-3.00" {
			ticinfo = "Билет выплачен с обменным билетом."
		} else if sum == "-4.00" {
			ticinfo = "Билет аннулирован."
		} else if sum == "0.00" {
			ticinfo = "Билет не выиграл."
		} else if sum == "-999" {
			ticinfo = "Ошибка при обработке запроса - " + err.Error()
		} else {
			//ticinfo = "Ваш виграш ${sum} грн."
			ticinfo = "Ваш виграш " + sum + " грн."
		}
		//fmt.Fprintf(w, "Form formAKchk called with ticket number %s reply as %s", strTicnum, ticinfo)
	valChkTicketPage:
		var strPage string = ""
		strPage = strPage + "<!DOCTYPE html>"
		strPage = strPage + "<html lang=\"en\">"
		strPage = strPage + "<head>"
		strPage = strPage + "<meta charset=\"utf-8\" />"
		strPage = strPage + "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />"
		strPage = strPage + "<title>Ticket info</title>"
		strPage = strPage + "<style>"
		strPage = strPage + "#ticinfo {"
		//strPage = strPage + "width: 70%;"
		strPage = strPage + "margin: 3% 3% 3% 3%;"
		strPage = strPage + "background-color: #dfdbdb;"
		strPage = strPage + "border: thick solid black;"
		strPage = strPage + "outline: dashed red;"
		strPage = strPage + "}"
		strPage = strPage + "#ticback {"
		strPage = strPage + "display: block;"
		strPage = strPage + "width: 10%;"
		strPage = strPage + "margin: 3% 3% 3% 3%;"
		strPage = strPage + "padding: 1% 1% 1% 1%;"
		strPage = strPage + "color: white;"
		strPage = strPage + "background-color: blue;"
		strPage = strPage + "border: thin solid black;"
		strPage = strPage + "border-radius: 15%;"
		strPage = strPage + "text-decoration:none;"
		strPage = strPage + "}"
		strPage = strPage + "#ticket {"
		strPage = strPage + "display: block;"
		strPage = strPage + "margin: 3% 3% 3% 3%;"
		strPage = strPage + "padding: 1% 1% 1% 1%;"
		strPage = strPage + "background-color: white;"
		strPage = strPage + "border: thin solid black;"
		strPage = strPage + "}"
		//strPage = strPage + ""
		strPage = strPage + "</style>"
		strPage = strPage + "</head>"
		strPage = strPage + "<body>"
		strPage = strPage + "<div id=\"ticinfo\">"
		strPage = strPage + "<a id=\"ticback\" href=\"/?t=" + strTicnum + "\">Back</a>"
		strPage = strPage + "<p id=\"ticket\">"
		strPage = strPage + ticinfo
		strPage = strPage + "</p>"
		strPage = strPage + "</div>"
		strPage = strPage + "<script>"
		strPage = strPage + "console.log('page body script started');"
		strPage = strPage + "console.log(document.getElementById('ticket').innerHTML);"
		strPage = strPage + "if (document.getElementById('ticket').innerHTML === 'Ожидайте ответа сервера...') document.location.reload();"
		strPage = strPage + "</script>"
		strPage = strPage + "</body>"
		strPage = strPage + "</html>"
		//strPage = strPage + ""
		return strPage
	} else { // if val == "Y" begin of validation case
		txnid = txnid + 1
		var strSearch string = ""
		var dtVar time.Time = time.Now()
		year, month, day := dtVar.Date()
		var strYear string = strconv.Itoa(year)        // e.g. "2021"
		var strMonth string = strconv.Itoa(int(month)) // e.g. "1" or "12"
		var strDay string = strconv.Itoa(day)          // e.g. "1" or "31"
		//fmt.Println(strYear, strMonth, strDay)
		if len(strMonth) == 1 {
			strMonth = "0" + strMonth
		}
		if len(strDay) == 1 {
			strDay = "0" + strDay
		}
		//fmt.Println(strYear, strMonth, strDay)
		// http://unl.works:38000/?agent=16&type=2&command=validation&date=20220601&txn_id=101&ticket_number=007-13725021-1736928&msisdn=0
		strSearch = "?agent=" + strAgent + "&type=2&command=validation&date=" + strYear + strMonth + strDay + "&txn_id=" + strconv.Itoa(txnid)
		strSearch = strSearch + "&ticket_number=" + strTicnum + "&msisdn=0"
		var reqStringVal string = urlpay + strSearch
		fmt.Println(reqStringVal)
		var bytRep []byte
		var err error
		var strXML string = ""
		var errMsg string = ""
		var result string = ""
		var txnID string = ""
		var wingame string = ""
		var game string = ""
		var gameName string = ""
		var date string = ""
		var time string = ""
		var agent string = ""
		var numOfDraws string = ""
		var board = [10]string{"", "", "", "", "", "", "", "", "", ""}
		var bType = [10]string{"", "", "", "", "", "", "", "", "", ""}
		var i int
		var sum string = ""
		var number string = ""
		var gguard string = ""
		var firstDraw string = ""
		var system string = ""
		var numUsed string = ""
		var stake string = ""
		var paid string = ""
		var tax string = ""
		var war_tax string = ""
		var exchange bool = false
		var sumex string = ""
		var numberex string = ""
		var pos1 int = -1
		var pos2 int = -1
		var ticinfo string = ""
		var strPage string = ""
		res, err := http.Get(reqStringVal)
		if err != nil {
			// err.Error() msg will be used later.
			//log.Fatal(err)
			//return sum
			//errMsg = "Connection failed"
		} else { // begin readAll(res.Body)
			bytRep, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				// err.Error() msg will be used later.
				//log.Fatal(err)
			} else { // begin response body processing.
				strXML = string(bytRep)
				fmt.Printf("%s\n", strXML)
				//<result>0</result>
				//01234567890
				pos1 = strings.Index(strXML, "<result>")
				pos2 = strings.Index(strXML, "</result>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
					errMsg = "Wrong server response XML format"
					goto ExitValTicket
				}
				result = strXML[pos1+8 : pos2] // get result string numerical code.
				if result != "0" {
					if result == "900" {
						errMsg = "Билет не выиграл."
					} else if result == "901" {
						errMsg = "Большой выигрыш!!!."
					} else if result == "902" {
						errMsg = "Билет уже выплачен."
					} else if result == "903" {
						//errMsg = "Билет выплачен с ошибкой в обменном билете." // NB! resut will be used later.
						goto ContWinTicProc
					} else {
						errMsg = "Server reply with result: " + result
					}
					goto ExitValTicket
				}
			ContWinTicProc:
				//ticinfo = strXML // for test only
				//<txn_id>1</txn_id>
				//0123456789
				pos1 = strings.Index(strXML, "<txn_id>")
				pos2 = strings.Index(strXML, "</txn_id>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
					errMsg = "Server response wrong txn_id format"
					goto ExitValTicket
				}
				txnID = strXML[pos1+8 : pos2]
				//<game>6</game>
				//0123456789
				pos1 = strings.Index(strXML, "<win_game>")
				pos2 = strings.Index(strXML, "</win_game>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+7)) {
					errMsg = "Server response wrong win game format"
					goto ExitValTicket
				}
				wingame = strXML[pos1+10 : pos2]
				//<number>298-75699630-0976316</number>
				//01234567890123456789012345678
				pos1 = strings.Index(strXML, "<number>")
				pos2 = strings.Index(strXML, "</number>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+28)) {
					errMsg = "Server response wrong number format"
					goto ExitValTicket
				}
				number = strXML[pos1+8 : pos2]
				//<sum>3.00</sum>
				//0123456789
				pos1 = strings.Index(strXML, "<sum>")
				pos2 = strings.Index(strXML, "</sum>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
					errMsg = "Server response wrong sum format"
					goto ExitValTicket
				}
				sum = strXML[pos1+5 : pos2]
				//<paid>3.00</paid>
				//0123456789
				pos1 = strings.Index(strXML, "<paid>")
				pos2 = strings.Index(strXML, "</paid>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+10)) {
					errMsg = "Server response wrong paid format"
					goto ExitValTicket
				}
				paid = strXML[pos1+6 : pos2]
				// <tax>5.76</tax>
				// 01234
				pos1 = strings.Index(strXML, "<tax>")
				pos2 = strings.Index(strXML, "</tax>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
					errMsg = "Server response wrong tax format"
					goto ExitValTicket
				}
				tax = strXML[pos1+5 : pos2]
				// <war_tax>0.48</war_tax>
				// 0123456789012
				pos1 = strings.Index(strXML, "<war_tax>")
				pos2 = strings.Index(strXML, "</war_tax>")
				if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+13)) {
					errMsg = "Server response wrong tax format"
					goto ExitValTicket
				}
				war_tax = strXML[pos1+9 : pos2]
				// <exchange_ticket>
				// 01234567890123456
				pos1 = strings.Index(strXML, "<exchange_ticket>")
				pos2 = strings.Index(strXML, "</exchange_ticket>")
				if (pos1 != -1) && !(result == "903") { // exchange ticket exists
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+17)) {
						errMsg = "Server response wrong exchange ticket format"
						goto ExitValTicket
					}
					exchange = true
					//<game>6</game>
					//0123456789
					pos1 = strings.Index(strXML, "<game>")
					pos2 = strings.Index(strXML, "</game>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+7)) {
						errMsg = "Server response wrong exchange game format"
						goto ExitValTicket
					}
					game = strXML[pos1+6 : pos2]
					//<date>25.10.18</date>
					//012345678901234
					pos1 = strings.Index(strXML, "<date>")
					pos2 = strings.Index(strXML, "</date>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
						errMsg = "Server response wrong exchange date format"
						goto ExitValTicket
					}
					date = strXML[pos1+6 : pos2]
					//<time>13:12:58</time>
					//012345678901234
					pos1 = strings.Index(strXML, "<time>")
					pos2 = strings.Index(strXML, "</time>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
						errMsg = "Server response wrong exchange time format"
						goto ExitValTicket
					}
					time = strXML[pos1+6 : pos2]
					//<agent>02501267</agent>
					//0123456789012345
					pos1 = strings.Index(strXML, "<agent>")
					pos2 = strings.Index(strXML, "</agent>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+15)) {
						errMsg = "Server response wrong exchange agent format"
						goto ExitValTicket
					}
					agent = strXML[pos1+7 : pos2]
					//<num_of_draws>1</num_of_draws>
					//0123456789012345
					pos1 = strings.Index(strXML, "<num_of_draws>")
					pos2 = strings.Index(strXML, "</num_of_draws>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+15)) {
						errMsg = "Server response wrong exchange num_of_draws format"
						goto ExitValTicket
					}
					numOfDraws = strXML[pos1+14 : pos2]
					//<gguard>425377</gguard>
					//012345678901234
					pos1 = strings.Index(strXML, "<gguard>")
					pos2 = strings.Index(strXML, "</gguard>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+14)) {
						errMsg = "Server response wrong exchange gguard format"
						goto ExitValTicket
					}
					gguard = strXML[pos1+8 : pos2]
					//<sum>3.00</sum>
					//0123456789
					pos1 = strings.LastIndex(strXML, "<sum>")
					pos2 = strings.LastIndex(strXML, "</sum>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9)) {
						errMsg = "Server response wrong exchange sum format"
						goto ExitValTicket
					}
					sumex = strXML[pos1+5 : pos2]
					//<number>298-75699630-0976316</number>
					//01234567890123456789012345678
					pos1 = strings.LastIndex(strXML, "<number>")
					pos2 = strings.LastIndex(strXML, "</number>")
					if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+28)) {
						errMsg = "Server response wrong exchange number format"
						goto ExitValTicket
					}
					numberex = strXML[pos1+8 : pos2]
					if game == "5" || game == "6" {
						//<first_draw>27.10.18</first_draw> <first_draw>05.06.11 - 14.06.11</first_draw>
						//012345678901234567890
						pos1 = strings.Index(strXML, "<first_draw>")
						pos2 = strings.Index(strXML, "</first_draw>")
						if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+20)) {
							errMsg = "Server response wrong exchange first_draw format"
							goto ExitValTicket
						}
						firstDraw = strXML[pos1+12 : pos2]
						//<system>7</system>
						//0123456789
						pos1 = strings.Index(strXML, "<system>")
						pos2 = strings.Index(strXML, "</system>")
						if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+9) {
							system = strXML[pos1+8 : pos2]
						}
					}
					if game == "2" || game == "4" {
						//<stake>1</stake>
						//012345678
						pos1 = strings.Index(strXML, "<stake>")
						pos2 = strings.Index(strXML, "</stake>")
						if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+8)) {
							errMsg = "Server response wrong exchange stake format"
							goto ExitValTicket
						}
						stake = strXML[pos1+7 : pos2]
					}
					if game == "2" {
						//<num_used>2</num_used>
						//012345678901
						pos1 = strings.Index(strXML, "<num_used>")
						pos2 = strings.Index(strXML, "</num_used>")
						if !((pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+11)) {
							errMsg = "Server response wrong exchange num_used format"
							goto ExitValTicket
						}
						numUsed = strXML[pos1+10 : pos2]
					}
					//board = "01_80" // for test only.
					//================== boards processing ===================================
					//<board1a>01 02 03 04 05 06 07 08 09 10 11 12</board1a>
					//<board1>01 02 03 04 05 06 07 08 09 10 11 12</board1>
					//<board1>123S</board1>
					//<board1a>15 18</board1a>
					//<board1>15 18</board1>
					//<board1>123S</board1>
					//0123456789012
					if strings.Index(strXML, "<board1") == -1 {
						errMsg = "Server response no boards info"
						goto ExitValTicket
					}
					var maxBords int = 6
					if game == "2" {
						maxBords = 2
					} else if game == "4" {
						maxBords = 10
					}
					var strBoardBeg string = ""
					var strBoardEnd string = ""
					for i = 1; i <= maxBords; i = i + 1 { // begin of boards processing
						strBoardBeg = "<board" + strconv.Itoa(i) + ">"
						strBoardEnd = "</board" + strconv.Itoa(i) + ">"
						pos1 = strings.Index(strXML, strBoardBeg)
						pos2 = strings.Index(strXML, strBoardEnd)
						if pos1 == -1 {
							strBoardBeg = "<board" + strconv.Itoa(i) + "a>"
							strBoardEnd = "</board" + strconv.Itoa(i) + "a>"
							pos1 = strings.Index(strXML, strBoardBeg)
							pos2 = strings.Index(strXML, strBoardEnd)
							if (pos1 != -1) && (pos2 != -1) {
								bType[i-1] = "АВТО"
							}
						}
						if (pos1 != -1) && (pos2 != -1) && (pos2 >= pos1+len(strBoardBeg)+2) {
							board[i-1] = strXML[pos1+len(strBoardBeg) : pos2]
						} else {
							break
						}
					} // end of boards processing.
				} // end of exchange ticket exists.
			} // end of response body processing.
		} // end readAll(res.Body).
	ExitValTicket:
		//ticinfo = "Выплата пока не реализована."
		if err != nil {
			errMsg = err.Error()
		}
		if errMsg != "" {
			if errMsg == "Server reply with result: 777" {
				ticinfo = "Ожидайте ответа сервера..."
			} else {
				ticinfo = errMsg
			}
			goto valTicketPage
		}
		ticinfo = ""
		ticinfo = ticinfo + "<li>Україньска Національна Лотерея</li>"
		if wingame == "2" {
			gameName = "Кено"
		} else if game == "4" {
			gameName = "Трійка"
		} else if game == "5" {
			gameName = "Максима"
		} else if game == "6" {
			gameName = "Супер Лото"
		} else {
			gameName = "Неведома"
		}
		ticinfo = ticinfo + "<li>Білет: " + number + "</li>"
		ticinfo = ticinfo + "<li>Гра: " + gameName + "</li>"
		ticinfo = ticinfo + "<li>Сума виграша: " + sum + "</li>"
		ticinfo = ticinfo + "<li>К выплате: " + paid + "</li>"
		ticinfo = ticinfo + "<li>Налог: " + tax + "</li>"
		ticinfo = ticinfo + "<li>Военный налог: " + war_tax + "</li>"
		ticinfo = ticinfo + "<li>txn_id: " + txnID + "</li>"
		_ = fmt.Sprint(agent, system, numUsed) // to avoid unused warnings.
		if exchange && !(result == "903") {    // ticinfo exchange case
			ticinfo = ticinfo + "<li>Обменный билет</li>"
			ticinfo = ticinfo + "<li>Дата: " + date + "</li>"
			ticinfo = ticinfo + "<li>Час: " + time + "</li>"
			if numOfDraws != "" {
				ticinfo = ticinfo + "<li>Розіграшей: " + numOfDraws + "</li>"
			}
			if firstDraw != "" {
				ticinfo = ticinfo + "<li>Розіграш: " + firstDraw + "</li>"
			} else {
				ticinfo = ticinfo + "<li>Перший розіграш: " + date + "</li>"
			}
			ticinfo = ticinfo + "<li>Комбінації:</li>"
			ticinfo = ticinfo + "<li>------------------------------------------------------</li>"
			for i = 0; i < len(board); i = i + 1 {
				if board[i] == "" {
					break
				}
				if game == "4" {
					// e.g. 012S or 112B or 765A or 123Y
					if board[i][3:] == "S" {
						ticinfo = ticinfo + "<li>" + board[i][0:3] + " Точний" + "</li>"
					} else if board[i][3:] == "B" {
						ticinfo = ticinfo + "<li>" + board[i][0:3] + " Довільний" + "</li>"
					} else if board[i][3:] == "A" {
						ticinfo = ticinfo + "<li>" + board[i][0:3] + " Точний + Довільний" + "</li>"
					} else {
						ticinfo = ticinfo + "<li>" + board[i][0:3] + " Система" + "</li>"
					}
				} else {
					ticinfo = ticinfo + "<li>" + board[i] + " " + bType[i] + "</li>"
				}
			}
			ticinfo = ticinfo + "<li>------------------------------------------------------</li>"
			if stake != "" {
				ticinfo = ticinfo + "<li>Ставка: " + stake + "</li>"
			}
			ticinfo = ticinfo + "<li>Сума: " + sumex + "</li>"
			ticinfo = ticinfo + "<li>Білет: " + decrGG(gguard) + " " + decrNum(numberex) + "</li>"
		} // end of ticinfo exchange case.
		// Create strPage:
	valTicketPage:
		//ticinfo = "Выплата пока не реализована."
		strPage = "<!DOCTYPE html>"
		strPage = strPage + "<html lang=\"en\">"
		strPage = strPage + "<head>"
		strPage = strPage + "<meta charset=\"utf-8\" />"
		strPage = strPage + "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />"
		strPage = strPage + "<title>Ticket info</title>"
		strPage = strPage + "<style>"
		strPage = strPage + "#ticinfo {"
		//strPage = strPage + "width: 70%;"
		strPage = strPage + "margin: 3% 3% 3% 3%;"
		strPage = strPage + "background-color: #dfdbdb;"
		strPage = strPage + "border: thick solid black;"
		strPage = strPage + "outline: dashed red;"
		strPage = strPage + "}"
		strPage = strPage + "#ticback {"
		strPage = strPage + "display: block;"
		strPage = strPage + "width: 10%;"
		strPage = strPage + "margin: 1% 3% 1% 3%;"
		strPage = strPage + "padding: 1% 1% 1% 1%;"
		strPage = strPage + "color: white;"
		strPage = strPage + "background-color: blue;"
		strPage = strPage + "border: thin solid black;"
		strPage = strPage + "border-radius: 15%;"
		strPage = strPage + "text-decoration:none;"
		strPage = strPage + "}"
		strPage = strPage + "#tichdr {"
		strPage = strPage + "margin: 1% 3% 1% 3%;"
		//strPage = strPage + "padding: 1% 1% 1% 1%;"
		strPage = strPage + "}"
		strPage = strPage + "#ticket {"
		strPage = strPage + "display: block;"
		strPage = strPage + "margin: 1% 3% 1% 3%;"
		strPage = strPage + "padding: 1% 1% 1% 1%;"
		strPage = strPage + "background-color: white;"
		strPage = strPage + "border: thin solid black;"
		strPage = strPage + "}"
		//strPage = strPage + ""
		strPage = strPage + "</style>"
		strPage = strPage + "</head>"
		strPage = strPage + "<body>"
		strPage = strPage + "<div id=\"ticinfo\">"
		strPage = strPage + "<a id=\"ticback\" href=\"/?t=" + strTicnum + "\">Back</a>"
		strPage = strPage + "<h3 id=\"tichdr\">Выплата билета</h3>"
		strPage = strPage + "<ul id=\"ticket\">"
		strPage = strPage + ticinfo
		strPage = strPage + "</ul>"
		strPage = strPage + "</div>"
		strPage = strPage + "<script>"
		strPage = strPage + "console.log('page body script started');"
		strPage = strPage + "console.log(document.getElementById('ticket').innerHTML);"
		strPage = strPage + "if (document.getElementById('ticket').innerHTML === 'Ожидайте ответа сервера...') document.location.reload();"
		strPage = strPage + "</script>"
		strPage = strPage + "</body>"
		strPage = strPage + "</html>"
		//strPage = strPage + ""
		return strPage
	} // end of validation case.
} // end of checkValTicket

func main() {
	// time.Now().UnixNano(), which yields a constantly-changing number.
	rand.Seed(time.Now().UnixNano())
	txnid = txnid + rand.Intn(100000)
	fmt.Println("Server starts with txnid=", txnid)
	fmt.Println("ListenAndServe http://localhost:8080/")
	http.HandleFunc("/", handlerReq)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decrNum(strEnc string) string {
	// Old format: decr(strEnc string, blnOdd bool) -> blnOdd is False for Even julian (0), or True for Odd julian (1).
	// strEnc e.g. "051-52749309-5226022"
	//             "01234567890123456789"
	var strDec string = ""
	var strTic string = ""
	var blnOdd bool = false
	var err error
	var intVar int = 0
	if len(strEnc) != 20 {
		return ""
	}
	if strEnc[3:4] != "-" || strEnc[12:13] != "-" {
		return ""
	}
	intVar, err = strconv.Atoi(strEnc[0:3])
	if err != nil {
		return ""
	}
	strTic = strTic + strEnc[0:3] + "-"
	if intVar%2 == 0 {
		blnOdd = false
	} else {
		blnOdd = true
	}
	_, err = strconv.Atoi(strEnc[4:12]) // Use _ to ignore returning value.
	if err != nil {
		return ""
	}
	csEnc := C.CString(strEnc[4:12])
	//intLen := C.strlenM(cs)
	//csDec := C.doDecrypt(csEnc, 0) // non needed and can not be freed later.
	if blnOdd {
		strDec = C.GoString(C.doDecrypt(csEnc, 1)) // second parameter is 1 (true) for Odd julian.
	} else {
		strDec = C.GoString(C.doDecrypt(csEnc, 0)) // second parameter is 0 (false) for Even julian.
	}
	C.free(unsafe.Pointer(csEnc))
	if strDec == "" {
		return ""
	}
	strTic = strTic + strDec + "-"
	_, err = strconv.Atoi(strEnc[13:]) // Use _ to ignore returning value.
	if err != nil {
		return ""
	}
	csEnc = C.CString(strEnc[13:])
	if blnOdd {
		strDec = C.GoString(C.doDecrypt(csEnc, 1)) // second parameter is 1 (true) for Odd julian.
	} else {
		strDec = C.GoString(C.doDecrypt(csEnc, 0)) // second parameter is 0 (false) for Even julian.
	}
	C.free(unsafe.Pointer(csEnc))
	if strDec == "" {
		return ""
	}
	strTic = strTic + strDec
	//intLen := C.strlenM(cs)
	//csDec := C.doDecrypt(csEnc, 0) // non needed and can not be freed later.
	//fmt.Println(strDec)
	//strDec = C.GoString(C.doDecrypt(csEnc, 1)) // Odd julian.
	//fmt.Println(strDec)
	//C.free(unsafe.Pointer(csEnc))
	//C.free(unsafe.Pointer(csDec)) // can not be freed later.
	return strTic
}

func decrGG(strEnc string) string {
	// strEnc e.g. "775388"
	//             "012345"
	var strDec string = ""
	//var err error
	//var intVar int = 0
	if len(strEnc) != 6 {
		return ""
	}
	_, err := strconv.Atoi(strEnc[0:]) // Use _ to ignore returning value.
	if err != nil {
		return ""
	}
	//intVar = intVar + 0 // artificially use intVar for successful compilation.
	csEnc := C.CString(strEnc[0:])
	strDec = C.GoString(C.doDecrypt(csEnc, 0)) // non needed and can not be freed later.
	C.free(unsafe.Pointer(csEnc))
	if strDec == "" {
		return ""
	}
	return strDec
}

func mainTest() {
	var strEnc string
	var strDec string = ""
	// ./ unld 775388 051-52749309-5226022
	//012969 051 -02831601-0108302
	//strEnc = "52749309"
	//strDec = decr(strEnc, true) // Second parameter is False for Even julian, or True for Odd julian.
	strEnc = "775388"
	strDec = decrGG(strEnc) // Second parameter is False for Even julian, or True for Odd julian.
	fmt.Println(strDec)
	strEnc = "051-52749309-5226022"
	strDec = decrNum(strEnc) // Second parameter is False for Even julian, or True for Odd julian.
	fmt.Println(strDec)
	// ./unld 785391 132-55516848-4260932
	//015079 132 - 01079742 - 1780861
	//strEnc = "55516848"
	//strDec = decr(strEnc, false) // Second parameter is False for Even julian, or True for Odd julian.
	strEnc = "785391"
	strDec = decrGG(strEnc) // Second parameter is False for Even julian, or True for Odd julian.
	fmt.Println(strDec)
	strEnc = "132-55516848-4260932"
	strDec = decrNum(strEnc) // Second parameter is False for Even julian, or True for Odd julian.
	fmt.Println(strDec)
}
