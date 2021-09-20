package regexp

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/dlclark/regexp2"
)

var (
	expressions = []string{
		"((\\{.*?\\})?((obc.client.dropped.points.sum.60)|(obc.client.enqueue.points.sum.60)))",
		//		"(?!(\\{.*?\\})?(((zippydb|ZippyDB)\\..*$)|(\\..*$)|(search\\.sigrid\\.predictor.*$)|(unicorn\\.(rackaggr|index|term|doc)\\..*$)))((\\{.*?\\})?((.*\\.num_errors\\.count\\.60)|(.*\\.num_requests\\.count\\.60)|(SREventBase\\.eventbase_busy_time\\.p5.60)|(SREventBase\\.eventbase_busy_time\\.p95.60)|(SREventBase\\.eventbase_busy_time\\.p99.60)|(SREventBase\\.eventbase_busy_time\\.sum\\.60)|(SREventBase\\.eventbase_idle_time\\.p99.60)|(SREventBase\\.eventbase_idle_time\\.sum\\.60)|(SREventBase\\.io_threads_count)|(SRSelection\\.eventbase_busy_time\\.p99.60)|(SRSelection\\.eventbase_busy_time\\.sum\\.60)|(SRSelection\\.eventbase_idle_time\\.p99.60)|(SRSelection\\.eventbase_idle_time\\.sum\\.60)))",
		"((\\{.*?\\})?((cpu_load.*\\.avg\\.60)|(cpu_util.*\\.avg\\.60)|(memory_free.*\\.avg\\.60)|(memory_usage.*\\.avg\\.60)|(periodictask:.+)|(thrift\\..*dropped_conns\\.count\\.60)|(thrift\\..*killed_tasks\\.count\\.60)|(thrift\\..*num_calls\\.sum\\.60)|(thrift\\..*num_exceptions\\.sum\\.60)|(thrift\\..*num_processed\\.sum\\.60)|(thrift\\..*process_delay\\.p99\\.60)|(thrift\\..*process_time\\.p99\\.60)|(thrift\\..*queue_timeouts\\.count\\.60)|(thrift\\..*rejected_conns\\.count\\.60)|(thrift\\..*server_overloaded\\.count\\.60)|(thrift\\..*time_process_us\\.avg\\.60)|(thrift\\..*timeout_tasks\\.count\\.60)|(thrift\\.accepted_connections\\.count\\.60)|(thrift\\.active_requests\\.avg\\.60)|(thrift\\.adaptive_ideal_rtt_us)|(thrift\\.adaptive_min_concurrency)|(thrift\\.adaptive_sampled_rtt_us)|(thrift\\.admission_control\\..*)|(thrift\\.eventbase_busy_time\\.p99.60)|(thrift\\.eventbase_idle_time\\.p99.60)|(thrift\\.max_requests)|(thrift\\.num_active_requests)|(thrift\\.num_busy_pool_workers)|(thrift\\.num_idle_pool_workers)|(thrift\\.queued_requests)|(thrift\\.queued_requests\\.avg\\.60)|(thrift\\.queuelag.*\\.sum\\.60)|(thrift\\.queues\\..*\\.60)|(thrift\\.received_requests\\.count\\.60)|(thrift\\.server_load)))",
	}

	strings = []string{
		"system.cpu-idle",
		"system.cpu-iowait",
		"system.cpu-nice",
		"system.cpu-sys",
		"system.cpu-softirq",
		"system.cpu-hardirq",
		"system.cpu-user",
		"system.cpu-busy-pct",
		"system.cpu-util-pct",
		"system.load-1",
		"system.load-5",
		"system.uptime",
		"system.mem_free",
		"system.mem_free_nobuffer_nocache",
		"system.mem_used",
		"system.mem-util-pct",
		"system.mem_slab",
		"system.mem_anon",
		"system.mem_kernel",
		"system.mem_total",
		"system.n_eth-rxbyt_s",
		"system.n_eth-txbyt_s",
		"system.n_eth-util-pct",
		"system.n_eth-rxpck_s",
		"system.n_eth-txpck_s",
		"system.n_eth-rxerr_s",
		"system.n_eth-txerr_s",
		"system.rx_errors",
		"system.tx_errors",
		"system.net.tcp.attempt_fails_per_s",
		"system.net.tcp.estab_rsts_per_s",
		"system.net.tcp.rsts_per_s",
		"system.net.tcp.rxmits_per_s",
		"system.net.tcp.socket_count",
		"system.net.tcp.socket_time_wait_count",
		"system.net.tcp6.socket_count",
		"system.net.udp.error_drop_per_s",
		"system.connection_count_estab",
		"system.connection_count_total",
		"system.swap_free",
		"system.swap_used",
		"system.swap_pagein_per_s",
		"system.swap_pageout_per_s",
	}
)

func BenchmarkStd(b *testing.B) {
	for idx, expr := range expressions {
		b.Run(fmt.Sprintf("expr_%d", idx), func(b *testing.B) {
			r := regexp.MustCompile(expr)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, s := range strings {
					r.MatchString(s)
				}
			}
		})
	}
}

func BenchmarkDlclark(b *testing.B) {
	for idx, expr := range expressions {
		b.Run(fmt.Sprintf("expr_%d", idx), func(b *testing.B) {
			r := regexp2.MustCompile(expr, regexp2.None)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				for _, s := range strings {
					r.MatchString(s)
				}
			}
		})
	}
}
