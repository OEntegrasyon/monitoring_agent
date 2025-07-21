// Copyright (C) 2025 Özgür Entegrasyon
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

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
