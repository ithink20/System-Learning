package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"net/url"
	"strconv"
	"strings"
)

type RequestPayload struct {
	requestId string
	cookie    string
	url       string
	userAgent string
	country   string
	referer   string
	clientIp  string
	ecData    string
	clientId  string
	hsg       string
	hsj       string
	hsjh      string
	headers   string
	body      string
	method    string
	nurl      string
	protocol  string
}

func bindTSSRequestV2(body string) {
	var requestBody []string
	_ = json.Unmarshal([]byte(body), &requestBody)
}

func bindTSSRequestV3(body string) {
	var requestBody []string
	_ = sonic.Unmarshal([]byte(body), &requestBody)
}

func bindRequestPayload(body string, lenArr []int, startIndex int) {
	req := &RequestPayload{}
	start := startIndex + 2
	end := start + lenArr[0]
	req.requestId = body[start:end]
	start = end + 3
	end = start + lenArr[1]
	req.cookie = body[start:end]
	start = end + 3
	end = start + lenArr[2]
	req.url = body[start:end]
	start = end + 3
	end = start + lenArr[3]
	req.userAgent = body[start:end]
	start = end + 3
	end = start + lenArr[4]
	req.country = body[start:end]
	start = end + 3
	end = start + lenArr[5]
	req.referer = body[start:end]
	start = end + 3
	end = start + lenArr[6]
	req.clientIp = body[start:end]
	start = end + 3
	end = start + lenArr[7]
	req.ecData = body[start:end]
	start = end + 3
	end = start + lenArr[8]
	req.clientId = body[start:end]
	start = end + 3
	end = start + lenArr[9]
	req.hsg = body[start:end]
	start = end + 3
	end = start + lenArr[10]
	req.hsj = body[start:end]
	start = end + 3
	end = start + lenArr[11]
	req.hsjh = body[start:end]
	start = end + 3
	end = start + lenArr[12]
	req.headers = body[start:end]
	start = end + 3
	end = start + lenArr[13]
	req.body = body[start:end]
	start = end + 3
	end = start + lenArr[14]
	req.method = body[start:end]
	start = end + 3
	end = start + lenArr[15]
	req.nurl = body[start:end]
	start = end + 3
	end = start + lenArr[16]
	req.protocol = body[start:end]
	//fmt.Printf("\n\ndata : %+v", req)
}

func CustomSplitV2(input string) (result []int, skipIndex int) {
	lenArrayEndIndex := strings.IndexByte(input, ',')
	lenArray := input[2 : lenArrayEndIndex-1]

	current := 0
	for index, char := range lenArray {
		if char == ';' {
			num, _ := strconv.Atoi(lenArray[current:index])
			result = append(result, num)
			current = index + 1
		}
	}
	num, _ := strconv.Atoi(lenArray[current:])
	result = append(result, num)
	return result, lenArrayEndIndex
}

var headerMap = make(map[string]string, 30)

