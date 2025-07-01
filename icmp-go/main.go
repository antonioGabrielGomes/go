package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

const (
	ProtocolICMP = 1
)

func main() {
	// Configuração
	target := "8.8.8.8"
	timeout := 3 * time.Second
	// count := 3

	// Resolve o endereço IP do alvo
	dst, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		log.Fatalf("Falha ao resolver IP: %v", err)
	}

	// Cria uma conexão ICMP
	conn, err := icmp.ListenPacket("ip4:icmp", "")
	if err != nil {
		log.Fatalf("Falha ao criar conexão ICMP: %v", err)
	}
	defer conn.Close()

	// Configura timeout
	err = conn.SetDeadline(time.Now().Add(timeout))
	if err != nil {
		log.Fatalf("Falha ao definir timeout: %v", err)
	}

	// Prepara a mensagem ICMP (Echo Request)
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff, // Usa o PID do processo como ID
			Seq:  1,
			Data: []byte("HELLO-PING"),
		},
	}

	// Serializa a mensagem para bytes
	msgBytes, err := msg.Marshal(nil)
	if err != nil {
		log.Fatalf("Falha ao serializar mensagem ICMP: %v", err)
	}

	// Envia o ping
	start := time.Now()
	_, err = conn.WriteTo(msgBytes, dst)
	if err != nil {
		log.Fatalf("Falha ao enviar ping: %v", err)
	}

	// Buffer para receber a resposta
	reply := make([]byte, 1500)
	n, peer, err := conn.ReadFrom(reply)
	if err != nil {
		log.Fatalf("Falha ao receber resposta: %v", err)
	}
	duration := time.Since(start)

	// Decodifica a resposta
	replyMsg, err := icmp.ParseMessage(ProtocolICMP, reply[:n])
	if err != nil {
		log.Fatalf("Falha ao decodificar resposta: %v", err)
	}

	// Verifica se é uma resposta do tipo Echo Reply
	switch replyMsg.Type {
	case ipv4.ICMPTypeEchoReply:
		echoReply, ok := replyMsg.Body.(*icmp.Echo)
		if !ok {
			log.Fatalf("Resposta ICMP inválida")
		}
		fmt.Printf("Resposta de %s: bytes=%d tempo=%v\n", peer, len(echoReply.Data), duration)
	default:
		log.Printf("Resposta inesperada: %+v\n", replyMsg)
	}
}
