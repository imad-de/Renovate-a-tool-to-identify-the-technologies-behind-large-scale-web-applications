package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"

	"text/tabwriter"

	"github.com/PuerkitoBio/goquery"
	"github.com/akamensky/argparse"

	"github.com/tealeg/xlsx"
)

func main() {
	// Create new parser object
	parser := argparse.NewParser("print", "Prints provided string to stdout")

	//add arguments
	u := parser.String("u", "string", &argparse.Options{Required: false, Help: "put URL"})
	l := parser.File("l", "log-file", os.O_RDWR, 0600, &argparse.Options{Required: false, Help: "put file"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	file_open(*u, l)
	// host(u, l)
	// url_form(*u, l)

}

// have valid format url
func formaturl(u string, l *os.File) string {
	var url string

	format := "https://"
	if strings.Contains(u, "http") {
		url = u
	} else {
		url = format + u

	}

	return url

}

// to display host ip from domain name
func ip_url(u string, l *os.File) string {
	var ur string
	ips, _ := net.LookupIP(host_validate(&u, l))
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			// fmt.Println("IPv4: ", ipv4)
			ur = ipv4.String()
		}
	}

	return ur
}

// display domain name
func host_validate(u *string, l *os.File) string {
	var s string
	var test string
	var ok string
	s = *u

	//remove "http://" from the url
	if strings.Contains(s, "http://") {
		s = strings.Replace(s, "http://", "", 3)

	}
	//remove "https://" from the url
	if strings.Contains(s, "https://") {
		s = strings.Replace(s, "https://", "", 3)

	}

	// display the domain name
	imad := strings.Split(s, "/")

	// remove www.
	if strings.Contains(imad[0], "www") {
		format := "."

		count := strings.Count(imad[0], ".")
		if count == 3 {
			imad = strings.Split(imad[0], ".")

			test = imad[1] + format + imad[2] + format + imad[3]
		}
		if count == 2 {
			imad = strings.Split(imad[0], ".")

			test = imad[1] + format + imad[2]

		}

		return test

	} else {

		ok = imad[0]

		return ok
	}
}

// for display url without directory(for Sharepoint in func new_cms)
func host(u *string, l *os.File) string {
	var s string
	var test string
	s = *u

	//remove "http://" from the url
	if strings.Contains(s, "http") {
		imad := strings.Split(s, "/")
		// fmt.Println(imad[2])
		test = "https://" + imad[2]
		// fmt.Println(test)

	} else {
		imad := strings.Split(s, "/")
		test = "https://" + imad[0]

	}

	return test
}

