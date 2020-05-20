package main

//func GetServerCreds() credentials.TransportCredentials {
//	cert, err := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool := x509.NewCertPool()
//	ca, err := ioutil.ReadFile("cert/ca.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool.AppendCertsFromPEM(ca)
//	creds := credentials.NewTLS(&tls.Config{
//		Certificates: []tls.Certificate{cert},   // 服务端证书
//		ClientAuth:   tls.RequireAndVerifyClientCert,
//		ClientCAs:    certPool,
//	})
//	return creds
//}
//
//func GetClientCreds() credentials.TransportCredentials {
//	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool := x509.NewCertPool()
//	ca, err := ioutil.ReadFile("cert/ca.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//	certPool.AppendCertsFromPEM(ca)
//	creds := credentials.NewTLS(&tls.Config{
//		Certificates: []tls.Certificate{cert},
//		ServerName:   "localhost",
//		RootCAs:      certPool,
//	})
//	return creds
//}
