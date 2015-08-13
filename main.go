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

func overview(theme gxui.Theme) gxui.Control {
	layout := theme.CreateLinearLayout()

	font := theme.DefaultFont()
	
	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Overview")

	layout.AddChild(label)

	button := theme.CreateButton()
	button.SetHorizontalAlignment(gxui.AlignCenter)
	button.SetText("Test")
	button.OnClick(func(gxui.MouseEvent) { test() })
	layout.AddChild(button)

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
