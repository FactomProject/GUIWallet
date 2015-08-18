package main

import (
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	//"github.com/google/gxui/gxfont"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"

	//"github.com/FactomProject/GUIWallet/FactoidAPI"
	"github.com/FactomProject/fctwallet/Wallet"

	"fmt"
)

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
	holder.AddPanel(overview(theme), "Overview")
	holder.AddPanel(send(theme), "Send")
	holder.AddPanel(receive(theme), "Receive")
	holder.AddPanel(transactions(theme), "Transactions")

	window := theme.CreateWindow(800, 450, "Factoid Wallet")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

/*
func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 75)
	if err != nil {
		panic(err)
	}

	window := theme.CreateWindow(800, 450, "Factoid Wallet")
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Factoid Wallet")

	window.AddChild(label)

	window.OnClose(driver.Terminate)
}*/

func test() {
	fmt.Printf("Test\n")
	tx, err := GetTransactions()
	fmt.Printf("%v, %v\n", tx, err)
}

func getAddressesString() string {
	addresses, err := GetAddresses()
	if err != nil {
		return err.Error()
	}
	answer := "FactoidAddresses:\n"
	for _, v := range addresses.FactoidAddresses {
		answer += fmt.Sprintf("%v\t%v\t%v\n", v.Name, v.Address, v.Balance)
	}
	answer += "ECAddresses:\n"
	for _, v := range addresses.ECAddresses {
		answer += fmt.Sprintf("%v\t%v\t%v\n", v.Name, v.Address, v.Balance)
	}
	return answer
}

func getBalancesString() string {
	balances, err := GetBalances()
	if err != nil {
		return err.Error()
	}
	answer := fmt.Sprintf("Factoid Balance:\t%v\n", balances.FactoidBalances)
	answer += fmt.Sprintf("EC Balance:\t%v\n", balances.ECBalances)
	return answer
}

func overview(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()
	layout.SetPadding(math.ZeroSpacing)
	layout.SetMargin(math.ZeroSpacing)
	layout.SetSize(math.Size{W: 400, H: 400})

	font := theme.DefaultFont()

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Overview")
	layout.AddChild(label)

	textBox := theme.CreateTextBox()
	textBox.SetFont(font)
	textBox.SetText("")
	textBox.SetDesiredWidth(800)

	button := theme.CreateButton()
	button.SetHorizontalAlignment(gxui.AlignCenter)
	button.SetText("Test")
	button.OnClick(func(gxui.MouseEvent) { test() })
	layout.AddChild(button)

	button2 := theme.CreateButton()
	button2.SetHorizontalAlignment(gxui.AlignCenter)
	button2.SetText("GetAddresses")
	button2.OnClick(func(gxui.MouseEvent) {
		balances := getAddressesString()
		textBox.SetText(balances)
	})
	layout.AddChild(button2)

	button3 := theme.CreateButton()
	button3.SetHorizontalAlignment(gxui.AlignCenter)
	button3.SetText("GetBalances")
	button3.OnClick(func(gxui.MouseEvent) {
		balances := getBalancesString()
		textBox.SetText(balances)
	})
	layout.AddChild(button3)

	layout.AddChild(textBox)

	return layout
}

func send(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()

	font := theme.DefaultFont()

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Send")

	layout.AddChild(label)

	return layout
}

func receive(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()

	font := theme.DefaultFont()

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("receive")

	layout.AddChild(label)

	return layout
}

func transactions(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()

	font := theme.DefaultFont()

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Transactions")

	layout.AddChild(label)

	return layout
}

func main() {
	initWallet()
	gl.StartDriver(appMain)
}

func initWallet() {
	fmt.Printf("\ninitWallet\n")
	keys, _ := Wallet.GetWalletNames()
	if len(keys) == 0 {
		for i := 1; i <= 10; i++ {
			name := fmt.Sprintf("%02d-Fountain", i)
			_, err := Wallet.GenerateFctAddress([]byte(name), 1, 1)
			if err != nil {
				fmt.Printf("\nError - %v\n", err)
				return
			}
		}
	}
	fmt.Printf("initWallet done\n")
}