// open file and print the final result on terminal and file.xlsx
func file_open(u string, l *os.File) {
	var w *tabwriter.Writer

	//display in file.xlsx
	filee := xlsx.NewFile()
	sheet, err := filee.AddSheet("technology detection")
	if err != nil {
		panic(err.Error())
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "url"
	cell = row.AddCell()
	cell.Value = "Valid_url"

	cell = row.AddCell()
	cell.Value = "ip adress"
	cell = row.AddCell()
	cell.Value = "server"
	cell = row.AddCell()
	cell.Value = "framework"
	cell = row.AddCell()
	cell.Value = "cms"

	err = filee.Save("re.xlsx")
	if err != nil {
		panic(err.Error())
	}

	//open file
	file, err := os.Open("urls.txt")

	// Reading from a file using scanner.
	scanner := bufio.NewScanner(file)

	var cnt = 1
	var txtlines []string

	//scanne the file
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	if err != nil {
		return
	}
	defer file.Close()

	fmt.Println("      +----------+-------------------------------------------------------------------------------+-------------------+--------------------------+------------------------------------------------------------+---------------------------------+---------------------------------+	")

	fmt.Println("      |     ID   |                            url                                                |        valid url  |           ip adress      |                       server                               |                framework        |                cms              |	")
	fmt.Println("      +----------|-------------------------------------------------------------------------------|-------------------|---------------+----------|------------------------------------------------------------|---------------------------------|---------------------------------|	")

	for _, eachline := range txtlines {

		u := eachline
		if err != nil {
			continue
		}

		resp, err := http.Get(formaturl(u, l))
		if err != nil {
			const padding = 6
			w = tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

			fmt.Fprintln(w, "           \t                                                                         \t             \t                    \t                                                      \t                           \t                           \t")

			fmt.Fprintln(w, "     ", cnt, "\t"+u+"\tFalse      \t \t\t\t\t")
			fmt.Fprintln(w, "      |----------|-------------------------------------------------------------------------------|-------------------|--------------------------|------------------------------------------------------------|---------------------------------|---------------------------------|")
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.Value = u
			cell = row.AddCell()
			cell.Value = "False"
			w.Flush()
			cnt++

		} else if resp.StatusCode == 522 {
			const padding = 6
			w = tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)

			fmt.Fprintln(w, "           \t                                                                         \t             \t                    \t                                                      \t                           \t                           \t")

			fmt.Fprintln(w, "     ", cnt, "\t"+u+"\tFalse      \t \t\t\t\t")
			fmt.Fprintln(w, "      |----------|-------------------------------------------------------------------------------|-------------------|--------------------------|------------------------------------------------------------|---------------------------------|---------------------------------|")
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.Value = u
			cell = row.AddCell()
			cell.Value = "False"

			err = filee.Save("re.xlsx")
			if err != nil {
				panic(err.Error())
			}
			w.Flush()
			cnt++
		} else {
			response2 := resp.Header.Get("server")

			response3 := resp.Header.Get("x-aspnet-version")

			// if err != nil {

			// }

			// display the server
			var server string
			if response2 != "" {
				server = response2
			} else {
				server = "cant'identify"
			}

			//display the CMS
			var cms string

			cms = new_cms(u, l)
			if cms == "" {
				cms = "can't identify"

			}

			// display the framework
			var framework string

			if Detect_PHP(u, l) {

				if headers(u, l) != "" {
					framework = headers(u, l)

				} else {
					framework = "PHP"
				}

			} else if Detect_ASP(u, l) {
				if response3 != "" {
					framework = " (aspnet/" + response3 + ")"

				} else if headers(u, l) != "" {
					framework = headers(u, l)

				} else {
					framework = "ASP.net"

				}

			} else if strings.Contains(cms, "Drupal") || strings.Contains(cms, "SPIP") || strings.Contains(cms, "Joomla") || strings.Contains(cms, "WordPress") || strings.Contains(cms, "Magento") {

				framework = "PHP"

			} else if strings.Contains(cms, "SharePoint") || strings.Contains(cms, "DotNetNuke") {
				framework = "ASP.net"

			} else {
				framework = "can't identify"

			}

			const padding = 6
			w = tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.AlignRight|tabwriter.Debug)
			fmt.Fprintln(w, "           \t                                                                         \t             \t                    \t                                                      \t                           \t                           \t")

			fmt.Fprintln(w, "     ", cnt, "\t"+u+"\tTrue      \t "+ip_url(u, l)+"\t"+server+"\t"+framework+"\t"+cms+"\t")
			fmt.Fprintln(w, "      |----------|-------------------------------------------------------------------------------|-------------------|--------------------------|------------------------------------------------------------|---------------------------------|---------------------------------|")
			row := sheet.AddRow()
			cell := row.AddCell()
			cell.Value = u
			cell = row.AddCell()
			cell.Value = "True"

			cell = row.AddCell()
			cell.Value = ip_url(u, l)
			cell = row.AddCell()
			cell.Value = server
			cell = row.AddCell()
			cell.Value = framework
			cell = row.AddCell()
			cell.Value = cms

			err = filee.Save("re.xlsx")
			if err != nil {
				panic(err.Error())
			}
			w.Flush()

			cnt++
		}

	}

	if err := recover(); err != nil {

	}

}

// print the header X-Powered-by
func headers(u string, l *os.File) string {
	urltest := formaturl(u, l)
	var response2 string

	resp, err := http.Get(urltest)
	if err == nil {
		response2 = resp.Header.Get("X-Powered-By")

		defer resp.Body.Close()

	} else {
		response2 = "no such host"
	}

	return response2

}

// find if the website use php
func php_html(u string, l *os.File) bool {

	doc, err := goquery.NewDocument(formaturl(u, l))
	var result []string
	var cnt = 0
	if err != nil {

	} else {

		doc.Find("[href]").Each(func(index int, item *goquery.Selection) {
			href, _ := item.Attr("href")

			// href that contains "https://" or "http://"
			if strings.Contains(href, "https://") || strings.Contains(href, "http://") {

				// display just repertory of the url
				href = strings.Replace(href, "https://", "", 3)
				href = strings.Replace(href, "http://", "", 3)
				imad := strings.Split(href, "/")

				result = imad[1:]

				justString := strings.Join(result, " ")

				// check if .php exist
				if strings.Contains(justString, "php") {
					cnt++

				}

			} else {
				// href that don't contains "https://" or "http://"
				if strings.Contains(href, "php") {
					cnt++

				}

			}

		})
	}
	// if cnt>0 so the web site use php
	if cnt > 0 {
		return true
	}
	return false

}

