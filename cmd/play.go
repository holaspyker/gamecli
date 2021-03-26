/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		play()
	},
}

type Question struct {
	Country string
	Capital []string
}

func init() {
	rootCmd.AddCommand(playCmd)

}

func play() {

	response, err := http.Get("http://localhost:8080/questions")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//fmt.Println(response.Body)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))
	var question []Question
	json.Unmarshal([]byte(responseData), &question)
	fmt.Println("what is the Capital of", question[0].Country)
	for _, c := range question[0].Capital {
		fmt.Print(c, " ")
	}
	fmt.Println()

	/*for _, s := range question {
		fmt.Println("what is the Capital of ", s.Country)
		for _, c := range s.Capital {
			fmt.Print(c, " ")
		}
		fmt.Println()

	}*/

}
