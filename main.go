package main

import (
	"strconv"
	"time"
)

func main() {
	cfg := LoadConfig("/etc/giysmon-agent.yaml")
	RabbitURL := "amqp://" + cfg.RabbitUser + ":" + cfg.RabbitPass + "@" + cfg.RabbitIP + ":" + cfg.RabbitPort + "/"
	sender := NewSender(RabbitURL)
	defer sender.Close()

	for {
		loop_count := cfg.Interval / 5
		cpu_total := float64(0)
		memory_total := float64(0)
		for i := 0; i < loop_count; i++ {
			cpu := getCPUUsage()
			memory := getMemoryUsage()
			cpu_total += float64(cpu)
			memory_total += float64(memory)
			time.Sleep(5 * time.Second)
		}
		cpu_avg := cpu_total / float64(loop_count)
		memory_avg := memory_total / float64(loop_count)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		data := "[" + timestamp + "] Mem: " + strconv.FormatFloat(memory_avg, 'f', 2, 64) + "% CPU: " + strconv.FormatFloat(cpu_avg, 'f', 2, 64) + "%"
		sender.Send(cfg.Hostname + ": " + data)
	}
}
