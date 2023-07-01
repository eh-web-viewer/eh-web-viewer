package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/antchfx/htmlquery"
	"github.com/eh-web-viewer/eh-web-viewer/mycurl"
)

func Test_Idx(t *testing.T) {
	mycurl.SetClient("proxy")
	query := `f_cats=767&f_search=tsukiyo`
	resp, _ := mycurl.Fetch("GET", BASE_URL+"?"+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	respText, _ := io.ReadAll(resp.Body)
	respStr := string(respText)
	log.Println(respStr)
	log.Println(reResults.FindString(respStr))
	log.Println(reNextPage.FindString(respStr))
	log.Println(rePrevPage.FindString(respStr)) // 麻了，传回去js里面处理吧。

	doc, err := htmlquery.Parse(strings.NewReader(respStr))
	if err != nil {
		t.Error(err)
	}
	list, err := htmlquery.QueryAll(doc, "//div[@class='gl1t']")
	if err != nil {
		t.Error(err)
	}
	log.Println(list)
}
func Test_fetch(t *testing.T) {
	mycurl.SetClient("proxy")
	query := `f_cats=767&f_search=tsukiyo`
	resp, _ := mycurl.Fetch("GET", BASE_URL+"?"+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	respText, _ := io.ReadAll(resp.Body)
	respStr := string(respText)
	fmt.Println(respStr)
}
func Test_Htmlquery(t *testing.T) {
	// mycurl.SetClient("proxy")
	// query := `f_cats=767&f_search=tsukiyo`
	// resp, _ := mycurl.Fetch("GET", BASE_URL+"?"+query,
	// 	map[string]string{"Cookie": COOKIE},
	// 	nil)
	// respText, _ := io.ReadAll(resp.Body)
	respStr := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
	<head>
	<title>ExHentai.org</title>
	<link rel="stylesheet" type="text/css" href="https://exhentai.org/z/0363/x.css" />
	<style type="text/css">
	@supports(display:grid){.gld{display:grid;grid-template-columns:repeat(5,1fr)}.gl1t{min-width:250px;max-width:400px}@media screen and (max-width:1360px) {.gld{grid-template-columns:repeat(4,1fr)}.gl1t:nth-child(8n+1),.gl1t:nth-child(8n+3),.gl1t:nth-child(8n+6),.gl1t:nth-child(8n+8){background:#363940}.gl1t:nth-child(8n+2),.gl1t:nth-child(8n+4),.gl1t:nth-child(8n+5),.gl1t:nth-child(8n+7){background:#3c414b}}@media screen and (max-width:1090px) {.gld{grid-template-columns:repeat(3,1fr)}.gl1t:nth-child(2n+1){background:#363940}.gl1t:nth-child(2n+2){background:#3c414b}}@media screen and (max-width:820px) {.gld{grid-template-columns:repeat(2,1fr)}.gl1t:nth-child(4n+1),.gl1t:nth-child(4n+4){background:#363940}.gl1t:nth-child(4n+2),.gl1t:nth-child(4n+3){background:#3c414b}}}
	</style>
	</head>
	<body>
	<script type="text/javascript">
	function popUp(URL,w,h) {
			window.open(URL,"_pu"+(Math.random()+"").replace(/0\./,""),"toolbar=0,scrollbars=0,location=0,statusbar=0,menubar=0,resizable=0,width="+w+",height="+h+",left="+((screen.width-w)/2)+",top="+((screen.height-h)/2));
			return false;
	}
	</script>
	<div id="nb" class="nosel"><div><a class="nbw" href="https://exhentai.org/">Front<span class="nbw1"> Page</span></a></div><div><a href="https://exhentai.org/watched">Watched</a></div><div><a href="https://exhentai.org/popular">Popular</a></div><div><a href="https://exhentai.org/torrents.php">Torrents</a></div><div><a href="https://exhentai.org/favorites.php">Fav<span class="nbw1">orite</span>s</a></div><div><a href="https://exhentai.org/uconfig.php">Settings</a></div><div><a href="https://upld.exhentai.org/upld/manage"><span class="nbw2">My </span>Uploads</a></div><div><a href="https://exhentai.org/mytags">My Tags</a></div></div>
	<div class="ido" style="max-width:1370px">
	<div id="toppane">
	<h1 class="ih">ExHentai.org - <a href="http://exhentai55ld2wyap5juskbm67czulomrouspdacjamjeloj7ugjbsad.onion">Now With Layers</a> &nbsp;<a href="https://en.wikipedia.org/wiki/Tor_(anonymity_network)">[?]</a></h1><div id="searchbox" class="idi"><form action="https://exhentai.org/" method="get" style="margin:0px; padding:0px"><input type="hidden" id="f_cats" name="f_cats" value="767" /><table class="itc"><tr><td><div id="cat_2" class="cs ct2" onclick="toggle_category(2)" data-disabled="1">Doujinshi</div></td><td><div id="cat_4" class="cs ct3" onclick="toggle_category(4)" data-disabled="1">Manga</div></td><td><div id="cat_8" class="cs ct4" onclick="toggle_category(8)" data-disabled="1">Artist CG</div></td><td><div id="cat_16" class="cs ct5" onclick="toggle_category(16)" data-disabled="1">Game CG</div></td><td><div id="cat_512" class="cs cta" onclick="toggle_category(512)" data-disabled="1">Western</div></td></tr><tr><td><div id="cat_256" class="cs ct9" onclick="toggle_category(256)">Non-H</div></td><td><div id="
	cat_32" class="cs ct6" onclick="toggle_category(32)" data-disabled="1">Image Set</div></td><td><div id="cat_64" class="cs ct7" onclick="toggle_category(64)" data-disabled="1">Cosplay</div></td><td><div id="cat_128" class="cs ct8" onclick="toggle_category(128)" data-disabled="1">Asian Porn</div></td><td><div id="cat_1" class="cs ct1" onclick="toggle_category(1)" data-disabled="1">Misc</div></td></tr></table><div><input type="text" id="f_search" name="f_search" placeholder="Search Keywords" value="tsukiyo" size="90" maxlength="200" /><input type="submit" value="Search" onclick="search_presubmit()" /><input type="button" value="Clear" onclick="top.location.href='https://exhentai.org/'; return false" /></div><div>[<a href="#" onclick="toggle_advsearch_pane(this); return false">Show Advanced Options</a>] &nbsp; &nbsp;[<a href="#" onclick="toggle_filesearch_pane(this); return false">Show File Search</a>]</div><div id="advdiv" style="display:none"></div></form></div><div id="fsdiv" class="idi" style="margin-top:10px
	; display:none"></div>
	<script type="text/javascript" src="https://exhentai.org/z/0363/ehg_index.c.js"></script>
	
	<script type="text/javascript">
	var ulhost = "https://upld.exhentai.org/upld/";
	</script>
	
	</div><div style="position:relative; z-index:2"><div id="rangebar"></div><div class="searchtext"><p>Found about 40 results. </p></div>
	<script type="text/javascript">
	var prevurl="";
	var nexturl="https://exhentai.org/?f_search=tsukiyo&f_cats=767&next=892481";
	var maxdate="2023-05-20";
	var mindate="2007-03-20";
	var rangeurl="https://exhentai.org/?f_search=tsukiyo&amp;f_cats=767";
	var rangemin=0;
	var rangemax=54;
	var rangespan=2;
	build_rangebar();
	</script>
	<div class="searchnav">
			<div></div>
			<div><span id="ufirst">&lt;&lt; First</span></div><div><span id="uprev">&lt; Prev</span></div><div id="ujumpbox" class="jumpbox"><a id="ujump" href="javascript:enable_jump_mode('u')">Jump/Seek</a></div><div><a id="unext" href="https://exhentai.org/?f_search=tsukiyo&amp;f_cats=767&amp;next=892481">Next &gt;</a></div><div><a id="ulast" href="https://exhentai.org/?f_search=tsukiyo&amp;f_cats=767&amp;prev=1">Last &gt;&gt;</a></div>
			<div><select onchange="document.location='https://exhentai.org/?f_search=tsukiyo&amp;f_cats=767&amp;inline_set=dm_'+this.value+''"><option value="m">Minimal</option><option value="p">Minimal+</option><option value="l">Compact</option><option value="e">Extended</option><option value="t" selected="selected">Thumbnail</option></select></div>
	</div><div class="itg gld"><div class="gl1t"><a href="https://exhentai.org/g/2557923/cd8d42df44/"><div class="gl4t glname glink">[Artist] Tsukiyo</div></a><div class="gl3t" style="height:340px;width:223px"><a href="https://exhentai.org/g/2557923/cd8d42df44/"><img style="height:375px;width:223px;top:-17px" alt="[Artist] Tsukiyo" title="[Artist] Tsukiyo" src="https://s.exhentai.org/t/d3/41/d341a2000e7f7aed8ea6dcc099948a527d687582-1647015-1790-3018-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=2557923&amp;t=cd8d42df44&amp;act=addfav',675,415)" id="posted_2557923">2023-05-20 03:47</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>398 pages</div><div class="gldown"><a href="https://exhentai.org/gallerytorrents.php?gid=2557923&amp;t=cd8d42df44" onclick="return popUp('https://exhentai.org/gallerytorrents.php?gid=2557923&amp;
	t=cd8d42df44', 610, 590)" rel="nofollow"><img src="https://exhentai.org/img/t.png" alt="T" title="Show torrents" /></a></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/2477276/81ec4e9067/"><div class="gl4t glname glink">(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase (Touhou Project)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/2477276/81ec4e9067/"><img style="height:352px;width:250px;top:-6px" alt="(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase (Touhou Project)" title="(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase (Touhou Project)" src="https://s.exhentai.org/t/31/f9/31f9aaeaf2daf84fda961fda1cbb3f2ba6707c6a-4477664-2135-3006-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=2477276&amp;t=81ec4e9067&amp;act=addfav',675,415)" id="posted_2477276">2023-02-25 14:02<
	/div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>26 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/2477267/454141e503/"><div class="gl4t glname glink">(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase | 相約月夜 (Touhou Project) [Chinese] [紅銀漢化組]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/2477267/454141e503/"><img style="height:353px;width:250px;top:-6px" alt="(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase | 相約月夜 (Touhou Project) [Chinese] [紅銀漢化組]" title="(C81) [Nekokobo (Kimura Kei)] Tsukiyonoban ni Machiawase | 相約月夜 (Touhou Project) [Chinese] [紅銀漢化組]" src="https://s.exhentai.org/t/91/23/91231f6133f3c7bd714d7d26590f801a1fa39663-2400765-1420-2000-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" on
	click="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=2477267&amp;t=454141e503&amp;act=addfav',675,415)" id="posted_2477267">2023-02-25 13:59</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>27 pages</div><div class="gldown"><a href="https://exhentai.org/gallerytorrents.php?gid=2477267&amp;t=454141e503" onclick="return popUp('https://exhentai.org/gallerytorrents.php?gid=2477267&amp;t=454141e503', 610, 590)" rel="nofollow"><img src="https://exhentai.org/img/t.png" alt="T" title="Show torrents" /></a></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/2326723/021d86a749/"><div class="gl4t glname glink">[tsukiyo rui, shiokonbu, light novel] Kaifuku Jutsushi no Yarinaoshi illust compliation</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/2326723/021d86a749/"><img style="height:354px;width:250px;top:-7px" alt="[tsukiyo rui, shiokonbu
	, light novel] Kaifuku Jutsushi no Yarinaoshi illust compliation" title="[tsukiyo rui, shiokonbu, light novel] Kaifuku Jutsushi no Yarinaoshi illust compliation" src="https://s.exhentai.org/t/92/eb/92eb0b363fba748c207c0cd434dbf7eefe7208a6-509757-424-600-png_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=2326723&amp;t=021d86a749&amp;act=addfav',675,415)" id="posted_2326723">2022-09-14 08:52</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>143 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1931335/bee57d099c/"><div class="gl4t glname glink">(Puniket 35) [Tsukiyo no Mahou (Various)] Okawari Hyakupai!! (Ojamajo Doremi)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="
	https://exhentai.org/g/1931335/bee57d099c/"><img style="height:367px;width:250px;top:-13px" alt="(Puniket 35) [Tsukiyo no Mahou (Various)] Okawari Hyakupai!! (Ojamajo Doremi)" title="(Puniket 35) [Tsukiyo no Mahou (Various)] Okawari Hyakupai!! (Ojamajo Doremi)" src="https://s.exhentai.org/t/66/a8/66a8462d590d42f72703eee2b5bacecb1e95ab84-911378-1669-2449-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1931335&amp;t=bee57d099c&amp;act=addfav',675,415)" id="posted_1931335">2021-06-11 11:08</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>114 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1831421/3f5f344b87/"><div class="gl4t glname glink">(C74) [Ukigusa. (Uki)] Tsukiyo no Sora Futari
	 (Touhou Project) [Spanish] {Nahu89 &amp; Sanji-Nayra}</div></a><div class="gl3t" style="height:176px;width:250px"><a href="https://exhentai.org/g/1831421/3f5f344b87/"><img style="height:176px;width:250px" alt="(C74) [Ukigusa. (Uki)] Tsukiyo no Sora Futari (Touhou Project) [Spanish] {Nahu89 &amp; Sanji-Nayra}" title="(C74) [Ukigusa. (Uki)] Tsukiyo no Sora Futari (Touhou Project) [Spanish] {Nahu89 &amp; Sanji-Nayra}" src="https://s.exhentai.org/t/30/02/3002679fe0860243fca2277a4f98f159d51865a6-1463980-2273-1600-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1831421&amp;t=3f5f344b87&amp;act=addfav',675,415)" id="posted_1831421">2021-01-23 00:57</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>21 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div><
	/div></div><div class="gl1t"><a href="https://exhentai.org/g/1792705/bb430e9c0a/"><div class="gl4t glname glink">[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Chinese] [靴下汉化组] [Digital]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1792705/bb430e9c0a/"><img style="height:348px;width:250px;top:-4px" alt="[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Chinese] [靴下汉化组] [Digital]" title="[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Chinese] [靴下汉化组] [Digital]" src="https://s.exhentai.org/t/c3/78/c3786fa1cfe8c1a43ac7b0cb1f5841c383cea919-603527-800-1113-png_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1792705&amp;t=bb430e9c0a&amp;act=addfav',675,415)" id="posted_1792705">2020-12-06 11:24</div></div><div><div class="ir" style="ba
	ckground-position:-16px -1px;opacity:1"></div><div>37 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1707379/948f4eeaec/"><div class="gl4t glname glink">[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Digital]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1707379/948f4eeaec/"><img style="height:348px;width:250px;top:-4px" alt="[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Digital]" title="[Kare no Iori (Akisaka Takumi)] Tsukiyo no Kaigou (Touhou Project) [Digital]" src="https://s.exhentai.org/t/10/26/102663879bbeb6099114547f7a8ae576797daf80-1077230-800-1113-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1707379&amp;t=948f4eeaec&amp;ac
	t=addfav',675,415)" id="posted_1707379">2020-08-14 05:59</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>34 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1539567/e5988dec41/"><div class="gl4t glname glink">(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | El asunto de la campana limitada de Año Nuevo (Touhou Project) [Spanish] {Paty Scans}</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1539567/e5988dec41/"><img style="height:354px;width:250px;top:-7px" alt="(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | El asunto de la campana limitada de Año Nuevo (Touhou Project) [Spanish] {Paty Scans}" title="(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | El asunto de la campana limitada de Año Nuevo (Touhou Project) [Spanish] {Paty
	 Scans}" src="https://s.exhentai.org/t/c5/d0/c5d0265b1797a8e21adb8c5658ac397f83d4ed6f-760356-1202-1700-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1539567&amp;t=e5988dec41&amp;act=addfav',675,415)" id="posted_1539567">2019-12-24 22:36</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>23 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1510856/282e17fe22/"><div class="gl4t glname glink">(Puniket 25) [Tsukiyo no Mahou (Various)] HazuGoyomi (Ojamajo Doremi)</div></a><div class="gl3t" style="height:172px;width:250px"><a href="https://exhentai.org/g/1510856/282e17fe22/"><img style="height:172px;width:250px" alt="(Puniket 25) [Tsukiyo no Mahou (Various)] HazuGoyomi (Ojamajo Doremi)" ti
	tle="(Puniket 25) [Tsukiyo no Mahou (Various)] HazuGoyomi (Ojamajo Doremi)" src="https://s.exhentai.org/t/77/1b/771bd8d81088e9148aa028d9ad1f57de4c919088-1026762-3025-2074-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1510856&amp;t=282e17fe22&amp;act=addfav',675,415)" id="posted_1510856">2019-11-02 15:06</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>38 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1449245/f081fc6b10/"><div class="gl4t glname glink">[Hanayuuzutsu (Kiduki Kaya)] Hime to Tsukiyo to Takaramono | The Princess a Moonlit Night and Treasure (Touhou Project) [English] [DB Scans] [Digital]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://ex
	hentai.org/g/1449245/f081fc6b10/"><img style="height:352px;width:250px;top:-6px" alt="[Hanayuuzutsu (Kiduki Kaya)] Hime to Tsukiyo to Takaramono | The Princess a Moonlit Night and Treasure (Touhou Project) [English] [DB Scans] [Digital]" title="[Hanayuuzutsu (Kiduki Kaya)] Hime to Tsukiyo to Takaramono | The Princess a Moonlit Night and Treasure (Touhou Project) [English] [DB Scans] [Digital]" src="https://s.exhentai.org/t/e9/b9/e9b98d215c14fc84647548ae779e11ceb51b9c2d-1664587-854-1200-png_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1449245&amp;t=f081fc6b10&amp;act=addfav',675,415)" id="posted_1449245">2019-07-19 22:00</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>24 pages</div><div class="gldown"><a href="https://exhentai.org/gallerytorrents.php?gid=1449245&amp;t=f081fc6b10" onclick="return popUp('https://exhentai.
	org/gallerytorrents.php?gid=1449245&amp;t=f081fc6b10', 610, 590)" rel="nofollow"><img src="https://exhentai.org/img/t.png" alt="T" title="Show torrents" /></a></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1371310/94d2879e73/"><div class="gl4t glname glink">(Koharu Komichi 2) [Iro wa Nioe do (okari)] Sakura wa Tsuki yo no Moto de Saku (Touhou Project)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1371310/94d2879e73/"><img style="height:355px;width:250px;top:-7px" alt="(Koharu Komichi 2) [Iro wa Nioe do (okari)] Sakura wa Tsuki yo no Moto de Saku (Touhou Project)" title="(Koharu Komichi 2) [Iro wa Nioe do (okari)] Sakura wa Tsuki yo no Moto de Saku (Touhou Project)" src="https://s.exhentai.org/t/9d/01/9d016a042a3b60c7ce9b52576a7429e98d68a228-3608642-1199-1700-png_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gall
	erypopups.php?gid=1371310&amp;t=94d2879e73&amp;act=addfav',675,415)" id="posted_1371310">2019-02-24 18:34</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>27 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1323656/211f2fa977/"><div class="gl4t glname glink">(Misora-machi Seijin no Tsudoi) [Tsukiyo no Mahou (Various)] Friends - Seijin Kinen Tokudaigou (Ojamajo Doremi)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1323656/211f2fa977/"><img style="height:371px;width:250px;top:-15px" alt="(Misora-machi Seijin no Tsudoi) [Tsukiyo no Mahou (Various)] Friends - Seijin Kinen Tokudaigou (Ojamajo Doremi)" title="(Misora-machi Seijin no Tsudoi) [Tsukiyo no Mahou (Various)] Friends - Seijin Kinen Tokudaigou (Ojamajo Doremi)" src="https://s.exhentai.org/t/5a/c3/5ac34ee8603c3a97a8f46b5df2733b04fa2ca
	518-1583640-2046-3031-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1323656&amp;t=211f2fa977&amp;act=addfav',675,415)" id="posted_1323656">2018-12-03 08:46</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>92 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1300177/9bbb9cee7f/"><div class="gl4t glname glink">(Puniket 36) [Tsukiyo no Mahou (Rapisu)] Ane ni Suki da to Tsutaeta Hi (Ojamajo Doremi)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1300177/9bbb9cee7f/"><img style="height:373px;width:250px;top:-16px" alt="(Puniket 36) [Tsukiyo no Mahou (Rapisu)] Ane ni Suki da to Tsutaeta Hi (Ojamajo Doremi)" title="(Puniket 36) [Tsukiyo no Mahou
	 (Rapisu)] Ane ni Suki da to Tsutaeta Hi (Ojamajo Doremi)" src="https://s.exhentai.org/t/c4/33/c433c3496206bd978e0a7727f0918caccbfa3020-882547-2049-3052-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1300177&amp;t=9bbb9cee7f&amp;act=addfav',675,415)" id="posted_1300177">2018-10-13 10:34</div></div><div><div class="ir" style="background-position:-16px -21px;opacity:1"></div><div>10 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1103253/c7c47e49e0/"><div class="gl4t glname glink">(Reitaisai 9) [Tsukiyomi (Porurin)] Shinreibyou VS Tengu (Touhou Project)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1103253/c7c47e49e0/"><img style="height:367px;width:250px;top:-13px" alt="(Reitai
	sai 9) [Tsukiyomi (Porurin)] Shinreibyou VS Tengu (Touhou Project)" title="(Reitaisai 9) [Tsukiyomi (Porurin)] Shinreibyou VS Tengu (Touhou Project)" src="https://s.exhentai.org/t/f5/cc/f5cc14183c071237bab94821a1ba7ed305fd5961-977018-1365-2000-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1103253&amp;t=c7c47e49e0&amp;act=addfav',675,415)" id="posted_1103253">2017-08-19 03:16</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>32 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1091523/5debdcdab3/"><div class="gl4t glname glink">(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | New Year Limited Bell Affair (Touhou Project) [English] [DB Scans]</div></a><div class="g
	l3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1091523/5debdcdab3/"><img style="height:354px;width:250px;top:-7px" alt="(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | New Year Limited Bell Affair (Touhou Project) [English] [DB Scans]" title="(C83) [Tsukiyo Gensou (Yuuki Eishi)] Ganjitsu Gentei Bonshou Jiken | New Year Limited Bell Affair (Touhou Project) [English] [DB Scans]" src="https://s.exhentai.org/t/b7/b2/b7b2997ad292fe2b3b50f9c6dd9764bcbc8ccb15-1372436-1414-2000-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1091523&amp;t=5debdcdab3&amp;act=addfav',675,415)" id="posted_1091523">2017-07-25 07:02</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>20 pages</div><div class="gldown"><a href="https://exhentai.org/gallerytorrents.php?gid=1091523&amp;t=5debdcdab3" onclick="re
	turn popUp('https://exhentai.org/gallerytorrents.php?gid=1091523&amp;t=5debdcdab3', 610, 590)" rel="nofollow"><img src="https://exhentai.org/img/t.png" alt="T" title="Show torrents" /></a></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1091522/8e260dcefd/"><div class="gl4t glname glink">(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken | Festival Limited Mirror Affair (Touhou Project) [English] [DB Scans]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1091522/8e260dcefd/"><img style="height:350px;width:250px;top:-5px" alt="(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken | Festival Limited Mirror Affair (Touhou Project) [English] [DB Scans]" title="(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken | Festival Limited Mirror Affair (Touhou Project) [English] [DB Scans]" src="https://s.exhentai.org/t/73/4a/734a59b731ba08a7c76cd4ef7ec5d20a3c0c2d75-686453-1079-1510-jpg_
	250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1091522&amp;t=8e260dcefd&amp;act=addfav',675,415)" id="posted_1091522">2017-07-25 07:01</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>25 pages</div><div class="gldown"><a href="https://exhentai.org/gallerytorrents.php?gid=1091522&amp;t=8e260dcefd" onclick="return popUp('https://exhentai.org/gallerytorrents.php?gid=1091522&amp;t=8e260dcefd', 610, 590)" rel="nofollow"><img src="https://exhentai.org/img/t.png" alt="T" title="Show torrents" /></a></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1075069/1213f09ce5/"><div class="gl4t glname glink">(Tsukiyomi no Utage 5) [S_size (Risunosuri)] Servant wa Yume o miru ka (Fate/stay night) [Chinese]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1075069/1213f09ce5
	/"><img style="height:357px;width:250px;top:-8px" alt="(Tsukiyomi no Utage 5) [S_size (Risunosuri)] Servant wa Yume o miru ka (Fate/stay night) [Chinese]" title="(Tsukiyomi no Utage 5) [S_size (Risunosuri)] Servant wa Yume o miru ka (Fate/stay night) [Chinese]" src="https://s.exhentai.org/t/d0/ab/d0abec79fa467354da01ccc57ee636cdc13d0203-1494103-1500-2136-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1075069&amp;t=1213f09ce5&amp;act=addfav',675,415)" id="posted_1075069">2017-06-14 14:55</div></div><div><div class="ir" style="background-position:-32px -1px;opacity:1"></div><div>38 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1071660/662c807fcb/"><div class="gl4t glname glink">(Tsukiyomi no Utage 5) [S_size (Risunosuri)]
	 Servant wa Yume o miru ka (Fate/stay night)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1071660/662c807fcb/"><img style="height:357px;width:250px;top:-8px" alt="(Tsukiyomi no Utage 5) [S_size (Risunosuri)] Servant wa Yume o miru ka (Fate/stay night)" title="(Tsukiyomi no Utage 5) [S_size (Risunosuri)] Servant wa Yume o miru ka (Fate/stay night)" src="https://s.exhentai.org/t/d0/ab/d0abec79fa467354da01ccc57ee636cdc13d0203-1494103-1500-2136-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1071660&amp;t=662c807fcb&amp;act=addfav',675,415)" id="posted_1071660">2017-06-06 02:02</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>38 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class=
	"gl1t"><a href="https://exhentai.org/g/1017924/e3d74cd35d/"><div class="gl4t glname glink">(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night) [Chinese]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1017924/e3d74cd35d/"><img style="height:358px;width:250px;top:-9px" alt="(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night) [Chinese]" title="(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night) [Chinese]" src="https://s.exhentai.org/t/db/f0/dbf0a233314ad30aff4a4ee71f9642d8274fe58f-604575-1200-1718-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1017924&amp;t=e3d74cd35d&amp;act=addfav',675,415)" id="posted_1017924">2017-01-14 12:45</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>22 pages</div><div class="gldown"><img src="https://exhentai.or
	g/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1015226/d4fa63cfb3/"><div class="gl4t glname glink">(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night) [Chinese]</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1015226/d4fa63cfb3/"><img style="height:354px;width:250px;top:-7px" alt="(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night) [Chinese]" title="(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night) [Chinese]" src="https://s.exhentai.org/t/f4/e1/f4e155b08cdaaccb9ba70622f0f21a3d743d395f-2089597-1200-1697-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1015226&amp;t=d4fa63cfb3&amp;act=addfav',675,415)" id="posted_1015226">2017-01-07 15:08</div></div><di
	v><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>33 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1011474/5c70a44dc2/"><div class="gl4t glname glink">(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1011474/5c70a44dc2/"><img style="height:358px;width:250px;top:-9px" alt="(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night)" title="(Tsukiyomi no Utage) [KP (Marimo)] M/S system (Fate/stay night)" src="https://s.exhentai.org/t/db/f0/dbf0a233314ad30aff4a4ee71f9642d8274fe58f-604575-1200-1718-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1011474&amp;t=5c70a44dc2&amp;act=addfav',67
	5,415)" id="posted_1011474">2016-12-30 04:26</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>22 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/1002510/48c948b4f8/"><div class="gl4t glname glink">(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/1002510/48c948b4f8/"><img style="height:354px;width:250px;top:-7px" alt="(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night)" title="(Tsukiyomi no Utage 4) [S_size (Risunosuri)] Trick Trick Chair (Fate/stay night)" src="https://s.exhentai.org/t/f4/e1/f4e155b08cdaaccb9ba70622f0f21a3d743d395f-2089597-1200-1697-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</d
	iv><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=1002510&amp;t=48c948b4f8&amp;act=addfav',675,415)" id="posted_1002510">2016-12-04 11:59</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>33 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/892484/526387d655/"><div class="gl4t glname glink">(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken (Touhou Project)</div></a><div class="gl3t" style="height:340px;width:250px"><a href="https://exhentai.org/g/892484/526387d655/"><img style="height:350px;width:250px;top:-5px" alt="(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken (Touhou Project)" title="(Reitaisai 6) [Tsukiyo Gensou (Yuuki Eishi)] Saijitsu Gentei Jouhari Jiken (Touhou Project)" src="https://s.exhentai.org/t/73/4a/734a59b731ba08a7c76cd4ef7ec5d20a3c0c2d75-686453-10
	79-1510-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=892484&amp;t=526387d655&amp;act=addfav',675,415)" id="posted_892484">2016-01-09 22:56</div></div><div><div class="ir" style="background-position:-16px -1px;opacity:1"></div><div>27 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div><div class="gl1t"><a href="https://exhentai.org/g/892481/d61e22787f/"><div class="gl4t glname glink">(C76) [Tsukiyo Gensou (Yuuki Eishi)] Oujo Gensou -atropa belladonna- (Touhou Project)</div></a><div class="gl3t" style="height:175px;width:250px"><a href="https://exhentai.org/g/892481/d61e22787f/"><img style="height:175px;width:250px" alt="(C76) [Tsukiyo Gensou (Yuuki Eishi)] Oujo Gensou -atropa belladonna- (Touhou Project)" title="(C76) [Tsukiyo Gensou (Yuuki Eishi)] Oujo Gensou -atropa bell
	adonna- (Touhou Project)" src="https://s.exhentai.org/t/6f/62/6f6218eb133e640b0a274dee8e09fa3aa2fdf97c-1311534-2155-1500-jpg_250.jpg" /></a></div><div class="gl5t"><div><div class="cs ct9" onclick="document.location='https://exhentai.org/non-h'">Non-H</div><div onclick="popUp('https://exhentai.org/gallerypopups.php?gid=892481&amp;t=d61e22787f&amp;act=addfav',675,415)" id="posted_892481">2016-01-09 22:39</div></div><div><div class="ir" style="background-position:0px -21px;opacity:1"></div><div>60 pages</div><div class="gldown"><img src="https://exhentai.org/img/td.png" alt="T" title="No torrents available" /></div></div></div></div></div><div class="searchnav"><div></div><div><span id="dfirst">&lt;&lt; First</span></div><div><span id="dprev">&lt; Prev</span></div><div id="djumpbox" class="jumpbox"><a id="djump" href="javascript:enable_jump_mode('d')">Jump/Seek</a></div><div><a id="dnext" href="https://exhentai.org/?f_search=tsukiyo&amp;f_cats=767&amp;next=892481">Next &gt;</a></div><div><a id="dlast" href="htt
	ps://exhentai.org/?f_search=tsukiyo&amp;f_cats=767&amp;prev=1">Last &gt;&gt;</a></div><div></div></div></div></div>
	<div class="dp" style="margin:0 auto 5px">
			<a href="https://exhentai.org/">Front Page</a>
	
			  &nbsp; <a href="http://exhentai55ld2wyap5juskbm67czulomrouspdacjamjeloj7ugjbsad.onion">Onion</a>
	</div>
	</body>
	</html>
`
	doc, err := htmlquery.Parse(strings.NewReader(respStr))
	if err != nil {
		t.Error(err)
	}
	list, err := htmlquery.QueryAll(doc, "//div[@class='gl1t']")
	if err != nil {
		t.Error(err)
	}

	dom := list[0]

	fmt.Printf("%+v\n\n", parseIndexNodeToSummary(dom))
}

func Test_error(t *testing.T) {

	f := func() (gs *GallerySummary) {
		// gs = &GallerySummary{}
		defer func() {
			if err := recover(); err != nil {
				gs.Error = fmt.Sprintf("RECOVERED: %v\n", err)
				return
			}
		}()
		gs = &GallerySummary{} // 这里也是可以的
		panic("123")
	}
	fmt.Println(f())
}

func Test_QueryIndex(t *testing.T) {
	mycurl.SetClient("proxy")

	r, err := queryIndex(`f_cats=767&f_search=tsukiyo`)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func Test_NullError(t *testing.T) {
	mycurl.SetClient("proxy")
	query := `` // why when query homepage, encounter a null ptr error

	index := &Index{}

	index.Query = query

	resp, err := mycurl.Fetch("GET", BASE_URL+query,
		map[string]string{"Cookie": COOKIE},
		nil)
	if err != nil {
		index.Error = err.Error()
		return
	}

	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		index.Error = err.Error()
		return
	}

	respStr := string(respText)
	index.Results = (reResults.FindString(respStr))
	index.NextPage = (reNextPage.FindString(respStr))
	index.PrevPage = (rePrevPage.FindString(respStr)) // 麻了，golang的re包太麻了，传回去js里面处理吧。

	doc, err := htmlquery.Parse(strings.NewReader(respStr))
	if err != nil {
		index.Error = err.Error()
		return
	}
	list, err := htmlquery.QueryAll(doc, "//div[@class='gl1t']")
	if err != nil {
		index.Error = err.Error()
		return
	}
	index.Galleries = make([]*GallerySummary, len(list))
	for k, v := range list {
		gs := &GallerySummary{}
		n := v

		url := htmlquery.FindOne(n, "//a")
		gs.Url = htmlquery.SelectAttr(url, "href")

		name := htmlquery.FindOne(n, "//div[@class='gl4t glname glink']")
		gs.Title = htmlquery.InnerText(name)

		img := htmlquery.FindOne(n, "//img")
		gs.Preview = htmlquery.SelectAttr(img, "src")

		gl5t := htmlquery.FindOne(n, "//div[@class='gl5t']")
		gl5t_1 := htmlquery.FindOne(gl5t, "div[1]")

		fmt.Println(htmlquery.OutputHTML(gl5t, true))
		fmt.Println(htmlquery.OutputHTML(gl5t_1, true))

		category := htmlquery.FindOne(gl5t_1, "//div[contains(@class, 'cs')]")
		gs.Category = htmlquery.InnerText(category)

		date := htmlquery.FindOne(gl5t_1, "div[2]")
		gs.Date = htmlquery.InnerText(date)

		gl5t_2 := htmlquery.FindOne(gl5t, "div[2]")
		rate := htmlquery.FindOne(gl5t_2, "div[1]")
		gs.Rate = htmlquery.SelectAttr(rate, "style")

		pages := htmlquery.FindOne(gl5t_2, "div[2]")
		gs.Pages = htmlquery.InnerText(pages)

		seed := htmlquery.FindOne(gl5t_2, "//img")
		// gs.Seeds = []string{htmlquery.SelectAttr(seed, "src")}
		gs.Seeds = htmlquery.SelectAttr(seed, "src")
		// Error       string   `json:"error"`        // error

		index.Galleries[k] = gs
	}

	fmt.Println(index)
	j, _ := json.Marshal(index)
	fmt.Println(string(j))
}

func TestRe(t *testing.T) {
	fmt.Println(reResults.FindString("123,4+ results."))
}