// find if the website use asp.net
func asp_html(u string, l *os.File) bool {
	doc, err := goquery.NewDocument(formaturl(u, l))
	var result []string
	var cnt = 0
	if err != nil {
	} else {

		doc.Find("[href]").Each(func(index int, item *goquery.Selection) {
			href, _ := item.Attr("href")

			// href that contains "https://"or "http://"
			if strings.Contains(href, "https://") || strings.Contains(href, "http://") {

				// display just repertory of the url
				href = strings.Replace(href, "https://", "", 3)
				href = strings.Replace(href, "http://", "", 3)

				imad := strings.Split(href, "/")

				result = imad[1:]

				justString := strings.Join(result, " ")

				// check if .aspx exist
				if strings.Contains(justString, ".aspx") {
					cnt++
				}

			} else {
				// href that don't contains "https://"or "http://"
				if strings.Contains(href, ".aspx") {
					cnt++

				}
			}
		})
	}
	// if cnt>0 so the web site use asp.net
	if cnt > 0 {

		return true
	}
	return false

}

//if {headers() return contain "php" or php_html() return true} return true
func Detect_PHP(u string, l *os.File) bool {

	if strings.Contains(headers(u, l), "PHP") {
		return true

	} else {
		if php_html(u, l) {
			return true

		}
	}

	return false
}

//if {headers() return contain "asp" or asp_html() return true} return true
func Detect_ASP(u string, l *os.File) bool {

	resp, err := http.Get(formaturl(u, l))
	if err != nil {

	}
	response3 := resp.Header.Get("x-aspnet-version")

	if strings.Contains(headers(u, l), "ASP") || response3 != "" {
		return true
	} else {
		if asp_html(u, l) {
			return true

		}
	}

	return false

}

