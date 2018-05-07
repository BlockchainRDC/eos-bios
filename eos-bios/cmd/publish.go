// Copyright © 2018 Alexandre Bourget <alex@eoscanada.com>

package cmd

import (
	"fmt"
	"os"

	bios "github.com/eoscanada/eos-bios"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish my discovery file to the seed network",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		info, ipfs := ipfsClient()

		fmt.Printf("Reading discovery file... ")
		if err := bios.ValidateDiscoveryFile(myDiscoveryFile); err != nil {
			fmt.Println("error")
			fmt.Fprintf(os.Stderr, "format invalid: %s", err)
			os.Exit(1)
		}

		fl, err := os.Open(myDiscoveryFile)
		if err != nil {
			fmt.Println("failed")
			fmt.Fprintf(os.Stderr, "error opening %q: %s\n", myDiscoveryFile, err)
			os.Exit(2)
		}
		defer fl.Close()
		fmt.Println("ok")

		// TODO: load private keys, or use a keos wallet if configured as such

		// load into `disco`..
		var disco *Discovery
		api.SignPushActions(
,			disco.NewUpdateDiscovery(disco),
		)

		// fmt.Printf("Publishing discovery file... ")
		// newObj, err := ipfs.Add(fl)
		// if err != nil {
		// 	fmt.Println("failed")
		// 	fmt.Fprintln(os.Stderr, "error adding content to ipfs:", err)
		// 	os.Exit(3)
		// }
		// fmt.Println("/ipfs/" + newObj + " published")

		// fmt.Printf("Updating IPNS link /ipns/" + info.ID + " ... ")
		// if err = ipfs.Publish("", newObj); err != nil {
		// 	fmt.Println("failed")
		// 	fmt.Fprintln(os.Stderr, "error publishing new ipns address:", err)
		// 	os.Exit(4)
		// }
		// fmt.Println("ok")
		// fmt.Println("")
		// fmt.Println("")
		// fmt.Println("You can now provide your IPNS link to your network in this form:")
		// fmt.Println("")
		// fmt.Println("    /ipns/" + info.ID)
		// fmt.Println("")
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)
	// publishCmd.PersistentFlags().StringVarP(&ipfsAPIAddress, "ipfs-api-address", "", "/ip4/127.0.0.1/tcp/5001", "API Endpoint of the local IPFS node, to publish content. Make sure it matches your running `ipfs` instance.")

	// for _, flag := range []string{"ipfs-api-address"} {
	// 	viper.BindPFlag(flag, publishCmd.Flags().Lookup(flag))
	// }
}