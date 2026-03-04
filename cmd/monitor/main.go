package main

import (
	"fmt"

	"github.com/backend-study-org/site-monitor/internal/checker"
)

func main() {
	fmt.Println("Site Monitor started")
	workingSites := []string{
		"https://google.com",
		"https://youtube.com",
		"https://yandex.ru",
		"https://vk.com",
		"https://mail.ru",
		"https://ok.ru",
		"https://avito.ru",
		"https://wildberries.ru", // почему-то не работает
		"https://ozon.ru",        // почему-то не работает
		"https://gismeteo.ru",
		"https://2gis.ru",
		"https://gosuslugi.ru",
		"https://rbc.ru",
		"https://kinopoisk.ru",
		"https://hh.ru",
	}
	notWorkingSites := []string{
		"http://nonexistent-12345-test-domain.com",
		"http://this-domain-should-not-exist-xyz.net",
		"http://invalid_domain",
		"http://256.256.256.256",
		"http://example.invalid",
		"http://no-such-hostname-abcdefg.local",
		"http://unknown-host-foo-bar.baz",
		"http://notregistered-tld.example.abc",
		"http://expired-domain-123456789.com",
		"http://test-domain-does-not-resolve.xyz",
		"http://unreachable-host-foo-bar-123.com",
		"http://connection-timeout-test-abcdef.com",
		"http://bad-protocol://example.com",
		"http://missing-tld-domain",
		"http://non-existent-subdomain.unknown-example.com",
	}
	sites := append(notWorkingSites, workingSites...)
	for _, site := range sites {
		res := checker.CheckSite(site)
		if res.Error != nil || res.Code != 200 {
			fmt.Printf("Site %s NOT ok %d\n", res.URL, res.Code)
		}
		if res.Code == 200 {
			fmt.Printf("Site %s ok\n", res.URL)
		}
	}
}
