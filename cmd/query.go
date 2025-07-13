package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

var BaseStyle = lipgloss.NewStyle().Bold(true).BorderStyle(lipgloss.ThickBorder()).BorderForeground(lipgloss.Color("63")).PaddingTop(1).PaddingBottom(1).Background(lipgloss.Color("#434343"))
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "query for a stock's price",
	Long:  "query for a stock's price, an arguments it required to quote a price",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load()
		finnhubClient := getFinnhubClient()
		res, _, _ := finnhubClient.Quote(context.Background(), args[0])

		var priceDifference = res.C - res.Pc
		var priceDifferncePercent = (res.C - res.Pc) * 100 / res.Pc

		s := fmt.Sprintf("💰 The current price for %s = %+v$  %+v$ ( %+v% % )", args[0], res.C, priceDifference, priceDifferncePercent)
		fmt.Printf("%s", BaseStyle.Render(s))

	},

}