func headerEncode() {
	headerMap["authority"] = "shopee.sg"
	headerMap["accept"] = "application/json"
	headerMap["accept-language"] = "en-US,en;q=0.9"
	headerMap["af-ac-enc-dat"] = "AAczLjMuMC0yAAABjD6brBwAABBxAzAAAAAAAAAAAv0CU/aTV0SD/jycgsufjcZ/eBvpl/82EyfDx/bVRcPaaRvYm5f/NhMnw8f21UXD2mkb2Jtz6Qf2dlDKbCRUtNVA1LitTX5OcPT+838L0LU4MCylvdIACsKvcJXpvNmfWLPMuUaM1ge+W/8o5zEdOjkDWLv4Q6rWYAfDda9H1fc6Ul12RgUgsteDIoKwUsp6DzP7CeeX/zYTJ8PH9tVFw9ppG9iboQav8dow1MpbrmR495Fj3/KMjlZPgg83do8+M30IKpMTeUAhbodJtmu4kLBVAN/I7e/J70LRwsmJCnZgrNgTptEcknYcoSs06Hl67Aq/+1YahsHLWBbnuRPUzzuWFKljS+mlj7cLs2/LGMg6ZjMMiBvMJk99llmOEsYl9f0oke64s5tFoZEFgYf5MgDhtow2nd9Z1tlxSFUixbTWYFNMdZTsCqmPT4gDej6dsQN6RDUK0rSbqHQTQYVGpQYIorZchlHodGMbEzp7tSm6o3sSjIFsOCRhd8QpUk3jJg+1RkYCbC376bN59TtJhIy6LbEskhNBLjm1h17osrvnSCzxcRCD9IPgxiPeN5y7wHMpZIvC2w9RceYESpV5+Sv44ymTu0cge3fcQ/CkPE2wl8xqLaHpNHRwWVPUnyo9gIPQIP7oTSYUoZmmcTW5TFjhYuOTLvIRZyZ1lLkuxgMzcDHyX0U7Y+Mj9wrMBk/bry68U4ZcqjywJ/hF/zOo/59kkDOYEO/JrcdTnLcImYhKI7pCaAJ4p8UTTWugnL0IvwdiQJdcqjywJ/hF/zOo/59kkDOYGJ+oW4DGwZziBCetwV+/fDcdf23jwtkc2gUdlgSvHDVRAl+rcrUAbUSko7rXqS2LBtjs/0ocwvppC0XgWUp3rNCBhaQafyRUEGkb3E20b3lSXDGzDMO3KD87PktFjVW6Btjs/0ocwvppC0XgWUp3rKXMztvUr1gIO+0VU5aM3aEak8TsSV5yLoMfz2Q6MeRuLMKiLwDcJj9E5jxrAy72JPGdAk6q1/J0XDbzwi8/cZ6tgI7hsxXAJUaTxr+iJecws5hKsi2KJB2sch1stvHct+uYA4d3ddemJNqsyOE63l4="
	headerMap["content-type"] = "application/json"
	headerMap["sec-ch-ua"] = `""Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"`
	headerMap["sec-ch-ua-mobile"] = "?0"
	headerMap["sec-ch-ua-platform"] = `"macOS"`
	headerMap["sec-fetch-dest"] = "empty"
	headerMap["sec-fetch-mode"] = "cors"
	headerMap["sec-fetch-site"] = "same-origin"
	headerMap["sz-token"] = "hXcc3pWHcu3B+IKvJebSwA==|Vz1wabxGPIyOkOcJiRY4LyCQ2G58ep1ikW7UaKRIt0oDAfD6csDKiMbYjnwgTUJpgr+Fwr8BYMZC7g==|mFsHXk9y6t4WmbrE|08|3"
	headerMap["x-api-source"] = "pc"
	headerMap["x-csrftoken"] = "cRCO3hhz3pfZH3PIJQJRWghUcl91gOJs"
	headerMap["x-requested-with"] = "XMLHttpRequest"
	headerMap["x-sap-access-f"] = "3.2.119.6.0|13|3.3.0-2_5.5.200_0_308|79b616b2c3904bb789f8317d145775cc182e3e694cf445|10900|100"
	headerMap["x-sap-access-s"] = "49EYjKaIqz2NU-HarNHNv8DZNICzdDvwpm_Z_bgByBQ="
	headerMap["x-sap-access-t"] = "1701857441"
	for key, val := range headerMap {
		headerMap[key] = url.QueryEscape(val)
	}
	//fmt.Printf("encodedMap: %v\n\n", headerMap["sec-ch-ua"])
}

func headerDecode() {
	for key, val := range headerMap {
		ans, _ := url.QueryUnescape(val)
		headerMap[key] = ans
	}
	//fmt.Printf("encodedMap: %v\n\n", headerMap["sec-ch-ua"])
}

