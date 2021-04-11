package gui

import (
	"github.com/gonutz/wui/v2"
)

const (
	font           = "Segoe UI"
	normalFontSize = -11
	bigFontSize    = -14
)

func Window() {
	windowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   font,
		Height: normalFontSize,
	})

	labelFont, _ := wui.NewFont(wui.FontDesc{
		Name:   font,
		Height: bigFontSize,
		Bold:   true,
	})

	normalFont, _ := wui.NewFont(wui.FontDesc{
		Name:   font,
		Height: normalFontSize,
	})

	startButtonFont, _ := wui.NewFont(wui.FontDesc{
		Name:   font,
		Height: normalFontSize,
		Bold:   true,
	})

	window := wui.NewWindow()
	window.SetFont(windowFont)
	window.SetInnerSize(310, 266)
	window.SetTitle("Go to sleep")
	window.SetHasMaxButton(false)
	window.SetResizable(false)

	methodPanelFont, _ := wui.NewFont(wui.FontDesc{
		Name:   font,
		Height: normalFontSize,
	})

	methodPanel := wui.NewPanel()
	methodPanel.SetFont(methodPanelFont)
	methodPanel.SetBounds(10, 12, 290, 64)
	methodPanel.SetBorderStyle(wui.PanelBorderSingleLine)
	window.Add(methodPanel)

	methodLabel := wui.NewLabel()
	methodLabel.SetFont(labelFont)
	methodLabel.SetBounds(10, 10, 150, 13)
	methodLabel.SetText("Sleep Method")
	methodPanel.Add(methodLabel)

	sleepMethodSelection := wui.NewComboBox()
	sleepMethodSelection.SetBounds(8, 30, 270, 21)
	sleepMethodSelection.SetItems([]string{
		"Hibernate",
		"Sleep",
	})
	sleepMethodSelection.SetSelectedIndex(0)
	methodPanel.Add(sleepMethodSelection)

	timePanel := wui.NewPanel()
	timePanel.SetFont(normalFont)
	timePanel.SetBounds(10, 84, 290, 98)
	timePanel.SetBorderStyle(wui.PanelBorderSingleLine)
	window.Add(timePanel)

	timeLabel := wui.NewLabel()
	timeLabel.SetFont(labelFont)
	timeLabel.SetBounds(10, 10, 150, 13)
	timeLabel.SetText("Time")
	timePanel.Add(timeLabel)

	comboBox1 := wui.NewComboBox()
	comboBox1.SetBounds(8, 30, 270, 21)
	comboBox1.SetItems([]string{
		"Countdown",
		"Time",
	})
	comboBox1.SetSelectedIndex(0)
	timePanel.Add(comboBox1)

	editLine1 := wui.NewEditLine()
	editLine1.SetBounds(8, 60, 270, 20)
	editLine1.SetText("02:00")
	timePanel.Add(editLine1)

	startButton := wui.NewButton()
	startButton.SetFont(startButtonFont)
	startButton.SetBounds(10, 212, 290, 35)
	startButton.SetText("Start")
	window.Add(startButton)

	window.Show()
}
