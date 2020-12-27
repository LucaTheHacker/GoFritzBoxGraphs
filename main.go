/*
 * GoFritzBoxGraphs
 * Copyright (C) 2020-2020 Dametto Luca <https://damettoluca.com>
 *
 * main.go is part of GoFritzBoxGraphs
 *
 * You should have received a copy of the GNU Affero General Public License v3.0 along with GoFritzBoxGraphs.
 * If not, see <https://github.com/LucaTheHacker/GoFritzBoxGraphs/blob/main/LICENSE>.
 */

package main

import (
	"github.com/LucaTheHacker/GoFritzBox"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Print("Fritz!Box endpoint (http://192.168.188.1): ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	endpoint := input.Text()
	if strings.Index(endpoint, "http") != 0 {
		endpoint = "http://" + endpoint
	}

	fmt.Print("Username: ")
	input = bufio.NewScanner(os.Stdin)
	input.Scan()
	username := input.Text()

	fmt.Print("Password: ")
	input = bufio.NewScanner(os.Stdin)
	input.Scan()
	password := input.Text()

	fmt.Print("\033[H\033[2J")

	log.Println("Trying to login...")
	session, err := GoFritzBox.Login(endpoint, username, password)
	if err != nil {
		panic(err)
	}
	log.Println("Login successful")

	log.Println("Starting download...")
	data, err := session.GetAssistanceData()
	if err != nil {
		panic(err)
	}

	log.Println("Download ended.")
	err = ioutil.WriteFile("FritzBoxExport", data, 0644)
	if err != nil {
		panic(err)
	}
	log.Println("Document saved, filename: 'FritzBoxExport'")
}