func headerUnMarshal() {
	//jsonString, _ := json.Marshal(headerMap)
	jsonString := "{\"accept\":\"application%2Fjson\",\"accept-language\":\"en-US%2Cen%3Bq%3D0.9\",\"af-ac-enc-dat\":\"AAczLjMuMC0yAAABjD6brBwAABBxAzAAAAAAAAAAAv0CU%2FaTV0SD%2FjycgsufjcZ%2FeBvpl%2F82EyfDx%2FbVRcPaaRvYm5f%2FNhMnw8f21UXD2mkb2Jtz6Qf2dlDKbCRUtNVA1LitTX5OcPT%2B838L0LU4MCylvdIACsKvcJXpvNmfWLPMuUaM1ge%2BW%2F8o5zEdOjkDWLv4Q6rWYAfDda9H1fc6Ul12RgUgsteDIoKwUsp6DzP7CeeX%2FzYTJ8PH9tVFw9ppG9iboQav8dow1MpbrmR495Fj3%2FKMjlZPgg83do8%2BM30IKpMTeUAhbodJtmu4kLBVAN%2FI7e%2FJ70LRwsmJCnZgrNgTptEcknYcoSs06Hl67Aq%2F%2B1YahsHLWBbnuRPUzzuWFKljS%2Bmlj7cLs2%2FLGMg6ZjMMiBvMJk99llmOEsYl9f0oke64s5tFoZEFgYf5MgDhtow2nd9Z1tlxSFUixbTWYFNMdZTsCqmPT4gDej6dsQN6RDUK0rSbqHQTQYVGpQYIorZchlHodGMbEzp7tSm6o3sSjIFsOCRhd8QpUk3jJg%2B1RkYCbC376bN59TtJhIy6LbEskhNBLjm1h17osrvnSCzxcRCD9IPgxiPeN5y7wHMpZIvC2w9RceYESpV5%2BSv44ymTu0cge3fcQ%2FCkPE2wl8xqLaHpNHRwWVPUnyo9gIPQIP7oTSYUoZmmcTW5TFjhYuOTLvIRZyZ1lLkuxgMzcDHyX0U7Y%2BMj9wrMBk%2Fbry68U4ZcqjywJ%2FhF%2FzOo%2F59kkDOYEO%2FJrcdTnLcImYhKI7pCaAJ4p8UTTWugnL0IvwdiQJdcqjywJ%2FhF%2FzOo%2F59kkDOYGJ%2BoW4DGwZziBCetwV%2B%2FfDcdf23jwtkc2gUdlgSvHDVRAl%2BrcrUAbUSko7rXqS2LBtjs%2F0ocwvppC0XgWUp3rNCBhaQafyRUEGkb3E20b3lSXDGzDMO3KD87PktFjVW6Btjs%2F0ocwvppC0XgWUp3rKXMztvUr1gIO%2B0VU5aM3aEak8TsSV5yLoMfz2Q6MeRuLMKiLwDcJj9E5jxrAy72JPGdAk6q1%2FJ0XDbzwi8%2FcZ6tgI7hsxXAJUaTxr%2BiJecws5hKsi2KJB2sch1stvHct%2BuYA4d3ddemJNqsyOE63l4%3D\",\"af-ac-enc-sz-token\":\"hXcc3pWHcu3B%2BIKvJebSwA%3D%3D%7CVz1wabxGPIyOkOcJiRY4LyCQ2G58ep1ikW7UaKRIt0oDAfD6csDKiMbYjnwgTUJpgr%2BFwr8BYMZC7g%3D%3D%7CmFsHXk9y6t4WmbrE%7C08%7C3\",\"authority\":\"shopee.sg\",\"content-type\":\"application%2Fjson\",\"sec-ch-ua\":\"%22%22Google+Chrome%22%3Bv%3D%22119%22%2C+%22Chromium%22%3Bv%3D%22119%22%2C+%22Not%3FA_Brand%22%3Bv%3D%2224%22\",\"sec-ch-ua-mobile\":\"%3F0\",\"sec-ch-ua-platform\":\"%22macOS%22\",\"sec-fetch-dest\":\"empty\",\"sec-fetch-mode\":\"cors\",\"sec-fetch-site\":\"same-origin\",\"sz-token\":\"hXcc3pWHcu3B%2BIKvJebSwA%3D%3D%7CVz1wabxGPIyOkOcJiRY4LyCQ2G58ep1ikW7UaKRIt0oDAfD6csDKiMbYjnwgTUJpgr%2BFwr8BYMZC7g%3D%3D%7CmFsHXk9y6t4WmbrE%7C08%7C3\",\"x-api-source\":\"pc\",\"x-csrftoken\":\"cRCO3hhz3pfZH3PIJQJRWghUcl91gOJs\",\"x-requested-with\":\"XMLHttpRequest\",\"x-sap-access-f\":\"3.2.119.6.0%7C13%7C3.3.0-2_5.5.200_0_308%7C79b616b2c3904bb789f8317d145775cc182e3e694cf445%7C10900%7C100\",\"x-sap-access-s\":\"49EYjKaIqz2NU-HarNHNv8DZNICzdDvwpm_Z_bgByBQ%3D\",\"x-sap-access-t\":\"1701857441\",\"x-sap-ri\":\"a24870658e47e2129e0ea13c0101d4be67e1e083990aebee5b7f\",\"x-sap-sec\":\"lhwd7dkInWi4uWi4qOiSuW94qOi4uW94uWiJuWi4XWY4u1i5uWikuWi4dVJGQB%2B4uWg4ufi4JWy4uvRdEYtQlAhzl55CRjJdXxofZMR%2By69iGcN0UPtHtgUXgFTGAiQ%2Fn3k4x%2FxZldWkQGDEhAmU2hrsIDyaB4TdYfYPZcxQKQKdgfaVl46BieSagj5IMLJglY87PmFqB%2FdRTdEGxwTTBA29zXSDI2bn7qcyuTy9csPYIqD37Bn4peRTZGdilIENAe3P4jIQyuQERZjQfmPk05dbm74DWWp2apf%2BEp4FhoFaNJ89%2FuYCpbKfgDc9iXFeJM%2FvA%2BDhfQSERoEDDnpNnuCTKen1iu3Q2AK%2Bldwy3%2FwAg73a4zmqRMcx8lURjyjkMHfmj7ZdPVonBZVcYFOk0L9OyrwWiKdfKlrZ4y5%2FTkd72GC%2FnGx1e2hQDtfeWpG9qK10T38rQWNeHuNe77rlwURR%2FDHO5y9WPDuYcdAjceW8CDpd%2FEn%2FgYLz9fazG8bc4T8JH23AW%2BZJXzQHfVxRPblMKacrZDK2%2BsM%2BDzn9XOhTS8JL%2BCwgcHPIEZI0hNAo0ib6fIyFNYRV%2Bl609mzmp%2FnlMtvIQAYwblTZXpuU1RZPxkPr5ipxWfAOCQgsennaEch5kCei9%2BN86kQptCuMWGwGtWi4uTygiTzWiXyzuWi4uvPGQ6SSuWi4vWi4u894uWfDjQVbtYJNQIAyBEL0lrAwn5s%2BRf%2B4uWizij%2BgFywWiWi4uWiSuWk4tWikuW%2B4uWiSuWi4vWi4u894uWgtYTNvgv%2FgaYA92gjLD02nfRGT4f%2B4uWigTTyijNbXFWi4uWf%3D\",\"x-shopee-language\":\"en\",\"x-sz-sdk-version\":\"3.3.0-2%261.6.8\"}"
	//fmt.Printf("%v", string(jsonString))
	_ = json.Unmarshal([]byte(jsonString), &headerMap)
}

