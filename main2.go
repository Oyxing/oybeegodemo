package main

// type A struct {
// 	Created  time.Time
// 	Time     string
// 	Hardware string
// }

// func main2() {
// 	haha := A{}

// 	a, e := d.CreateHardware()
// 	if e != nil {
// 		fmt.Println(e)
// 		return
// 	}
// 	haha.Hardware = a
// 	haha.Time = "aaaaa"
// 	aescode := "hgfedcba87654321"
// 	w, _ := json.Marshal(haha)

// 	//生成秘钥对
// 	bits := 2048
// 	err := d.GenerateRSAKey(bits, "1.pem", "2.pem")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("RSA密码生成成功")

// 	//加密
// 	ecode, err := d.En_crypt_AES_Base64(w, []byte(aescode))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	x, err := d.RSA_Encrypt([]byte(aescode), "2.pem")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//解密
// 	y, err := d.RSA_Decrypt(x, "1.pem")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	str, err := d.De_crypt_Base64_AES(ecode, y)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//验证字符串
// 	dd := A{}
// 	json.Unmarshal([]byte(str), &dd)

// 	if dd.Hardware == a {
// 		fmt.Println(1)
// 	} else {
// 		fmt.Println(2)
// 	}

// }