// detect witch cms is used and the version
func new_cms(u string, l *os.File) string {
	var test string
	var href string
	doc, err := goquery.NewDocument(formaturl(u, l))
	if err != nil {
	} else {

		doc.Find("[content]").Each(func(index int, item *goquery.Selection) {
			href, _ = item.Attr("content")

			if strings.Contains(href, "SharePoint") {

				test = "SharePoint"

			} else if strings.Contains(href, "DotNetNuke") {
				test = "DotNetNuke"

			} else if strings.Contains(href, "SPIP") {

				test = "SPIP"

			} else if strings.Contains(href, "Drupal") {

				imad := strings.Split(href, "(")
				imad = imad[:1]

				test = strings.Join(imad, ", ")

			} else if strings.Contains(href, "Joomla") {

				test = "Joomla"

			} else if strings.Contains(href, "WordPress") {
				if strings.HasPrefix(href, "WordPress") {

					test = href

				}

			}

		})
		// check for the version of joomla
		if test == "" || strings.Contains(test, "Joomla") {
			var url string
			if strings.HasSuffix(formaturl(u, l), "/") {
				url = formaturl(u, l) + "administrator/manifests/files/joomla.xml"

			} else {
				url = formaturl(u, l) + "/administrator/manifests/files/joomla.xml"

			}
			url = formaturl(u, l) + "/administrator/manifests/files/joomla.xml"

			resp, err := http.Get(url)
			if err != nil {
			} else {
				defer resp.Body.Close()
				html, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				b := string([]byte(html))
				if strings.Contains(b, "Joomla! Core") {

					re := regexp.MustCompile(`<version>(.*)</version>`)
					match := re.FindStringSubmatch(b)
					if len(match) > 1 {
						test = "Joomla " + match[1]
					} else {
					}
				}
			}

		}

		// check for the version of Sharepoint
		if test == "SharePoint" {
			var url string

			url = host(&u, l) + "/_vti_pvt/service.cnf"

			resp, err := http.Get(url)
			if err != nil {
			} else {
				defer resp.Body.Close()
				html, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				b := string([]byte(html))
				if strings.Contains(b, "vti_extenderversion:SR|14.") {
					test = "SharePoint 2010"
				} else if strings.Contains(b, "vti_extenderversion:SR|15.") {
					test = "SharePoint 2013"
				} else if strings.Contains(b, "vti_extenderversion:SR|16.0.0.4") || strings.Contains(b, "vti_extenderversion:SR|16.0.0.5") {
					test = "SharePoint 2016"
				} else if strings.Contains(b, "vti_extenderversion:SR|16.0.0.1") {
					test = "SharePoint 2019"
				}
			}

		}

		// check for the version of drupal
		if test == "" || strings.Contains(test, "Drupal") {
			var url string
			if strings.HasSuffix(formaturl(u, l), "/") {
				url = formaturl(u, l) + "CHANGELOG.txt"

			} else {
				url = formaturl(u, l) + "/CHANGELOG.txt"

			}

			resp, err := http.Get(url)
			if err != nil {
			} else {
				defer resp.Body.Close()
				html, err := ioutil.ReadAll(resp.Body)
				if err != nil {
				}
				b := string([]byte(html))
				if strings.HasPrefix(b, "Drupal") {

					user := b[:strings.IndexByte(b, ',')]
					test = user

				}
			}

		}
		if test == "" {

			doc.Find("[class]").Each(func(index int, item *goquery.Selection) {
				href, _ := item.Attr("class")
				if strings.Contains(href, "spip_") {
					test = "SPIP"

				}

			})
		}

		// check for the version of magento
		if test == "" {
			var url string
			if strings.HasSuffix(formaturl(u, l), "/") {
				url = formaturl(u, l) + "magento_version"

			} else {
				url = formaturl(u, l) + "/magento_version"

			}

			resp, err := http.Get(url)
			if err != nil {
			} else {
				defer resp.Body.Close()
				html, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					panic(err)
				}
				b := string([]byte(html))

				if strings.Contains(b, "html") {

				} else {
					if strings.Contains(b, "Magento") {
						test = b

					}

				}
			}

		}
		if test == "" {

			doc.Find("[type]").Each(func(index int, item *goquery.Selection) {
				href, _ := item.Attr("type")
				if strings.Contains(href, "magento-init") {
					test = "Magento"
				}

			})
		}

		if test == "" {

			doc.Find("[id]").Each(func(index int, item *goquery.Selection) {
				href, _ := item.Attr("id")
				if strings.Contains(href, "dnn_") {
					test = "DotNetNuke"
				}

			})
		}
		if test == "" {

			doc.Find("[href]").Each(func(index int, item *goquery.Selection) {
				href, _ := item.Attr("href")

				// check for this string in href

				if strings.Contains(href, "/Portals/0") || strings.Contains(href, "DesktopModules") {
					test = "DotNetNuke"
				} else if strings.Contains(href, "SharePoint") || strings.Contains(href, "_layouts") {
					test = "SharePoint"
				} else if strings.Contains(href, "wp-content") || strings.Contains(href, "wp-includes") || strings.Contains(href, "wp-admin") {
					test = "WordPress"
				} else if strings.Contains(href, "joomla") {
					test = "Joomla"
				} else if strings.Contains(href, "drupal") {
					test = "Drupal"
				} else if strings.Contains(href, "Magento") {
					test = "Magento"
				} else if strings.Contains(href, "spip.php") {
					test = "SPIP"
				}

			})
		}

		if test == "" {
			doc.Find("[src]").Each(func(index int, item *goquery.Selection) {
				href, _ := item.Attr("src")

				// check for this string in src

				if strings.Contains(href, "/Portals/0") || strings.Contains(href, "DesktopModules") {
					test = "DotNetNuke"
				} else if strings.Contains(href, "SharePoint") || strings.Contains(href, "_layouts") {
					fmt.Println("test", href)
					test = "SharePoint"
				} else if strings.Contains(href, "wp-content") || strings.Contains(href, "wp-includes") || strings.Contains(href, "wp-admin") {
					test = "WordPress"
				} else if strings.Contains(href, "joomla") {
					test = "Joomla"
				} else if strings.Contains(href, "drupal") {
					test = "Drupal"
				} else if strings.Contains(href, "Magento") {
					test = "Magento"
				} else if strings.Contains(href, "spip.php") {
					test = "SPIP"
				}

			})

		}

	}
	return test
}

//if {cms_new return wordpress} return true
func wordpress(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "WordPress") {

		return true
	}

	return false

}

//if {cms_new return joomla} return true
func joomla(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "Joomla") {

		return true
	}

	return false

}

//if {cms_new return drupal} return true
func drupal(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "Drupal") {

		return true
	}

	return false

}

//if {cms_new return magento} return true
func magento(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "Magento") {

		return true
	}

	return false
}

//if {cms_new return spip} return true
func SPIP(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "SPIP") {

		return true
	}

	return false

}

//if {cms_new return DotNetNuke} return true
func DotNetNuke(u string, l *os.File) bool {

	if strings.Contains(new_cms(u, l), "DotNetNuke") {

		return true
	}

	return false

}

//if {cms_new return SharePoint} return true
func SharePoint(u string, l *os.File) bool {
	if strings.Contains(new_cms(u, l), "SharePoint") {

		return true
	}

	return false
}
