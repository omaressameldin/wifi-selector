package cmd

import (
	"fmt"
	"os/exec"

	"github.com/omaressameldin/wifi-selector/network"
	"github.com/omaressameldin/wifi-selector/utils"
	"github.com/spf13/cobra"
)

func deleteWifi(selectedNet network.SavedNetwork) {
	c := exec.Command(
		"bash",
		"-c",
		fmt.Sprintf("nmcli connection delete \"%v\"", selectedNet.Name),
	)
	o, err := c.Output()
	utils.Must(err)
	fmt.Println(string(o))
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete selected wifi network",
	Run: func(cmd *cobra.Command, args []string) {
		wifiLists := utils.GetSavedWifis()
		selected := utils.SelectFromList("Select wifi network you want to delete", wifiLists, "💣")
		deleteWifi(wifiLists[selected])
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
