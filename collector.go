package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

func getMemoryUsage() float64 {
	data, _ := os.ReadFile("/proc/meminfo")
	lines := strings.Split(string(data), "\n")

	memTotal := 0
	memAvailable := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			memTotal, _ = strconv.Atoi(fields[1]) // in KB
		} else if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			memAvailable, _ = strconv.Atoi(fields[1])
		}
	}

	if memTotal == 0 {
		return 0
	}
	used := memTotal - memAvailable
	return (float64(used) / float64(memTotal)) * 100
}

func getCPUUsage() float64 {
	startIdle, startTotal := readCPU()
	time.Sleep(200 * time.Millisecond)
	endIdle, endTotal := readCPU()

	idleTicks := float64(endIdle - startIdle)
	totalTicks := float64(endTotal - startTotal)

	if totalTicks == 0 {
		return 0
	}
	return (1.0 - idleTicks/totalTicks) * 100
}

func readCPU() (idle, total int) {
	data, _ := os.ReadFile("/proc/stat")
	fields := strings.Fields(strings.Split(string(data), "\n")[0])
	user, _ := strconv.Atoi(fields[1])
	nice, _ := strconv.Atoi(fields[2])
	system, _ := strconv.Atoi(fields[3])
	idle, _ = strconv.Atoi(fields[4])
	iowait, _ := strconv.Atoi(fields[5])
	irq, _ := strconv.Atoi(fields[6])
	softirq, _ := strconv.Atoi(fields[7])

	total = user + nice + system + idle + iowait + irq + softirq
	return idle, total
}