func main() {
	body := "[\"10;1160;462;137;2;49;14;0;0;1;138;32;1053;0;3;37;8\",\"vikas-test\",\"__LOCALE__null=SG; csrftoken=X2wOktygIduHwe2Zds5mMnIXyiB6mbwI; _gcl_au=1.1.1648155970.1690254295; SPC_T_IV=RGp1ZTBtUXF3Z1cxZjNPcg==; SPC_SI=Lu6wZAAAAABtQWRlblExbJRSqgIAAAAAU09vSDU4ZVk=; SPC_F=VrnSi26BC9OoaxewpO0q60KDmpJgG1ko; REC_T_ID=0812da37-2a98-11ee-b0ab-3e0427df20f5; SPC_R_T_ID=DJLbTOEPIPqE9ATl3cBPQLqPIjErTfCKBmGz5o9g7luvHZ2iZzTAezcLUWv2P+oG6TDFp+jFdNuQgtTf2ChrTrTp33K8CH61vFi7wFOuKdrrvaYJU59VV8ZTsgeQKx0lUhT4QJB9WkO8ZnLLgJci+qHgm/d12VqOclzjO9VwSQM=; SPC_R_T_IV=RGp1ZTBtUXF3Z1cxZjNPcg==; SPC_T_ID=DJLbTOEPIPqE9ATl3cBPQLqPIjErTfCKBmGz5o9g7luvHZ2iZzTAezcLUWv2P+oG6TDFp+jFdNuQgtTf2ChrTrTp33K8CH61vFi7wFOuKdrrvaYJU59VV8ZTsgeQKx0lUhT4QJB9WkO8ZnLLgJci+qHgm/d12VqOclzjO9VwSQM=; _QPWSDCXHZQA=6ffaf0c6-86a7-4f45-fde0-23d9edd08e08; AMP_TOKEN=%24NOT_FOUND; _gid=GA1.2.872370811.1690254360; _dc_gtm_UA-125099498-1=1; shopee_webUnique_ccd=FxuNEaNRJXycjOEBANcitQ%3D%3D%7Cg8aFyUDF%2FHHpjNHbQHXQm0grj%2F%2FJOjXixuW4%2B4F8RD3IPm7UXYfsfkmDnRArx9UfPCqL%2FbNFqPPImeJx%7CFtcd62SjXM98Xeah%7C08%7C3; ds=290425dd546d9437721b5cded7f92b40; _ga=GA1.2.460650685.1690254296; _ga_CGXK257VSB=GS1.1.1690254295.1.1.1690254385.45.0.0; _ga_PN56VNNPQX=GS1.2.1690254360.1.1.1690254392.0.0.0\",\"https://demo.studio.sg/api/v4/search/test-vikas-custom-path?by=relevancy\\u0026extra_params=%7B%22global_search_session_id%22%3A%22gs-4e4593e7-5f20-49ba-a8eb-e707724cf1d6%22%2C%22search_session_id%22%3A%22ss-e24f7a98-3dfa-43ed-a631-04a63a7542c5%22%7D\\u0026keyword=test_postman_at_testtest019\\u0026limit=20\\u0026newest=0\\u0026order=desc\\u0026page_type=search\\u0026scenario=PAGE_GLOBAL_SEARCH\\u0026version=2\\u0026view_session_id=8cf19899-c954-4eaa-94df-0942e402de9f\",\"Mozilla/5.0 (Linux; Android 8.0.0; SM-G955U Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Mobile Safari/537.36\",\"SG\",\"https://demo.studio.sg/search?keyword=hinthint009\",\"143.92.111.111\",\"\",\"\",\"0\",\"771,4865-4866-4867-49199-49195-49200-49196-52393-52392-49161-49171-49162-49172-156-157-47-53-10,0-23-65281-10-11-35-13-51-45-43,29-23-24,0\",\"dda262729e5413660ec0e6a8d4279860\",\"{\\\"sec-ch-ua-platform\\\":\\\"\\\\\\\"Android\\\\\\\"\\\",\\\"sz-token\\\":\\\"\\\\/KHejAJT2i7\\\\/R7G7fTnkTA==|BzJe4zN+ST7s1dicSxA10B4YU2usLOr\\\\/U5SG8vKoTEGDbet4TlJANyXYQQJtaeolfDjwVFDJtxpLeQrl|HSF0grphDm4dZFxp|08|3\\\",\\\"host\\\":\\\"test.shopee.sg\\\",\\\"sec-fetch-dest\\\":\\\"empty\\\",\\\"x-api-source\\\":\\\"rweb\\\",\\\"accept-language\\\":\\\"en,zh-CN;q=0.9,zh;q=0.8,ja;q=0.7\\\",\\\"content-type\\\":\\\"application\\\\/json\\\",\\\"connection\\\":\\\"keep-alive\\\",\\\"x-csrftoken\\\":\\\"aI9t9VIyuiAasYEQstZ4D71A0oDw6jSb\\\",\\\"x-shopee-language\\\":\\\"en\\\",\\\"sec-fetch-site\\\":\\\"same-origin\\\",\\\"accept-encoding\\\":\\\"gzip, deflate, br\\\",\\\"x-requested-with\\\":\\\"XMLHttpRequest\\\",\\\"af-ac-enc-sz-token\\\":\\\"\\\\/KHejAJT2i7\\\\/R7G7fTnkTA==|BzJe4zN+ST7s1dicSxA10B4YU2usLOr\\\\/U5SG8vKoTEGDbet4TlJANyXYQQJtaeolfDjwVFDJtxpLeQrl|HSF0grphDm4dZFxp|08|3\\\",\\\"x-sz-sdk-version\\\":\\\"2.9.2-2\\u00261.4.1\\\",\\\"sec-fetch-mode\\\":\\\"cors\\\",\\\"sec-ch-ua-mobile\\\":\\\"?1\\\",\\\"authority\\\":\\\"test.shopee.sg\\\",\\\"sec-ch-ua\\\":\\\"\\\\\\\"Not.A\\\\/Brand\\\\\\\";v=\\\\\\\"8\\\\\\\", \\\\\\\"Chromium\\\\\\\";v=\\\\\\\"114\\\\\\\", \\\\\\\"Google Chrome\\\\\\\";v=\\\\\\\"114\\\\\\\"\\\",\\\"accept\\\":\\\"application\\\\/json\\\"}\",\"\",\"GET\",\"/api/v4/search/test-vikas-custom-path\",\"HTTP/1.1\"]"
	req, index := CustomSplitV2(body)
	bindRequestPayload(body, req, index)
}
