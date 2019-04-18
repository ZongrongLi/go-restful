package view

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang/glog"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/tiancai110a/go-rpc/service"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(ctx context.Context, resp *service.Resp) {
	fmt.Println("health check")
	message := "OK"
	glog.Info("ok")
	resp.Add("status", message)
}

// DiskCheck checks the disk usage.
func DiskCheck(ctx context.Context, resp *service.Resp) {
	u, _ := disk.Usage("/")

	usedMB := int64(u.Used) / MB
	usedGB := int64(u.Used) / GB
	totalMB := int64(u.Total) / MB
	totalGB := int64(u.Total) / GB
	usedPercent := int64(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusOK
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	glog.Infof("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%% status:%d", text, usedMB, usedGB, totalMB, totalGB, usedPercent, status)
	resp.Add("disk_status", text)
	resp.Add("Free_space(MB)", strconv.FormatInt(usedMB, 10))
	resp.Add("total_space(MB)", strconv.FormatInt(totalMB, 10))
	resp.Add("usedPercent", strconv.FormatInt(usedPercent, 10))
}

// CPUCheck checks the cpu usage.
func CPUCheck(ctx context.Context, resp *service.Resp) {
	cores, _ := cpu.Counts(false)

	a, _ := load.Avg()
	l1 := a.Load1
	l5 := a.Load5
	l15 := a.Load15

	status := http.StatusOK
	text := "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	glog.Infof("%s - Load average: %.2f, %.2f, %.2f | Cores: %d status:%d", text, l1, l5, l15, cores, status)

	resp.Add("cpu_status", text)
	resp.Add("Load_average_l1", strconv.FormatFloat(l1, 'f', 2, 64))
	resp.Add("Load_average_l5", strconv.FormatFloat(l5, 'f', 2, 64))
	resp.Add("Load_average_l15", strconv.FormatFloat(l15, 'f', 2, 64))
	resp.Add("cores", strconv.FormatInt(int64(cores), 10))
}

// RAMCheck checks the disk usage.
func RAMCheck(ctx context.Context, resp *service.Resp) {
	u, _ := mem.VirtualMemory()

	usedMB := int64(u.Used) / MB
	usedGB := int64(u.Used) / GB
	totalMB := int64(u.Total) / MB
	totalGB := int64(u.Total) / GB
	usedPercent := int64(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	glog.Infof("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%% status: %d", text, usedMB, usedGB, totalMB, totalGB, usedPercent, status)
	resp.Add("ram_status", text)
	resp.Add("Free_space(MB)", strconv.FormatInt(usedMB, 10))
	resp.Add("total_space(MB)", strconv.FormatInt(totalMB, 10))
	resp.Add("usedPercent", strconv.FormatInt(usedPercent, 10))
}
