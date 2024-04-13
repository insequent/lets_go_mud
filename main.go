package main

import (
	"flag"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"

	"github.com/insequent/lets_go_mud/telnet"
)

func main() {
	config := flag.String("c", "", "The config file to use for settings")
	host := flag.String("s", "3k.org", "The MUD server to the connect to")
	port := flag.Int("p", 3000, "The port the MUD server is listening on")

	flag.Parse()

	if *config != "" {
		if err := parseConfig(*config); err != nil {
			log.Fatalf("Failed to parse given configuration file (%s): %v", *config, err)
		}
	}

	client, err := telnet.NewClient(*host, *port)
	if err != nil {
		log.Fatalf("Failed to initialize telnet client: %v", err)
	}

	if err := client.Dial(); err != nil {
		log.Fatalf("Failed to dial remote server %s: %v", *host, err)
	}

	//client.StartAndListen()
	screen()
}

func parseConfig(config string) error {
	return nil
}

func screen() {
	if _, err := tea.NewProgram(screen.NewModel(), tea.WithAltScreen()).Run(); err != nil {
		fmt.Println("Error while running program:", err)
		os.Exit(1)
	}
}
