func (comm *HTTPCommunicator) Deliver(/*parameters*/) {
 // Transforma bytes para request
 // Altera o host alvo
 actualRequestRecovered.Header.Set("Host",
  "http://"+comm.toAddr+actualRequestRecovered.RequestURI)
 actualRequestRecovered.URL.Host = comm.toAddr
 actualRequestRecovered.Host = comm.toAddr
 actualRequestRecovered.RequestURI = ""
 actualRequestRecovered.URL.Scheme = "http"
 // Executa request
}

func (comm *HTTPCommunicator) Listen(/*parameters*/)
 // Transforma request em bytes
 // Envia ao ordenador
 handle(bufferedRequestHolder.Bytes())
}
