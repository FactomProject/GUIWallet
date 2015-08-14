package main

import (

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	//"github.com/google/gxui/gxfont"
	"github.com/google/gxui/samples/flags"
	"github.com/google/gxui/math"

	"github.com/FactomProject/GUIWallet/FactoidAPI"
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
	FactoidAPI.GetAddresses()
}

func getBalances() string {
	return FactoidAPI.GetAddresses()
}

func overview(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()

	font := theme.DefaultFont()
	
	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Overview")
	layout.AddChild(label)
	
	textBox := theme.CreateTextBox()
	textBox.SetFont(font)
	textBox.SetText("")
	textBox.SetSize(math.Size{W: 300, H: 300})

	button := theme.CreateButton()
	button.SetHorizontalAlignment(gxui.AlignCenter)
	button.SetText("Test")
	button.OnClick(func(gxui.MouseEvent) { test() })
	layout.AddChild(button)

	button2 := theme.CreateButton()
	button2.SetHorizontalAlignment(gxui.AlignCenter)
	button2.SetText("GetBalances")
	button2.OnClick(func(gxui.MouseEvent) {
			balances:=getBalances()
			textBox.SetText(balances)
			textBox.SetSize(math.Size{W: 400, H: 400})
		})

	layout.AddChild(button2)
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
	gl.StartDriver(appMain)
}
