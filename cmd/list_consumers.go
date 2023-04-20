/*
 * Copyright (c) Blue Safespring AB - Jan Johansson <jj@safespring.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package cmd

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/safespring/cloutility-api-client/cloutapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// consumersCmd represents the consumers command
var consumersCmd = &cobra.Command{
	Use:   "consumers",
	Short: "list consumers will list the available consumers / consumption-units",
	Long: `
The command 'list consumers' will list all the available consumers / consumption-units 
for the current user account`,
	Run: func(cmd *cobra.Command, args []string) {
		listConsumers()
	},
}

func listConsumers() {
	client, err := cloutapi.Init(
		context.Background(),
		viper.GetString("client_id"),
		viper.GetString("client_origin"),
		viper.GetString("username"),
		viper.GetString("password"),
		viper.GetString("url"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	twriter := new(tabwriter.Writer)
	twriter.Init(os.Stdout, 8, 8, 1, '\t', 0)
	defer twriter.Flush()

	user, err := client.GetUser()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "ID", "Name", "Creation date", "Url")
	fmt.Fprintf(twriter, "%s\t%s\t%s\t%s\n", "--", "----", "-------------", "---")

	cUnits, err := client.GetConsumers(user.BusinessUnit.ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, cUnit := range cUnits {
		fmt.Fprintf(twriter, "%v\t%s\t%s\t%s\n", cUnit.ID, cUnit.Name, cUnit.CreatedDate.Format(time.ANSIC), cUnit.Href)
	}
}

func init() {
	listCmd.AddCommand(consumersCmd)
}
